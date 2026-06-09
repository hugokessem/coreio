package utils

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Config struct {
	MaxRetries int
	Timeout    time.Duration
}

var once sync.Once
var httpClient *http.Client

func GetHTTPClient() *http.Client {
	once.Do(func() {
		dialer := &net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}

		transport := &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          500,
			MaxIdleConnsPerHost:   200,
			MaxConnsPerHost:       500,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     false,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS12,
			},
		}

		httpClient = &http.Client{
			Timeout:   30 * time.Second,
			Transport: transport,
		}
	})

	return httpClient
}

func DoPostWithRetry(
	url string,
	body string,
	cfg Config,
	headers map[string]string,
) (*http.Response, error) {

	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = 3
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 30 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	client := GetHTTPClient()
	var lastErr error
	for attempt := 0; attempt < cfg.MaxRetries; attempt++ {

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			url,
			strings.NewReader(body),
		)
		if err != nil {
			return nil, fmt.Errorf(
				"create request: %w",
				err,
			)
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)

		// Success
		if err == nil {

			if !shouldRetryStatus(resp.StatusCode) {
				return resp, nil
			}

			lastErr = fmt.Errorf(
				"received retryable status code: %d",
				resp.StatusCode,
			)

			drainAndClose(resp.Body)

		} else {
			lastErr = err

			// Context timeout/cancel
			if errors.Is(err, context.DeadlineExceeded) ||
				errors.Is(err, context.Canceled) {
				return nil, err
			}
		}

		// Last attempt
		if attempt == cfg.MaxRetries-1 {
			break
		}

		backoff := calculateBackoff(attempt)

		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, fmt.Errorf(
		"request failed after %d retries: %w",
		cfg.MaxRetries,
		lastErr,
	)
}

func shouldRetryStatus(status int) bool {
	switch status {
	case http.StatusInternalServerError, // 500
		http.StatusBadGateway,         // 502
		http.StatusServiceUnavailable, // 503
		http.StatusGatewayTimeout:     // 504
		return true

	default:
		return false
	}
}

func calculateBackoff(attempt int) time.Duration {

	base := time.Second * time.Duration(1<<attempt)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	jitter := time.Duration(
		r.Intn(500),
	) * time.Millisecond

	return base + jitter
}

func drainAndClose(body io.ReadCloser) {
	if body == nil {
		return
	}

	_, _ = io.Copy(io.Discard, body)
	_ = body.Close()
}

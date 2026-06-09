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
	MaxRetries     int
	Timeout        time.Duration
	RetryOnTimeout bool
}

var (
	once       sync.Once
	httpClient *http.Client
	rng        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetHTTPClient() *http.Client {
	once.Do(func() {
		dialer := &net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 60 * time.Second,
		}

		transport := &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DialContext:       dialer.DialContext,
			ForceAttemptHTTP2: true,

			MaxIdleConns:        2000,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     1000,

			IdleConnTimeout:       120 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 65 * time.Second,

			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}

		// Use context timeout as the single source of truth.
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   0,
		}
	})

	return httpClient
}

func DoPost(
	ctx context.Context,
	url string,
	body string,
	cfg Config,
	headers map[string]string,
) (*http.Response, error) {

	if cfg.Timeout <= 0 {
		cfg.Timeout = 70 * time.Second
	}

	client := GetHTTPClient()

	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		strings.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, classifyError(err)
	}

	return resp, nil
}

func DoGetWithRetry(
	ctx context.Context,
	url string,
	cfg Config,
	headers map[string]string,
) (*http.Response, error) {

	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = 3
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 30 * time.Second
	}

	client := GetHTTPClient()

	var lastErr error

	for attempt := 0; attempt < cfg.MaxRetries; attempt++ {

		ctx, cancel := context.WithTimeout(
			ctx,
			cfg.Timeout,
		)

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			url,
			nil,
		)
		if err != nil {
			cancel()
			return nil, fmt.Errorf("create request: %w", err)
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)

		cancel()

		if err == nil {
			if !shouldRetryStatus(resp.StatusCode) {
				return resp, nil
			}

			lastErr = fmt.Errorf(
				"retryable status code: %d",
				resp.StatusCode,
			)

			drainAndClose(resp.Body)
		} else {

			if errors.Is(err, context.Canceled) {
				return nil, err
			}

			if errors.Is(err, context.DeadlineExceeded) {
				if !cfg.RetryOnTimeout {
					return nil, err
				}
			}

			lastErr = err

			if !isRetryableNetworkError(err) {
				return nil, err
			}
		}

		if attempt < cfg.MaxRetries-1 {
			time.Sleep(calculateBackoff(attempt))
		}
	}

	return nil, fmt.Errorf(
		"request failed after %d attempts: %w",
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

func isRetryableNetworkError(err error) bool {
	var netErr net.Error

	if errors.As(err, &netErr) {
		return true
	}

	return false
}

func calculateBackoff(attempt int) time.Duration {
	base := time.Second * time.Duration(1<<attempt)

	jitter := time.Duration(
		rng.Intn(500),
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

func classifyError(err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return fmt.Errorf("request timeout: %w", err)

	case errors.Is(err, context.Canceled):
		return fmt.Errorf("request canceled: %w", err)

	default:
		return err
	}
}

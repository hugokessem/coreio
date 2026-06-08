package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"math"
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

var (
	client *http.Client
	once   sync.Once
)

func getClient() *http.Client {
	once.Do(func() {
		dialer := &net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		client = &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy:                 http.ProxyFromEnvironment,
				DialContext:           dialer.DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          10,
				MaxIdleConnsPerHost:   100,
				MaxConnsPerHost:       500,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     false,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					MinVersion:         tls.VersionTLS12,
				},
			},
		}
	})

	return client
}

func DoPostWithRetry(url string, xmlBody string, config Config, headers map[string]string) (*http.Response, error) {
	var resp *http.Response
	var err error

	client := getClient()
	for attempt := 0; attempt < config.MaxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
		req, reqErr := http.NewRequestWithContext(
			ctx,
			"POST",
			url,
			strings.NewReader(xmlBody),
		)

		if reqErr != nil {
			cancel() // Ensure we cancel the context to free resources
			return nil, fmt.Errorf("failed to create request: %w", reqErr)
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err = client.Do(req)
		cancel() // Ensure we cancel the context to free resources
		if resp == nil {
			return nil, fmt.Errorf("response is nil: %w", err)
		}

		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}

		// Clean up response if we’ll retry!
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		// Exponential backoff with jitter
		backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		sleep := backoff + jitter

		fmt.Printf("Attempt %d/%d failed: %v, retrying in %v\n", attempt+1, config.MaxRetries, err, sleep)
		time.Sleep(sleep)
	}

	return nil, fmt.Errorf("request failed after %d retries: %w", config.MaxRetries, err)
}

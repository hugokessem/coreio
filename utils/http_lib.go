package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	MaxRetries int
	Timeout    time.Duration
}

func DoPostWithRetry(url string, xmlBody string, config Config, headers map[string]string) (*http.Response, error) {
	var resp *http.Response
	var err error

	client := &http.Client{
		Timeout: config.Timeout * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				MinVersion:         tls.VersionTLS13,
			},
			DisableKeepAlives: true,
			IdleConnTimeout:   10 * time.Second,
		},
	}

	for attempt := 0; attempt < config.MaxRetries; attempt++ {
		req, reqErr := http.NewRequest("POST", url, strings.NewReader(xmlBody))
		if reqErr != nil {
			return nil, fmt.Errorf("failed to create request: %w", reqErr)
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err = client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil // success or 4xx error (no retry)
		}

		// Clean up response if we’ll retry
		if resp != nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}

		// Exponential backoff with jitter
		backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		sleep := backoff + jitter

		fmt.Printf("Attempt %d/%d failed: %v, retrying in %v\n", attempt+1, config.MaxRetries, err, sleep)
		time.Sleep(sleep)
	}

	return nil, fmt.Errorf("request failed after %d retries: %w", config.MaxRetries, err)
}

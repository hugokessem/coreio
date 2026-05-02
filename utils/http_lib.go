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

// tlsConfig := &tls.Config{
//         // Set the minimum version to TLS 1.2
//         MinVersion: tls.VersionTLS12,

//         // Optional: Force ONLY TLS 1.2 (ignoring TLS 1.3)
//         // MaxVersion: tls.VersionTLS12,

//         // Optional: Specify secure cipher suites for TLS 1.2
//         CipherSuites: []uint16{
//             tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
//             tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
//         },
//     }

//     // Apply to an HTTP Transport
//     transport := &http.Transport{
//         TLSClientConfig: tlsConfig,
//     }

//     client := &http.Client{Transport: transport}

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
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS12,
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
		if resp == nil {
			return nil, fmt.Errorf("response is nil: %w", err)
		}

		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}

		// Clean up response if we’ll retry!
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

package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseResponse(resp *http.Response, target interface{}, successStatus ...int) error {
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusNoContent {
		log.Println("received empty response")
		return nil
	}
	statusIsSuccess := false

	for _, status := range successStatus {
		if resp.StatusCode == status {
			statusIsSuccess = true
			break
		}
	}

	if !statusIsSuccess {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		log.Printf("API call failed with status %d: %s", resp.StatusCode, string(responseBody))
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}

func NewRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

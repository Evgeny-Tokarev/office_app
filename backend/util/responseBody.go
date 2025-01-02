package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func OutputUnknownBody(body io.ReadCloser) {
	fmt.Println("trying to output body...")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close body: %v", err)
		}
	}(body)

	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	var raw json.RawMessage
	err = json.Unmarshal(b, &raw)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	fmt.Println(string(raw))
}

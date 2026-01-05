package cmdutil

import (
	"encoding/json"
	"fmt"
	"log"

	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/spf13/cobra"
)

func UnmarshalPayload(cmd *cobra.Command, body interface{}) error {
	var payload []byte
	var err error
	payloadFile, _ := cmd.Flags().GetString("payload-file")
	payloadRaw, _ := cmd.Flags().GetString("payload")

	if payloadFile != "" {
		payload, err = os.ReadFile(payloadFile)
	} else if payloadRaw != "" {
		payload = []byte(payloadRaw)
	} else {
		return fmt.Errorf("either --payload or --payload-file is required")
	}

	if err != nil {
		return fmt.Errorf("failed to read payload: %w", err)
	}
	err = json.Unmarshal(payload, body)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request body: %w", err)
	}
	return nil
}

func HandleError(err error, format string) {
	if err == nil {
		return
	}
	if openapiErr, ok := err.(datadog.GenericOpenAPIError); ok {
		log.Fatalf(format+": %v\n%s", err, string(openapiErr.Body()))
	}
	log.Fatalf(format+": %v", err)
}

func FormatJSON(res interface{}, filterType string) string {
	if filterType == "" {
		s, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal: %v", err)
		}
		return string(s)
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("failed to marshal for filtering: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatalf("failed to unmarshal for filtering: %v", err)
	}

	if data, ok := m["data"].([]interface{}); ok {
		var filteredData []interface{}
		for _, item := range data {
			if obj, ok := item.(map[string]interface{}); ok {
				t, ok := obj["type"].(string)
				if ok && t == filterType {
					filteredData = append(filteredData, obj)
				} else if !ok {
					// If there is no type field, don't filter it out
					filteredData = append(filteredData, obj)
				}
			} else {
				// If item is not an object, don't filter it out
				filteredData = append(filteredData, item)
			}
		}
		m["data"] = filteredData
	} else if obj, ok := m["data"].(map[string]interface{}); ok {
		// Single object data
		if t, ok := obj["type"].(string); ok && t != filterType {
			m["data"] = nil
		}
	}

	delete(m, "included")

	s, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}
	return string(s)
}

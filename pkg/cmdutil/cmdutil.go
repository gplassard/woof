package cmdutil

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintJSON(res interface{}, filterType string) {
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
				if t, ok := obj["type"].(string); ok && t == filterType {
					filteredData = append(filteredData, obj)
				}
			}
		}
		m["data"] = filteredData
	}

	delete(m, "included")

	s, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}
	fmt.Println(string(s))
}

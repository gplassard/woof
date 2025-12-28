package cmdutil

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintJSON(res interface{}, filterType string) {
	if filterType == "" {
		s, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal: %v", err)
		}
		fmt.Println(string(s))
		return
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
	fmt.Println(string(s))
}

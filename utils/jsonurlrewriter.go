package utils

import (
	"encoding/json"
	"strings"
)

// RewriteURLsInJSON takes a JSON byte slice and rewrites all occurrences of the targetURL with replacementURL.
func RewriteURLsInJSON(jsonData []byte, targetURL, replacementURL string) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	// Define the recursive function to search and replace URLs in the decoded JSON.
	var replaceURLs func(data interface{})
	replaceURLs = func(data interface{}) {
		switch v := data.(type) {
		case map[string]interface{}:
			for key, value := range v {
				if strValue, ok := value.(string); ok && strings.Contains(strValue, targetURL) {
					v[key] = strings.ReplaceAll(strValue, targetURL, replacementURL)
				} else {
					replaceURLs(value)
				}
			}
		case []interface{}:
			for i, item := range v {
				replaceURLs(item)
				v[i] = item
			}
		}
	}

	// Apply the URL replacement.
	replaceURLs(data)

	// Re-encode the modified data to JSON.
	return json.Marshal(data)
}

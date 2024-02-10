package utils

import (
	"strings"
)

// RewriteURLsInJSONStringBased takes a JSON byte slice and rewrites all occurrences of the targetURL with replacementURL, attempting to preserve the order of elements.
func RewriteURLsInJSONStringBased(jsonData []byte, targetURL, replacementURL string) ([]byte, error) {
	// This is a simplified example that directly manipulates the JSON as a string.
	// For real applications, more sophisticated parsing and replacement logic might be required.
	jsonString := string(jsonData)
	// A simple string replacement; this might need to be more complex to avoid false matches.
	modifiedString := strings.ReplaceAll(jsonString, targetURL, replacementURL)

	return []byte(modifiedString), nil
}

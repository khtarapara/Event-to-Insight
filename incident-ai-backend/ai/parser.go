package ai

import (
	"strings"
)

func ParseAIClassification(content string) (string, string) {
	lines := strings.Split(content, "\n")
	var severity, category string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Severity:") {
			severity = strings.TrimSpace(strings.TrimPrefix(line, "Severity:"))
		}
		if strings.HasPrefix(line, "Category:") {
			category = strings.TrimSpace(strings.TrimPrefix(line, "Category:"))
		}
	}
	return severity, category
}

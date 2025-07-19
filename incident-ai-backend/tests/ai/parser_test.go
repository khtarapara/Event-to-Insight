package tests

import (
	"fmt"
	"incident-ai-backend/ai"
	"testing"
)

func TestParseAIClassification(t *testing.T) {
	content := `
	Severity: High
	Category: Software
	`
	severity, category := ai.ParseAIClassification(content)
	fmt.Println(severity, category)
	if severity != "High" || category != "Software" {
		t.Error("ParseAIClassification failed")
	}
}

package tests

import (
	"incident-ai-backend/ai"
	"testing"
)

func TestMakePromptForEventClassification(t *testing.T) {
	title := "Test Incident"
	description := "This is a test incident description."
	expectedPrompt := `
	Given the following incident report:

	Title: Test Incident
	Description: This is a test incident description.

	Classify the severity (Low, Medium, High, Critical) and category (Network, Software, Hardware, Security).
	Return result in this format:
	Severity: <value>
	Category: <value>
	`

	prompt := ai.MakePromptForEventClassification(title, description)
	if prompt != expectedPrompt {
		t.Errorf("Expected prompt:\n%s\nGot:\n%s", expectedPrompt, prompt)
	}
}

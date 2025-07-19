package ai

import "fmt"

func MakePromptForEventClassification(title, description string) string {
	return fmt.Sprintf(`
	Given the following incident report:

	Title: %s
	Description: %s

	Classify the severity (Low, Medium, High, Critical) and category (Network, Software, Hardware, Security).
	Return result in this format:
	Severity: <value>
	Category: <value>
	`, title, description)
}

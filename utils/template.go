package utils

import (
	"bytes"
	"text/template"
)

// TemplateString is a helper function to render a template string with the given data
func TemplateString(base string, data any) (string, error) {
	// Build the template with the given base
	tmp, err := template.New("template").Parse(base)
	if err != nil {
		return "", err
	}

	// Render the template with the given data
	var out bytes.Buffer
	err = tmp.Execute(&out, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

package template

import (
	"bytes"
	"html/template"
)

// Get a template object from a string
func GetTemplate(text string) *template.Template {
	t := template.Must(template.New("template").Parse(text))

	return t
}

// Get document based on template and data
func CreateDocument(t *template.Template, d map[string]interface{}) string {
	var b bytes.Buffer
	t.Execute(&b, d)

	return b.String()
}

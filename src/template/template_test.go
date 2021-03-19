package template

import "testing"

func TestTemplate(t *testing.T) {
	//Given
	templateText := `This is a testing document
	string value: {{ .stringValue }}`

	// When
	template := GetTemplate(templateText)

	// Then
	if template == nil {
		t.Errorf("Template was nil")
	}
}

func TestDocument(t *testing.T) {
	//Given
	templateText := `This is a testing document
	string value: {{ .stringValue }}`

	expectedDocument := `This is a testing document
	string value: testingValue`

	template := GetTemplate(templateText)

	data := map[string]interface{}{
		"stringValue": "testingValue",
	}

	// When
	d := CreateDocument(template, data)

	// Then
	if d != expectedDocument {
		t.Errorf("Document did not match expectation")
	}
}

func TestDocumentWithList(t *testing.T) {
	//Given
	templateText := `This is a testing document
	{{ range .things -}}
	  - {{ . }}
  	{{ end -}}`

	expectedDocument := `This is a testing document
	- thingA
  	- thingB
  	`

	template := GetTemplate(templateText)

	data := map[string]interface{}{
		"things": []string{"thingA", "thingB"},
	}

	// When
	d := CreateDocument(template, data)

	// Then
	if d != expectedDocument {
		t.Errorf("Document did not match expectation")
	}
}

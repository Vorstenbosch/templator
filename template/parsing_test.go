package template

import "testing"

func TestAllRequiredVariablesWithList(t *testing.T) {
	//Given
	templateText := `This is a testing document
	{{ range .things -}}
	  - {{ . }}
  	{{ end -}}`

	template := GetTemplate(templateText)

	// When
	rv := GetAllRequiredVariables(template)

	// Then
	if len(rv) != 1 {
		t.Errorf("Required variables was not as expected")
	}
}

func TestAllRequiredVariables(t *testing.T) {
	//Given
	templateText := `This is a testing document
	string value: {{ .stringValue1 }}
	another value: {{ .stringValue2 }}
	and double values: {{ .stringValue1 }}-{{ .stringValue2 }}`

	template := GetTemplate(templateText)

	// When
	rv := GetAllRequiredVariables(template)

	// Then
	if len(rv) != 2 {
		t.Errorf("Required variables was not as expected")
	}
}

package yamldata

import "testing"

func TestDocument(t *testing.T) {
	//Given
	d := map[string]string{
		"value1": "string",
		"list1":  "array",
	}

	expectedOutcome := "list1:\n- '...'\nvalue1: '...'\n"

	// When
	b, e := PrepareYamlFromData(d)

	// Then
	if e != nil {
		t.Errorf("Error returned during data preparations")
	}

	if string(b) != expectedOutcome {
		t.Errorf("Resulting YAML file not as expected. '%v' is not '%v'", string(b), expectedOutcome)
	}
}

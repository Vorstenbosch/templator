package template

import (
	"html/template"
	"regexp"
	"text/template/parse"
)

// Get all requied variables to ask to the user
func GetAllRequiredVariables(t *template.Template) map[string]string {
	variables := map[string]string{}

	for _, n := range t.Tree.Root.Nodes {
		switch v := n.(type) {
		case *parse.ActionNode:
			variable := getVariableName(v.String())
			variables[variable] = "string"
		case *parse.RangeNode:
			variable := getVariableName(v.BranchNode.String())
			variables[variable] = "array"
		default:
			// Unsupported type
		}
	}

	return variables
}

func getVariableName(s string) string {
	// To identify variables
	regex := regexp.MustCompile(`{{.*?\.([a-zA-z0-9]+)?}}`)

	matches := regex.FindAllStringSubmatch(s, -1)
	if matches != nil {
		// Expecting the variable name in the first group
		return matches[0][1]
	}

	return ""
}

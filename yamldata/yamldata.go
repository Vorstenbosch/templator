package yamldata

import (
	"gopkg.in/yaml.v2"
)

func PrepareYamlFromData(d map[string]string) ([]byte, error) {
	m := createYamlDataMap(d)

	b, err := yaml.Marshal(m)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func createYamlDataMap(d map[string]string) map[string]interface{} {
	m := map[string]interface{}{}

	for k, v := range d {
		switch v {
		case "string":
			m[k] = "..."
		case "array":
			m[k] = []string{"..."}
		}
	}

	return m
}

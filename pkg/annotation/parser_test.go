package annotation

import (
	"sigs.k8s.io/yaml"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	target := &struct {
		ID   int64    `json:"id"`
		Name string   `yaml:"name"`
		List []int64  `json:"list"`
		Test []string `json:"test"`
	}{}

	ann := Annotation{
		Name: "test",
		Args: map[string]string{
			"id":   "999",
			"name": "test",
			"list": "[1,2,3]",
			"test": "[test]",
		},
	}

	err := Unmarshal(ann, target)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(target)

	result, err := yaml.Marshal(target)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}

package pkg

import "gopkg.in/yaml.v3"

type IYaml interface {
	ToMap(content string) (map[string]interface{}, error)
}

type y struct {
}

func NewYaml() IYaml {
	return &y{}
}

func (y y) ToMap(content string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal([]byte(content), &result)
	return result, err
}

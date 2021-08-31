package configer

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type yamlLoader struct{}

func (p *yamlLoader) FromFile(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, out)
	return err
}

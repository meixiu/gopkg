package configer

import (
	"encoding/json"
	"io/ioutil"
)

type jsonLoader struct{}

func (p *jsonLoader) FromFile(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, out)
	return err
}

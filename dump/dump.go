package dump

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	yaml "gopkg.in/yaml.v2"
)

// echo
func Echo(args ...interface{}) {
	fmt.Println(args...)
}

// 打印并退出
func Exit(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

// DumpJSON
func JSON(item ...interface{}) {
	for _, v := range item {
		b, err := json.MarshalIndent(v, "", "  ")
		if err == nil {
			fmt.Println(string(b))
		}
	}
}

// DumpYAML
func YAML(item ...interface{}) {
	for _, v := range item {
		b, err := yaml.Marshal(v)
		if err == nil {
			fmt.Println(string(b))
		}
	}
}

func Dump(item ...interface{}) {
	spew.Dump(item...)
}

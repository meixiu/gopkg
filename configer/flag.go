package configer

import "flag"

type File struct {
	filename *string
}

func (f *File) Name() string {
	return *f.filename
}

func BindFlag(name string) *File {
	c := flag.String(name, "config/local.yaml", "Configuration file path")
	return &File{filename: c}
}

func ParseFlag() {
	flag.Parse()
}

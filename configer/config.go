package configer

import "strings"

type (
	DefaultConfig struct {
		Name       string      `json:"name" yaml:"name"`
		Env        string      `json:"env" yaml:"env"`
		Debug      bool        `json:"debug" yaml:"debug"`
		ServerAddr string      `json:"server_addr" yaml:"server_addr"`
		Db         DbConfig    `json:"db" yaml:"db"`
		Redis      RedisConfig `json:"redis" yaml:"redis"`
		Log        LogConfig   `json:"log" yaml:"log"`
	}

	DbConfig struct {
		Driver      string `json:"driver" yaml:"driver"`
		Master      string `json:"master" yaml:"master"`
		Slave       string `json:"slave" yaml:"slave"`
		Prefix      string `json:"prefix" yaml:"prefix"`
		MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn"`
		MaxIdleConn int    `json:"max_idle_conn" yaml:"max_idle_conn"`
	}

	RedisConfig struct {
		Addr     string `json:"addr" yaml:"addr"`
		Password string `json:"password" yaml:"password"`
		Db       int    `json:"db" yaml:"db"`
	}

	LogConfig struct {
		AccessLog string `json:"access_log" yaml:"access_log"`
		ErrorLog  string `json:"error_log" yaml:"error_log"`
	}
)

type Loader interface {
	FromFile(filename string, out interface{}) error
}

func NewLoader(t string) Loader {
	t = strings.ToLower(t)
	if t == "yaml" {
		return &yamlLoader{}
	} else if t == "json" {
		return &jsonLoader{}
	} else {
		return nil
	}
}

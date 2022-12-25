package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}

	Server struct {
		Port string
	}
}

func NewConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		//toml 파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
			return c
		}
	}
}

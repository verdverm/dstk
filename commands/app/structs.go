package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Type    string
	Name    string
	Basedir string
}

func ReadAppConfig(name string) *AppConfig {
	home := os.Getenv("HOME")
	fn := home + "/.dstk/apps/" + name + "/config.ini"

	cfg := new(AppConfig)

	data, err := ioutil.ReadFile(fn)
	panicErr(err)

	err = json.Unmarshal(data, cfg)
	panicErr(err)

	return cfg
}

func WriteAppConfig(name string, cfg *AppConfig) {
	home := os.Getenv("HOME")
	dir := home + "/.dstk/apps/" + name
	fn := dir + "/config.ini"

	data, err := json.MarshalIndent(cfg, "", "  ")
	panicErr(err)

	os.MkdirAll(dir, 0755)
	err = ioutil.WriteFile(fn, data, 0644)
	panicErr(err)
}

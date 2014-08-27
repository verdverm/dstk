package cluster

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ClusterConfig struct {
	Type    string
	Name    string
	DataDir string

	Tool       string
	Status     string
	MasterHost string
	MasterPort string
	Nodes      []string
}

func ReadClusterConfig(name string) *ClusterConfig {
	home := os.Getenv("HOME")
	fn := home + "/.dstk/clusters/" + name + "/config.ini"

	cfg := new(ClusterConfig)

	data, err := ioutil.ReadFile(fn)
	panicErr(err)

	err = json.Unmarshal(data, cfg)
	panicErr(err)

	return cfg
}

func WriteClusterConfig(name string, cfg *ClusterConfig) {
	home := os.Getenv("HOME")
	dir := home + "/.dstk/clusters/" + name
	fn := dir + "/config.ini"

	data, err := json.MarshalIndent(cfg, "", "  ")
	panicErr(err)

	os.MkdirAll(dir, 0755)
	err = ioutil.WriteFile(fn, data, 0644)
	panicErr(err)
}

package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/robfig/config"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var cfgStr = `# dstk - global config

# DEFAULT SETTINGS
datadir: ${HOME}/.dstk/data
databases: neo4j postgresql couchdb
slaves: 3

[neo4j]
replicas: 1
host: "127.0.0.1"
port: "7474"
user: ""
pass: ""

[postgresql]
replicas: 1
host: "127.0.0.1"
port: "5432"
user: ""
pass: ""

[couchdb]
replicas: 1
host: "127.0.0.1"
port: "5984"
user: ""
pass: ""
`

type DbSettings struct {
	DbType    string
	DbName    string
	DbDataDir string
	Replicas  int
	Host      string
	Port      string
	User      string
	Pass      string
}

type DstkConfig struct {
	DataDir    string
	Slaves     int
	Databases  []string
	DbSettings map[string]*DbSettings

	Clusters map[string]*ClusterConfig
}

var (
	CFGMAP *config.Config
	CONFIG *DstkConfig
)

func InitDstk(c *cli.Context) error {
	CONFIG = newDefaultDstkConfig()
	readDstkConfigFile()
	readDstkClusterDir()
	return nil
}

func SetupDstk(c *cli.Context) {
	fmt.Println("Setting up DSTK config in ~/.dstk")
	// setup a '.dstk' folder in home directory
	// subdirs '.dstk/clusters/<name>'
	var err error

	home := os.Getenv("HOME")
	err = os.MkdirAll(home+"/.dstk/clusters", 0755)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// create .dstk/config.ini
	fn := home + "/.dstk/config.ini"
	err = ioutil.WriteFile(fn, []byte(cfgStr), 0644)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// maybe a 'sqlite' database?

}

func newDefaultDstkConfig() *DstkConfig {
	cfg := &DstkConfig{
		Databases:  make([]string, 0, 4),
		DbSettings: make(map[string]*DbSettings),
		Clusters:   make(map[string]*ClusterConfig),
	}
	return cfg
}

func readDstkConfigFile() {
	home := os.Getenv("HOME")
	fn := home + "/.dstk/config.ini"
	cfg, err := config.ReadDefault(fn)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	// fmt.Println(cfg.SectionOptions("neo4j"))

	CFGMAP = cfg

	// convert cfg -> CONFIG
	convertConfigMapToStruct()
}

func convertConfigMapToStruct() {
	var err error

	// read defaults first
	CONFIG.DataDir, err = CFGMAP.String("default", "datadir")
	checkPanic(err)
	CONFIG.Slaves, err = CFGMAP.Int("default", "slaves")
	checkPanic(err)

	dbs, err := CFGMAP.String("default", "databases")
	checkPanic(err)
	CONFIG.Databases = strings.Fields(dbs)

	// read database defaults
	for _, db := range CONFIG.Databases {
		convertDbConfigMapToStruct(db)
	}

	// check for overrides by clustername
}

func convertDbConfigMapToStruct(section string) {
	// fmt.Println("DB-->", section)
	var err error
	dbs := new(DbSettings)
	dbs.DbType = section

	dbs.DbDataDir, err = CFGMAP.String(section, "datadir")
	checkPanic(err)

	dbs.Replicas, err = CFGMAP.Int(section, "replicas")
	checkPanic(err)

	dbs.Host, err = CFGMAP.String(section, "host")
	checkPanic(err)
	dbs.Port, err = CFGMAP.String(section, "port")
	checkPanic(err)
	dbs.User, err = CFGMAP.String(section, "user")
	checkPanic(err)
	dbs.Pass, err = CFGMAP.String(section, "pass")
	checkPanic(err)

	CONFIG.DbSettings[section] = dbs
}

func readDstkClusterDir() {
	home := os.Getenv("HOME")
	cluster_dirs, err := ioutil.ReadDir(home + "/.dstk/clusters")
	checkPanic(err)

	for _, C := range cluster_dirs {
		if !C.IsDir() {
			panic("non-directory trying to be read")
		}
		cname := C.Name()
		// fmt.Println("Reading cluster config:", cname)
		CONFIG.Clusters[cname] = readClusterConfig(cname)
	}
}

func PrintConfigValues(c *cli.Context) {

}

func SetConfigValue(c *cli.Context) {

}

func GetConfigValue(c *cli.Context) {

}

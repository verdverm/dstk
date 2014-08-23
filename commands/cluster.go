package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"

	"github.com/verdverm/dstk/commands/cluster"
)

type ClusterConfig struct {
	Type    string
	Name    string
	DataDir string

	Tool       string
	Status     string
	MasterHost string
	MasterPort string
	Slaves     []string
}

func readClusterConfig(name string) *ClusterConfig {
	home := os.Getenv("HOME")
	fn := home + "/.dstk/clusters/" + name + "/config.ini"

	cfg := new(ClusterConfig)

	data, err := ioutil.ReadFile(fn)
	checkPanic(err)

	err = json.Unmarshal(data, cfg)
	checkPanic(err)

	return cfg
}

func writeClusterConfig(name string, cfg *ClusterConfig) {
	home := os.Getenv("HOME")
	dir := home + "/.dstk/clusters/" + name
	fn := dir + "/config.ini"

	data, err := json.MarshalIndent(cfg, "", "  ")
	checkPanic(err)

	os.MkdirAll(dir, 0755)
	err = ioutil.WriteFile(fn, data, 0644)
	checkPanic(err)
}

func CreateCluster(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		tool = c.Args()[0]
	} else if len(c.Args()) == 2 {
		tool = c.Args()[0]
		cname = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to LaunchCluster")
	}

	prov := c.GlobalString("provider")
	switch prov {
	case "docker":
		fmt.Println("Launching Docker Cluster")

		ccfg, ok := CONFIG.Clusters[cname]
		// Start brand new cluster
		if !ok {
			ccfg = new(ClusterConfig)
			ccfg.DataDir = CONFIG.DataDir
			ccfg.Name = cname
			ccfg.Type = "docker"
			ccfg.Status = "HALTED"
			ccfg.MasterHost = "127.0.0.1"
			ccfg.MasterPort = "10000"

			writeClusterConfig(cname, ccfg)
			CONFIG.Clusters[cname] = ccfg
		}

		if ccfg.Status == "RUNNING" {
			fmt.Println("Cluster already running")
			return
		} else {
			ccfg.Status = "STARTING"
			writeClusterConfig(cname, ccfg)
		}

		ccfg.Tool = tool
		switch ccfg.Tool {
		case "hadoop":
			cluster.LaunchDockerHadoopCluster()
		case "spark":
			cluster.LaunchDockerSparkCluster()
		default:
			panic("Unknown tool type")
		}

		ccfg.Status = "RUNNING"
		writeClusterConfig(cname, ccfg)

	case "vagrant", "gce", "aws":
		fmt.Println("provider", prov, "not available yet")
		return
	default:
		fmt.Println("provider", prov, "unknown")
		return
	}
}

func StartCluster(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		cname = c.Args()[0]
	} else if len(c.Args()) == 2 {
		cname = c.Args()[0]
		tool = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to LaunchCluster")
	}

	prov := c.GlobalString("provider")
	switch prov {
	case "docker":
		fmt.Println("Launching Docker Cluster")

		ccfg, ok := CONFIG.Clusters[cname]
		// Start brand new cluster
		if !ok {
			panic("Unknown cluster " + cname)
		}

		if ccfg.Status == "RUNNING" {
			fmt.Println("Cluster already running")
			return
		} else {
			ccfg.Status = "STARTING"
			writeClusterConfig(cname, ccfg)
		}

		if tool != "" {
			ccfg.Tool = tool
		}

		switch ccfg.Tool {
		case "hadoop":
			cluster.LaunchDockerHadoopCluster()
		case "spark":
			cluster.LaunchDockerSparkCluster()
		default:
			panic("Unknown tool type")
		}

		ccfg.Status = "RUNNING"
		writeClusterConfig(cname, ccfg)

	case "vagrant", "gce", "aws":
		fmt.Println("provider", prov, "not available yet")
		return
	default:
		fmt.Println("provider", prov, "unknown")
		return
	}
}

func StopCluster(c *cli.Context) {
	cname := c.GlobalString("clustername")
	if len(c.Args()) > 0 {
		cname = c.Args().First()
	}

	ccfg, ok := CONFIG.Clusters[cname]
	if !ok {
		panic("Couldn't find cluster in ClusterMap")
	}
	ccfg.Status = "STOPPING"
	writeClusterConfig(cname, ccfg)

	prov := ccfg.Type
	switch prov {
	case "docker":
		cluster.DestroyDockerCluster()
		ccfg.Status = "HALTED"
		writeClusterConfig(cname, ccfg)

	case "vagrant", "gce", "aws":
		fmt.Println("provider", prov, "not available yet")
		return
	default:
		fmt.Println("provider", prov, "unknown")
		return
	}

}

func DestroyCluster(c *cli.Context) {
	cname := c.GlobalString("clustername")
	if len(c.Args()) > 0 {
		cname = c.Args().First()
	}

	ccfg, ok := CONFIG.Clusters[cname]
	if !ok {
		panic("Couldn't find cluster in ClusterMap")
	}

	if ccfg.Status != "HALTED" {
		StopCluster(c)
	}

	// remove directory
	home := os.Getenv("HOME")
	fn := home + "/.dstk/clusters/" + cname
	err := os.RemoveAll(fn)
	checkPanic(err)

	// remove configuration
	delete(CONFIG.Clusters, cname)

}

func PrintClusterStatus(c *cli.Context) {
	fmtstr := "%-12s  %-12s  %-12s  %s:%s\n"
	fmt.Printf(fmtstr, "Name", "Type", "Status", "Host", "Port")
	printr := func(ccfg *ClusterConfig) {
		fmt.Printf(fmtstr, ccfg.Name, ccfg.Type, ccfg.Status, ccfg.MasterHost, ccfg.MasterPort)
	}

	cluster_name := c.Args().First()
	if cluster_name == "all" || cluster_name == "" {
		if len(CONFIG.Clusters) == 0 {
			fmt.Println("No known clusters!")
			return
		}
		for _, ccfg := range CONFIG.Clusters {
			printr(ccfg)
		}
		return
	}
	ccfg, ok := CONFIG.Clusters[cluster_name]
	if !ok {
		fmt.Println("Unknown cluster:", cluster_name)
		return
	}
	printr(ccfg)

}

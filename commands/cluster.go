package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/verdverm/dstk/commands/cluster"
	"github.com/verdverm/dstk/commands/jobs"
)

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
			ccfg = new(cluster.ClusterConfig)
			ccfg.DataDir = CONFIG.DataDir
			ccfg.Name = cname
			ccfg.Type = "docker"
			ccfg.Status = "HALTED"

			cluster.WriteClusterConfig(cname, ccfg)
			CONFIG.Clusters[cname] = ccfg
		}

		if ccfg.Status == "RUNNING" {
			fmt.Println("Cluster already running")
			return
		} else {
			ccfg.Status = "STARTING"
			cluster.WriteClusterConfig(cname, ccfg)
		}

		ccfg.Tool = tool
		ccfg.Nodes = make([]string, 0, CONFIG.Nodes)
		for i := 0; i < CONFIG.Nodes; i++ {
			ccfg.Nodes = append(ccfg.Nodes, fmt.Sprintf("dstk-node-%02d", i))
		}
		cluster.LaunchDockerCluster(ccfg)

		ccfg.Status = "RUNNING"
		cluster.WriteClusterConfig(cname, ccfg)

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
			cluster.WriteClusterConfig(cname, ccfg)
		}

		if tool != "" {
			ccfg.Tool = tool
		}

		cluster.LaunchDockerCluster(ccfg)

		ccfg.Status = "RUNNING"
		cluster.WriteClusterConfig(cname, ccfg)

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
	cluster.WriteClusterConfig(cname, ccfg)

	prov := ccfg.Type
	switch prov {
	case "docker":
		cluster.DestroyDockerCluster(ccfg)
		ccfg.Status = "HALTED"
		cluster.WriteClusterConfig(cname, ccfg)

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
	printr := func(ccfg *cluster.ClusterConfig) {
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

func RunClusterJob(c *cli.Context) {
	cluster_name := c.GlobalString("clustername")
	job_name := c.GlobalString("jobname")
	cmd_name := c.Args().First()
	println("running", cmd_name, "(", job_name, ") on", cluster_name)

	jobs.SparkSubmit()
}

func LoginClusterNode(c *cli.Context) {
	cluster_name := c.GlobalString("clustername")
	node_name := c.Args().First()
	println("login: ", node_name, "on", cluster_name)

	// ccfg, ok := CONFIG.Clusters[cname]
	// if !ok {
	// 	panic("Couldn't find cluster in ClusterMap")
	// }

	cluster.LoginDockerNode(node_name)
}

func ClusterAmbariShell(c *cli.Context) {
	cname := c.GlobalString("clustername")

	ccfg, ok := CONFIG.Clusters[cname]
	if !ok {
		panic("Couldn't find cluster in ClusterMap")
	}

	cluster.LaunchDockerAmbariShell(ccfg)
}

package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/verdverm/dstk/src/dstk/commands/cluster"
)

func LaunchCluster(c *cli.Context) {
	// cname := c.GlobalString("clustername")
	// if len(c.Args()) > 0 {
	// 	cname = c.Args().First()
	// }

	prov := c.GlobalString("provider")
	switch prov {
	case "docker":
		cluster.LaunchDockerCluster()
	case "vagrant", "gce", "aws":
		fmt.Println("provider", prov, "not available yet")
		return
	default:
		fmt.Println("provider", prov, "unknown")
		return
	}
}

func DestroyCluster(c *cli.Context) {
	// cname := c.GlobalString("clustername")
	// if len(c.Args()) > 0 {
	// 	cname = c.Args().First()
	// }

	prov := c.GlobalString("provider")
	switch prov {
	case "docker":
		cluster.DestroyDockerCluster()
	case "vagrant", "gce", "aws":
		fmt.Println("provider", prov, "not available yet")
		return
	default:
		fmt.Println("provider", prov, "unknown")
		return
	}
}

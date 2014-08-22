package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/verdverm/dstk/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "dstk"
	app.Version = "0.0.1"
	app.Usage = "Data Science ToolKit!"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "provider",
			Value:  "docker",
			Usage:  "cluster provider [docker] [...,vagrant,gce,aws]",
			EnvVar: "DTSK_PROVIDER",
		},
		cli.StringFlag{
			Name:   "clustername,n",
			Value:  "default",
			Usage:  "cluster name",
			EnvVar: "DTSK_CLUSTER",
		},
		cli.IntFlag{
			Name:   "slaves",
			Value:  1,
			Usage:  "number of slave nodes",
			EnvVar: "DTSK_SLAVES",
		},
	}

	app.Before = commands.InitDstk

	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Usage:  "setup the dstk configuration",
			Action: commands.SetupDstk,
		},
		{
			Name:   "config",
			Usage:  "print or alter dstk config values",
			Action: commands.PrintConfigValues,
			Subcommands: []cli.Command{
				{
					Name:   "set",
					Usage:  "dstk config set <key> <value>",
					Action: commands.SetConfigValue,
				},
				{
					Name:   "get",
					Usage:  "dstk config get <key> <value>",
					Action: commands.GetConfigValue,
				},
			},
		},
		{
			Name:  "cluster",
			Usage: "options for cluster management",
			Subcommands: []cli.Command{
				{
					Name:   "launch",
					Usage:  "launch a new named cluster",
					Action: commands.LaunchCluster,
				},
				{
					Name:   "destroy",
					Usage:  "destroy a running cluster",
					Action: commands.DestroyCluster,
				},
				{
					Name:   "status",
					Usage:  "print status of clusters",
					Action: commands.PrintClusterStatus,
				},
			},
		},
		{
			Name:  "spark",
			Usage: "options for spark cluster management",
			Subcommands: []cli.Command{
				{
					Name:  "run",
					Usage: "run a job on a spark cluster",
					Action: func(c *cli.Context) {
						cluster_name := c.String("cluster")
						job_name := c.Args().First()
						println("running", job_name, "on", cluster_name)
					},
				},
				{
					Name:  "stop",
					Usage: "stop a running job on a spark cluster",
					Action: func(c *cli.Context) {
						cluster_name := c.String("cluster")
						job_name := c.Args().First()
						println("stopping", job_name, "on", cluster_name)
					},
				},
				{
					Name:  "status",
					Usage: "print status of spark jobs on cluster(s)",
					Action: func(c *cli.Context) {
						println("spark cluster status")
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

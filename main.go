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
			Name:   "provider,p",
			Value:  "docker",
			Usage:  "cluster provider [docker] [...,vagrant,gce,aws]",
			EnvVar: "DTSK_PROVIDER",
		},
		cli.IntFlag{
			Name:   "slaves,s",
			Value:  1,
			Usage:  "number of slave nodes",
			EnvVar: "DTSK_SLAVES",
		},
		cli.StringFlag{
			Name:   "clustername,c",
			Value:  "default",
			Usage:  "cluster name",
			EnvVar: "DTSK_CLUSTER",
		},
		cli.StringFlag{
			Name:   "jobname,j",
			Value:  "default",
			Usage:  "cluster name",
			EnvVar: "DTSK_CLUSTER",
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
					Usage:  "config set <key> <value>",
					Action: commands.SetConfigValue,
				},
				{
					Name:   "get",
					Usage:  "config get <key> <value>",
					Action: commands.GetConfigValue,
				},
			},
		},
		{
			Name:   "login",
			Usage:  "login <node_name>",
			Action: commands.LoginClusterNode,
		},
		{
			Name:  "cluster",
			Usage: "options for cluster management",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create <tool> <clustername> | tool:[hadoop,spark]",
					Action: commands.CreateCluster,
				},
				{
					Name:   "start",
					Usage:  "cluster start <clustername>",
					Action: commands.StartCluster,
				},
				{
					Name:   "stop",
					Usage:  "cluster stop <clustername>",
					Action: commands.StopCluster,
				},
				{
					Name:   "destroy",
					Usage:  "cluster destroy <clustername>",
					Action: commands.DestroyCluster,
				},
				{
					Name:   "status",
					Usage:  "cluster status <clustername>",
					Action: commands.PrintClusterStatus,
				},
				{
					Name:   "ambari",
					Usage:  "cluster ambari",
					Action: commands.ClusterAmbariShell,
				},
			},
		},
		{
			Name:  "jobs",
			Usage: "options for jobs management",
			Subcommands: []cli.Command{
				{
					Name:   "run",
					Usage:  "jobs run cmd args...",
					Action: commands.RunClusterJob,
				},
				{
					Name:  "stop",
					Usage: "jobs stop <jobname>",
					Action: func(c *cli.Context) {
						cluster_name := c.GlobalString("clustername")
						job_name := c.Args().First()
						println("stopping", job_name, "(", job_name, ")on", cluster_name)
					},
				},
				{
					Name:  "status",
					Usage: "jobs status <clustername>",
					Action: func(c *cli.Context) {
						println("jobs status <clustername>")
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

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
			Name:  "ambari",
			Usage: "dstk -c <clustername> ambari <cmd> <args>",
			Subcommands: []cli.Command{
				{
					Name:   "shell",
					Usage:  "ambari shell <nodename>",
					Action: commands.ClusterAmbariShell,
				},
				{
					Name:   "blueprint",
					Usage:  "ambari blueprint <subcmd>",
					Action: commands.ClusterAmbariBlueprint,
				},
				{
					Name:   "command",
					Usage:  "ambari command <cmd>",
					Action: commands.ClusterAmbariCommand,
				},
			},
		},
		{
			Name:  "hdfs",
			Usage: "dstk -c <clustername> hdfs",
			Subcommands: []cli.Command{
				{
					Name:   "ls",
					Usage:  "hdfs ls <path>",
					Action: commands.ClusterHdfsLs,
				},
				{
					Name:   "cpin",
					Usage:  "hdfs cpin <src> <dest>",
					Action: commands.ClusterHdfsCpin,
				},
				{
					Name:   "cpout",
					Usage:  "hdfs cpout <src> <dest>",
					Action: commands.ClusterHdfsCpout,
				},
				{
					Name:   "mv",
					Usage:  "hdfs mv <dir>",
					Action: commands.ClusterHdfsMv,
				},
				{
					Name:   "rm",
					Usage:  "hdfs rm <dir>",
					Action: commands.ClusterHdfsRm,
				},
				{
					Name:   "mkdir",
					Usage:  "hdfs mkdir <dir>",
					Action: commands.ClusterHdfsMkdir,
				},
				{
					Name:   "chown",
					Usage:  "hdfs chown <dir> <owner> <group>",
					Action: commands.ClusterHdfsChown,
				},
				{
					Name:   "chmod",
					Usage:  "hdfs chmod <file,dir> <oct-permissions>",
					Action: commands.ClusterHdfsChmod,
				},
			},
		},
		{
			Name:  "cluster",
			Usage: "dstk -c <clustername> cluster <cmd> <args>",
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

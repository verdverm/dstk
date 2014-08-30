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
			Value:  "test",
			Usage:  "cluster name",
			EnvVar: "DTSK_CLUSTER",
		},
		cli.StringFlag{
			Name:   "jobname,j",
			Value:  "default_job",
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
			Name:   "addhosts",
			Usage:  "addhosts <clustername> (to /etc/hosts)",
			Action: commands.WriteClusterHosts,
		},
		{
			Name:   "removehosts",
			Usage:  "removehosts <clustername> (from /etc/hosts)",
			Action: commands.RemoveClusterHosts,
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
					Usage:  "hdfs cpin <filename> <hdfspath>",
					Action: commands.ClusterHdfsCpin,
				},
				{
					Name:   "cpout",
					Usage:  "hdfs cpout <hdfspath> <filename>",
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
			Name:  "app",
			Usage: "options for dstk apps",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "dstk app create <name> <args...?>",
					Action: commands.CreateDstkApp,
				},
				{
					Name:   "build",
					Usage:  "dstk app build - builds the current directory",
					Action: commands.BuildDstkApp,
				},
				{
					Name:   "upload",
					Usage:  "dstk app upload <clustername>- uploads the current directory",
					Action: commands.UploadDstkApp,
				},
				{
					Name:   "run",
					Usage:  "dstk app run <appname> <classpath> <appargs...>",
					Action: commands.RunDstkApp,
				},
			},
		},
	}

	app.Run(os.Args)
}

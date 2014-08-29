package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"

	"github.com/verdverm/dstk/commands/app"
	"github.com/verdverm/dstk/commands/cluster"
	"github.com/verdverm/dstk/commands/jobs"
)

func CreateDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		tool = c.Args()[0]
	} else if len(c.Args()) == 2 {
		tool = c.Args()[0]
		cname = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to CreateDsktApp")
	}
}

func BuildDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		tool = c.Args()[0]
	} else if len(c.Args()) == 2 {
		tool = c.Args()[0]
		cname = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to BuildDstkApp")
	}
}

func UploadDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		tool = c.Args()[0]
	} else if len(c.Args()) == 2 {
		tool = c.Args()[0]
		cname = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to UploadDstkApp")
	}
}

func RunDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var tool string
	if len(c.Args()) == 1 {
		tool = c.Args()[0]
	} else if len(c.Args()) == 2 {
		tool = c.Args()[0]
		cname = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to RunDstkApp")
	}
}

package commands

import (
	"fmt"
	"os"
	// "os/exec"
	"path/filepath"

	"github.com/codegangsta/cli"

	"github.com/verdverm/dstk/commands/app"
	// "github.com/verdverm/dstk/commands/cluster"
	// "github.com/verdverm/dstk/commands/jobs"
)

func CreateDstkApp(c *cli.Context) {
	// cname := c.GlobalString("clustername")
	// var tool string
	// if len(c.Args()) == 1 {
	// 	tool = c.Args()[0]
	// } else if len(c.Args()) == 2 {
	// 	tool = c.Args()[0]
	// 	cname = c.Args()[1]
	// } else {
	// 	fmt.Println("Error: bad args to CreateDsktApp")
	// }
}

func BuildDstkApp(c *cli.Context) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	acfg := new(app.AppConfig)
	acfg.Basedir = cwd
	acfg.Name = filepath.Base(cwd)
	acfg.Type = "spark"

	fmt.Println(*acfg)

	app.BuildDstkScalaApp(acfg)

}

func UploadDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	if len(c.Args()) == 1 {
		cname = c.Args()[0]
	}

	ccfg, ok := CONFIG.Clusters[cname]
	if !ok {
		panic("Couldn't find cluster in ClusterMap")
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	appname := filepath.Base(cwd)
	fn := appname + ".jar"
	filename := cwd + "/" + fn
	path := "/user/root/apps/" + fn

	cmd := "curl -X PUT -L 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		path + "?user.name=hdfs&op=CREATE&permission=0644&overwrite=true' -T " + filename + "| jq '.'"
	println(cmd)
	exec_command("/bin/bash", "-c", cmd)
}

func RunDstkApp(c *cli.Context) {
	cname := c.GlobalString("clustername")
	var appname, classpath string
	if len(c.Args()) >= 2 {
		appname = c.Args()[0]
		classpath = c.Args()[1]
	} else {
		fmt.Println("Error: bad args to RunDstkApp")
	}

	appargs := ""
	for i := 2; i < len(c.Args()); i++ {
		appargs += fmt.Sprintf(" %s", c.Args()[i])
	}

	println(cname, appname, classpath, appargs)
	app.RunDstkSparkApp(cname, appname, classpath, appargs)
}

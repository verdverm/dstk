package app

import (
	"fmt"
)

func RunDstkSparkApp(clustername, appname, classpath, appargs string) {
	cmd := fmt.Sprintf("sudo docker-enter dstk-node-00 /tmp/spark-submit.sh %s %s %s", appname, classpath, appargs)
	fmt.Println(cmd)
	exec_command("/bin/bash", "-c", cmd)
}

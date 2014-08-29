package jobs

import (
	"fmt"
	"os"
	"os/exec"

	// "github.com/fsouza/go-dockerclient"
)

var endpoint = "unix:///var/run/docker.sock"

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func exec_command(program string, args ...string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

var (
	JAVA_HOME       = "/usr/jdk64/jdk1.7.0_45"
	HADOOP_CONF_DIR = "/etc/hadoop"
)

func SparkSubmit() {
	// build user directory with docker

	// copy jar to hdfs

	// run spark submit script
	exec_command("/bin/bash", "-c", "sudo docker-enter dstk-node-00 /tmp/spark-submit.sh")

}

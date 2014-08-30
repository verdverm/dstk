package app

import (
	"fmt"
	"os"
	"os/exec"
	// "time"

	"github.com/fsouza/go-dockerclient"
)

var endpoint = "unix:///var/run/docker.sock"

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
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

func BuildDstkScalaApp(acfg *AppConfig) {
	// docker run -it \
	//   -v $(pwd):/app/usercode \
	//   verdverm/dstk-app

	name := "dstk-app-build-scala"
	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image:      "verdverm/dstk-app",
			Entrypoint: []string{"/app/scripts/build-scala.sh"},
			Env: []string{
				"APPNAME=" + acfg.Name,
				"APPSRCDIR=" + acfg.Basedir + "/src",
			},
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-app",
		Privileged:      true,
		Binds: []string{
			acfg.Basedir + ":/app/usercode",
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating", name, "container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting", name, "container")
	err = client.StartContainer(name, sopts)
	panicErr(err)
	fmt.Println("  - Successfully started", name, "docker")

	aopts := docker.AttachToContainerOptions{
		Container: name,
		Logs:      true,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		Stream:    true,

		InputStream:  os.Stdin,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
	}

	fmt.Println("  - Attaching to ", name, "docker")
	err = client.AttachToContainer(aopts)
	panicErr(err)

	ropts := docker.RemoveContainerOptions{
		ID: name,
	}

	err = client.RemoveContainer(ropts)
	panicErr(err)
}

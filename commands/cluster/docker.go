package cluster

import (
	"fmt"
	"os"
	"os/exec"
	"time"

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

func LoginDockerNode(node_name string) {
	exec_command("/bin/bash", "-c", "sudo docker-enter "+node_name)
}

func LaunchDockerAmbariShell(ccfg *ClusterConfig) {
	cmd_str := fmt.Sprintf("sudo docker-enter dstk-node-00 AMBARI_HOST=%s /tmp/dstk-ambari-shell.sh", ccfg.MasterHost)
	exec_command("/bin/bash", "-c", cmd_str)
}

func LaunchDockerAmbariCommand(cmd string, ccfg *ClusterConfig) {
	cmd_str := fmt.Sprintf("sudo docker-enter dstk-node-00 AMBARI_HOST=%s /tmp/dstk-ambari-shell.sh %s", ccfg.MasterHost, cmd)
	exec_command("/bin/bash", "-c", cmd_str)
}

func LaunchDockerCluster(ccfg *ClusterConfig) {
	switch ccfg.Tool {
	case "hadoop", "spark":
		LaunchDockerHadoopCluster(len(ccfg.Nodes))
		ccfg.MasterHost = MASTERIP
		ccfg.MasterPort = "8080"

	default:
		panic("Unknown tool type")
	}
}

func LaunchDockerHadoopCluster(numnodes int) {
	fmt.Println("Launching local docker hadoop cluster:")

	// launchDockerDatabases()
	launchDockerHadoopFirst()
	for i := 1; i < numnodes; i++ {
		launchDockerHadoopNode(i)
	}

	wait := 10
	fmt.Printf("\033[sSleeping %2ds for everything to come up", wait)
	for i := 1; i <= wait; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("\033[uSleeping %2ds for everything to come up", wait-i)
	}

	fmt.Println("Installing Cluster")
	installDockerHadoopCluster(numnodes)

	wait_str := fmt.Sprintf("docker logs -f dstk-cluster-installer")
	exec_command("/bin/bash", "-c", wait_str)

	fmt.Println("Post-install Scripts")
	cmd_str := fmt.Sprintf("sudo docker-enter dstk-node-00 /tmp/mkdir-hadoop.sh")
	exec_command("/bin/bash", "-c", cmd_str)

}

func DestroyDockerCluster(ccfg *ClusterConfig) {
	fmt.Println("Destroying local docker cluster:")

	// use docker API to determine whats running
	// and then destroy that ?

	for i := 0; i < len(ccfg.Nodes); i++ {
		fmt.Printf("destroying dstk-node-%02d\n", i)
		destroyDocker(fmt.Sprintf("dstk-node-%02d", i))
	}
	destroyDocker("dstk-cluster-installer")
	// destroyDocker("dtsk-postgresql")
	// destroyDocker("dtsk-couchdb")
	// destroyDocker("dtsk-neo4j")
}

func launchDockerDatabases() {
	launchDockerPostgresql()
	launchDockerCouchDB()
	launchDockerNeo4j()
}

func launchDockerCouchDB() {
	name := "dstk-couchdb"

	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image: "klaemo/couchdb",
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("5984"): {},
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "klaemo/couchdb",
		Privileged:      true,
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("5984"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "5984",
				},
			},
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating CouchDB container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting CouchDB container")
	err = client.StartContainer(name, sopts)
	panicErr(err)

	fmt.Println("  - Successfully started CouchDB docker")

}

func launchDockerPostgresql() {
	name := "dstk-postgresql"

	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image: "paintedfox/postgresql",
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("5432"): {},
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "paintedfox/postgresql",
		Privileged:      true,
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("5432"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "2345",
				},
			},
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating Postgresql container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting Postgresql container")
	err = client.StartContainer(name, sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Postgresql docker")
}

func launchDockerNeo4j() {
	name := "dstk-neo4j"

	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image: "tpires/neo4j",
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("7474"): {},
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "tpires/neo4j",
		Privileged:      true,
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("7474"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "7474",
				},
			},
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating Neo4j container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting Neo4j container")
	err = client.StartContainer(name, sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Neo4j docker")
}

var MASTERIP = ""

// var BLUEPRINT = "myblueprint"

var BLUEPRINT = "multi-node-hdfs-yarn"

func launchDockerHadoopFirst() {
	name := fmt.Sprintf("dstk-node-%02d", 0)

	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image:      "verdverm/dstk-spark",
			Entrypoint: []string{"/usr/local/serf/bin/start-serf-agent.sh"},
			Cmd:        []string{"--tag", "ambari-server=true"},
			CpuShares:  4,
			Hostname:   name,
			Domainname: "iassic.com",
			// Dns:        []string{"127.0.0.1"},
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("8080"): {},
				docker.Port("7373"): {},
				docker.Port("7946"): {},
			},
			// Env: []string{
			// 	"KEYCHAIN=$KEYCHAIN",
			// },
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-spark",
		Privileged:      true,
		Dns:             []string{"127.0.0.1"},
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

	cont, err := client.InspectContainer(name)
	panicErr(err)
	MASTERIP = cont.NetworkSettings.IPAddress
}

func launchDockerHadoopNode(id int) {
	name := fmt.Sprintf("dstk-node-%02d", id)

	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image:      "verdverm/dstk-spark",
			Entrypoint: []string{"/usr/local/serf/bin/start-serf-agent.sh"},
			Cmd:        []string{"--log-level debug"},
			CpuShares:  4,
			Hostname:   name,
			Domainname: "iassic.com",
			// Dns:        []string{"127.0.0.1"},
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("8080"): {},
				docker.Port("7373"): {},
				docker.Port("7946"): {},
			},
			Env: []string{
				"SERF_JOIN_IP=" + MASTERIP,
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-spark",
		Privileged:      true,
		Dns:             []string{"127.0.0.1"},
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
}

func installDockerHadoopCluster(numnodes int) {
	name := "dstk-cluster-installer"
	// create options
	copts := docker.CreateContainerOptions{
		Name: name,
		Config: &docker.Config{
			Image:      "verdverm/dstk-spark",
			Entrypoint: []string{"/bin/sh"},
			Cmd:        []string{"-c", "/tmp/install-cluster.sh"},
			Hostname:   name,
			Domainname: "iassic.com",
			Env: []string{
				"BLUEPRINT=" + BLUEPRINT,
				fmt.Sprintf("EXPECTED_HOST_COUNT=%d", numnodes),
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-spark",
		Privileged:      true,
		Dns:             []string{"127.0.0.1"},
		Links:           []string{"dstk-node-00:ambariserver"},
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

	// ropts := docker.RemoveContainerOptions{
	// 	ID:    name,
	// 	Force: true,
	// }
	// err = client.RemoveContainer(ropts)
	// panicErr(err)

}

func destroyDocker(name string) {
	// remove options
	ropts := docker.RemoveContainerOptions{
		ID:            name,
		RemoveVolumes: true,
		Force:         true,
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("Removing", name, "container")
	err = client.RemoveContainer(ropts)
	checkErr(err)
}

func DockerHdfsLs(dir string, ccfg *ClusterConfig) {
	cmd := "curl -s 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" + dir + "?op=LISTSTATUS' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}

func DockerHdfsCpin(ccfg *ClusterConfig) {

}
func DockerHdfsCpout(ccfg *ClusterConfig) {

}
func DockerHdfsMv(src, dest string, ccfg *ClusterConfig) {
	fmt.Println("HDFS MV: ", src, dest)
	cmd := "curl -s -X PUT 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		src + "?user.name=hdfs&op=RENAME&destination=" + dest + "' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}
func DockerHdfsRm(dir string, ccfg *ClusterConfig) {
	cmd := "curl -s -X DELETE 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		dir + "?user.name=hdfs&op=DELETE&recursive=true' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}
func DockerHdfsMkdir(dir string, ccfg *ClusterConfig) {
	cmd := "curl -s -X PUT 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		dir + "?user.name=hdfs&op=MKDIRS&permission=0755' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}
func DockerHdfsChmod(dir, perm string, ccfg *ClusterConfig) {
	cmd := "curl -s -X PUT 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		dir + "?user.name=hdfs&op=SETPERMISSION&permission=" + perm + "' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}
func DockerHdfsChown(dir, owner, group string, ccfg *ClusterConfig) {
	cmd := "curl -s -X PUT 'http://" + ccfg.MasterHost + ":50070/webhdfs/v1" +
		dir + "?user.name=hdfs&op=SETOWN&owner=" + owner + "&group=" + group + "' | jq '.'"
	exec_command("/bin/bash", "-c", cmd)
}

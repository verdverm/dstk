package cluster

import (
	"fmt"
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

func LaunchDockerCluster() {
	fmt.Println("Launching local docker cluster:")

	launchDockerDatabases()
	launchDockerSparks()

	fmt.Println("Sleeping 10s for everything to come up")
	time.Sleep(10 * time.Second)
}

func DestroyDockerCluster() {
	fmt.Println("Destroying local docker cluster:")

	destroyDocker("spark-slave")
	destroyDocker("spark-master")
	destroyDocker("neo4j")
	destroyDocker("couchdb")
	destroyDocker("postgresql")
}

func launchDockerDatabases() {
	launchDockerPostgresql()
	launchDockerCouchDB()
	launchDockerNeo4j()
}

func launchDockerSparks() {
	launchDockerSparkMaster()
	launchDockerSparkSlaves()
}

func launchDockerCouchDB() {
	// create options
	copts := docker.CreateContainerOptions{
		Name: "couchdb",
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
	err = client.StartContainer("couchdb", sopts)
	panicErr(err)

	fmt.Println("  - Successfully started CouchDB docker")

}

func launchDockerPostgresql() {
	// create options
	copts := docker.CreateContainerOptions{
		Name: "postgresql",
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
					HostPort: "5432",
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
	err = client.StartContainer("postgresql", sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Postgresql docker")
}

func launchDockerNeo4j() {
	// create options
	copts := docker.CreateContainerOptions{
		Name: "neo4j",
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
	err = client.StartContainer("neo4j", sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Neo4j docker")
}

func launchDockerSparkMaster() {

}

func launchDockerSparkSlaves() {

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

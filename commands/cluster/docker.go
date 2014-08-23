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

func LaunchDockerHadoopCluster() {
	fmt.Println("Launching local docker hadoop cluster:")

	launchDockerDatabases()
	launchDockerHadoopMaster()
	launchDockerHadoopSlaves()

	fmt.Println("Sleeping 10s for everything to come up")
	time.Sleep(10 * time.Second)
}

func LaunchDockerSparkCluster() {
	fmt.Println("Launching local docker spark cluster:")

	launchDockerDatabases()
	launchDockerSparkMaster()
	launchDockerSparkSlaves()

	fmt.Println("Sleeping 10s for everything to come up")
	time.Sleep(10 * time.Second)
}

func DestroyDockerCluster() {
	fmt.Println("Destroying local docker cluster:")

	destroyDocker("spark-slave")
	destroyDocker("spark-master")
	destroyDocker("hadoop-slave")
	destroyDocker("hadoop-master")
	destroyDocker("neo4j")
	destroyDocker("couchdb")
	destroyDocker("postgresql")
}

func launchDockerDatabases() {
	launchDockerPostgresql()
	launchDockerCouchDB()
	launchDockerNeo4j()
}

func launchDockerMaster() {
	launchDockerHadoopMaster()
	// launchDockerSparkMaster()
}

func launchDockerSlaves() {
	launchDockerHadoopSlaves()
	// launchDockerSparkSlaves()
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

func launchDockerHadoopMaster() {

	// create options
	copts := docker.CreateContainerOptions{
		Name: "hadoop-master",
		Config: &docker.Config{
			Image: "verdverm/dstk-hadoop",
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("8088"):  {},
				docker.Port("8032"):  {},
				docker.Port("50070"): {},
				docker.Port("50075"): {},
				docker.Port("50090"): {},
			},
			Env: []string{
				"NODE_TYPE=single",
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-hadoop",
		Privileged:      true,
		NetworkMode:     "host",
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("8088"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "8088",
				},
			},
			docker.Port("8032"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "8032",
				},
			},

			docker.Port("50070"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "50070",
				},
			},
			docker.Port("50075"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "50075",
				},
			},
			docker.Port("50090"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "50090",
				},
			},
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating Hadoop-Master container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting Hadoop-Master container")
	err = client.StartContainer("hadoop-master", sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Hadoop-Master docker")
}

func launchDockerHadoopSlaves() {

}

func launchDockerSparkMaster() {
	// create options
	copts := docker.CreateContainerOptions{
		Name: "spark-master",
		Config: &docker.Config{
			Image: "verdverm/dstk-spark",
			ExposedPorts: map[docker.Port]struct{}{
				// hadoop related
				docker.Port("8088"): {},
				docker.Port("8032"): {},

				// spark related
				docker.Port("8080"): {},
				docker.Port("7077"): {},
				docker.Port("22"):   {},
			},
		},
	}

	// start options for:
	sopts := &docker.HostConfig{
		ContainerIDFile: "verdverm/dstk-spark",
		Privileged:      true,
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("8088"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "8088",
				},
			},
			docker.Port("8032"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "8032",
				},
			},

			docker.Port("8080"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "8080",
				},
			},
			docker.Port("7077"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "7077",
				},
			},
			docker.Port("22"): {
				docker.PortBinding{
					HostIp:   "0.0.0.0",
					HostPort: "2222",
				},
			},
		},
	}

	client, err := docker.NewClient(endpoint)
	panicErr(err)

	fmt.Println("  - Creating Spark-Master container")
	_, err = client.CreateContainer(copts)
	panicErr(err)

	fmt.Println("  - Starting Spark-Master container")
	err = client.StartContainer("spark-master", sopts)
	panicErr(err)

	fmt.Println("  - Successfully started Spark-Master docker")
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

dstk - Data Science ToolKit
===========================

A tool for using other data science tools,
built on Docker and Hadoop.

It's a little rough right now, expect things to change.

### Installing

```
go get github.com/verdverm/dstk
cd $GOPATH/github.com/verdverm/dstk/dockers
./build.sh
cd ..
dstk setup
```

This will create a `.dstk` directory in the user's home directory.
There are some config files there you can checkout.

Some third party codes in use:

- [github.com/codegangsta/cli](https://github.com/codegangsta/cli)
- [github.com/fsouza/go-dockerclient](https://github.com/fsouza/go-dockerclient)
- [github.com/jpetazzo/nsenter](https://github.com/jpetazzo/nsenter)
- several from [github.com/sequenceiq](https://github.com/sequenceiq)
- verdverm/dstk-spark docker is derived from sequenciq/ambari.


### CLI help command

is most likely the most up-to-date source right now.

`dstk help`
`dstk <command> help`


### Launch a Hadoop/Yarn/Spark cluster

`dtsk cluster create hadoop <clustername>`
`dstk addhosts <clustername>`

The cluster has a hadoop2.5/yarn setup with a Spark on Hadoop2 binary included.

### Teardown a Cluster

`dstk removehosts <clustername>`
`dstk cluster destroy <clustername>`

### Get Cluster Status

`dstk cluster status <clustername>`

Not overly informative at this point, but the master IP is there.

### Open Ambari WebUI

open your browser to `<masterip>:8080`.

The user:pass is `admin:admin`

### Login to a node

`dstk -c <clustername> login <nodename>`

Environment variables from the Dockerfiles are missing on login.
requires sudo access for `docker-enter (nsenter)`.

### Open Ambari Shell

`dstk -c <clustername> cluster ambari`

### Run the SparkPi example

`dstk -c <clustername> -j <jobname> jobs run <command>`

currently the command is ignored... sorry.
It's hardcoded to run the SparkPi example temporarily.

### TODO

Help out! Here are some things that need doing

- add custom blueprint feature
- abstract provider concept
  -- Golang interface and CLI making use of provider tools
  -- docker only
  -- vagrant / vmware
  -- GCE / AWS
  -- DigitalOcean / Rackspace
- run arbitrary command
- run arbitrary script
- build and run application from directory
  -- subdirs [data,src]
  -- config file
  -- eamples of these
- convert cluster meta-data storage to Sqlite
- use Ambari / Hadoop APIs instead of shelling out
- adding / removing nodes
  -- test fault tolerance
  -- resize cluster
  -- install script for single node?
  -- serf
- add custom serf handlers
  -- serf plugin handle script
- use [github.com/nsf/termbox-go](https://github.com/nsf/termbox-go) for cli



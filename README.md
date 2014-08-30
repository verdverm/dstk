dstk - Data Science ToolKit
===========================

A tool for simplifying data science tasks,
built on Docker, Hadoop, Spark...

Right now you can:

- Create and run a local Hadoop2 / Yarn / Spark, docker based cluster.
- Build, upload, and run Spark applications in Scala.
- Perform many hdfs functions.

It's a little early right now, expect things to change.

### Help command

is most likely the most up-to-date source for command usage.
Don't quote me though.

`dstk help`
`dstk <command> help`

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
- [github.com/sequenceiq/...](https://github.com/sequenceiq)
- `verdverm/dstk-spark` docker is derived from `sequenciq/ambari`.


## Cluster Commands

### Launch a Hadoop/Yarn/Spark cluster

`dtsk cluster create hadoop <clustername>`
`dstk addhosts <clustername>`

The cluster has a Hadoop2.5 / Yarn setup, with a Spark 1.0.2 on Hadoop2 binary included.

### Teardown a Cluster

`dstk removehosts <clustername>`
`dstk cluster destroy <clustername>`

### Get Cluster Status

`dstk cluster status <clustername>`

Not overly informative at this point, but the master IP is there.

### Add and Remove hostnames

These commands will add hostname / ip information to your `/etc/hosts` file.
This is necessary for some of the`hdfs` functions and useful for the WebUI.

`dstk addhosts <clustername>`
`dstk removehosts <clustername>`

I think removehosts is broken right now...

### Open Ambari WebUI

open your browser to `<masterip>:8080`.

The user:pass is `admin:admin`

### Login to a node

`dstk -c <clustername> login <nodename>`

Environment variables from the Dockerfiles are missing on login.
requires sudo access for `docker-enter (nsenter)`.

### Open Ambari Shell

`dstk -c <clustername> cluster ambari`


## App Commands

### Build an app

Directory layout should be
```
appname
  \ src
    \ appname
      - appname.scala
  - MANIFEST.MF
```

From the appname directory, run `dstk app build`

### Upload to a cluster

From the appname directory, run `dstk app upload <clustername>`

### Run an application

From anywhere, run

`dstk -c <clustername> app run <appname> <appclass> <appargs...>`

To run the wordcount example

`dstk -c <clustername> app run wordcount wordcount.WordCount /user/root/in.txt /user/root/out.txt`

Output can be obtained by running

`dstk -c <clustername> hdfs cpout /user/root/in.txt/part-00000 out0.txt`
`dstk -c <clustername> hdfs cpout /user/root/in.txt/part-00001 out1.txt`


# TODO

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



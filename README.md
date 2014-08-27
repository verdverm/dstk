dstk - Data Science ToolKit
===========================

A tool for using other data science tools,
built on Docker and Hadoop.

It's a little rough right now

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

### Launch a Hadoop/Yarn/Spark cluster

`dtsk cluster create hadoop <name>`

### Teardown a Cluster

`dstk cluster destroy <name>`

### Get Cluster Status

`dstk cluster status <name>`

Not overly informative at this point, but the master IP is there.

### Open Ambari WebUI

open your browser to `<masterip>:8080`.

The user:pass is `admin:admin`

### Login to a node

`dstk -c <clustername> login <nodename>`

Environment variables from the Dockerfiles are missing on login...

### Open Ambari Shell

`dstk -c <clustername> cluster ambari`

### Run the SparkPi example

`dstk -c <clustername> -j <jobname> jobs run <command>`

currently the command is ignored... sorry

###

#!/bin/bash

# usecache=""
usecache="--no-cache"

# build base docker
docker build $usecache -t verdverm/dtsk-base base

# build java docker
# http://www.oracle.com/technetwork/java/javase/downloads/index.html
wget -nv --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" \
    http://download.oracle.com/otn-pub/java/jdk/8u20-b26/jdk-8u20-linux-x64.tar.gz
docker build $usecache -t verdverm/dtsk-java java

# build hadoop docker
# http://hadoop.apache.org/releases.html
curl http://mirror.reverse.net/pub/apache/hadoop/common/hadoop-2.4.1/hadoop-2.4.1.tar.gz > apache/hadoop/hadoop-2.4.1.tar.gz
docker build $usecache -t verdverm/dtsk-hadoop hadoop

# build spark docker
# http://spark.apache.org/downloads.html
curl http://mirror.reverse.net/pub/apache/spark/spark-1.0.2/spark-1.0.2-bin-hadoop2.tgz > apache/spark/spark-1.0.2-bin-hadoop2.tgz
docker build $usecache -t verdverm/dtsk-spark spark

#!/bin/bash

set -e

# usecache="--no-cache"

# build base docker
# ===================
# docker build $usecache -t verdverm/dstk-base base


# build java docker
# ===================
cd java
# rm *.tar.gz
# http://www.oracle.com/technetwork/java/javase/downloads/index.html
# wget --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" \
#     http://download.oracle.com/otn-pub/java/jdk/8u20-b26/jdk-8u20-linux-x64.tar.gz
# mv jdk-8u20-linux-x64.tar.gz java/jdk-8u20-linux-x64.tar.gz
# tar xzf jdk-8u20-linux-x64.tar.gz
# docker build $usecache -t verdverm/dstk-java .
cd ..


# build hadoop docker
# ===================
# http://hadoop.apache.org/releases.html
cd hadoop
# rm *.tar.gz
# curl http://mirror.reverse.net/pub/apache/hadoop/common/hadoop-2.5.0/hadoop-2.5.0.tar.gz | tar xz
# install flume 1.5.9
# curl http://www.bizdirusa.com/mirrors/apache/flume/1.5.0/apache-flume-1.5.0-bin.tar.gz | tar xz
# install sqoop 1.99.3
# curl http://mirrors.sonic.net/apache/sqoop/1.99.3/sqoop-1.99.3-bin-hadoop200.tar.gz | tar xz
# docker build $usecache -t verdverm/dstk-hadoop .
cd ..

# build spark docker
# ===================
cd spark
# rm *.tar.gz
# http://spark.apache.org/downloads.html
# curl http://mirror.reverse.net/pub/apache/spark/spark-1.0.2/spark-1.0.2-bin-hadoop2.tgz | tar xz
# docker build $usecache -t verdverm/dstk-spark .
cd ..

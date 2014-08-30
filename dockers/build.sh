#!/bin/bash

set -e

# usecache="--no-cache"

# build spark docker
# ===================
cd spark
# # http://spark.apache.org/downloads.html
# curl http://mirror.reverse.net/pub/apache/spark/spark-1.0.2/spark-1.0.2-bin-hadoop2.tgz | tar xz
# docker build $usecache -t verdverm/dstk-spark .
cd ..

# build app docker
# ===================
cd app
# scala
# http://www.scala-lang.org/download/all.html
# curl http://downloads.typesafe.com/scala/2.11.2/scala-2.11.2.tgz | tar xz
docker build $usecache -t verdverm/dstk-app .
cd ..

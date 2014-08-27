#!/bin/bash


# export AMBARI_HOST=172.17.0.28:8080
export JAVA_HOME=/usr/jdk64/jdk1.7.0_45

echo "Ambari Host: $AMBARI_HOST"

cd /tmp
/usr/jdk64/jdk1.7.0_45/bin/java -jar ambari-shell.jar --ambari.host=$AMBARI_HOST

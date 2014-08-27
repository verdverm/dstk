#!/bin/bash

su hdfs << EOF

export JAVA_HOME="/usr/jdk64/jdk1.7.0_45"
export HADOOP_CONF_DIR="/etc/hadoop/conf"
export YARN_CONF_DIR=$HADOOP_CONF_DIR

export PATH=$PATH:$JAVA_HOME/bin

echo $JAVA_HOME
echo $HADOOP_CONF_DIR
echo $YARN_CONF_DIR
echo $PATH
echo "USER: $USER"

# hadoop fs -mkdir /user
hadoop fs -mkdir /user/root
hadoop fs -chown root:root /user/root
EOF

#!/bin/bash

bindir="/apache/spark-1.0.1-bin-hadoop2/sbin"

service ssh start
ssh-keygen

function start_namenode() {
    # Format the namenode directory (DO THIS ONLY ONCE, THE FIRST TIME)
    $HADOOP_PREFIX/bin/hdfs namenode -format
    # Start the namenode daemon
    $HADOOP_PREFIX/sbin/hadoop-daemon.sh start namenode
}

function start_resourcemanager() {
    # Start the resourcemanager daemon
    $HADOOP_PREFIX/sbin/yarn-daemon.sh start resourcemanager
}

function start_workernode() {
    # Start the nodemanager daemon
    $HADOOP_PREFIX/sbin/yarn-daemon.sh start nodemanager
    # Start the datanode daemon
    $HADOOP_PREFIX/sbin/hadoop-daemon.sh start datanode
}


if [ "$NODE_TYPE" == "master" ]; then
    start_namenode
    start_resourcemanager
    $bindir/start-master.sh
fi

if [ "$NODE_TYPE" == "worker" ]; then
    start_workernode
    $bindir/start-slaves.sh
fi

if [ "$NODE_TYPE" == "single" ]; then
    start_namenode
    start_resourcemanager
    start_workernode
    $bindir/start-all.sh
fi

while :
do
    echo "Press [CTRL+C] to stop.."
    sleep 3600
done


# $HADOOP_PREFIX/bin/hadoop jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar org.apache.hadoop.yarn.applications.distributedshell.Client --jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar --shell_command date --num_containers 2 --master_memory 1024

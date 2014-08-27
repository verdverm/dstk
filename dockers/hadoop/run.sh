#!/bin/bash

function start_namenode() {
    # Format the namenode directory (DO THIS ONLY ONCE, THE FIRST TIME)
    $HADOOP_PREFIX/bin/hdfs namenode -format $CLUSTER_NAME
    # Start the namenode daemon
    $HADOOP_PREFIX/sbin/hadoop-daemon.sh --script hdfs start namenode
    # Start the MR JobHistory server
    $HADOOP_PREFIX/sbin/mr-jobhistory-daemon.sh start historyserver
}

function start_resourcemanager() {
    # Start the resourcemanager daemon
    $HADOOP_YARN_HOME/sbin/yarn-daemon.sh start resourcemanager
    # Start the WebAppProxy server
    $HADOOP_YARN_HOME/sbin/yarn-daemon.sh start proxyserver
}

function start_workernode() {
    # Start the datanode daemon
    $HADOOP_PREFIX/sbin/hadoop-daemon.sh --script hdfs start datanode
    # Start the nodemanager daemon
    $HADOOP_YARN_HOME/sbin/yarn-daemon.sh start nodemanager
}


if [ "$NODE_TYPE" == "master" ]; then
    start_namenode
    start_resourcemanager
    start_proxy_servers
fi

if [ "$NODE_TYPE" == "namenode" ]; then
    start_namenode
    start_resourcemanager
    start_proxy_servers
fi

if [ "$NODE_TYPE" == "resourcemanager" ]; then
    start_namenode
    start_resourcemanager
    start_proxy_servers
fi

if [ "$NODE_TYPE" == "worker" ]; then
    start_workernode
fi

if [ "$NODE_TYPE" == "single" ]; then
    start_namenode
    start_resourcemanager
    start_proxy_servers
    start_workernode
fi

while :
do
    echo "Press [CTRL+C] to stop.."
    sleep 3600
done


# $HADOOP_PREFIX/bin/hadoop jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar org.apache.hadoop.yarn.applications.distributedshell.Client --jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar --shell_command date --num_containers 2 --master_memory 1024

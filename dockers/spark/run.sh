#!/bin/bash

bindir="/apache/spark-1.0.1-bin-hadoop2/sbin"

service ssh start
ssh-keygen

if [ "$NODE_TYPE" == "master" ]; then
    $bindir/start-master.sh
fi

if [ "$NODE_TYPE" == "worker" ]; then
    $bindir/start-slaves.sh
fi

if [ "$NODE_TYPE" == "single" ]; then
    $bindir/start-all.sh
fi

while :
do
    echo "Press [CTRL+C] to stop.."
    sleep 3600
done


# $HADOOP_PREFIX/bin/hadoop jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar org.apache.hadoop.yarn.applications.distributedshell.Client --jar $HADOOP_PREFIX/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-2.4.1.jar --shell_command date --num_containers 2 --master_memory 1024

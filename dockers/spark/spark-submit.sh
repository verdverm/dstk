#!/bin/bash

export JAVA_HOME="/usr/jdk64/jdk1.7.0_45"
export HADOOP_CONF_DIR="/etc/hadoop/conf"
export YARN_CONF_DIR=$HADOOP_CONF_DIR

export PATH=$PATH:$JAVA_HOME/bin

echo $JAVA_HOME
echo $HADOOP_CONF_DIR
echo $YARN_CONF_DIR
echo $PATH

sparkbinary=/apache/spark-1.0.2-bin-hadoop2/bin/spark-submit


cd /tmp
curl -L 'http://172.17.0.209:50070/webhdfs/v1/user/root/wordcount.jar?user.name=hdfs&op=OPEN' > wordcount.jar


 $sparkbinary --class wordcount.WordCount \
    --verbose \
    --master yarn-cluster \
    --num-executors 2 \
    --driver-memory 1g \
    --executor-memory 1g \
    --executor-cores 1 \
    wordcount.jar




# cd /apache/spark-1.0.2-bin-hadoop2

# ./bin/spark-submit --class org.apache.spark.examples.SparkPi \
#     --verbose \
#     --master yarn-cluster \
#     --num-executors 2 \
#     --driver-memory 1g \
#     --executor-memory 1g \
#     --executor-cores 1 \
#     lib/spark-examples*.jar \
#     10

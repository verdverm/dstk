FROM verdverm/dstk-java
MAINTAINER Tony Worm verdverm@gmail.com
# adapted from
# http://www.alexjf.net/blog/distributed-systems/hadoop-yarn-installation-definitive-guide

RUN apt-get update && apt-get install --no-install-recommends -y rsync

ENV HADOOP_PREFIX       /apache/hadoop-2.5.0
ENV HADOOP_HOME         $HADOOP_PREFIX
ENV HADOOP_COMMON_HOME  $HADOOP_PREFIX
ENV HADOOP_CONF_DIR     $HADOOP_PREFIX/etc/hadoop
ENV HADOOP_HDFS_HOME    $HADOOP_PREFIX
ENV HADOOP_MAPRED_HOME  $HADOOP_PREFIX
ENV HADOOP_YARN_HOME    $HADOOP_PREFIX

# hack so that HADOOP_IDENT_STRING actually has a value...
# stupid env handling in (hadoop-env.sh & hadoop-daemon.sh)
ENV USER v2

RUN mkdir -p /apache

ADD hadoop-2.5.0        /apache/hadoop-2.5.0
ADD files/hdfs-site.xml /apache/hadoop-2.5.0/etc/hadoop/hdfs-site.xml
ADD files/core-site.xml /apache/hadoop-2.5.0/etc/hadoop/core-site.xml
ADD files/yarn-site.xml /apache/hadoop-2.5.0/etc/hadoop/yarn-site.xml

ADD apache-flume-1.5.0-bin     /apache/apache-flume-1.5.0-bin
ADD sqoop-1.99.3-bin-hadoop200 /apache/sqoop-1.99.3-bin-hadoop200

ADD run.sh /apache/run.sh
CMD /apache/run.sh

EXPOSE 8088
EXPOSE 8032

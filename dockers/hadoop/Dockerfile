FROM verdverm/java:latest
MAINTAINER Tony Worm verdverm@gmail.com
# adapted from
# http://www.alexjf.net/blog/distributed-systems/hadoop-yarn-installation-definitive-guide

# install hadoop 2.4.1
RUN mkdir -p /apache
WORKDIR /apache
# RUN curl http://mirror.reverse.net/pub/apache/hadoop/common/hadoop-2.4.1/hadoop-2.4.1.tar.gz | tar xz
ADD hadoop-2.4.1.tar.gz /apache/hadoop-2.4.1.tar.gz
RUN tar xzf hadoop-2.4.1.tar.gz

ENV HADOOP_PREFIX       /apache/hadoop-2.4.1
ENV HADOOP_HOME         $HADOOP_PREFIX
ENV HADOOP_COMMON_HOME  $HADOOP_PREFIX
ENV HADOOP_CONF_DIR     $HADOOP_PREFIX/etc/hadoop
ENV HADOOP_HDFS_HOME    $HADOOP_PREFIX
ENV HADOOP_MAPRED_HOME  $HADOOP_PREFIX
ENV HADOOP_YARN_HOME    $HADOOP_PREFIX

# hack so that HADOOP_IDENT_STRING actually has a value...
# stupid env handling in (hadoop-env.sh & hadoop-daemon.sh)
ENV USER v2

ADD files/hdfs-site.xml /apache/hadoop-2.4.1/etc/hadoop/hdfs-site.xml
ADD files/core-site.xml /apache/hadoop-2.4.1/etc/hadoop/core-site.xml
ADD files/yarn-site.xml /apache/hadoop-2.4.1/etc/hadoop/yarn-site.xml
ADD run.sh              /apache/run.sh


# install flume 1.5.9
RUN curl http://www.bizdirusa.com/mirrors/apache/flume/1.5.0/apache-flume-1.5.0-bin.tar.gz | tar xz

# install sqoop 1.99.3
RUN curl http://mirrors.sonic.net/apache/sqoop/1.99.3/sqoop-1.99.3-bin-hadoop200.tar.gz | tar xz


# http://tez.apache.org/install.html
# http://apache.mirrors.lucidnetworks.net/incubator/tez/tez-0.4.1-incubating/tez-0.4.1-incubating-src.tar.gz

CMD /apache/run.sh

EXPOSE 8088
EXPOSE 8032

# FROM verdverm/dstk-hadoop
FROM sequenceiq/ambari:latest
MAINTAINER Tony Worm verdverm@gmail.com

RUN yum install -y sudo

RUN mkdir /apache
ADD spark-1.0.2-bin-hadoop2 /apache/spark-1.0.2-bin-hadoop2

# # ADD spark startup to serf handlers
# ADD serf /usr/local/serf

ADD myblueprint /tmp/myblueprint
ADD install-cluster.sh /tmp/install-cluster.sh
ADD mkdir-hadoop.sh /tmp/mkdir-hadoop.sh
ADD dstk-ambari-shell.sh /tmp/dstk-ambari-shell.sh
ADD spark-submit.sh /tmp/spark-submit.sh
# # CMD /apache/run-spark.sh

# EXPOSE 8080 7077

# from sequenceiq/ambari
ENTRYPOINT ["/usr/local/serf/bin/start-serf-agent.sh"]
CMD ["--log-level", "debug"]

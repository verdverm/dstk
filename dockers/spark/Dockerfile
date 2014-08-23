FROM verdverm/dstk-hadoop
MAINTAINER Tony Worm verdverm@gmail.com

RUN apt-get install --no-install-recommends -y openssh-server

ADD spark-1.0.2-bin-hadoop2 /apache/spark-1.0.2-bin-hadoop2

EXPOSE 8080 7077 22

ADD run.sh /apache/run-spark.sh
CMD /apache/run-spark.sh


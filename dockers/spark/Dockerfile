FROM verdverm/hadoop:latest
MAINTAINER Tony Worm verdverm@gmail.com

WORKDIR /apache

# install spark 1.0.2
ADD spark-1.0.2-bin-hadoop2.tgz /apache/spark-1.0.2-bin-hadoop2.tgz
# RUN curl http://mirror.reverse.net/pub/apache/spark/spark-1.0.2/spark-1.0.2-bin-hadoop2.tgz > spark-1.0.2-bin-hadoop2.tgz
RUN tar zxf spark-1.0.2-bin-hadoop2.tgz

# setup spark configuration
WORKDIR /apache/spark-1.0.2-bin-hadoop2
# RUN echo "localhost" >> conf/slaves

RUN apt-get install --no-install-recommends -y openssh-server

EXPOSE 8080 7077 22

ADD run.sh /apache/run-spark.sh
CMD /apache/run-spark.sh


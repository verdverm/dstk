FROM verdverm/dstk-spark:latest
MAINTAINER Tony Worm verdverm@gmail.com

ADD scala-2.11.2 /usr/lib/scala
ENV SCALA_HOME /usr/lib/scala
ENV PATH $PATH:/usr/lib/scala/bin:

RUN mkdir -p /jobs/usercode
# RUN cp scala-library.jar /jobs/scala-library.jar

ADD scripts /jobs/scripts

VOLUME ["/jobs/usercode"]

ENTRYPOINT ["/jobs/scripts/main.sh"]

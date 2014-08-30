FROM verdverm/dstk-spark:latest
MAINTAINER Tony Worm verdverm@gmail.com

ADD scala-2.11.2 /usr/lib/scala
ENV SCALA_HOME /usr/lib/scala
ENV PATH $PATH:/usr/lib/scala/bin:

RUN mkdir -p /app/usercode
# RUN cp scala-library.jar /app/scala-library.jar

ADD scripts /app/scripts

VOLUME ["/app/usercode"]

ENTRYPOINT ["/app/scripts/main.sh"]

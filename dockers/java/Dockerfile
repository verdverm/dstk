FROM verdverm/dstk-base
MAINTAINER Tony Worm verdverm@gmail.com

RUN apt-get update

# install openjdk-7 java runtime
# RUN apt-get install --no-install-recommends -y \
#     openjdk-7-jdk

# install oracle-jdk-8 java runtime
RUN mkdir -p /usr/lib/jvm
ADD jdk1.8.0_20 /usr/lib/jvm/jdk1.8.0_20

WORKDIR /usr/lib/jvm
RUN ln -s /usr/lib/jvm/jdk1.8.0_20/bin/* /usr/local/bin

ENV JAVA_HOME /usr/lib/jvm/jdk1.8.0_20

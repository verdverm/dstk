FROM verdverm/dstk-base
MAINTAINER Tony Worm verdverm@gmail.com

RUN apt-get update

# install openjdk-7 java runtime
# RUN apt-get install --no-install-recommends -y \
#     openjdk-7-jdk

# install oracle-jdk-8 java runtime
RUN mkdir -p /usr/lib/jvm
# RUN wget -nv --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" \
#     http://download.oracle.com/otn-pub/java/jdk/8u11-b12/jdk-8u20-linux-x64.tar.gz
ADD jdk-8u20-linux-x64.tar.gz /usr/lib/jvm/jdk-8u20-linux-x64.tar.gz

WORKDIR /usr/lib/jvm
RUN ls
RUN cat jdk-8u20-linux-x64.tar.gz | tar xz
# RUN rm jdk-8u20-linux-x64.tar.gz
RUN ln -s /usr/lib/jvm/jdk1.8.0_20/bin/* /usr/local/bin

ENV JAVA_HOME /usr/lib/jvm/jdk1.8.0_20

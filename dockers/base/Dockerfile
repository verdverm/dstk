FROM debian:wheezy
MAINTAINER Tony Worm verdverm@gmail.com

RUN apt-get update && apt-get upgrade -y

# base stuff
RUN apt-get install --no-install-recommends -y \
  build-essential \
  ca-certificates \
  wget \
  curl \
  git \
  mercurial \
  net-tools \
  ssh

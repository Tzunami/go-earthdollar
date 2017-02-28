FROM ubuntu:wily
MAINTAINER caktux

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get upgrade -q -y && \
    apt-get dist-upgrade -q -y && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 923F6CA9 && \
    echo "deb http://ppa.launchpad.net/earthdollar/earthdollar/ubuntu wily main" | tee -a /etc/apt/sources.list.d/earthdollar.list && \
    apt-get update && \
    apt-get install -q -y ged

EXPOSE 8811
EXPOSE 20203

ENTRYPOINT ["/usr/bin/geth"]

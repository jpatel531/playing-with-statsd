FROM ubuntu:14.04
MAINTAINER Jamie Patel

RUN apt-get update
RUN apt-get install -yqq git nodejs devscripts debhelper systemd dh-systemd npm nodejs-legacy
RUN mkdir ~/build

RUN cd /root/build && git clone https://github.com/etsy/statsd.git
RUN cd /root/build/statsd && dpkg-buildpackage
RUN cd /root/build && dpkg -i statsd*.deb

ADD conf/localConfig.js /etc/statsd/localConfig.js

WORKDIR /root/build/statsd/bin

CMD ["nodejs", "statsd", "/etc/statsd/localConfig.js", ">> statsd.log"]
FROM golang:1.9.2-alpine3.7

RUN wget -O /tmp/cf.tar "https://packages.cloudfoundry.org/stable?release=linux64-binary&version=6.33.1&source=github-rel"
RUN tar -zxvf /tmp/cf.tar
RUN mv cf /usr/bin

RUN mkdir -p /opt/resource
RUN mkdir -p /opt/plugin

ADD protocol /opt/plugin/
RUN cf install-plugin /opt/plugin/protocol -f

ADD in /opt/resource
ADD out /opt/resource
ADD check /opt/resource

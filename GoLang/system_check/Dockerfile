FROM ubuntu:22.04

RUN apt-get update && apt-get install -y wget && apt-get install -y sysstat && apt-get install -y iproute2 && apt-get install -y mtr-tiny && apt-get install -y vim
RUN wget https://dl.google.com/go/go1.22.5.linux-amd64.tar.gz && tar -xvf go1.22.5.linux-amd64.tar.gz
RUN mv go /usr/local
RUN rm -rf go1.22.5.linux-amd64.tar.gz
RUN mkdir -pv /home/golang/system_check/
COPY main.go /home/golang/system_check/main.go
FROM ubuntu:latest

RUN apt-get update && apt-get install -y dpkg

COPY myprogram.deb /tmp/

RUN dpkg -i /tmp/myprogram.deb

ENTRYPOINT ["/usr/bun/lab"]

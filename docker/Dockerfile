FROM alpine:edge

# Setup GO
RUN apk update &&\
    apk add --no-cache curl bash git gcc musl-dev ca-certificates go &&\
    git clone https://go.googlesource.com/go goroot --progress --verbose &&\
    cd goroot/src &&\
    ./make.bash &&\
    apk del go gcc musl-dev curl bash &&\
    mv /goroot/bin/go /usr/bin/go &&\
    mv /goroot/bin/gofmt /usr/bin/gofmt


RUN apk add --no-cache btrfs-progs e2fsprogs e2fsprogs-extra iptables xfsprogs xz
RUN apk add --no-cache docker iptables bash btrfs-progs
RUN apk add openrc --no-cache

RUN rc-update add docker boot\
  && mkdir /run/openrc\
  && touch /run/openrc/softlevel

# Setup docker in docker
#
#COPY entrypoint.sh /usr/local/bin/
#RUN apk add --no-cache docker iptables &&\
#    wget https://raw.githubusercontent.com/docker/docker/3b5fac462d21ca164b3778647420016315289034/hack/dind -O /usr/local/bin/dind &&\
#    chmod +x /usr/local/bin/dind &&\
#    chmod +x /usr/local/bin/entrypoint.sh
# https://github.com/yobasystems/alpine-docker/blob/master/alpine-dind-amd64/Dockerfile
VOLUME /var/lib/dockers
VOLUME /var/lib/docker
VOLUME /var/lib/docker/btrfs
EXPOSE 2375
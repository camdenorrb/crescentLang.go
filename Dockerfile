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


ENV DIND_COMMIT 3b5fac462d21ca164b3778647420016315289034

RUN apk add docker iptables --no-cache
# https://github.com/yobasystems/alpine-docker/blob/master/alpine-dind-amd64/Dockerfile
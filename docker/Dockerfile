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

# Setup docker in docker
COPY entrypoint.sh /usr/local/bin/
RUN apk add --no-cache docker iptables &&\
    wget https://raw.githubusercontent.com/docker/docker/3b5fac462d21ca164b3778647420016315289034/hack/dind -O /usr/local/bin/dind &&\
    chmod +x /usr/local/bin/dind &&\
    chmod +x /usr/local/bin/entrypoint.sh

VOLUME /var/lib/docker
EXPOSE 2375

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD []
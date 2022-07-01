FROM ubuntu:18.04

WORKDIR /home/5gvinni

RUN apt-get update \
	&& apt-get install --no-install-recommends -y git wget net-tools iputils-ping libpcap-dev build-essential ca-certificates

COPY  5gvinni-stgutg /home/5gvinni/5gvinni-stgutg
#RUN git clone https://github.com/UPM-RSTI/5gvinni-stgutg 

RUN wget --no-check-certificate https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz \
	&& tar -C /usr/local -zxvf go1.12.9.linux-amd64.tar.gz \
	&& mkdir -p ~/go/{bin,pkg,src}

ENV GOPATH=/root/go:/home/5gvinni/5gvinni-stgutg
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOPATH/bin:$GOROOT/bin
ENV GO111MODULE=off

RUN go get -u github.com/aead/cmac \
	&& go get -u github.com/antonfisher/nested-logrus-formatter \
	&& go get -u github.com/calee0219/fatal \
	&& go get -u github.com/dgrijalva/jwt-go \
	&& go get -u github.com/ghedo/go.pkt/capture/pcap \
	&& go get -u github.com/gin-gonic/gin \
	&& go get -u github.com/ishidawataru/sctp \
	&& go get -u golang.org/x/net/ipv4 \
	&& go get -u gopkg.in/yaml.v2

RUN cd /home/5gvinni/5gvinni-stgutg \
	&& go build src/stg-utg.go

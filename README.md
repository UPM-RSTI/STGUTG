# 5gvinni-stgutg


## Installation with Docker

1. Download all github 5gvinni-stgutg repository.
2. Extract dockerfile and docker-compose files from folder 5gvinni-stgutg. This is because Dockerfile needs copy 5gvinni-stgutg folder to docker container.

Default path tree (e.g /home/5gvinni-stgutg)
  - src
  - utils
  - Dockerfile
  - README.md
  - config.yaml
  - docker-compose.yml
  - init.sh
  - stg-utg

Path tree for docker (e.g /home)
  - 5gvinni-stgutg
  - Dockerfile
  - docker-compose.yml

## Installation without Docker

### 1. Download and install GO 1.12.9

`wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz`

`sudo tar -C /usr/local -zxvf go1.12.9.linux-amd64.tar.gz`

`mkdir -p ~/go/{bin,pkg,src}`

### 2. Clone project

`git clone git@github.com:UPM-RSTI/5gvinni-stgutg.git`


### 3. Configure Environment variables

`export GOPATH=/home/user/go:/home/user/5gvinni-stgutg` (Or the pathds where the go folder and the cloned project are stored)

`export GOROOT=/usr/local/go`

`export PATH=$PATH:$GOPATH/bin:$GOROOT/bin`

`export GO111MODULE=off`

This configuration can be added to ~/.bashrc.

### 4. Install dependencies

`go get github.com/aead/cmac`

`go get github.com/antonfisher/nested-logrus-formatter`

`go get github.com/calee0219/fatal`

`go get github.com/dgrijalva/jwt-go`

`apt-get install build-essential`

`apt-get install libpcap-dev`

`go get github.com/ghedo/go.pkt/capture/pcap`

`go get github.com/gin-gonic/gin`

`go get github.com/ishidawataru/sctp`

`go get golang.org/x/net/ipv4`

`go get gopkg.in/yaml.v2`


### 5. Build executable

`go build src/stg-utg.go`


### 6. Configure and run

`nano src/config.yaml`

`stg-utg` or

`stg-utg -t` for testing mode

# STGUTG


## Installation with Docker

1. Download all github stgutg repository.
2. Extract dockerfile and docker-compose files from folder stgutg. This is because Dockerfile needs copy stgutg folder to docker container.

Default path tree (e.g /home/stgutg)
  - src
  - utils
  - Dockerfile
  - README.md
  - config.yaml
  - docker-compose.yml
  - init.sh
  - stg-utg

Path tree for docker (e.g /home)
  - stgutg
  - Dockerfile
  - docker-compose.yml

## Installation without Docker

### 1. Download and install GO 1.12.9

`wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz`

`sudo tar -C /usr/local -zxvf go1.12.9.linux-amd64.tar.gz`

`mkdir -p ~/go/{bin,pkg,src}`

### 2. Clone project

`git clone git@github.com:UPM-RSTI/STGUTG.git`


### 3. Configure Environment variables

`export GOPATH=/home/user/go:/home/user/STGUTG` (Or the paths where the go folder and the cloned project are stored)

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

### 4.b Select dependencies
`cd ~/go/src/github.com/gin-gonic/gin`

`git checkout v1.7.0`

`go get github.com/free5gc/nas/nasMessage`

`go get github.com/free5gc/openapi/models`

`go get github.com/golang/protobuf/proto`

`cd ~/go/src/github.com/free5gc/nas`

`git checkout v1.0.1`

`cd ~/go/src/github.com/free5gc/openapi`

`git checkout v1.0.2`

`go get github.com/free5gc/logger_conf`

`go get github.com/free5gc/logger_util`

`go get github.com/free5gc/ngap`

`go get github.com/free5gc/util_3gpp`

### 5. Build executable

`go build src/stg-utg.go`


### 6. Configure and run

`nano src/config.yaml`

`stg-utg` or

`stg-utg -t` for testing mode

## Example scenario with Free5gc-Compose v3.0.5

![](esquemagit.png)

This is a network scenario in which we are going to use Free5gc-Compose v3.0.5 and the STGUTG to give Internet access to a virtual machine. The scenario consist on 3 VMs as it is represented on the picture.

### 1. Install Free5gc-compose v3.0.5
 You should start by installing Docker and Docker-compose, then you should be able to download the Free5gc v3.0.5 repository to start adapting the docker-compose.yaml and the NFs configuration files. https://github.com/free5gc/free5gc-compose/releases/tag/v3.0.5
 
### 2. Configure the docker-compose.yaml
First of all erase the declaration of the containers upf1, upf2 and n3iwf and all their mentions on the rest of the file.

Second step in the AMF section add

```
ports :
  - "48412:38412/sctp"
```
  
in the UPF section add
```
ports :
  - " 2152:2152/ udp "
  - " 2152:2152/ tcp "
command : sh -c "chmod +x upf-iptables.sh && ./upf-iptables.sh && ./free5gc-upfd -f ..config/upfcfg.yaml"
volumes :
  - ./config/upfcfgb.yaml:/free5gc/config/upfcfg.yaml
  - ./config/upf-iptables.sh:/free5gc/free5gc-upfd/upf-iptables.sh
```
After this we should create an script with the following sentences for the iptables of upfb container. We save it in the config/ folder of the Free5gc-compose folder.
```
# !/ bin / sh
iptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE
iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
IPTABLES -I FORWARD 1 -j ACCEPT
```

Finally on the NRF section we add to the "depends on:" the upfb
```
  depends_on:
    - db
    - free5gc-upf-b
```
### 2. Configure the NFs of the Free5gc-Compose

Go to /config/smfcfg.yaml and add our UTG IP address, which is 192.168.43.2, in the gnb1 section. 
```
gNB1 : # the name of the node
  type : AN # the type of the node ( AN or UPF )
  an_ip : 192.168.43.2
```

### 3. Add the Network interfaces to the Free5gc VM and the IPtables rules

Add this IPs with ifconfig to the network interfaces that you created to connect the Free5gc machine with the STGUTG machine, in our case they are ens5 and ens4.
```
sudo ifconfig ens4 192.168.43.3 netmask 255.255.255.0 up
sudo ifconfig ens5 192.168.42.3 netmask 255.255.255.0 up
```
Then execute this commands on the terminal to allow the Free5gc-Compose to reach Internet, in our case the interface to do it was ens6. Be sure to use the one that fits your scenario.

```
sudo sysctl -w net . ipv4 . ip_forward =1
sudo iptables -t nat -A POSTROUTING -o <Interface that reaches internet> -j MASQUERADE
sudo systemctl stop ufw
sudo iptables -I FORWARD 1 -j ACCEPT
```
### 4. Configure the STGUTG

First add the IPs of the picture to the network interfaces that you have previosuly created to interconnect both machines. In our case they are 
```
sudo ifconfig enp0s9 192.168.43.2 netmask 255.255.255.0 up
sudo ifconfig enp0s8 192.168.42.2 netmask 255.255.255.0 up
sudo ifconfig enp0s10 60.60.0.3 netmask 255.255.255.0 up
```
Now we are going to modify the config.yaml to match the network configuration.

As you will see, this IP addresses are already written in the config, so you can already identify which IP is which by looking on the picture.
```
# Network Functions
  amf_ngap: 192.168.42.3
  amf_port: 48412

  upf_gtp: 192.168.43.3
  upf_port: 2152

  gnb_gtp: 192.168.43.2
  gnbg_port: 9487

  gnb_ngap: 192.168.42.2
  gnbn_port: 2152

  ue_ori: 60.60.0.1
```

The next part to configure are the MAC addresses and network interfaces. You need to add on src_iface the interface of the STGUTG VM that faces the UE VM which in our case is enp0s10, then on dst_iface add the interface of the STGUTG VM that faces the Free5gc VM on the user plane (enp0s9 in this example).
Lastly you need to add JUST the eth_src, which is the MAC address of the interface of the STGUTG VM that faces the UE VM which in our case is enp0s10. eth_dst is not needed. Lastly you can decide how many UE VM do you want to give coberage in the parameter ue_number. For this example 1 is correct since we only have one VM as UE.

```
 # Interaces and MACs for traffic
  src_iface: "enp0s10"
  dst_iface: "enp0s9"
  eth_src: "fa163e2d99eb"
  eth_dst: "fa163ef6eaed"

  # Number of UEs to use in traffic mode
  ue_number: 1
```

### 4. Configure the UE VM

In this VM you just need to create an interface to reach the STGUTG VM and make a default route to make all the traffic reach the STGUTG. It is important to have an MTU of 1400, since our STGUTG cant apply packet fragmentation.

```
sudo ifconfig enp0s9 192.168.43.2 netmask 255.255.255.0 mtu 1400 up
sudo ip route add default via 60.60.0.3
```

### 5. Launch the scenario

Once you have checked that the VMs are able to ping each other it is time to establish the 5G Internet access for the UE VM.

First step:  Go to the Free5gc VM and execute in the Free5gc-compose folder

```
sudo docker-compose build
sudo docker-compose start free5gc-upf-b
sudo docker-compose up
```

Second Step: subscribe the UEs, go to a web browser and access the VM of the Free5gc by searching http://<ip of free5gc VM>:5000

Then introduce the credentials admin/free5gc.
Once you are logged in just add a new subscriber with out modifying any parameter.

Third Step: Go to the STGUTG VM and execute
```
sudo ./stg-utg
```

Forth Step: Use internet on the UE VM to check the connectivity. If you ping to 8.8.8.8 and the time is around 10 ms it means that you are correctly using the 5G scenario we have just deploying.





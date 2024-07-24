
![](5gtc.png) 

# 5G Traffic Connector

5G Traffic Connector is a software created for the generation of both signal and user traffic to be sent to a 5G network core. It is based on implementations from the [Free5GC](https://www.free5gc.org/) project and is distributed under an Apache 2.0 license.

Developed by [UPM RSTI Research group](https://blogs.upm.es/rsti).

## Installation 

### 1. Build executable

```
go build
```

### 2. Configure and run

```
nano src/config.yaml
```
```
stgutgmain 
```
or
```
stgutgmain -t 
```
for testing mode

---

## Example: Deployment scenario with Open5GS

![](esquemagit.png)

This is a network scenario in which we are going to use Open5GS and the STGUTG to give Internet access to a virtual machine. The scenario consists of 3 VMs as it is represented in the picture. 

[Open5GS](https://open5gs.org/) is an open-source project for 5th generation (5G) mobile core networks, which intends to implement the 5G core network (5GC) defined in 3GPP Release 17 (R17). In this example, we use the NFs implemented in Open5GS to deploy a 5G core and then test the STGUTG software.

### 1. Install Open5GS

Open5GS is a repository for the NFs implemented by 5G. As a prerequisite, MongoDB should be installed.

This is a guide for the instalation of Open5GS: [Open5GS Quickstart](https://open5gs.org/open5gs/docs/guide/01-quickstart/).

Once the tools are installed, clone the [Open5GS](https://open5gs.org/) to start adapting the NFs configuration files. 
 
### 2. Configure Open5GS

1. Configure the AMF (amf.yaml) with the PLMN required. Additionally, it is important to change the IP address indicating the location of the NF, as well as adapt the sd to the one used in the connector.

```
  sd: 010203
```

```
ngap:
  - addr: 192.168.61.4
```

```
  mcc: 001
  mnc: 01
```

2. Configure the UPF (upf.yaml), indicating the IP address of the NF and the subnet of the N6.

```
gtpu:
  - addr: 192.168.61.4
```

```
subnet:
  - addr: 10.46.0.1/16
```

3. Register the subscriber information through the WebUI.

```
IMSI: 001010000000001
Subscriber Key (K): 465B5CE8 B199B49F AA5F0A2E E238A6BC
Operator Key (OPc): E8ED289D EBA952E4 283B54E8 8E6183CA
```

4. Activate the IPv4 and IPv6 forwarding and add the NAT rule.

```
sudo sysctl -w net.ipv4.ip_forward=1
sudo sysctl -w net.ipv6.conf.all.forwarding=1
sudo iptables -t nat -A POSTROUTING -s 10.46.0.0/16 ! -o ogstun -j MASQUERADE
sudo ip6tables -t nat -A POSTROUTING -s 2001:db8:cafe::/48 ! -o ogstun -j MASQUERADE
```

5. Create the TUN.

```
sudo ip tuntap add name ogstun mode tun 
sudo ip addr add 10.46.0.1/16 dev ogstun 
sudo ip addr add 2001:db8:cafe::1/48 dev ogstun 
sudo ip link set ogstun up
```

### 3. Configure STGUTG

1. Add the IPs of the figure to the network interfaces (previously created to interconnect both machines). In othe example the addresses are: 

```
sudo ifconfig enp0s8 10.45.0.4 netmask 255.255.255.0 up
```
```
sudo ifconfig enp0s9 192.168.61.3 netmask 255.255.255.0 up
```

2. Modify the config.yaml file to match the network configuration (the example IP addresses are already written in the config file, so in this case, IP addresses can be matched with the ones shown in the figure).

```
# Network Functions
  amf_ngap_ip: 192.168.61.4
  amf_ngap_port: 38412 #48412

  gnb_gtp_ip: 192.168.61.3

  stg_ngap_ip: 192.168.61.3
  stg_ngap_port: 9487
```

3. Adapt the UE and gNB configuration to match with the Open5GS ones.

```
# UE
  initial_imsi: "001010000000001"
  mcc: "001"
  mnc: "01"

# GNB data
  gnb_id: "\x00\x01\x02"
  gnb_bitlength: 24
  gnb_name: "open5gs"

# UE Authentication Data
  k: "465B5CE8B199B49FAA5F0A2EE238A6BC"
  opc: "E8ED289DEBA952E4283B54E88E6183CA"
  op: "E8ED289DEBA952E4283B54E88E6183CA"


  sst: 1
  sd: "010203"
```

4. Configure MAC addresses and network interfaces.

- src_iface is the interface of the STGUTG VM that faces the UE VM (enp0s8 in the example).
- dst_iface is the interface of the STGUTG VM that faces the Open5GS VM on the user plane (enp0s9 in the example).
- ue_number is the number of UEs to be emulated (it depends on the number of VMs deployed as UEs).

```
  # Interaces for traffic
  src_iface: "enp0s8"
  dst_iface: "enp0s9"

  # Number of UEs to use in traffic mode
  ue_number: 1
```

### 4. UE VM configuration 

Create an interface to reach the STGUTG VM and make a default route to make all the traffic reach the STGUTG. Set MTU to 1400B.


```
sudo ifconfig enp0s3 10.45.0.3 netmask 255.255.255.0 mtu 1400 up
```
```
sudo ip route add default via 10.45.0.4
```

### 5. Run the scenario

1. In Open5GS VM, execute in the Open5GS folder the following commands. This will start the NFs of the 5G core:

```
sudo systemctl restart open5gs-mmed
sudo systemctl restart open5gs-sgwcd
sudo systemctl restart open5gs-smfd
sudo systemctl restart open5gs-amfd
sudo systemctl restart open5gs-sgwud
sudo systemctl restart open5gs-upfd
sudo systemctl restart open5gs-hssd
sudo systemctl restart open5gs-pcrfd
sudo systemctl restart open5gs-nrfd
sudo systemctl restart open5gs-scpd
sudo systemctl restart open5gs-seppd
sudo systemctl restart open5gs-ausfd
sudo systemctl restart open5gs-udmd
sudo systemctl restart open5gs-pcfd
sudo systemctl restart open5gs-nssfd
sudo systemctl restart open5gs-bsfd
sudo systemctl restart open5gs-udrd
sudo systemctl restart open5gs-webui
```

2.  run the STGUTG software:
```
sudo ./stgutgmain
```

3. Use the UE VM to send traffic through the core to any Internet-based service (ping to 8.8.8.8 should suffice to test if the configuration is successful).

---

![](logorsti.png) 

#!/usr/bin/env bash


## Configuracion
NETIF=enp0s3
AMF=192.168.0.100
#GNBN=10.100.255.39
GNBN=10.100.200.4
SMFP=192.168.2.1
UPFP=192.168.2.2
GNBG=10.200.200.2
UPFG=192.168.0.101
UE=60.60.0.101
#UENET=60.60.0.0
UENET=172.100.100.0
#UENET=60.60.0.0
UEMSK=24
UEIF=enp0s9
LOCALHOST=127.0.0.1
GNBAN=192.188.2.3 #?
PINGSRC=60.60.0.1
PINGDST=8.8.8.8
#PINGDST=192.168.0.39


# Se obtiene la IP de la interfaz NETIF
IPLINE=$(ifconfig $NETIF | grep "inet ")
GI=$(echo $IPLINE | cut -d " " -f 2)

# Se habilita el ip forwarding
echo "> Habilitando IP Forwarding"
sysctl -w net.ipv4.ip_forward=1

# Se configuran las interfaces de red
echo "> Configurando interfaces de red"

#ifconfig enp0s3:amf $AMF up
ifconfig enp0s8:amf $AMF netmask 255.255.255.0
ifconfig enp0s8:amf up
echo ">> AMF: $AMF"

#ifconfig enp0s3:gnbn $GNBN up
#echo ">> GNBN: $GNBN"

ifconfig enp0s3:smfp $SMFP up
echo ">> SMFP: $SMFP"

ifconfig enp0s3:upfp $UPFP up
echo ">> UPFP: $UPFP"

#fconfig enp0s3:gnbg $GNBG up
#echo ">> GNBG: $GNBG"

#ifconfig enp0s3:upfg $UPFG up
ifconfig enp0s9:upfg $UPFG netmask 255.255.255.0
ifconfig enp0s9:upfg up
echo ">> UPFG: $UPFG"

ifconfig enp0s3:ue $UE up
echo ">> UE: $UE"

# Borramos la tabla de nat por si acaso
echo "> Borrando reglas NAT"
iptables -t nat -F

# Se añade forwarding a iptables
iptables -A FORWARD -j ACCEPT

# Se añade una regla de NAT para la interfaz de salida
echo "> Añadiendo NAT sobre interfaz de red"
iptables -t nat -A POSTROUTING -o $NETIF -j MASQUERADE
#iptables -t nat -A POSTROUTING -s $UENET/$UEMSK -j SNAT --to-source $GI
iptables -t nat -nvL POSTROUTING

# Se deshabilita el firewall
echo "> Deshabilitando firewall"
systemctl stop ufw
systemctl disable ufw

CURDATE=$(date '+%Y-%m-%d-%H-%M-%S')
cd config

echo "> Configurando amfcfg.conf"
#cp amfcfg.conf amfcfg.conf.$CURDATE

# Dirección AMF
NGAPLINE=$(awk '/ngapIpList/{ print NR; exit }' amfcfg.conf)
NGAPLINE2=$((NGAPLINE+1))
sed -i "${NGAPLINE2}s/-.*/- ${AMF}/" amfcfg.conf

echo "> Configurando smfcfg.conf"
#cp smfcfg.conf smfcfg.conf.$CURDATE

# Dirección PFCP
NGAPLINE=$(awk '/pfcp/{ print NR; exit }' smfcfg.conf)
NGAPLINE2=$((NGAPLINE+1))
sed -i "${NGAPLINE2}s/addr.*/addr: ${SMFP}/" smfcfg.conf

#Dirección gNB1
NGAPLINE=$(awk '/an_ip/{ print NR; exit }' smfcfg.conf)
sed -i "${NGAPLINE}s/an_ip.*/an_ip: ${GNBAN}/" smfcfg.conf

# Dirección UPF
NGAPLINE=$(awk '/node_id/{ print NR; exit }' smfcfg.conf)
sed -i "${NGAPLINE}s/node_id.*/node_id: ${UPFP}/" smfcfg.conf

# Dirección red UE
NGAPLINE=$(awk '/ue_subnet/{ print NR; exit }' smfcfg.conf)
sed -i "${NGAPLINE}s/ue_subnet.*/ue_subnet: ${UENET}\/${UEMSK}/" smfcfg.conf

echo "> Configurando upfcfg.yaml"
cd ../src/upf/build/config
#cp upfcfg.yaml upfcfg.yaml.$CURDATE

NGAPLINE=$(awk '/pfcp/{ print NR; exit }' upfcfg.yaml)
NGAPLINE2=$((NGAPLINE+1))
sed -i "${NGAPLINE2}s/addr.*/addr: ${UPFP}/" upfcfg.yaml

NGAPLINE=$(awk '/gtpu/{ print NR; exit }' upfcfg.yaml)
NGAPLINE2=$((NGAPLINE+1))
sed -i "${NGAPLINE2}s/addr.*/addr: ${UPFG}/" upfcfg.yaml

NGAPLINE=$(awk '/cidr/{ print NR; exit }' upfcfg.yaml)
sed -i "${NGAPLINE}s/cidr.*/cidr: ${UENET}\/${UEMSK}/" upfcfg.yaml

NGAPLINE=$(awk '/natifname:/{ print NR; exit }' upfcfg.yaml)
sed -i "${NGAPLINE}s/natifname.*/natifname: ${UEIF}/" upfcfg.yaml

echo "> Configurando registration_test.go"
cd ../../../test
#cp registration_test.go registration_test.go.$CURDATE
sed -i "s/AMF string.*/AMF string = \"$AMF\"/" registration_test.go
sed -i "s/GNBN string.*/GNBN string = \"$GNBN\"/" registration_test.go
sed -i "s/UPFG string.*/UPFG string = \"$UPFG\"/" registration_test.go
sed -i "s/SRC string.*/SRC string = \"$PINGSRC\"/" registration_test.go
sed -i "s/DST string.*/DST string = \"$PINGDST\"/" registration_test.go
sed -i "s/GNBG string.*/GNBG string = \"$GNBG\"/" registration_test.go

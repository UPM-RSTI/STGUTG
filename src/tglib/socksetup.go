package tglib

import (
	"fmt"
	"net"
	"syscall"
)

type EthSocketConn struct {
	Iface *net.Interface
	Addr  syscall.SockaddrLinklayer
	Fd    int
}

type IPSocketConn struct {
	Fd int
}

func NewEthSocketConn(ifname string) (EthSocketConn, error) {
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		return EthSocketConn{}, fmt.Errorf("Get link by name: %s", err)
	}

	addr := syscall.SockaddrLinklayer{
		Ifindex: iface.Index,
	}

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0x0300) // 0x0300 = syscall.ETH_P_ALL
	if err != nil {
		return EthSocketConn{}, fmt.Errorf("Create Ethernet raw socket: %s", err)
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		return EthSocketConn{}, fmt.Errorf("Bind Ethernet raw socket: %s", err)
	}

	socketConn := EthSocketConn{
		Iface: iface,
		Addr:  addr,
		Fd:    fd,
	}

	return socketConn, nil
}

func NewIPSocketConn() (IPSocketConn, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		return IPSocketConn{}, fmt.Errorf("Create IP raw socket: %s", err)
	}

	socketConn := IPSocketConn{
		Fd: fd,
	}

	return socketConn, nil
}

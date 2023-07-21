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
		return EthSocketConn{}, fmt.Errorf("Create raw socket: %s", err)
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		return EthSocketConn{}, fmt.Errorf("Bind raw socket: %s", err)
	}

	// Setup a timeout of 3 seconds for listening to the raw socket
	syscall.SetsockoptTimeval(fd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &(syscall.Timeval{Sec: 3, Usec: 0}))

	socketConn := EthSocketConn{}
	socketConn.Iface = iface
	socketConn.Addr = addr
	socketConn.Fd = fd

	return socketConn, nil
}

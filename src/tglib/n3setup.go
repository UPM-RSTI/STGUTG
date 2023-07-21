package tglib

import (
	"fmt"
	"net"
	"syscall"
)

func ConnectToUpf(gnbPort int) (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return 0, fmt.Errorf("Create UPF raw socket: %s", err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		return 0, fmt.Errorf("Config ReuseAddr: %s", err)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", gnbPort))
	if err != nil && udpAddr.IP != nil {
		return 0, fmt.Errorf("Resolve local address: %s", err)
	}

	err = syscall.Bind(fd, &syscall.SockaddrInet4{Port: udpAddr.Port})
	if err != nil && udpAddr.IP != nil {
		return 0, fmt.Errorf("Bind socket: %s", err)
	}

	// Setup a timeout of 3 seconds for listening to the raw socket
	syscall.SetsockoptTimeval(fd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &(syscall.Timeval{Sec: 3, Usec: 0}))

	return fd, nil
}

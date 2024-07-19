package tglib

import (
	"fmt"
	"net"
	"os"

	"github.com/ishidawataru/sctp"
)

const ngapPPID uint32 = 0x3c000000

func getNgapIp(amfIP, ranIP string, amfPort, ranPort int) (amfAddr, ranAddr *sctp.SCTPAddr, err error) {
	ips := []net.IPAddr{}
	if ip, err1 := net.ResolveIPAddr("ip", amfIP); err1 != nil {
		err = fmt.Errorf("error resolving address '%s': %v", amfIP, err1)
		return nil, nil, err
	} else {
		ips = append(ips, *ip)
	}
	amfAddr = &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    amfPort,
	}
	ips = []net.IPAddr{}
	if ip, err1 := net.ResolveIPAddr("ip", ranIP); err1 != nil {
		err = fmt.Errorf("error resolving address '%s': %v", ranIP, err1)
		return nil, nil, err
	} else {
		ips = append(ips, *ip)
	}
	ranAddr = &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    ranPort,
	}
	return amfAddr, ranAddr, nil
}

func ConnectToAmf(amfIP, stgIP string, amfPort, stgPort int) (*sctp.SCTPConn, error) {
	amfAddr, ranAddr, err := getNgapIp(amfIP, stgIP, amfPort, stgPort)
	if err != nil {
		return nil, err
	}
	conn, err := sctp.DialSCTP("sctp", ranAddr, amfAddr)
	if err != nil {
		return nil, err
	}
	info, err := conn.GetDefaultSentParam()
	if err != nil {
		fmt.Printf("conn GetDefaultSentParam error in ConntectToAmf: %+v", err)
		os.Exit(1)
	}
	info.PPID = ngapPPID
	err = conn.SetDefaultSentParam(info)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

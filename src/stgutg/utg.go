package stgutg

// UTG
// Functions that manage the User Traffic Generation capabilites. It includes
// capturing traffic from clients connected to the connector and injecting the
// traffic in the GTP tunnel.
// It also provides a function to capture and forward the traffic sent to the
// client.
// Version: 0.9
// Date: 9/6/21

import (
	"context"
	"sync"
	"tglib"

	"bytes"
	"encoding/hex"
	"syscall"
)

// ListenForResponses
// Function that keeps listening in the network interface connected to the UPF (dst)
// and captures the traffic to the client app. It decapsulates the packet from the GTP
// tunnel, then checks the destination IP and looks up the corresponding MAC address
// in the system ARP table. It then builds the Eth header and sends the packet back to
// the client.
func ListenForResponses(ethSocketConn tglib.EthSocketConn, upfFD int, ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	// TODO: this should be a configuration parameter?
	table := GetARPTable("/proc/net/arp")

	rcvBuf := make([]byte, 1500)

	for {

		select {
		case <-ctx.Done():
			return
		default:
		}

		udpPayloadSize, _, err := syscall.Recvfrom(upfFD, rcvBuf, 0)
		if err != nil {
			// TODO: Timeout introduced in socket generates errors, filter them
			// fmt.Printf("Error reading udp socket: %s\n", err)
			continue
		}

		enc_b := rcvBuf[:udpPayloadSize]

		gtp_hdr_size := 8
		if enc_b[0]&4 != 0 {
			gtp_hdr_size += 4 + int(enc_b[12])*4
		} else if enc_b[0]&3 != 0 {
			gtp_hdr_size += 4
		}
		enc_b = enc_b[gtp_hdr_size:]
		ManageError("Error capturing receiving traffic", err)

		// TODO: This will fail if no ARP table entry has been previously added
		// for the IP in enc_b.
		eth_dst, err := hex.DecodeString(GetMAC(GetIP(enc_b), table))
		ManageError("Error retrieving Dst MAC address", err)

		frame := bytes.Join([][]byte{eth_dst, []byte(ethSocketConn.Iface.HardwareAddr), []byte("\x08\x00"), enc_b}, nil)

		err = syscall.Sendto(ethSocketConn.Fd, frame, 0, &(ethSocketConn.Addr))
		ManageError("Sendto", err)

	}
}

// SendTraffic
// Function that captures traffic in the interface connected to the client or clients that
// generate the user traffic (src), emulating the UEs.
// It checks the source IP address to determine the TEID to use when adding the GTP
// header and then sends the traffic to the UPF.
func SendTraffic(upfFD int, ethSocketConn tglib.EthSocketConn, teidUpfIPs map[[4]byte]TeidUpfIp, ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	data := make([]byte, 1500)

	for {

		select {
		case <-ctx.Done():
			return
		default:
		}

		frameSize, _, err := syscall.Recvfrom(ethSocketConn.Fd, data, 0)
		if err != nil {
			// TODO: Timeout introduced in socket generates errors, filter them
			// fmt.Printf("Error receiving traffic: %s\n", err)
			continue
		}

		if bytes.Equal(data[0:6], []byte(ethSocketConn.Iface.HardwareAddr)) && bytes.Equal(data[12:14], []byte{8, 0}) {

			ethFrame := data[:frameSize]

			//fmt.Println(ethFrame)

			src_ip := ethFrame[14+12 : 14+16]

			//fmt.Println(src_ip)
			teid := teidUpfIPs[([4]byte)(src_ip)].teid

			gtpHdr, err := tglib.BuildGTPv1Header(false, 0, false, 0, false, 0, uint16(len(ethFrame[14:])), teid)
			ManageError("Error capturing and sending traffic", err)

			err = syscall.Sendto(upfFD, append(gtpHdr, ethFrame[14:]...), 0, teidUpfIPs[([4]byte)(src_ip)].upfAddr)
			ManageError("Error capturing and sending traffic", err)

		}
	}
}

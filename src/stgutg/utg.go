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
  "tglib"

  "encoding/hex"
  "net"
  //"fmt"

  "github.com/ghedo/go.pkt/capture/pcap"
  "github.com/ghedo/go.pkt/layers"
  "github.com/ghedo/go.pkt/filter"
  "github.com/ghedo/go.pkt/packet"
)

// ListenForResponses
// Function that keeps listening in the network interface connected to the UPF (dst)
// and captures the traffic to the client app. It decapsulates the packet from the GTP
// tunnel, then checks the destination IP and looks up the corresponding MAC address
// in the system ARP table. It then builds the Eth header and sends the packet back to
// the client.
func ListenForResponses (src *pcap.Handle, dst *pcap.Handle, srcMac string, gnb_gtp string){

	f,err := filter.Compile("not arp and dst "+gnb_gtp,packet.Eth,true)
  ManageError("Error capturing receiving traffic",err)

   // TODO: this should be a configuration parameter?
  table := GetARPTable("/proc/net/arp")

	dst.Activate()
	for {

  		buf,err := dst.Capture()

      ManageError("Error capturing receiving traffic",err)

  		if f.Match(buf) {

  			p,err := layers.UnpackAll(buf,packet.Eth)
        ManageError("Error capturing receiving traffic",err)
  			ip_p := p.Payload()
  			udp_p := ip_p.Payload()
  			tgp_p := udp_p.Payload()
  			enc_b,err := layers.Pack(tgp_p)
        ManageError("Error capturing receiving traffic",err)

        // TODO: This will fail if no ARP table entry has been previously added
        // for the IP in enc_b.
        eth_hdr,err := hex.DecodeString(GetMAC(GetIP(enc_b),table)+srcMac+"0800")
        ManageError("Error capturing receiving traffic",err)

  			src.Inject(append(eth_hdr,enc_b[8:]...))
  		}
	}
}

// SendTraffic
// Function that captures traffic in the interface connected to the client or clients that
// generate the user traffic (src), emulating the UEs.
// It checks the source IP address to determine the TEID to use when adding the GTP
// header and then sends the traffic to the UPF.
func SendTraffic (upfConn *net.UDPConn,src *pcap.Handle, teids []Ipteid) {

  f,_ := filter.Compile("not arp",packet.Eth,true)

  i:=0
  for {
		buf,err := src.Capture()
    ManageError("Error capturing and sending traffic",err)

		if f.Match(buf) {


      p,err := layers.UnpackAll(buf,packet.Eth)
      ManageError("Error capturing and sending traffic",err)
      ip_p := p.Payload()
      enc_ipp, err := layers.Pack(ip_p)
      ManageError("Error capturing and sending traffic",err)
      src_ip := net.IP(enc_ipp[12:16])
      //fmt.Println(src_ip)
      teid := GetTEID(src_ip, teids)

			gtpHdr,err := tglib.BuildGTPv1Header(false, 0, false, 0, false, 0, uint16(len(buf[14:])), teid)
      ManageError("Error capturing and sending traffic",err)

      _, err = upfConn.Write(append(gtpHdr,buf[14:]...))
      ManageError("Error capturing and sending traffic",err)

      i++
		}
	}
}

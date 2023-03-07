package stgutg

// Service
// Functions to request services through a PDU session
// Version: 0.9
// Date: 9/6/21

import (
	"free5gc/lib/nas"
	"free5gc/lib/nas/nasMessage"
	"free5gc/lib/nas/nasTestpacket"
	"free5gc/lib/ngap"
	"tglib"

	"strconv"
	"strings"
	"time"

	"github.com/ishidawataru/sctp"
)

// ServiceRequest
// Function that, given a PDU session, allows for the request of services.
// TODO: Requires improvement (it sometimes fails with free5gc)
func ServiceRequest(pdu []byte, ue *tglib.RanUeContext, conn *sctp.SCTPConn, gnb_gtp string) []byte {
	var recvMsg = make([]byte, 2048)

	ueSupi := strings.Split(ue.Supi, "-")[1]
	supiInt, _ := strconv.Atoi(ueSupi)
	pduId := int64(supiInt % 1e4)

	/*
	  pduSessionIDList := []int64{pduId}
		sendMsg,err := tglib.GetUEContextReleaseRequest(ue.AmfUeNgapId,
	                                                  ue.RanUeNgapId,
	                                                  pduSessionIDList)
	  ManageError("Error in service Request",err)

		_, err = conn.Write(sendMsg)
	  ManageError("Error in service Request",err)

	  n, err := conn.Read(recvMsg)
	  ManageError("Error in service Request",err)

	  _, err = ngap.Decoder(recvMsg[:n])
	  ManageError("Error in service Request",err)

	  sendMsg, err = tglib.GetUEContextReleaseComplete(ue.AmfUeNgapId,
	                                                   ue.RanUeNgapId,
	                                                   nil)
	  ManageError("Error in service Request",err)

		_, err = conn.Write(sendMsg)
	  ManageError("Error in service Request",err)


	  time.Sleep(1 * time.Second)
	*/

	pdu = nasTestpacket.GetServiceRequest(nasMessage.ServiceTypeData)

	pdu, err := tglib.EncodeNasPduWithSecurity(ue,
		pdu,
		nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		true,
		false)
	ManageError("Error in service Request", err)

	sendMsg, err := tglib.GetInitialUEMessage(ue.RanUeNgapId,
		pdu,
		"") //fe0000000001
	ManageError("Error in service Request", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error in service Request", err)

	n, err := conn.Read(recvMsg)
	ManageError("Error in service Request", err)

	_, err = ngap.Decoder(recvMsg[:n])
	ManageError("Error in service Request", err)

	sendMsg, err = tglib.GetInitialContextSetupResponseForServiceRequest(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pduId, // The function from tglib has been changed to include pduId parameter (see Packet.go and Build.go)
		gnb_gtp)

	_, err = conn.Write(sendMsg)
	ManageError("Error in service Request", err)

	time.Sleep(1 * time.Second)

	return pdu

}

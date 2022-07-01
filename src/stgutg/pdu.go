package stgutg

// PDU
// Functions to manage the PDU sessions, including establishment, modification
// and release.
// Version: 0.9
// Date: 9/6/21

import (
  "free5gc/lib/openapi/models"
  "free5gc/lib/nas/nasTestpacket"
  "free5gc/lib/nas/nasMessage"
  "free5gc/lib/nas"
  "free5gc/lib/ngap"

  "tglib"
  "fmt"
  "time"
  "encoding/binary"
  "net"

  "strconv"
  "strings"

  "github.com/ishidawataru/sctp"
)

// EstablishPDU
// Function that establishes a new PDU session for a given UE.
// It requres a previously generated UE and an active SCTP connection with an AMF.
// It returns a tuple of assigned IP for the UE and the corresponding TEID.
func EstablishPDU(sst int32, sd string, pdu []byte, ue *tglib.RanUeContext, conn *sctp.SCTPConn, gnb_gtp string) Ipteid {

  var recvMsg = make([]byte, 2048)
  sNssai := models.Snssai{
		Sst: sst,
		Sd: sd,
	}

  var it Ipteid

  ueSupi := strings.Split(ue.Supi, "-") [1]
	supiInt, _ := strconv.Atoi(ueSupi)
	pduId := int64(supiInt % 1e4) //TODO: Check if it works with large SUPI numbers

	pdu = nasTestpacket.GetUlNasTransport_PduSessionEstablishmentRequest(uint8(pduId),
                                                                       nasMessage.ULNASTransportRequestTypeInitialRequest,
                                                                       "internet",
                                                                       &sNssai)

	pdu, err := tglib.EncodeNasPduWithSecurity(ue,
                                             pdu,
                                             nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
                                             true,
                                             false)
  ManageError("Error establishing PDU",err)


	sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
                                              ue.RanUeNgapId,
                                              pdu)
  ManageError("Error establishing PDU",err)
	_, err = conn.Write(sendMsg)
  ManageError("Error establishing PDU",err)

  n, err := conn.Read(recvMsg)
  ManageError("Error establishing PDU",err)

	msg, err := ngap.Decoder(recvMsg[:n])
  ManageError("Error establishing PDU",err)


  // Recover assigned IP and TEID for the session.
  // TODO: This is an ad-hoc solution for free5gc. Check if it works for other cores.
  bteid := msg.InitiatingMessage.Value.PDUSessionResourceSetupRequest.ProtocolIEs.List[2].Value.PDUSessionResourceSetupListSUReq.List[0].PDUSessionResourceSetupRequestTransfer[13:17]
  bip := msg.InitiatingMessage.Value.PDUSessionResourceSetupRequest.ProtocolIEs.List[2].Value.PDUSessionResourceSetupListSUReq.List[0].PDUSessionNASPDU.Value[41:45]
  teid := binary.BigEndian.Uint32(bteid)
  ip := net.IP(bip)
  it.ueip = ip
  it.teid = teid

  sendMsg, err = tglib.GetPDUSessionResourceSetupResponse(ue.AmfUeNgapId,
                                                          ue.RanUeNgapId,
                                                          pduId, // The function from tglib has been changed to include pduId parameter (see Packet.go and Build.go)
                                                          gnb_gtp)
  ManageError("Error establishing PDU",err)

  _,err = conn.Write(sendMsg)
  ManageError("Error establishing PDU",err)


  return it
}

// ReleasePDU
// Function that removes a previously established PDU session.
// It uses the SST and SD values to determine the session to delete.
func ReleasePDU (sst int32, sd string, ue *tglib.RanUeContext, conn *sctp.SCTPConn) []byte{
  sNssai := models.Snssai{
    Sst: sst,
    Sd: sd,
  }

  ueSupi := strings.Split(ue.Supi, "-") [1]
  supiInt, err := strconv.Atoi(ueSupi)
  ManageError("Error releasing PDU",err)

  pduId := int64(supiInt % 1e4) //TODO: Check if it works with large SUPI numbers
  pdu := nasTestpacket.GetUlNasTransport_PduSessionReleaseRequest(uint8(pduId))
  pdu, err = tglib.EncodeNasPduWithSecurity(ue,
                                            pdu,
                                            nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
                                            true,
                                            false)
  ManageError("Error releasing PDU",err)

  sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
                                              ue.RanUeNgapId,
                                              pdu)
  ManageError("Error releasing PDU",err)

  _, err = conn.Write(sendMsg)
  ManageError("Error releasing PDU",err)

  time.Sleep(100 * time.Millisecond)

  sendMsg,err = tglib.GetPDUSessionResourceReleaseResponse(ue.AmfUeNgapId,
                                                           ue.RanUeNgapId,
                                                           pduId)
  ManageError("Error releasing PDU",err)

  _, err = conn.Write(sendMsg)
  ManageError("Error releasing PDU",err)

  time.Sleep(10 * time.Millisecond)

  pdu = nasTestpacket.GetUlNasTransport_PduSessionReleaseComplete(uint8(pduId),
                                                                  nasMessage.ULNASTransportRequestTypeInitialRequest,
                                                                  "internet",
                                                                  &sNssai)

  pdu, err = tglib.EncodeNasPduWithSecurity(ue,
                                            pdu,
                                            nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
                                            true,
                                            false)
  ManageError("Error releasing PDU",err)

  sendMsg, err = tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
                                             ue.RanUeNgapId,
                                             pdu)
  ManageError("Error releasing PDU",err)

  _,err = conn.Write(sendMsg)
  ManageError("Error releasing PDU",err)

  time.Sleep(1 * time.Second)

  return pdu
}

// ModifyPDU
// Function that modifies existing PDU session.
// It uses the SST and SD values to determine the session to modify.
// TODO: Requires improvement (currently fails with free5gc)
func ModifyPDU (sst int32, sd string, ue *tglib.RanUeContext, conn *sctp.SCTPConn) []byte{
  var recvMsg = make([]byte, 2048)
  sNssai := models.Snssai{
    Sst: sst,
    Sd: sd,
  }

  ueSupi := strings.Split(ue.Supi, "-") [1]
  supiInt, err := strconv.Atoi(ueSupi)
  ManageError("Error modifying PDU",err)

  pduId := int64(supiInt % 1e4) //TODO: Check if it works with large SUPI numbers

  pdu := nasTestpacket.GetUlNasTransport_PduSessionModificationRequest(uint8(pduId),
                                                                       nasMessage.ULNASTransportRequestTypeExistingPduSession,
                                                                       "internet",
                                                                       &sNssai)

  pdu, err = tglib.EncodeNasPduWithSecurity(ue,
                                            pdu,
                                            nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
                                            true,
                                            false)
  ManageError("Error modifying PDU",err)

  sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
                                              ue.RanUeNgapId,
                                              pdu)
  ManageError("Error modifying PDU",err)

  _,err = conn.Write(sendMsg)
  ManageError("Error modifying PDU",err)

  n, err := conn.Read(recvMsg)
  fmt.Println("Modified PDU",n)
  ManageError("Error establishing PDU",err)

  return pdu
}

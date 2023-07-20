package stgutg

// PDU
// Functions to manage the PDU sessions, including establishment, modification
// and release.
// Version: 0.9
// Date: 9/6/21

import (
	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasTestpacket"
	"github.com/free5gc/ngap"
	"github.com/free5gc/openapi/models"

	"encoding/binary"
	"fmt"
	"net"
	"tglib"
	"time"

	"strconv"
	"strings"

	"github.com/ishidawataru/sctp"
)

// 3GPP TS 24.501 8.3.2.1
// If the length of the element is not fixed it is represented with a negative number:
// -1 and -2 for elements whose length is encoded with 1 and 2 bytes respectively
var PDUSessionEstablishmentAcceptOptionalElementsLength = map[byte]int{
	0x59: 2,
	0x29: -1,
	0x56: 2,
	0x22: -1,
	0x75: -2,
	0x78: -2,
	0x79: -2,
	0x7B: -2,
	0x25: -1,
	0x17: -1,
	0x18: 4,
	0x77: -2,
	0x66: -1,
	0x1F: 3,
}

// 3GPP TS 24.501 8.3.2.1
// These optional element have a total length of 1 byte and their ID's
// are coded with octets 4 to 7 of that byte
var PDUSessionEstablishmentAcceptOptionalElementsHalfByte = []byte{
	0x80,
	0xC0,
}

// EstablishPDU
// Function that establishes a new PDU session for a given UE.
// It requres a previously generated UE and an active SCTP connection with an AMF.
// It returns a tuple of assigned IP for the UE and the corresponding TEID.
func EstablishPDU(sst int32, sd string, pdu []byte, ue *tglib.RanUeContext, conn *sctp.SCTPConn, gnb_gtp string, teidUpfIPs map[[4]byte]TeidUpfIp) {

	var recvMsg = make([]byte, 2048)
	sNssai := models.Snssai{
		Sst: sst,
		Sd:  sd,
	}

	ueSupi := strings.Split(ue.Supi, "-")[1]
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
	ManageError("Error establishing PDU", err)

	sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error establishing PDU", err)
	_, err = conn.Write(sendMsg)
	ManageError("Error establishing PDU", err)

	n, err := conn.Read(recvMsg)
	ManageError("Error establishing PDU", err)

	msg, err := ngap.Decoder(recvMsg[:n])
	ManageError("Error establishing PDU", err)

	// Recover assigned IP and TEID for the session.
	// Only works if 5G-EEA0 is used as cypher

	PDUSessionResourceSetupItemSUReq := msg.InitiatingMessage.Value.PDUSessionResourceSetupRequest.ProtocolIEs.List[2].Value.PDUSessionResourceSetupListSUReq.List[0]

	bip := DecodePDUSessionNASPDU(PDUSessionResourceSetupItemSUReq.PDUSessionNASPDU.Value)
	bteid, bupfip := DecodePDUSessionResourceSetupRequestTransfer(PDUSessionResourceSetupItemSUReq.PDUSessionResourceSetupRequestTransfer)

	teid := binary.BigEndian.Uint32(bteid)
	upfip := net.IP(bupfip)

	teidUpfIPs[bip] = TeidUpfIp{teid, upfip}

	sendMsg, err = tglib.GetPDUSessionResourceSetupResponse(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pduId, // The function from tglib has been changed to include pduId parameter (see Packet.go and Build.go)
		gnb_gtp)
	ManageError("Error establishing PDU", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error establishing PDU", err)
}

// ReleasePDU
// Function that removes a previously established PDU session.
// It uses the SST and SD values to determine the session to delete.
func ReleasePDU(sst int32, sd string, ue *tglib.RanUeContext, conn *sctp.SCTPConn) []byte {
	sNssai := models.Snssai{
		Sst: sst,
		Sd:  sd,
	}

	ueSupi := strings.Split(ue.Supi, "-")[1]
	supiInt, err := strconv.Atoi(ueSupi)
	ManageError("Error releasing PDU", err)

	pduId := int64(supiInt % 1e4) //TODO: Check if it works with large SUPI numbers
	pdu := nasTestpacket.GetUlNasTransport_PduSessionReleaseRequest(uint8(pduId))
	pdu, err = tglib.EncodeNasPduWithSecurity(ue,
		pdu,
		nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		true,
		false)
	ManageError("Error releasing PDU", err)

	sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error releasing PDU", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error releasing PDU", err)

	time.Sleep(100 * time.Millisecond)

	sendMsg, err = tglib.GetPDUSessionResourceReleaseResponse(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pduId)
	ManageError("Error releasing PDU", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error releasing PDU", err)

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
	ManageError("Error releasing PDU", err)

	sendMsg, err = tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error releasing PDU", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error releasing PDU", err)

	time.Sleep(1 * time.Second)

	return pdu
}

// ModifyPDU
// Function that modifies existing PDU session.
// It uses the SST and SD values to determine the session to modify.
// TODO: Requires improvement (currently fails with free5gc)
func ModifyPDU(sst int32, sd string, ue *tglib.RanUeContext, conn *sctp.SCTPConn) []byte {
	var recvMsg = make([]byte, 2048)
	sNssai := models.Snssai{
		Sst: sst,
		Sd:  sd,
	}

	ueSupi := strings.Split(ue.Supi, "-")[1]
	supiInt, err := strconv.Atoi(ueSupi)
	ManageError("Error modifying PDU", err)

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
	ManageError("Error modifying PDU", err)

	sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error modifying PDU", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error modifying PDU", err)

	n, err := conn.Read(recvMsg)
	fmt.Println("Modified PDU", n)
	ManageError("Error establishing PDU", err)

	return pdu
}

// DecodePDUSessionResourceSetupRequestTransfer
// Function that extracts UPF IP address and TEID from a given PDUSessionResourceSetupRequestTransfer
func DecodePDUSessionResourceSetupRequestTransfer(PDUSessionResourceSetupRequestTransfer []byte) ([]byte, []byte) {
	var bteid []byte = nil
	var bupfip []byte = nil

	offset := 3 //  We skip number of protocolIEs as we are only interested in the first or second one

	for offset < len(PDUSessionResourceSetupRequestTransfer) {

		if int(binary.BigEndian.Uint16(PDUSessionResourceSetupRequestTransfer[offset:offset+2])) != 139 {
			offset += 3 + int(PDUSessionResourceSetupRequestTransfer[offset+3]) + 1
		} else {
			offset += 3

			UPTrasportLayerInfoLength := int(PDUSessionResourceSetupRequestTransfer[offset])
			offset += 1

			UPTransportLayerInfo := PDUSessionResourceSetupRequestTransfer[offset : offset+UPTrasportLayerInfoLength]

			bteid = UPTransportLayerInfo[UPTrasportLayerInfoLength-4:]
			bupfip = UPTransportLayerInfo[UPTrasportLayerInfoLength-8 : UPTrasportLayerInfoLength-4]

			break
		}
	}

	return bteid, bupfip
}

// DecodePDUSessionNASPDU
// Function that extracts UE IP address from a given PDUSessionNASPDU message
func DecodePDUSessionNASPDU(PDUSessionNASPDU []byte) [4]byte {
	var bip [4]byte

	plainNAS5GSMessage := PDUSessionNASPDU[7:]

	payloadContainerLength := binary.BigEndian.Uint16(plainNAS5GSMessage[4:6])
	payloadContainerPlainNAS5GSMessage := plainNAS5GSMessage[6 : 6+payloadContainerLength]

	QoSRulesLength := binary.BigEndian.Uint16(payloadContainerPlainNAS5GSMessage[5:7])

	opElements := payloadContainerPlainNAS5GSMessage[5+2+QoSRulesLength+7:]
	length := len(opElements)
	index := 0
	var opElementID byte
	var opElementLength int
	_, _ = opElementID, opElementLength // Needed to fix "declared but not used" error

outerloop:
	for index < length {
		opElementID = opElements[index]

		if opElementID == 0x29 { // PDU Address
			bip = ([4]byte)(opElements[index+3 : index+7])
			index += 7
			break outerloop
		}

		for _, id := range PDUSessionEstablishmentAcceptOptionalElementsHalfByte {
			if opElementID&0xF0 == id {
				index += 1
				continue outerloop
			}
		}

		opElementLength = PDUSessionEstablishmentAcceptOptionalElementsLength[opElementID]

		if opElementLength > 0 {
			index += opElementLength
		} else if opElementLength == -1 { // 1 byte long length indicator
			opElementLength = int(opElements[index+1])
			index += 1 + 1 + opElementLength
		} else if opElementLength == -2 { // 2 bytes long length indicator
			opElementLength = int(binary.BigEndian.Uint16(opElements[index+1 : index+1+2]))
			index += 1 + 2 + opElementLength
		}
	}

	return bip
}

func printHex(bytearray []byte) {
	for _, byteelement := range bytearray {
		fmt.Printf("%02X ", byteelement)
	}
	fmt.Print("\n\n")
}

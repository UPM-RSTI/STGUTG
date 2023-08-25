package stgutg

// UE
// Functions to manage UEs to register in the core. Incldes the creation of new
// UEs, its registration and deregistration.
// Version: 0.9
// Date: 9/6/21

import (
	"free5gclib/nas"
	"free5gclib/nas/nasMessage"
	"free5gclib/nas/nasTestpacket"
	"free5gclib/nas/security"
	"free5gclib/ngap"
	"free5gclib/ngap/ngapType"

	"strconv"
	"strings"
	"time"

	"tglib"

	"github.com/ishidawataru/sctp"
)

// CreateUE
// Function that cretes an new UE structure containing an specific SUPI and
// currently, common security parameters (from the configuration file)
func CreateUE(imsi int, K string, OPC string, OP string) *tglib.RanUeContext {

	ranUeNgapId := imsi % 1e4
	supiStr := strconv.Itoa(imsi)
	supi := "imsi-" + supiStr

	ue := tglib.NewRanUeContext(supi,
		int64(ranUeNgapId),
		security.AlgCiphering128NEA0,
		security.AlgIntegrity128NIA2)
	ue.AuthenticationSubs = tglib.GetAuthSubscription(K, OPC, OP)

	return ue
}

// RegisterUE
// Function that, given a UE context, allows its registration in the core.
// The UE information should have been previously stored in the core database
// (subscriptor information)
func RegisterUE(ue *tglib.RanUeContext, mnc string, conn *sctp.SCTPConn) (*tglib.RanUeContext, []byte, *ngapType.NGAPPDU) {
	var recvMsg = make([]byte, 2048)

	mobileIdentity5GS := EncodeSuci([]byte(strings.TrimPrefix(ue.Supi, "imsi-")), len(mnc))

	ueSecurityCapability := ue.GetUESecurityCapability()

	registrationRequest := nasTestpacket.GetRegistrationRequest(nasMessage.RegistrationType5GSInitialRegistration,
		*mobileIdentity5GS,
		nil,
		ueSecurityCapability,
		nil,
		nil,
		nil)

	sendMsg, err := tglib.GetInitialUEMessage(ue.RanUeNgapId, registrationRequest, "")
	ManageError("Error in registering new UE", err)
	_, err = conn.Write(sendMsg)
	ManageError("Error in registering new UE", err)

	n, err := conn.Read(recvMsg)
	ManageError("Error in registering new UE", err)
	ngapMsg, err := ngap.Decoder(recvMsg[:n])
	ManageError("Error in registering new UE", err)

	nasPdu := tglib.GetNasPdu(ue,
		ngapMsg.InitiatingMessage.Value.DownlinkNASTransport)
	rand := nasPdu.AuthenticationRequest.GetRANDValue()
	resStat := ue.DeriveRESstarAndSetKey(ue.AuthenticationSubs,
		rand[:],
		"5G:mnc093.mcc208.3gppnetwork.org")

	ue.AmfUeNgapId = ngapMsg.InitiatingMessage.Value.DownlinkNASTransport.ProtocolIEs.List[0].Value.AMFUENGAPID.Value

	pdu := nasTestpacket.GetAuthenticationResponse(resStat, "")

	sendMsg, err = tglib.GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	_, err = conn.Write(sendMsg)
	ManageError("Error in registering new UE", err)

	n, err = conn.Read(recvMsg)
	ManageError("Error in registering new UE", err)
	_, err = ngap.Decoder(recvMsg[:n])
	ManageError("Error in registering new UE", err)

	registrationRequestWith5GMM := nasTestpacket.GetRegistrationRequest(nasMessage.RegistrationType5GSInitialRegistration,
		*mobileIdentity5GS,
		nil,
		ueSecurityCapability,
		ue.Get5GMMCapability(),
		nil,
		nil)

	pdu = nasTestpacket.GetSecurityModeComplete(registrationRequestWith5GMM)
	pdu, err = tglib.EncodeNasPduWithSecurity(ue,
		pdu,
		nas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext,
		true,
		true)
	ManageError("Error in registering new UE", err)

	sendMsg, err = tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error in registering new UE", err)
	_, err = conn.Write(sendMsg)
	ManageError("Error in registering new UE", err)

	n, err = conn.Read(recvMsg)
	ManageError("Error in registering new UE", err)

	ngapPdu, err := ngap.Decoder(recvMsg[:n])
	ManageError("Error in registering new UE", err)

	sendMsg, err = tglib.GetInitialContextSetupResponse(ue.AmfUeNgapId, ue.RanUeNgapId)
	ManageError("Error in registering new UE", err)
	_, err = conn.Write(sendMsg)
	ManageError("Error in registering new UE", err)

	pdu = nasTestpacket.GetRegistrationComplete(nil)

	pdu, err = tglib.EncodeNasPduWithSecurity(ue,
		pdu,
		nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		true,
		false)
	ManageError("Error in registering new UE", err)

	sendMsg, err = tglib.GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	ManageError("Error in registering new UE", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error in registering new UE", err)

	return ue, pdu, ngapPdu
}

// DeregisterUE
// Function that deregisters a given UE from the core.
func DeregisterUE(ue *tglib.RanUeContext, mnc string, conn *sctp.SCTPConn) {
	var recvMsg = make([]byte, 2048)

	mobileIdentity5GS := EncodeSuci([]byte(strings.TrimPrefix(ue.Supi, "imsi-")), len(mnc))

	pdu := nasTestpacket.GetDeregistrationRequest(nasMessage.AccessType3GPP,
		0,
		0x04,
		*mobileIdentity5GS)

	pdu, err := tglib.EncodeNasPduWithSecurity(ue,
		pdu,
		nas.SecurityHeaderTypeIntegrityProtectedAndCiphered,
		true,
		false)
	ManageError("Error deregistering UE", err)

	sendMsg, err := tglib.GetUplinkNASTransport(ue.AmfUeNgapId,
		ue.RanUeNgapId,
		pdu)
	ManageError("Error deregistering UE", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error deregistering UE", err)

	time.Sleep(500 * time.Millisecond)

	n, err := conn.Read(recvMsg)
	ManageError("Error deregistering UE", err)

	_, err = ngap.Decoder(recvMsg[:n]) //ngapPdu
	ManageError("Error deregistering UE", err)

	n, err = conn.Read(recvMsg)
	ManageError("Error deregistering UE", err)

	_, err = ngap.Decoder(recvMsg[:n]) //ngapPdu
	ManageError("Error deregistering UE", err)

	sendMsg, err = tglib.GetUEContextReleaseComplete(ue.AmfUeNgapId, ue.RanUeNgapId, nil)
	ManageError("Error deregistering UE", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error deregistering UE", err)
}

package stgutg

// NGSetup
// Functions to manage the inital NGAP setup
// Version: 0.9
// Date: 9/6/21

import (
	"free5gclib/ngap"

	"tglib"

	"github.com/ishidawataru/sctp"
)

// ManageNGSetup
// Generates and sends the initial setup request to the configured AMF.
// Receives the response from the AMF and decodes it.
func ManageNGSetup(conn *sctp.SCTPConn, gnbId string, bitlength uint64, name string) {
	var recvMsg = make([]byte, 2048)

	sendMsg, err := tglib.GetNGSetupRequest([]byte(gnbId), bitlength, name)
	ManageError("Error in NG Setup", err)

	_, err = conn.Write(sendMsg)
	ManageError("Error in NG Setup", err)

	n, err := conn.Read(recvMsg)
	ManageError("Error in NG Setup", err)

	_, err = ngap.Decoder(recvMsg[:n]) //ngapPdu?
	ManageError("Error in NG Setup", err)
}

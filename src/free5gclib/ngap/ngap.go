package ngap

import (
	"free5gclib/aper"
	"free5gclib/ngap/ngapType"
)

// Decoder is to decode raw data to NGAP pdu pointer with PER Aligned
func Decoder(b []byte) (pdu *ngapType.NGAPPDU, err error) {
	pdu = &ngapType.NGAPPDU{}

	err = aper.UnmarshalWithParams(b, pdu, "valueExt,valueLB:0,valueUB:2")
	return
}

// Encoder is to NGAP pdu to raw data with PER Aligned
func Encoder(pdu ngapType.NGAPPDU) ([]byte, error) {
	return aper.MarshalWithParams(pdu, "valueExt,valueLB:0,valueUB:2")
}

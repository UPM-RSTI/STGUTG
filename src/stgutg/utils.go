package stgutg

// Utils
// Auxiliar utility functions used by the rest of programs in this package.
// Version: 0.9
// Date: 9/6/21

import (
	"free5gclib/nas/nasMessage"
	"free5gclib/nas/nasType"

	"gopkg.in/yaml.v2"

	"fmt"
	"os"
)

type Conf struct {
	Configuration struct {
		AmfNgapIP                 string `yaml:"amf_ngap_ip"`
		AmfNgapPort               int    `yaml:"amf_ngap_port"`
		Gnb_gtp                   string `yaml:"gnb_gtp"`
		Gnbg_port                 int    `yaml:"gnbg_port"`
		StgNgapIP                 string `yaml:"stg_ngap_ip"`
		StgNgapPort               int    `yaml:"stg_ngap_port"`
		Upf_port                  int    `yaml:"upf_port"`
		Gnb_id                    string `yaml:"gnb_id"`
		Gnb_bitlength             uint64 `yaml:"gnb_bitlength"`
		Gnb_name                  string `yaml:"gnb_name"`
		Initial_imsi              int    `yaml:"initial_imsi"`
		Mnc                       string `yaml:"mnc"`
		K                         string `yaml:"k"`
		OPC                       string `yaml:"opc"`
		OP                        string `yaml:"op"`
		SST                       int32  `yaml:"sst"`
		SD                        string `yaml:"sd"`
		DLIface                   string `yaml:"downlink_iface"`
		ULIface                   string `yaml:"uplink_iface"`
		UeNumber                  int    `yaml:"ue_number"`
		Test_ue_registation       int    `yaml:"ue_registration"`
		Test_ue_pdu_establishment int    `yaml:"ue_pdu"`
		Test_ue_service           int    `yaml:"ue_service"`
		Test_ue_pdu_release       int    `yaml:"ue_pdu_release"`
		Test_ue_deregistration    int    `yaml:"ue_deregistration"`
	}
}

func (c *Conf) GetConfiguration() Conf {

	yamlFile, _ := os.ReadFile("config.yaml")
	yaml.Unmarshal(yamlFile, c)

	return *c
}

func GetMode(args []string) int {

	if len(args) == 1 {
		return 1
	} else if len(args) == 2 {
		testMode := os.Args[1]
		if testMode == "-t" {
			return 2
		} else {
			fmt.Println("Usage: stg-utg [-t]")
		}
	} else {
		fmt.Println("Usage: stg-utg [-t]")
	}

	return 0
}

// hexCharToByte
// Function that transforms a hex character into a byte
func hexCharToByte(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}

// EncodeSuci
// Function that generates a SUCI from an IMSI.
func EncodeSuci(imsi []byte, mncLen int) *nasType.MobileIdentity5GS {
	var msin []byte
	suci := nasType.MobileIdentity5GS{
		Buffer: []uint8{nasMessage.SupiFormatImsi<<4 |
			nasMessage.MobileIdentity5GSTypeSuci, 0x0, 0x0, 0x0, 0xf0, 0xff, 0x00, 0x00},
	}

	//mcc & mnc
	suci.Buffer[1] = hexCharToByte(imsi[1])<<4 | hexCharToByte(imsi[0])
	if mncLen > 2 {
		suci.Buffer[2] = hexCharToByte(imsi[3])<<4 | hexCharToByte(imsi[2])
		suci.Buffer[3] = hexCharToByte(imsi[5])<<4 | hexCharToByte(imsi[4])
		msin = imsi[6:]
	} else {
		suci.Buffer[2] = 0xf<<4 | hexCharToByte(imsi[2])
		suci.Buffer[3] = hexCharToByte(imsi[4])<<4 | hexCharToByte(imsi[3])
		msin = imsi[5:]
	}

	for i := 0; i < len(msin); i += 2 {
		suci.Buffer = append(suci.Buffer, 0x0)
		j := len(suci.Buffer) - 1
		if i+1 == len(msin) {
			suci.Buffer[j] = 0xf<<4 | hexCharToByte(msin[i])
		} else {
			suci.Buffer[j] = hexCharToByte(msin[i+1])<<4 | hexCharToByte(msin[i])
		}
	}
	suci.Len = uint16(len(suci.Buffer))
	return &suci
}

// ManageError
// Generic function that prints an error and exits the program.
func ManageError(message string, err error) {
	if err != nil {
		fmt.Println(message+":", err)
		os.Exit(1)
	}
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

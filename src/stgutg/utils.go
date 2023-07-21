package stgutg

// Utils
// Auxiliar utility functions used by the rest of programs in this package.
// Version: 0.9
// Date: 9/6/21

import (
	"free5gc/lib/nas/nasMessage"
	"free5gc/lib/nas/nasType"
	"syscall"

	"gopkg.in/yaml.v2"

	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Conf struct {
	Configuration struct {
		Amf_ngap                  string `yaml:"amf_ngap"`
		Amf_port                  int    `yaml:"amf_port"`
		Gnb_gtp                   string `yaml:"gnb_gtp"`
		Gnbg_port                 int    `yaml:"gnbg_port"`
		Gnb_ngap                  string `yaml:"gnb_ngap"`
		Gnbn_port                 int    `yaml:"gnbn_port"`
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
		SrcIface                  string `yaml:"src_iface"`
		UeNumber                  int    `yaml:"ue_number"`
		Test_ue_registation       int    `yaml:"ue_registration"`
		Test_ue_pdu_establishment int    `yaml:"ue_pdu"`
		Test_ue_service           int    `yaml:"ue_service"`
		Test_ue_pdu_release       int    `yaml:"ue_pdu_release"`
		Test_ue_deregistration    int    `yaml:"ue_deregistration"`
	}
}

// TeidUpfIp
// Structure to store tuples of IP addresses and TEIDs.
type TeidUpfIp struct {
	teid    uint32
	upfAddr *syscall.SockaddrInet4
}

func (c *Conf) GetConfiguration() Conf {

	yamlFile, _ := ioutil.ReadFile("config.yaml")
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
		fmt.Println(message, ":", err)
		os.Exit(1)
	}
}

// GetARPTable
// Function that reads a file containing the system's ARP table and returns
// a variable with its contents.
func GetARPTable(arptfile string) [][]string {

	var arptable [][]string

	arpfile, _ := os.Open(arptfile)

	arpscanner := bufio.NewScanner(arpfile)
	arpscanner.Split(bufio.ScanLines)

	arpscanner.Scan()

	for arpscanner.Scan() {

		arpline := strings.Fields(arpscanner.Text())
		arptable = append(arptable, arpline)
	}

	return arptable
}

// GetMAC
// Function that reads a variable containing an ARP table and looks up for
// the MAC corresponding to a given IP address.
func GetMAC(ip string, table [][]string) string {
	mac := "0"
	for _, line := range table {
		if ip == line[0] {
			mac = strings.ReplaceAll(line[3], ":", "")
		}
	}
	return mac
}

// GetIP
// Function that extracts an IP address from a buffer of bytes. Used to
// get the IP address of the encapsulated packet received from the GTP tunnel.
func GetIP(ip_buffer []byte) string {

	ip := fmt.Sprintf("%d", ip_buffer[16])
	for i := 17; i < 20; i++ {
		ip = ip + "." + fmt.Sprintf("%d", ip_buffer[i])
	}
	return ip
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

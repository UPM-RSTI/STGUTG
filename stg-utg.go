package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"stgutg"
	"tglib"

	"github.com/Rotchamar/xdp_gtp/xdpgtp"
	"github.com/cilium/ebpf/link"
)

type clientInfo struct {
	clientIP net.IP
	teid     uint32
	upfIP    net.IP
}

func main() {

	var ueList []*tglib.RanUeContext
	var pduList [][]byte

	var c stgutg.Conf
	c.GetConfiguration()

	mode := stgutg.GetMode(os.Args)

	if mode == 1 {
		fmt.Println("TRAFFIC MODE")
		fmt.Println("----------------------")

		fmt.Println(">> Setting up data plane interfaces")
		clientIface, err := net.InterfaceByName(c.Configuration.DLIface)
		if err != nil {
			stgutg.ManageError("Error obtaining client-facing address information", err)
		}

		upfIface, err := net.InterfaceByName(c.Configuration.ULIface)
		if err != nil {
			stgutg.ManageError("Error obtaining UPF-facing address information", err)
		}

		xgtp, err := xdpgtp.NewXDPGTP(link.XDPGenericMode)
		if err != nil {
			stgutg.ManageError("Error creating XDPGTP object", err)
		}
		defer xgtp.Close()

		err = xgtp.AttachClientFacingProgramToInterface(clientIface.Index)
		if err != nil {
			stgutg.ManageError("Error attaching client-facing program", err)
		}
		defer xgtp.DetachProgramFromInterface(clientIface.Index)

		err = xgtp.AttachUpfFacingProgramToInterface(upfIface.Index)
		if err != nil {
			stgutg.ManageError("Error attaching UPF-facing program", err)
		}
		defer xgtp.DetachProgramFromInterface(upfIface.Index)

		fmt.Println(">> Connecting to AMF")
		conn, err := tglib.ConnectToAmf(c.Configuration.AmfNgapIP,
			c.Configuration.StgNgapIP,
			c.Configuration.AmfNgapPort,
			c.Configuration.StgNgapPort)
		stgutg.ManageError("Error in connection to AMF", err)

		imsi := c.Configuration.Initial_imsi

		fmt.Println(">> Managing NG Setup")
		stgutg.ManageNGSetup(conn,
			c.Configuration.Gnb_id,
			imsi,
			c.Configuration.Mnc,
			c.Configuration.Gnb_bitlength,
			c.Configuration.Gnb_name)

		for i := 0; i < c.Configuration.UeNumber; i++ {

			fmt.Println(">> Creating new UE with IMSI:", imsi)
			ue := stgutg.CreateUE(imsi,
				i,
				c.Configuration.K,
				c.Configuration.OPC,
				c.Configuration.OP)

			fmt.Println(">> Registering UE with IMSI:", imsi)
			ue, pdu, _ := stgutg.RegisterUE(ue,
				c.Configuration.Mnc,
				c.Configuration.Mcc,
				conn)

			ueList = append(ueList, ue)
			pduList = append(pduList, pdu)

			time.Sleep(1 * time.Second)
		}

		for i := range pduList {
			fmt.Println(">> Establishing PDU session for", ueList[i].Supi)

			clientip, teid, upfip := stgutg.EstablishPDU(
				c.Configuration.SST,
				c.Configuration.SD,
				ueList[i],
				conn,
				c.Configuration.Gnb_gtp,
			)

			if !xgtp.UpfIsRegistered(upfip) {
				xgtp.AddUpf(upfip)
			}

			if !xgtp.ClientIsRegistered(clientip) {
				xgtp.AddClient(clientip, teid, upfip)
			}

			time.Sleep(1 * time.Second)
		}

		clientsIPs, teids, upfsIPs := xgtp.GetClients()

		registeredClients := make([]clientInfo, len(clientsIPs))

		for idx := range clientsIPs {
			registeredClients[idx] = clientInfo{clientsIPs[idx], teids[idx], upfsIPs[idx]}
		}

		fmt.Println(registeredClients)

		var stopProgram = make(chan os.Signal)

		// Program interrupted
		sig := <-stopProgram
		fmt.Println("\n>> Exiting program:", sig, "found")

		for _, ue := range ueList {
			fmt.Println(">> Releasing PDU session for", ue.Supi)
			stgutg.ReleasePDU(c.Configuration.SST,
				c.Configuration.SD,
				ue,
				conn)
			time.Sleep(1 * time.Second)
		}

		for _, ue := range ueList {
			fmt.Println(">> Deregistering UE", ue.Supi)
			stgutg.DeregisterUE(ue,
				c.Configuration.Mnc,
				conn)
			time.Sleep(2 * time.Second)
		}

		time.Sleep(1 * time.Second)
		conn.Close()

		os.Exit(0)

	} else if mode == 2 {
		fmt.Println("TEST MODE")
		fmt.Println("----------------------")

		pdu_establishment_number := stgutg.Min(c.Configuration.Test_ue_registation,
			c.Configuration.Test_ue_pdu_establishment)

		service_request_number := stgutg.Min(pdu_establishment_number,
			c.Configuration.Test_ue_service)

		pdu_release_number := stgutg.Min(pdu_establishment_number,
			c.Configuration.Test_ue_pdu_release)

		ue_deregistration_number := stgutg.Min(c.Configuration.Test_ue_registation,
			c.Configuration.Test_ue_deregistration)

		fmt.Println(">> Configured tests:")
		fmt.Println("> Registering UEs:", c.Configuration.Test_ue_registation)
		fmt.Println("> PDU sessions to establish:", pdu_establishment_number)
		fmt.Println("> Services to request: ", service_request_number)
		fmt.Println("> PDU sessions to release:", pdu_release_number)
		fmt.Println("> Deregistering UEs:", ue_deregistration_number)
		fmt.Println("----------------------")

		fmt.Println(">> Connecting to AMF")
		conn, err := tglib.ConnectToAmf(c.Configuration.AmfNgapIP,
			c.Configuration.StgNgapIP,
			c.Configuration.AmfNgapPort,
			c.Configuration.StgNgapPort)
		stgutg.ManageError("Error in connection to AMF", err)

		imsi := c.Configuration.Initial_imsi

		fmt.Println(">> Managing NG Setup")
		stgutg.ManageNGSetup(conn,
			c.Configuration.Gnb_id,
			imsi,
			c.Configuration.Mnc,
			c.Configuration.Gnb_bitlength,
			c.Configuration.Gnb_name)

		for i := 0; i < c.Configuration.Test_ue_registation; i++ {
			fmt.Println(">> [ UE REGISTRATION TEST", i+1, "]")

			fmt.Println(">> Creating new UE with IMSI:", imsi)
			ue := stgutg.CreateUE(imsi,
				i,
				c.Configuration.K,
				c.Configuration.OPC,
				c.Configuration.OP)

			fmt.Println(">> Registering UE with IMSI:", imsi)
			ue, pdu, _ := stgutg.RegisterUE(ue,
				c.Configuration.Mnc,
				c.Configuration.Mcc,
				conn)

			ueList = append(ueList, ue)
			pduList = append(pduList, pdu)

			time.Sleep(1 * time.Second)
		}

		for i := 0; i < pdu_establishment_number; i++ {
			fmt.Println(">> [ PDU ESTABLISHMENT TEST", i+1, "]")

			fmt.Println(">> Establishing PDU session for", ueList[i].Supi)
			stgutg.EstablishPDU(
				c.Configuration.SST,
				c.Configuration.SD,
				ueList[i],
				conn,
				c.Configuration.Gnb_gtp,
			)

			time.Sleep(1 * time.Second)
		}

		for i := 0; i < service_request_number; i++ {
			fmt.Println(">> [ SERVICE REQUEST TEST", i+1, "]")

			fmt.Println(">> Requesting service for", ueList[i].Supi)
			stgutg.ServiceRequest(pduList[i],
				ueList[i],
				conn,
				c.Configuration.Gnb_gtp)

			time.Sleep(1 * time.Second)
		}

		for i := 0; i < pdu_release_number; i++ {
			fmt.Println(">> [ PDU RELEASE TEST", i+1, "]")

			fmt.Println(">> Releasing PDU session for", ueList[i].Supi)
			stgutg.ReleasePDU(c.Configuration.SST,
				c.Configuration.SD,
				ueList[i],
				conn)
			time.Sleep(1 * time.Second)
		}

		for i := 0; i < ue_deregistration_number; i++ {
			fmt.Println(">> [ UE DEREGISTRATION TEST", i+1, "]")

			fmt.Println(">> Deregistering UE", ueList[i].Supi)
			stgutg.DeregisterUE(ueList[i],
				c.Configuration.Mnc,
				conn)
			time.Sleep(1 * time.Second)

		}

		fmt.Println(">> All tests finished")
		conn.Close()

		os.Exit(0)
	}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"stgutg"
	"tglib"
)

func main() {

	var ueList []*tglib.RanUeContext
	var pduList [][]byte
	teidUpfIPs := make(map[[4]byte]stgutg.TeidUpfIp)

	var c stgutg.Conf
	c.GetConfiguration()

	mode := stgutg.GetMode(os.Args)

	if mode == 1 {
		fmt.Println("TRAFFIC MODE")
		fmt.Println("----------------------")

		fmt.Println(">> Connecting to AMF")
		conn, err := tglib.ConnectToAmf(c.Configuration.Amf_ngap,
			c.Configuration.Gnb_ngap,
			c.Configuration.Amf_port,
			c.Configuration.Gnbn_port)
		stgutg.ManageError("Error in connection to AMF", err)

		fmt.Println(">> Managing NG Setup")
		stgutg.ManageNGSetup(conn,
			c.Configuration.Gnb_id,
			c.Configuration.Gnb_bitlength,
			c.Configuration.Gnb_name)

		for i := 0; i < c.Configuration.UeNumber; i++ {
			imsi := c.Configuration.Initial_imsi + i

			fmt.Println(">> Creating new UE with IMSI:", imsi)
			ue := stgutg.CreateUE(imsi,
				c.Configuration.K,
				c.Configuration.OPC,
				c.Configuration.OP)

			fmt.Println(">> Registering UE with IMSI:", imsi)
			ue, pdu, _ := stgutg.RegisterUE(ue,
				c.Configuration.Mnc,
				conn)

			ueList = append(ueList, ue)
			pduList = append(pduList, pdu)

			time.Sleep(1 * time.Second)
		}

		i := 0
		for _, pdu := range pduList {
			fmt.Println(">> Establishing PDU session for", ueList[i].Supi)
			stgutg.EstablishPDU(c.Configuration.SST,
				c.Configuration.SD,
				pdu,
				ueList[i],
				conn,
				c.Configuration.Gnb_gtp,
				teidUpfIPs)

			i++
			time.Sleep(1 * time.Second)
		}

		fmt.Println(teidUpfIPs)

		fmt.Println(">> Connecting to UPF")
		upfConn, err := tglib.ConnectToUpf(c.Configuration.Gnb_gtp,
			c.Configuration.Upf_gtp,
			c.Configuration.Gnbg_port,
			c.Configuration.Upf_port)
		stgutg.ManageError("Error in connection to UPF", err)

		fmt.Println(">> Opening traffic interfaces")
		ethSocketConn, err := tglib.NewEthSocketConn(c.Configuration.SrcIface)
		stgutg.ManageError("Error creating Ethernet socket", err)

		var stopProgram = make(chan os.Signal)
		signal.Notify(stopProgram, syscall.SIGTERM)
		signal.Notify(stopProgram, syscall.SIGINT)

		go func() {
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
			upfConn.Close()

			time.Sleep(1 * time.Second)
			os.Exit(0)
		}()

		fmt.Println(">> Listening to traffic responses")
		go stgutg.ListenForResponses(ethSocketConn,
			upfConn)

		fmt.Println(">> Waiting for traffic to send (Press Ctrl+C to quit)")
		stgutg.SendTraffic(upfConn, ethSocketConn, teidUpfIPs)

		time.Sleep(2 * time.Second)

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
		conn, err := tglib.ConnectToAmf(c.Configuration.Amf_ngap,
			c.Configuration.Gnb_ngap,
			c.Configuration.Amf_port,
			c.Configuration.Gnbn_port)
		stgutg.ManageError("Error in connection to AMF", err)

		fmt.Println(">> Managing NG Setup")
		stgutg.ManageNGSetup(conn,
			c.Configuration.Gnb_id,
			c.Configuration.Gnb_bitlength,
			c.Configuration.Gnb_name)

		for i := 0; i < c.Configuration.Test_ue_registation; i++ {
			fmt.Println(">> [ UE REGISTRATION TEST", i+1, "]")

			imsi := c.Configuration.Initial_imsi + i

			fmt.Println(">> Creating new UE with IMSI:", imsi)
			ue := stgutg.CreateUE(imsi,
				c.Configuration.K,
				c.Configuration.OPC,
				c.Configuration.OP)

			fmt.Println(">> Registering UE with IMSI:", imsi)
			ue, pdu, _ := stgutg.RegisterUE(ue,
				c.Configuration.Mnc,
				conn)

			ueList = append(ueList, ue)
			pduList = append(pduList, pdu)

			time.Sleep(1 * time.Second)
		}

		for i := 0; i < pdu_establishment_number; i++ {
			fmt.Println(">> [ PDU ESTABLISHMENT TEST", i+1, "]")

			fmt.Println(">> Establishing PDU session for", ueList[i].Supi)
			stgutg.EstablishPDU(c.Configuration.SST,
				c.Configuration.SD,
				pduList[i],
				ueList[i],
				conn,
				c.Configuration.Gnb_gtp,
				teidUpfIPs)

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

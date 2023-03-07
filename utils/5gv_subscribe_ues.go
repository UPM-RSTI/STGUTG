package main

import (
	"free5gc/lib/CommonConsumerTestData/UDM/TestGenAuthData"
	"free5gc/lib/MongoDBLibrary"
	"free5gc/lib/nas/security"
	"free5gc/src/test"

	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

type conf struct {
	Configuration struct {
		Name         string `yaml:"name"`
		Ipaddr       string `yaml:"ipaddr"`
		Port         string `yaml:"port"`
		Initial_imsi int    `yaml:"initial_imsi"`
		SplmnId      string `yaml:"splmnId"`
		Number       int    `yaml:"number"`
	}
}

func (c *conf) GetConfiguration() *conf {

	yamlFile, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(yamlFile, c)

	return c
}

func main() {
	var c conf

	c.GetConfiguration()

	MongoDBLibrary.SetMongoDB(c.Configuration.Name, "mongodb://"+c.Configuration.Ipaddr+":"+c.Configuration.Port)

	for i := 0; i < c.Configuration.Number; i++ {
		imsi := c.Configuration.Initial_imsi + i
		ue := test.NewRanUeContext("imsi-"+strconv.Itoa(imsi), 1, security.AlgCiphering128NEA0, security.AlgIntegrity128NIA2)

		ue.AmfUeNgapId = 1
		ue.AuthenticationSubs = test.GetAuthSubscription(TestGenAuthData.MilenageTestSet19.K,
			TestGenAuthData.MilenageTestSet19.OPC,
			TestGenAuthData.MilenageTestSet19.OP)

		servingPlmnId := "20893"

		test.InsertAuthSubscriptionToMongoDB(ue.Supi, ue.AuthenticationSubs)

		test.GetAuthSubscriptionFromMongoDB(ue.Supi)

		amData := test.GetAccessAndMobilitySubscriptionData()
		test.InsertAccessAndMobilitySubscriptionDataToMongoDB(ue.Supi, amData, servingPlmnId)
		test.GetAccessAndMobilitySubscriptionDataFromMongoDB(ue.Supi, servingPlmnId)

		smfSelData := test.GetSmfSelectionSubscriptionData()
		test.InsertSmfSelectionSubscriptionDataToMongoDB(ue.Supi, smfSelData, servingPlmnId)
		test.GetSmfSelectionSubscriptionDataFromMongoDB(ue.Supi, servingPlmnId)

		smSelData := test.GetSessionManagementSubscriptionData()
		test.InsertSessionManagementSubscriptionDataToMongoDB(ue.Supi, servingPlmnId, smSelData)
		test.GetSessionManagementDataFromMongoDB(ue.Supi, servingPlmnId)

		amPolicyData := test.GetAmPolicyData()
		test.InsertAmPolicyDataToMongoDB(ue.Supi, amPolicyData)
		test.GetAmPolicyDataFromMongoDB(ue.Supi)

		smPolicyData := test.GetSmPolicyData()
		test.InsertSmPolicyDataToMongoDB(ue.Supi, smPolicyData)
		test.GetSmPolicyDataFromMongoDB(ue.Supi)

		fmt.Println("imsi-" + strconv.Itoa(imsi) + " registered")
	}

}

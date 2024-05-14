package tglib

import (
	"encoding/binary"
	"encoding/hex"
	"free5gclib/CommonConsumerTestData/UDM/TestGenAuthData"
	"free5gclib/CommonConsumerTestData/UDR/TestRegistrationProcedure"
	"free5gclib/UeauCommon"
	"free5gclib/nas/nasMessage"
	"free5gclib/nas/nasType"
	"free5gclib/nas/security"
	"free5gclib/openapi/models"
	"regexp"

	"github.com/wmnsk/milenage"

	"github.com/calee0219/fatal"
)

type RanUeContext struct {
	Supi               string
	RanUeNgapId        int64
	AmfUeNgapId        int64
	ULCount            security.Count
	DLCount            security.Count
	CipheringAlg       uint8
	IntegrityAlg       uint8
	KnasEnc            [16]uint8
	KnasInt            [16]uint8
	Kamf               []uint8
	AuthenticationSubs models.AuthenticationSubscription
}

func GetAuthSubscription(k, opc, op string) models.AuthenticationSubscription {
	var authSubs models.AuthenticationSubscription
	authSubs.PermanentKey = &models.PermanentKey{
		PermanentKeyValue: k,
	}
	authSubs.Opc = &models.Opc{
		OpcValue: opc,
	}
	authSubs.Milenage = &models.Milenage{
		Op: &models.Op{
			OpValue: op,
		},
	}
	authSubs.AuthenticationManagementField = "8000"

	authSubs.SequenceNumber = TestGenAuthData.MilenageTestSet19.SQN
	authSubs.AuthenticationMethod = models.AuthMethod__5_G_AKA
	return authSubs
}

func GetAccessAndMobilitySubscriptionData() (amData models.AccessAndMobilitySubscriptionData) {
	return TestRegistrationProcedure.TestAmDataTable[TestRegistrationProcedure.FREE5GC_CASE]
}

func GetSmfSelectionSubscriptionData() (smfSelData models.SmfSelectionSubscriptionData) {
	return TestRegistrationProcedure.TestSmfSelDataTable[TestRegistrationProcedure.FREE5GC_CASE]
}

func GetSessionManagementSubscriptionData() (smfSelData models.SessionManagementSubscriptionData) {
	return TestRegistrationProcedure.TestSmSelDataTable[TestRegistrationProcedure.FREE5GC_CASE]
}

func GetAmPolicyData() (amPolicyData models.AmPolicyData) {
	return TestRegistrationProcedure.TestAmPolicyDataTable[TestRegistrationProcedure.FREE5GC_CASE]
}

func GetSmPolicyData() (smPolicyData models.SmPolicyData) {
	return TestRegistrationProcedure.TestSmPolicyDataTable[TestRegistrationProcedure.FREE5GC_CASE]
}

func NewRanUeContext(supi string, ranUeNgapId int64, cipheringAlg, integrityAlg uint8) *RanUeContext {
	ue := RanUeContext{}
	ue.RanUeNgapId = ranUeNgapId
	ue.Supi = supi
	ue.CipheringAlg = cipheringAlg
	ue.IntegrityAlg = integrityAlg
	return &ue
}

func (ue *RanUeContext) DeriveRESstarAndSetKey(
	authSubs models.AuthenticationSubscription, autn [16]uint8, rand []byte, snName string, mnc string, mcc string) []byte {

	sqn := make([]byte, 8)
	copy(sqn[2:], autn[0:6])

	amf, err := hex.DecodeString(authSubs.AuthenticationManagementField)
	if err != nil {
		fatal.Fatalf("DecodeString error: %+v", err)
	}

	// Run milenage
	opc := make([]byte, 16)
	_ = opc
	k, err := hex.DecodeString(authSubs.PermanentKey.PermanentKeyValue)
	if err != nil {
		fatal.Fatalf("DecodeString error: %+v", err)
	}
	var mil *milenage.Milenage
	if authSubs.Opc.OpcValue == "" {
		opStr := authSubs.Milenage.Op.OpValue
		var op []byte
		op, err = hex.DecodeString(opStr)
		if err != nil {
			fatal.Fatalf("DecodeString error: %+v", err)
		}

		mil = milenage.New(k, op, rand, binary.LittleEndian.Uint64(sqn), binary.LittleEndian.Uint16(amf))
	} else {
		opc, err = hex.DecodeString(authSubs.Opc.OpcValue)
		if err != nil {
			fatal.Fatalf("DecodeString error: %+v", err)
		}

		mil = milenage.NewWithOPc(k, opc, rand, binary.LittleEndian.Uint64(sqn), binary.LittleEndian.Uint16(amf))
	}

	// Generate MAC_A, MAC_S
	mil.F1()
	mil.F1Star(sqn, amf)

	// Generate CK, IK, AK
	_, ck, ik, ak, err := mil.F2345()
	if err != nil {
		fatal.Fatalf("ComputeRES error: %+v", err)
	}

	key := append(ck, ik...)

	ue.DerivateKamf(key, snName, autn[0:6], ak)
	ue.DerivateAlgKey()

	resStar, err := mil.ComputeRESStar(mcc, mnc)
	if err != nil {
		fatal.Fatalf("ComputeRES error: %+v", err)
	}

	return resStar
}

func (ue *RanUeContext) DerivateKamf(key []byte, snName string, SQN, AK []byte) {

	FC := UeauCommon.FC_FOR_KAUSF_DERIVATION
	P0 := []byte(snName)
	// SQNxorAK := make([]byte, 6)
	// for i := 0; i < len(SQN); i++ {
	// 	SQNxorAK[i] = SQN[i] ^ AK[i]
	// }
	// P1 := SQNxorAK
	P1 := SQN
	Kausf := UeauCommon.GetKDFValue(key, FC, P0, UeauCommon.KDFLen(P0), P1, UeauCommon.KDFLen(P1))
	P0 = []byte(snName)
	Kseaf := UeauCommon.GetKDFValue(Kausf, UeauCommon.FC_FOR_KSEAF_DERIVATION, P0, UeauCommon.KDFLen(P0))

	supiRegexp, err := regexp.Compile("(?:imsi|supi)-([0-9]{5,15})")
	if err != nil {
		fatal.Fatalf("regexp Compile error: %+v", err)
	}
	groups := supiRegexp.FindStringSubmatch(ue.Supi)

	P0 = []byte(groups[1])
	L0 := UeauCommon.KDFLen(P0)
	P1 = []byte{0x00, 0x00}
	L1 := UeauCommon.KDFLen(P1)

	ue.Kamf = UeauCommon.GetKDFValue(Kseaf, UeauCommon.FC_FOR_KAMF_DERIVATION, P0, L0, P1, L1)
}

// Algorithm key Derivation function defined in TS 33.501 Annex A.9
func (ue *RanUeContext) DerivateAlgKey() {
	// Security Key
	P0 := []byte{security.NNASEncAlg}
	L0 := UeauCommon.KDFLen(P0)
	P1 := []byte{ue.CipheringAlg}
	L1 := UeauCommon.KDFLen(P1)

	kenc := UeauCommon.GetKDFValue(ue.Kamf, UeauCommon.FC_FOR_ALGORITHM_KEY_DERIVATION, P0, L0, P1, L1)
	copy(ue.KnasEnc[:], kenc[16:32])

	// Integrity Key
	P0 = []byte{security.NNASIntAlg}
	L0 = UeauCommon.KDFLen(P0)
	P1 = []byte{ue.IntegrityAlg}
	L1 = UeauCommon.KDFLen(P1)

	kint := UeauCommon.GetKDFValue(ue.Kamf, UeauCommon.FC_FOR_ALGORITHM_KEY_DERIVATION, P0, L0, P1, L1)
	copy(ue.KnasInt[:], kint[16:32])
}

func (ue *RanUeContext) GetUESecurityCapability() (UESecurityCapability *nasType.UESecurityCapability) {
	UESecurityCapability = &nasType.UESecurityCapability{
		Iei:    nasMessage.RegistrationRequestUESecurityCapabilityType,
		Len:    2,
		Buffer: []uint8{0x00, 0x00},
	}
	switch ue.CipheringAlg {
	case security.AlgCiphering128NEA0:
		UESecurityCapability.SetEA0_5G(1)
	case security.AlgCiphering128NEA1:
		UESecurityCapability.SetEA1_128_5G(1)
	case security.AlgCiphering128NEA2:
		UESecurityCapability.SetEA2_128_5G(1)
	case security.AlgCiphering128NEA3:
		UESecurityCapability.SetEA3_128_5G(1)
	}

	switch ue.IntegrityAlg {
	case security.AlgIntegrity128NIA0:
		UESecurityCapability.SetIA0_5G(1)
	case security.AlgIntegrity128NIA1:
		UESecurityCapability.SetIA1_128_5G(1)
	case security.AlgIntegrity128NIA2:
		UESecurityCapability.SetIA2_128_5G(1)
	case security.AlgIntegrity128NIA3:
		UESecurityCapability.SetIA3_128_5G(1)
	}

	return
}

func (ue *RanUeContext) Get5GMMCapability() (capability5GMM *nasType.Capability5GMM) {
	return &nasType.Capability5GMM{
		Iei:   nasMessage.RegistrationRequestCapability5GMMType,
		Len:   1,
		Octet: [13]uint8{0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	}
}

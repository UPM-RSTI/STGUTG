package nasConvert

import (
	"free5gclib/nas/nasMessage"
	"free5gclib/nas/nasType"
	"free5gclib/openapi/models"
)

func SpareHalfOctetAndNgksiToModels(ngKsiNas nasType.SpareHalfOctetAndNgksi) (ngKsiModels models.NgKsi) {

	switch ngKsiNas.GetTSC() {
	case nasMessage.TypeOfSecurityContextFlagNative:
		ngKsiModels.Tsc = models.ScType_NATIVE
	case nasMessage.TypeOfSecurityContextFlagMapped:
		ngKsiModels.Tsc = models.ScType_MAPPED
	}

	ngKsiModels.Ksi = int32(ngKsiNas.GetNasKeySetIdentifiler())
	return
}

func SpareHalfOctetAndNgksiToNas(ngKsiModels models.NgKsi) (ngKsiNas nasType.SpareHalfOctetAndNgksi) {

	switch ngKsiModels.Tsc {
	case models.ScType_NATIVE:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagNative)
	case models.ScType_MAPPED:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagMapped)
	}

	ngKsiNas.SetSpareHalfOctet(0)
	ngKsiNas.SetNasKeySetIdentifiler(uint8(ngKsiModels.Ksi))
	return
}

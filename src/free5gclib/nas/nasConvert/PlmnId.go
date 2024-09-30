package nasConvert

import (
	"strconv"

	"free5gclib/nas/logger"
	"free5gclib/openapi/models"
)

func PlmnIDToNas(plmnID models.PlmnId) []uint8 {
	var plmnNas []uint8

	var mccDigit1, mccDigit2, mccDigit3 int
	if mccDigitTmp, err := strconv.Atoi(string(plmnID.Mcc[0])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc error: %+v", err)
	} else {
		mccDigit1 = mccDigitTmp
	}
	if mccDigitTmp, err := strconv.Atoi(string(plmnID.Mcc[1])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc error: %+v", err)
	} else {
		mccDigit2 = mccDigitTmp
	}
	if mccDigitTmp, err := strconv.Atoi(string(plmnID.Mcc[2])); err != nil {
		logger.ConvertLog.Warnf("atoi mcc error: %+v", err)
	} else {
		mccDigit3 = mccDigitTmp
	}

	var mncDigit1, mncDigit2, mncDigit3 int
	if mncDigitTmp, err := strconv.Atoi(string(plmnID.Mnc[0])); err != nil {
		logger.ConvertLog.Warnf("atoi mnc error: %+v", err)
	} else {
		mncDigit1 = mncDigitTmp
	}
	if mncDigitTmp, err := strconv.Atoi(string(plmnID.Mnc[1])); err != nil {
		logger.ConvertLog.Warnf("atoi mnc error: %+v", err)
	} else {
		mncDigit2 = mncDigitTmp
	}
	mncDigit3 = 0x0f
	if len(plmnID.Mnc) == 3 {
		if mncDigitTmp, err := strconv.Atoi(string(plmnID.Mnc[2])); err != nil {
			logger.ConvertLog.Warnf("atoi mn error: %+v", err)
		} else {
			mncDigit3 = mncDigitTmp
		}
	}

	plmnNas = []uint8{
		uint8((mccDigit2 << 4) | mccDigit1),
		uint8((mncDigit3 << 4) | mccDigit3),
		uint8((mncDigit2 << 4) | mncDigit1),
	}

	return plmnNas
}

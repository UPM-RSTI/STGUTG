package nasConvert

import (
	"encoding/hex"
	"free5gclib/nas/logger"
	"free5gclib/openapi/models"
)

func SnssaiToNas(snssai models.Snssai) []uint8 {
	var buf []uint8

	if snssai.Sd == "" {
		buf = append(buf, 0x01)
		buf = append(buf, uint8(snssai.Sst))
	} else {
		buf = append(buf, 0x04)
		buf = append(buf, uint8(snssai.Sst))
		if byteArray, err := hex.DecodeString(snssai.Sd); err != nil {
			logger.ConvertLog.Warnf("Decode snssai.sd failed: %+v", err)
		} else {
			buf = append(buf, byteArray...)
		}
	}
	return buf
}

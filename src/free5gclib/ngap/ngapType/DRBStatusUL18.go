package ngapType

import "free5gclib/aper"

// Need to import "free5gclib/aper" if it uses "aper"

type DRBStatusUL18 struct {
	ULCOUNTValue              COUNTValueForPDCPSN18                          `aper:"valueExt"`
	ReceiveStatusOfULPDCPSDUs *aper.BitString                                `aper:"sizeLB:1,sizeUB:131072,optional"`
	IEExtension               *ProtocolExtensionContainerDRBStatusUL18ExtIEs `aper:"optional"`
}

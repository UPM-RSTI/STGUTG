package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   SliceOverloadList
	SliceOverloadResponse               *OverloadResponse                                       `aper:"valueLB:0,valueUB:1,optional"`
	SliceTrafficLoadReductionIndication *TrafficLoadReductionIndication                         `aper:"optional"`
	IEExtensions                        *ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs `aper:"optional"`
}

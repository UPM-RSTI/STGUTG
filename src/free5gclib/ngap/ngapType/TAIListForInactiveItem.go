package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type TAIListForInactiveItem struct {
	TAI          TAI                                                     `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForInactiveItemExtIEs `aper:"optional"`
}

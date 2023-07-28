package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type TAIListForPagingItem struct {
	TAI          TAI                                                   `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs `aper:"optional"`
}

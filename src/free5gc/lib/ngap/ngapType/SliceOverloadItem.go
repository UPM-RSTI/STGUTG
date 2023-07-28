package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type SliceOverloadItem struct {
	SNSSAI       SNSSAI                                             `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceOverloadItemExtIEs `aper:"optional"`
}

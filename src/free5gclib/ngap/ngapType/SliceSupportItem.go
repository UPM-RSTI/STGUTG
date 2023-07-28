package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type SliceSupportItem struct {
	SNSSAI       SNSSAI                                            `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerSliceSupportItemExtIEs `aper:"optional"`
}

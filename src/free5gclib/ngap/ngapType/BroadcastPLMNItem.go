package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type BroadcastPLMNItem struct {
	PLMNIdentity        PLMNIdentity
	TAISliceSupportList SliceSupportList
	IEExtensions        *ProtocolExtensionContainerBroadcastPLMNItemExtIEs `aper:"optional"`
}

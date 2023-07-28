package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs `aper:"optional"`
}

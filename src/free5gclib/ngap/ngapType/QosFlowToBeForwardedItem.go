package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs `aper:"optional"`
}

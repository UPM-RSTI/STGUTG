package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowAcceptedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAcceptedItemExtIEs `aper:"optional"`
}

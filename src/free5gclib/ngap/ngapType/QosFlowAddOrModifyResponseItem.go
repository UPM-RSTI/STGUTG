package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowSetupResponseItemSURes struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowSetupResponseItemSUResExtIEs `aper:"optional"`
}

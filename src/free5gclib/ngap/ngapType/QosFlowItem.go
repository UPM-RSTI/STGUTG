package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowItem struct {
	QosFlowIdentifier QosFlowIdentifier
	Cause             Cause                                        `aper:"valueLB:0,valueUB:5"`
	IEExtensions      *ProtocolExtensionContainerQosFlowItemExtIEs `aper:"optional"`
}

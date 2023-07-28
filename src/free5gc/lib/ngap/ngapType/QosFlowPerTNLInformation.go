package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation UPTransportLayerInformation `aper:"valueLB:0,valueUB:1"`
	AssociatedQosFlowList       AssociatedQosFlowList
	IEExtensions                *ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation                             `aper:"valueLB:0,valueUB:1,optional"`
	QosFlowToBeForwardedList      *QosFlowToBeForwardedList                                `aper:"optional"`
	DataForwardingResponseDRBList *DataForwardingResponseDRBList                           `aper:"optional"`
	IEExtensions                  *ProtocolExtensionContainerHandoverCommandTransferExtIEs `aper:"optional"`
}

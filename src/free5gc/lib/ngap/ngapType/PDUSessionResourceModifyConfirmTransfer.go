package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList  QosFlowModifyConfirmList
	TNLMappingList            *TNLMappingList                                                          `aper:"optional"`
	QosFlowFailedToModifyList *QosFlowList                                                             `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs `aper:"optional"`
}

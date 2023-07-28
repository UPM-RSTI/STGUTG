package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceSetupResponseTransfer struct {
	QosFlowPerTNLInformation           QosFlowPerTNLInformation                                                 `aper:"valueExt"`
	AdditionalQosFlowPerTNLInformation *QosFlowPerTNLInformation                                                `aper:"valueExt,optional"`
	SecurityResult                     *SecurityResult                                                          `aper:"valueExt,optional"`
	QosFlowFailedToSetupList           *QosFlowList                                                             `aper:"optional"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs `aper:"optional"`
}

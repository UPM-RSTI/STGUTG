package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type PDUSessionResourceReleaseResponseTransfer struct {
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceReleaseResponseTransferExtIEs `aper:"optional"`
}

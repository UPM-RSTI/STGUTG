package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type HandoverPreparationUnsuccessfulTransfer struct {
	Cause        Cause                                                                    `aper:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerHandoverPreparationUnsuccessfulTransferExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type AMFTNLAssociationToRemoveItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation                                    `aper:"valueLB:0,valueUB:1"`
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs `aper:"optional"`
}

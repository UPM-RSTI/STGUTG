package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type XnExtTLAItem struct {
	IPsecTLA     *TransportLayerAddress                        `aper:"optional"`
	GTPTLAs      *XnGTPTLAs                                    `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerXnExtTLAItemExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type EPSTAI struct {
	PLMNIdentity PLMNIdentity
	EPSTAC       EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type TAI struct {
	PLMNIdentity PLMNIdentity
	TAC          TAC
	IEExtensions *ProtocolExtensionContainerTAIExtIEs `aper:"optional"`
}

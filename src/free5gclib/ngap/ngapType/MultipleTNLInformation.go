package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type MultipleTNLInformation struct {
	TNLInformationList TNLInformationList
	IEExtensions       *ProtocolExtensionContainerMultipleTNLInformationExtIEs `aper:"optional"`
}

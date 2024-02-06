package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type ERABInformationItem struct {
	ERABID       ERABID
	DLForwarding *DLForwarding                                        `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerERABInformationItemExtIEs `aper:"optional"`
}

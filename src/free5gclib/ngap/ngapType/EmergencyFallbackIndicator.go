package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type EmergencyFallbackIndicator struct {
	EmergencyFallbackRequestIndicator EmergencyFallbackRequestIndicator
	EmergencyServiceTargetCN          *EmergencyServiceTargetCN                                   `aper:"optional"`
	IEExtensions                      *ProtocolExtensionContainerEmergencyFallbackIndicatorExtIEs `aper:"optional"`
}

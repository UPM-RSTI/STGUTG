package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          EmergencyAreaID
	CancelledCellsInEAIEUTRA CancelledCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs `aper:"optional"`
}

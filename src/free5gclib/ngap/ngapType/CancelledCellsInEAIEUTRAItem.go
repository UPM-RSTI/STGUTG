package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInEAIEUTRAItemExtIEs `aper:"optional"`
}

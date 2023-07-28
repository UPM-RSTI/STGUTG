package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs `aper:"optional"`
}

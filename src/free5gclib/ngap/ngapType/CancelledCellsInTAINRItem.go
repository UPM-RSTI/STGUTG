package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAINRItemExtIEs `aper:"optional"`
}

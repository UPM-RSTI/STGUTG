package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastEUTRAItemExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs `aper:"optional"`
}

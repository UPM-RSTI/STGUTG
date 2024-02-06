package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type TAICancelledNRItem struct {
	TAI                   TAI `aper:"valueExt"`
	CancelledCellsInTAINR CancelledCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAICancelledNRItemExtIEs `aper:"optional"`
}

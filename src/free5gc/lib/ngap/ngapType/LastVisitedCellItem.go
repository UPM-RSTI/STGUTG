package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation                           `aper:"valueLB:0,valueUB:4"`
	IEExtensions               *ProtocolExtensionContainerLastVisitedCellItemExtIEs `aper:"optional"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

const (
	LastVisitedCellInformationPresentNothing int = iota /* No components present */
	LastVisitedCellInformationPresentNGRANCell
	LastVisitedCellInformationPresentEUTRANCell
	LastVisitedCellInformationPresentUTRANCell
	LastVisitedCellInformationPresentGERANCell
	LastVisitedCellInformationPresentChoiceExtensions
)

type LastVisitedCellInformation struct {
	Present          int
	NGRANCell        *LastVisitedNGRANCellInformation `aper:"valueExt"`
	EUTRANCell       *LastVisitedEUTRANCellInformation
	UTRANCell        *LastVisitedUTRANCellInformation
	GERANCell        *LastVisitedGERANCellInformation
	ChoiceExtensions *ProtocolIESingleContainerLastVisitedCellInformationExtIEs
}

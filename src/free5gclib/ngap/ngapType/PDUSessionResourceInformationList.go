package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceInformationList */
/* PDUSessionResourceInformationItem */
type PDUSessionResourceInformationList struct {
	List []PDUSessionResourceInformationItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToModifyListModCfm */
/* PDUSessionResourceFailedToModifyItemModCfm */
type PDUSessionResourceFailedToModifyListModCfm struct {
	List []PDUSessionResourceFailedToModifyItemModCfm `aper:"valueExt,sizeLB:1,sizeUB:256"`
}

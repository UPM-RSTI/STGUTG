package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct UL_NGU_UP_TNLModifyList */
/* ULNGUUPTNLModifyItem */
type ULNGUUPTNLModifyList struct {
	List []ULNGUUPTNLModifyItem `aper:"valueExt,sizeLB:0,sizeUB:4"`
}

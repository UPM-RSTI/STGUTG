package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct ServedGUAMIList */
/* ServedGUAMIItem */
type ServedGUAMIList struct {
	List []ServedGUAMIItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct XnGTP_TLAs */
/* TransportLayerAddress */
type XnGTPTLAs struct {
	List []TransportLayerAddress `aper:"sizeLB:1,sizeUB:16"`
}

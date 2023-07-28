package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

const (
	UPTransportLayerInformationPresentNothing int = iota /* No components present */
	UPTransportLayerInformationPresentGTPTunnel
	UPTransportLayerInformationPresentChoiceExtensions
)

type UPTransportLayerInformation struct {
	Present          int
	GTPTunnel        *GTPTunnel `aper:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerUPTransportLayerInformationExtIEs
}

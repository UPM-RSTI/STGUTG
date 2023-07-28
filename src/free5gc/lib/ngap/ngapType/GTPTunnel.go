package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type GTPTunnel struct {
	TransportLayerAddress TransportLayerAddress
	GTPTEID               GTPTEID
	IEExtensions          *ProtocolExtensionContainerGTPTunnelExtIEs `aper:"optional"`
}

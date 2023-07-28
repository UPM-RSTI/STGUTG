package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type SecurityContext struct {
	NextHopChainingCount NextHopChainingCount
	NextHopNH            SecurityKey
	IEExtensions         *ProtocolExtensionContainerSecurityContextExtIEs `aper:"optional"`
}

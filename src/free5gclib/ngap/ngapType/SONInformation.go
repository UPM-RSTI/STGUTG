package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

const (
	SONInformationPresentNothing int = iota /* No components present */
	SONInformationPresentSONInformationRequest
	SONInformationPresentSONInformationReply
	SONInformationPresentChoiceExtensions
)

type SONInformation struct {
	Present               int
	SONInformationRequest *SONInformationRequest
	SONInformationReply   *SONInformationReply `aper:"valueExt"`
	ChoiceExtensions      *ProtocolIESingleContainerSONInformationExtIEs
}

package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells *AssistanceDataForRecommendedCells                       `aper:"valueExt,optional"`
	PagingAttemptInformation          *PagingAttemptInformation                                `aper:"valueExt,optional"`
	IEExtensions                      *ProtocolExtensionContainerAssistanceDataForPagingExtIEs `aper:"optional"`
}

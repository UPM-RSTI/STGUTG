package ngapType

// Need to import "free5gclib/aper" if it uses "aper"

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR                             `aper:"optional"`
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA                          `aper:"optional"`
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs `aper:"optional"`
}

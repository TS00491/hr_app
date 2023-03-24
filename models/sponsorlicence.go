package models

type SponsorLicence struct {
	SponsorLicenceStatus string              `json:"sponsorlicencestatus"`
	BRPImage             string              `json:"brpimage"`
	VisaImage            string              `json:"visaimage"`
	VisaStartDate        string              `json:"visastartdate"`
	DateCOSAssigned      string              `json:"datecosassigned"`
	Qualifications       []Qualifications    `json:"qualifications"`
	SkilledWorkerVisa    []SkilledWorkerVisa `json:"skilledWorkerVisa"`
}

func NewSponsorLicence() *SponsorLicence {
	return &SponsorLicence{}
}

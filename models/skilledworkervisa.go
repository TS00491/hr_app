package models

type SkilledWorkerVisa struct {
	AdvertisingVacancyLink string  `json:"advertisingVacancyLink"`
	UKJobOffer             bool    `json:"ukJobOffer"`
	validCOSRef            bool    `json:"validCOSRef"`
	CoSRefNum              float64 `json:"cosRefNum"`
	IELTSBand              string  `json:"IELTSBand"`
	IELTSImg               string  `json:"IELTSImg"`
}

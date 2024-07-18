package models

type ResponseRoom struct {
	Data DataRoom `json:"data"`
}

type DataRoom struct {
	RoomGroups []RoomGroup `json:"roomGroups"`
}

type RoomGroup struct {
	MasterRoomTypeID            int                   `json:"masterRoomTypeId"`
	MasterRoomTypeName          string                `json:"masterRoomTypeName"`
	MasterRoomTypeAlternateName string                `json:"masterRoomTypeAlternateName"`
	MasterRoomTypeEnglishName   string                `json:"masterRoomTypeEnglishName"`
	CheapestRoomToken           string                `json:"cheapestRoomToken"`
	CheapestRoomTokens          []string              `json:"cheapestRoomTokens"`
	Rooms                       []RoomsData           `json:"rooms"`
	BedType                     string                `json:"bedType"`
	RoomSize                    float64               `json:"roomSize"`
	SizeInfo                    SizeInfo              `json:"sizeInfo"`
	Facilities                  []int                 `json:"facilities"`
	OccupancyInfo               OccupancyInfo         `json:"occupancyInfo"`
	Features                    []Feature             `json:"features"`
	Images                      []Image               `json:"images"`
	SuitableFor                 []string              `json:"suitableFor"`
	AllowChildren               bool                  `json:"allowChildren"`
	RoomReviewInformation       RoomReviewInformation `json:"roomReviewInformation"`
	IsRecommended               bool                  `json:"isRecommended"`
}

type RoomsData struct {
	RoomToken                          string                `json:"roomToken"`
	UID                                string                `json:"uid"`
	PC                                 PC                    `json:"pc"`
	BreakfastIncluded                  bool                  `json:"breakfastIncluded"`
	Benefits                           []RoomBenefit         `json:"benefits"`
	BadgeType                          int                   `json:"badgeType"`
	Badges                             []string              `json:"badges"`
	RoomOccupancyDescription           string                `json:"roomOccupancyDescription"`
	RemainRoom                         int                   `json:"remainRoom"`
	RoomInfoComponent                  []RoomInfoComponent   `json:"roomInfoComponent"`
	ProviderText                       ProviderText          `json:"providerText"`
	CancellationPolicyType             int                   `json:"cancellationPolicyType"`
	CurrentCancellationPolicyType      int                   `json:"currentCancellationPolicyType"`
	CancellationPolicyTitle            string                `json:"cancellationPolicyTitle"`
	CancellationPolicyTitleTemplate    Template              `json:"cancellationPolicyTitleTemplate"`
	CancellationPolicy                 string                `json:"cancellationPolicy"`
	CancellationInfo                   CancellationInfo      `json:"cancellationInfo"`
	IsBnpl                             bool                  `json:"isBnpl"`
	IsAgency                           bool                  `json:"isAgency"`
	IsMultiHotelEligible               bool                  `json:"isMultiHotelEligible"`
	IsPromoCodeEligible                bool                  `json:"isPromoCodeEligible"`
	IsVipDiscountEligible              bool                  `json:"isVipDiscountEligible"`
	IsFullyChargeAtAgoda               bool                  `json:"isFullyChargeAtAgoda"`
	IsCreditCardRequired               bool                  `json:"isCreditCardRequired"`
	IsApplePayAvailable                bool                  `json:"isApplePayAvailable"`
	IsImportantInformationTextRequired bool                  `json:"isImportantInformationTextRequired"`
	MasterRoomTypeID                   int                   `json:"masterRoomTypeId"`
	RoomName                           string                `json:"roomName"`
	EnglishRoomName                    string                `json:"englishRoomName"`
	RoomView                           string                `json:"roomView"`
	RoomSize                           float64               `json:"roomSize"`
	IsPromotionEligible                bool                  `json:"isPromotionEligible"`
	MaxRoomOccupancy                   int                   `json:"maxRoomOccupancy"`
	IsExtrabedAvaiable                 bool                  `json:"isExtrabedAvaiable"`
	HasSurcharge                       bool                  `json:"hasSurcharge"`
	MinimumFitRooms                    int                   `json:"minimumFitRooms"`
	NumberOfGuestsWithoutRoom          int                   `json:"numberOfGuestsWithoutRoom"`
	IsBOR                              bool                  `json:"isBOR"`
	OccupancyInfo                      OccupancyInfo         `json:"occupancyInfo"`
	HasBookingFee                      bool                  `json:"hasBookingFee"`
	PaymentInfo                        PaymentInfo           `json:"paymentInfo"`
	SuggestedRoomQuantity              int                   `json:"suggestedRoomQuantity"`
	DiscountMessages                   []string              `json:"discountMessages"`
	IsFreeChildrenSuggestionEligible   bool                  `json:"isFreeChildrenSuggestionEligible"`
	ChannelID                          int                   `json:"channelId"`
	StackedChannels                    StackedChannels       `json:"stackedChannels"`
	IsEasyCancel                       bool                  `json:"isEasyCancel"`
	RoomIdentifiers                    string                `json:"roomIdentifiers"`
	RareFindType                       int                   `json:"rareFindType"`
	RoomOfferNameType                  string                `json:"roomOfferNameType"`
	HasAnyBreakfastRelatedBenefits     bool                  `json:"hasAnyBreakfastRelatedBenefits"`
	IsTaiwanCampaignEligible           bool                  `json:"isTaiwanCampaignEligible"`
	IsGoLocalEligible                  bool                  `json:"isGoLocalEligible"`
	CartEligible                       bool                  `json:"cartEligible"`
	PricingDisplaySummary              PricingDisplaySummary `json:"pricingDisplaySummary"`
	IsCartRestricted                   bool                  `json:"isCartRestricted"`
	AllowMultipleBooking               bool                  `json:"allowMultipleBooking"`
	PapiRoomIdentifier                 string                `json:"papiRoomIdentifier"`
	RoomTypeID                         int                   `json:"roomTypeId"`
	BfURL                              string                `json:"bfUrl"`
	RoomOccupancy                      int                   `json:"roomOccupancy"`
	CancellationPolicyCode             string                `json:"cancellationPolicyCode"`
	RateplanID                         int                   `json:"rateplanID"`
	PromotionID                        int                   `json:"promotionID"`
	RoomNumAdult                       int                   `json:"roomNumAdult"`
	RoomNumChild                       int                   `json:"roomNumChild"`
	RoomDMCID                          string                `json:"roomDMCId"`
	RateCategory                       RateCategory          `json:"rateCategory"`
	Price                              float64               `json:"price"`
	PricePerNight                      float64               `json:"pricePerNight"`
	PriceType                          int                   `json:"priceType"`
	CrossedPrice                       float64               `json:"crossedPrice"`
	IsInclusive                        bool                  `json:"isInclusive"`
	PromotionType                      int                   `json:"promotionType"`
	PromotionSavings                   float64               `json:"promotionSavings"`
	MaxChildren                        int                   `json:"maxChildren"`
	LastBookedTimeStamp                string                `json:"lastBookedTimeStamp"`
	HasSummaryData                     bool                  `json:"HasSummaryData"`
	HasTaxPerPerson                    bool                  `json:"hasTaxPerPerson"`
	IsRequiredAddress                  bool                  `json:"isRequiredAddress"`
}

type PC struct {
	A                  float64     `json:"a"`
	I                  string      `json:"i"`
	E                  string      `json:"e"`
	T                  int         `json:"t"`
	S                  int         `json:"s"`
	Cor                float64     `json:"cor"`
	Ccor               float64     `json:"ccor"`
	Tf                 float64     `json:"tf"`
	TfPRPN             float64     `json:"tfPRPN"`
	PromotionPricePeek float64     `json:"promotionPricePeek"`
	OriginalTotal      float64     `json:"originalTotal"`
	PseudoCoupon       float64     `json:"pseudoCoupon"`
	PriceFreezePrice   PriceFreeze `json:"priceFreezePrice"`
}

type PriceFreeze struct {
	Exclusive float64 `json:"exclusive"`
	Inclusive float64 `json:"inclusive"`
}

type RoomBenefit struct {
	ID           int     `json:"id"`
	DisplayText  string  `json:"displayText"`
	Available    bool    `json:"available"`
	TargetType   string  `json:"targetType"`
	BenefitValue float64 `json:"benefitValue"`
}

type RoomInfoComponent struct {
	Title                string                `json:"title"`
	SubSectionComponents []SubSectionComponent `json:"subSectionComponents"`
	SectionType          int                   `json:"sectionType"`
	Type                 int                   `json:"type"`
}

type SubSectionComponent struct {
	SubTitle string `json:"subTitle"`
	Des      string `json:"des"`
	Icontype int    `json:"icontype"`
}

type ProviderText struct {
	Room    ProviderTextDetail `json:"room"`
	Payment ProviderTextDetail `json:"payment"`
}

type ProviderTextDetail struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Note     string `json:"note"`
}

type Template struct {
	Template string `json:"template"`
}

type CancellationInfo struct {
	Type                 int      `json:"type"`
	CurrentType          int      `json:"currentType"`
	FreeCancellationDate string   `json:"freeCancellationDate"`
	TitleTemplate        Template `json:"titleTemplate"`
	DescriptionTemplate  Template `json:"descriptionTemplate"`
	NoShowPolicy         string   `json:"noShowPolicy"`
	Policies             []string `json:"policies"`
	CancellationCode     string   `json:"cancellationCode"`
}

type OccupancyInfo struct {
	MaxOccupancy                   int `json:"maxOccupancy"`
	ExtraBed                       int `json:"extraBed"`
	MaxFreeChildren                int `json:"maxFreeChildren"`
	PerOfferFreeChildrenAndInfants int `json:"perOfferFreeChildrenAndInfants"`
}

type PaymentInfo struct {
	IsPrepay     bool `json:"isPrepay"`
	PaymentModel int  `json:"paymentModel"`
}

type StackedChannels struct {
	ID                int           `json:"id"`
	ChannelBreakDowns []interface{} `json:"channelBreakDowns"`
}

type PricingDisplaySummary struct {
	PerBook         PricingDetail `json:"perBook"`
	PerRoomPerBook  PricingDetail `json:"perRoomPerBook"`
	PerRoomPerNight PricingDetail `json:"perRoomPerNight"`
	PerNight        PricingDetail `json:"perNight"`
}

type PricingDetail struct {
	ChargeTotal              RoomPriceDetail `json:"chargeTotal"`
	RebateTotal              RoomPriceDetail `json:"rebateTotal"`
	RebateExtraBed           RoomPriceDetail `json:"rebateExtraBed"`
	DisplayTotal             RoomPriceDetail `json:"displayTotal"`
	PseudoCoupon             RoomPriceDetail `json:"pseudoCoupon"`
	OriginalTotal            RoomPriceDetail `json:"originalTotal"`
	CrossedOut               RoomPriceDetail `json:"crossedOut"`
	PayToAgoda               RoomPriceDetail `json:"payToAgoda"`
	PayAtHotel               RoomPriceDetail `json:"payAtHotel"`
	IgnoreDownliftAmount     RoomPriceDetail `json:"ignoreDownliftAmount"`
	AutoAppliedPromoDiscount RoomPriceDetail `json:"autoAppliedPromoDiscount"`
}

type RoomPriceDetail struct {
	Exclusive    float64 `json:"exclusive"`
	AllInclusive float64 `json:"allInclusive"`
}

type RateCategory struct {
	ID                 int         `json:"id"`
	StayPackageType    int         `json:"stayPackageType"`
	CheckIn            interface{} `json:"checkIn"`
	CheckOut           interface{} `json:"checkOut"`
	IsChildRateEnabled bool        `json:"isChildRateEnabled"`
}

type SizeInfo struct {
	Size            float64 `json:"size"`
	Unit            string  `json:"unit"`
	FullDescription string  `json:"fullDescription"`
}

type Feature struct {
	Name   string `json:"name"`
	Text   string `json:"text"`
	Symbol string `json:"symbol"`
}

type Image struct {
	URL          string `json:"url"`
	Caption      string `json:"caption"`
	ID           int    `json:"id"`
	TypeID       int    `json:"typeId"`
	UploadedDate string `json:"uploadedDate"`
	Blurhash     string `json:"blurhash"`
}

type RoomReviewInformation struct {
	Score     float64 `json:"score"`
	ScoreText string  `json:"scoreText"`
}

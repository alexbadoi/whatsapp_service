package models

type ResponseHotels struct {
	Data DataHotels `json:"data"`
}

type DataHotels struct {
	Filters    Filters    `json:"filters"`
	Properties []Property `json:"properties"`
}

type Filters struct {
	PaymentOptions []FilterData `json:"paymentOptions"`
	BedRooms       []FilterData `json:"bedRooms"`
	Facilities     []FilterData `json:"facilities"`
	PropertyTypes  []FilterData `json:"propertyType"`
	RoomOffers     []FilterData `json:"roomOffers"`
	RoomAmneties   []FilterData `json:"roomAmneties"`
}

type FilterData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	FilterKey string `json:"filterKey"`
}

type Property struct {
	PropertyID         int64       `json:"propertyId"`
	PropertyResultType string      `json:"propertyResultType"`
	MetaLab            MetaLab     `json:"metaLab"`
	SoldOut            interface{} `json:"soldOut"`
	Content            Content     `json:"content"`
	Pricing            Pricing     `json:"pricing"`
}

type MetaLab struct {
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	AttributeID int    `json:"attributeId"`
	Value       string `json:"value"`
}

type Content struct {
	PropertyID            int64                 `json:"propertyId"`
	Features              Features              `json:"features"`
	InformationSummary    InformationSummary    `json:"informationSummary"`
	Highlight             Highlight             `json:"highlight"`
	NonHotelAccommodation NonHotelAccommodation `json:"nonHotelAccommodation"`
	Reviews               Reviews               `json:"reviews"`
	FamilyFeatures        FamilyFeatures        `json:"familyFeatures"`
	PropertyEngagement    PropertyEngagement    `json:"propertyEngagement"`
	Images                Images                `json:"images"`
	LocalInformation      LocalInformation      `json:"localInformation"`
	RateCategories        RateCategories        `json:"rateCategories"`
}

type Features struct {
	HotelFacilities []HotelFacility `json:"hotelFacilities"`
}

type HotelFacility struct {
	ID   int         `json:"id"`
	Name interface{} `json:"name"`
}

type InformationSummary struct {
	LocaleName    string        `json:"localeName"`
	DefaultName   string        `json:"defaultName"`
	PropertyType  string        `json:"propertyType"`
	Accommodation Accommodation `json:"accommodation"`
	Remarks       interface{}   `json:"remarks"`
	Address       Address       `json:"address"`
	Rating        float32       `json:"rating"`
}

type Accommodation struct {
	AccommodationName string `json:"accommodationName"`
	AccommodationType int    `json:"accommodationType"`
}

type Address struct {
	Country Country `json:"country"`
	City    City    `json:"city"`
	Area    Area    `json:"area"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Highlight struct {
	Distance                      []interface{} `json:"distance"`
	CityCenter                    interface{}   `json:"cityCenter"`
	HasNearbyPublicTransportation bool          `json:"hasNearbyPublicTransportation"`
}

type NonHotelAccommodation struct {
	MasterRooms       []interface{} `json:"masterRooms"`
	HostLevel         interface{}   `json:"hostLevel"`
	IsRareFind        interface{}   `json:"isRareFind"`
	SupportedLongStay interface{}   `json:"supportedLongStay"`
}

type Reviews struct {
	Cumulative struct {
		ReviewCount int     `json:"reviewCount"`
		Score       float64 `json:"score"`
	} `json:"cumulative"`
}

type FamilyFeatures struct {
	IsFamilyRoom             bool `json:"isFamilyRoom"`
	IsInterConnectingRoom    bool `json:"isInterConnectingRoom"`
	IsInfantCottageAvailable bool `json:"isInfantCottageAvailable"`
	HasKidsClub              bool `json:"hasKidsClub"`
	HasKidsPool              bool `json:"hasKidsPool"`
}

type PropertyEngagement struct {
	PeopleLooking int    `json:"peopleLooking"`
	LastBooking   string `json:"lastBooking"`
	TodayBooking  string `json:"todayBooking"`
}

type Images struct {
	HotelImages []HotelImage `json:"hotelImages"`
}

type HotelImage struct {
	ID       int64       `json:"id"`
	Caption  string      `json:"caption"`
	URLs     []ImageURL  `json:"urls"`
	GroupID  string      `json:"groupId"`
	TypeID   int         `json:"typeId"`
	Blurhash interface{} `json:"blurhash"`
}

type ImageURL struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LocalInformation struct {
	Landmarks struct {
		TopLandmark interface{} `json:"topLandmark"`
	} `json:"landmarks"`
	HasAirportTransfer bool `json:"hasAirportTransfer"`
}

type RateCategories struct {
	EscapeRateCategories interface{} `json:"escapeRateCategories"`
}

type Pricing struct {
	HotelID                      int64                         `json:"hotelId"`
	IsReady                      bool                          `json:"isReady"`
	IsAvailable                  bool                          `json:"isAvailable"`
	IsEasyCancel                 bool                          `json:"isEasyCancel"`
	IsSuggested                  bool                          `json:"isSuggested"`
	Payment                      Payment                       `json:"payment"`
	PricingMessages              []PricingMessage              `json:"pricingMessages"`
	Benefits                     []int                         `json:"benefits"`
	Offers                       []Offer                       `json:"offers"`
	RoomBundle                   interface{}                   `json:"roomBundle"`
	SuggestedRoomQuantity        int                           `json:"suggestedRoomQuantity"`
	SuggestPriceType             SuggestPriceType              `json:"suggestPriceType"`
	LoyaltyDisplay               LoyaltyDisplay                `json:"loyaltyDisplay"`
	CheapestStayPackageRatePlans []CheapestStayPackageRatePlan `json:"cheapestStayPackageRatePlans"`
}

type Payment struct {
	PayLater     PayLater     `json:"payLater"`
	Cancellation Cancellation `json:"cancellation"`
	NoCreditCard NoCreditCard `json:"noCreditCard"`
	PayAtHotel   PayAtHotel   `json:"payAtHotel"`
}

type PayLater struct {
	IsEligible bool   `json:"isEligible"`
	AuthDate   string `json:"authDate"`
	ChargeDate string `json:"chargeDate"`
}

type Cancellation struct {
	Code             string `json:"code"`
	CancellationType string `json:"cancellationType"`
}

type NoCreditCard struct {
	IsEligible bool `json:"isEligible"`
}

type PayAtHotel struct {
	IsEligible bool `json:"isEligible"`
}

type PricingMessage struct {
	Location int   `json:"location"`
	IDs      []int `json:"ids"`
}

type Offer struct {
	BundleType   string      `json:"bundleType"`
	BundleDetail interface{} `json:"bundleDetail"`
	RoomOffers   []RoomOffer `json:"roomOffers"`
}

type RoomOffer struct {
	Room Room `json:"room"`
}

type Room struct {
	UID             string          `json:"uid"`
	SupplierID      int             `json:"supplierId"`
	AvailableRooms  int             `json:"availableRooms"`
	Benefits        []Benefit       `json:"benefits"`
	Payment         Payment         `json:"payment"`
	LocalVoucher    interface{}     `json:"localVoucher"`
	Campaign        interface{}     `json:"campaign"`
	IsPromoEligible bool            `json:"isPromoEligible"`
	Occupancy       int             `json:"occupancy"`
	AgodaCash       AgodaCash       `json:"agodaCash"`
	Cashback        interface{}     `json:"cashback"`
	Discount        Discount        `json:"discount"`
	Channel         Channel         `json:"channel"`
	Promotions      interface{}     `json:"promotions"`
	LoyaltyDisplay  interface{}     `json:"loyaltyDisplay"`
	BookingDuration interface{}     `json:"bookingDuration"`
	CorInfo         CorInfo         `json:"corInfo"`
	Pricing         []RoomPricing   `json:"pricing"`
	Packaging       Packaging       `json:"packaging"`
	StayPackageType int             `json:"stayPackageType"`
	PricingMessages PricingMessages `json:"pricingMessages"`
}

type Benefit struct {
	ID          int     `json:"id"`
	Value       float32 `json:"value"`
	Remark      string  `json:"remark"`
	Unit        int     `json:"unit"`
	TargetType  int     `json:"targetType"`
	Description string  `json:"description"`
}

type AgodaCash struct {
	ShowBadge    bool   `json:"showBadge"`
	GiftcardGuid string `json:"giftcardGuid"`
	DayToEarn    int    `json:"dayToEarn"`
	EarnID       int    `json:"earnId"`
	ExpiryDay    int    `json:"expiryDay"`
}

type Discount struct {
	PseudoCoupon    interface{} `json:"pseudoCoupon"`
	ChannelDiscount interface{} `json:"channelDiscount"`
	Deals           []string    `json:"deals"`
}

type Channel struct {
	ID int `json:"id"`
}

type CorInfo struct {
	CorBreakdown CorBreakdown `json:"corBreakdown"`
}

type CorBreakdown struct {
	TaxExPRPN []CorItem `json:"taxExPRPN"`
	TaxInPRPN []CorItem `json:"taxInPRPN"`
}

type CorItem struct {
	CmsID      int     `json:"cmsId"`
	ID         int     `json:"id"`
	IsDiscount bool    `json:"isDiscount"`
	Price      float64 `json:"price"`
}

type RoomPricing struct {
	Currency               string                 `json:"currency"`
	Price                  PriceDetail            `json:"price"`
	ChannelDiscountSummary ChannelDiscountSummary `json:"channelDiscountSummary"`
	PromotionPricePeek     interface{}            `json:"promotionPricePeek"`
}

type PriceDetail struct {
	PerBook         PriceBreakdown `json:"perBook"`
	PerRoomPerNight PriceBreakdown `json:"perRoomPerNight"`
	PerNight        PriceBreakdown `json:"perNight"`
	TotalDiscount   int            `json:"totalDiscount"`
}

type PriceBreakdown struct {
	Exclusive PriceInfo `json:"exclusive"`
	Inclusive PriceInfo `json:"inclusive"`
}

type PriceInfo struct {
	Display              float64     `json:"display"`
	ChargeTotal          float64     `json:"chargeTotal"`
	OriginalPrice        float64     `json:"originalPrice"`
	CrossedOutPrice      float64     `json:"crossedOutPrice"`
	RebatePrice          float64     `json:"rebatePrice"`
	CashbackPrice        interface{} `json:"cashbackPrice"`
	DisplayAfterCashback interface{} `json:"displayAfterCashback"`
	PseudoCouponPrice    float64     `json:"pseudoCouponPrice"`
}

type ChannelDiscountSummary struct {
	ChannelID                int                        `json:"channelId"`
	ChannelDiscountBreakdown []ChannelDiscountBreakdown `json:"channelDiscountBreakdown"`
}

type ChannelDiscountBreakdown struct {
	ChannelID       int     `json:"channelId"`
	Display         bool    `json:"display"`
	DiscountPercent float32 `json:"discountPercent"`
}

type Packaging struct {
	Token   Token         `json:"token"`
	Pricing []interface{} `json:"pricing"`
}

type Token struct {
	ClientToken      string `json:"clientToken"`
	InterSystemToken string `json:"interSystemToken"`
}

type PricingMessages struct {
	Formatted []FormattedMessage `json:"formatted"`
}

type FormattedMessage struct {
	Location int    `json:"location"`
	Texts    []Text `json:"texts"`
}

type Text struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

type SuggestPriceType struct {
	SuggestPrice string `json:"suggestPrice"`
	ApplyType    string `json:"applyType"`
}

type LoyaltyDisplay struct {
	Items []interface{} `json:"items"`
}

type CheapestStayPackageRatePlan struct {
	StayPackageType int `json:"stayPackageType"`
	RatePlanID      int `json:"ratePlanId"`
}

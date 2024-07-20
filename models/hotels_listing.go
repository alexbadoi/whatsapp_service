package models

import (
	"fmt"
	"reflect"
	"strconv"
)

type HotelListing struct {
	PropertyID    int64    // From Property.ID
	Name          string   // From Content.InformationSummary.LocaleName or DefaultName
	Rating        float64  // From Content.InformationSummary.Rating (convert to float64)
	LocationId    int64    // From Content.InformationSummary.Address.City.ID
	Location      string   // Combine Content.InformationSummary.Address.City.Name and Country.Name
	TotalReviews  int      // From Content.Reviews.Cumulative.ReviewCount
	ReviewScore   float64  // From Content.Reviews.Cumulative.Score
	Amenities     []string // From Content.Features.HotelFacilities (need to map IDs to names)
	PricePerNight float64  // From Pricing.Offers[0].RoomOffers[0].Room.Pricing[0].Price.PerNight.Inclusive.Display
	ImageURL      []string // From Content.Images.HotelImages[0].URLs (find the one with key "thumbnail")
	Currency      string   // From Pricing.Offers[0].RoomOffers[0].Room.Pricing[0].Price.PerNight.Inclusive.Currency
}

func GetHotelListing(property Property, price float64, filters Filters, currency string) HotelListing {
	return HotelListing{
		PropertyID:    property.PropertyID,
		Name:          property.Content.InformationSummary.LocaleName,
		Rating:        float64(property.Content.InformationSummary.Rating),
		Location:      fmt.Sprintf("%s, %s", property.Content.InformationSummary.Address.City.Name, property.Content.InformationSummary.Address.Country.Name),
		TotalReviews:  property.Content.Reviews.Cumulative.ReviewCount,
		ReviewScore:   property.Content.Reviews.Cumulative.Score,
		Amenities:     getAmenities(property.Content.Features.HotelFacilities, filters),
		PricePerNight: price,
		ImageURL:      getImageURL(property.Content.Images.HotelImages),
		Currency:      currency,
	}
}

func getAmenities(facilities []HotelFacility, filters Filters) []string {
	var amenities []string
	values := reflect.ValueOf(filters)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Type().Kind() == reflect.Slice {
			sliceValue := values.Field(i).Interface().([]FilterData)
			for _, filter := range sliceValue {
				for _, facility := range facilities {
					if filter.ID == facility.ID {
						amenities = append(amenities, filter.Name)
					}
				}
			}
		}
	}
	if len(amenities) > 3 {
		totLen := len(amenities)
		amenities = amenities[:3]
		amenities = append(amenities, strconv.Itoa(totLen-3)+" +")
	}
	return amenities
}

func getImageURL(images []HotelImage) []string {
	var urls []string
	for _, img := range images {
		for _, url := range img.URLs {
			urls = append(urls, url.Value)
		}
	}
	return urls
}

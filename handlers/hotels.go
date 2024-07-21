package handlers

import (
	"encoding/json"
	"hotel/models"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetHotels(c *fiber.Ctx) error {
	id := c.Query("id")
	checkIn := c.Query("checkIn")
	checkOut := c.Query("checkOut")
	adults := c.Query("adult")
	room := c.Query("room")
	searchUrl := url + "search-overnight?id=" + id + "&checkinDate=" + checkIn + "&checkoutDate=" + checkOut + "&adult=" + adults
	searchUrl += "&room=" + room
	rating := c.Query("starRating")
	childrens := c.Query("children")
	guestReviews := c.Query("guestReviews")
	kidsStayForFree := c.Query("kidsStayForFree")
	if rating != "" {
		searchUrl += "&starRating=" + rating
	}
	paymentOptions := c.Query("paymentOptions")
	if paymentOptions != "" {
		searchUrl += "&paymentOptions=" + paymentOptions
	}
	facilities := c.Query("facilities")
	if facilities != "" {
		searchUrl += "&facilities=" + facilities
	}
	bedrooms := c.Query("bedrooms")
	if bedrooms != "" {
		searchUrl += "&bedrooms=" + bedrooms
	}
	if guestReviews != "" {
		searchUrl += "&guestReviews=" + guestReviews
	}
	if kidsStayForFree != "" {
		searchUrl += "&kidsStayForFree=" + kidsStayForFree
	}
	if childrens != "" {
		searchUrl += "&childAges=" + childrens
	}
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		log.Println(err)
		return c.SendString("Internal server error")
	}
	req.Header.Set("x-rapidapi-key", key)
	log.Println(searchUrl)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
		log.Println(string(body))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	var hotelResponse models.ResponseHotels
	err = json.NewDecoder(resp.Body).Decode(&hotelResponse)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	var hotelList []models.HotelListing
	for _, property := range hotelResponse.Data.Properties {
		price := 0.0
		currency := "USD"
		if len(property.Pricing.Offers) > 0 {
			price = property.Pricing.Offers[0].RoomOffers[0].Room.Pricing[0].Price.PerNight.Inclusive.Display
			currency = property.Pricing.Offers[0].RoomOffers[0].Room.Pricing[0].Currency
		}
		hotelList = append(hotelList, models.GetHotelListing(property, price, hotelResponse.Data.Filters, currency))
	}
	return c.JSON(hotelList)
}

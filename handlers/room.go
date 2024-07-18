package handlers

import (
	"encoding/json"
	"hotel/models"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetRooms(c *fiber.Ctx) error {
	id := c.Query("id")
	checkIn := c.Query("checkIn")
	checkOut := c.Query("checkOut")
	room := c.Query("room")
	adults := c.Query("adult")
	currency := c.Query("currency")
	searchUrl := url + "room-prices?propertyId=" + id + "&checkinDate=" + checkIn + "&checkoutDate=" + checkOut + "&adult=" + adults
	searchUrl += "&room=" + room
	searchUrl += "&currency=" + currency
	childrens := c.Query("children")
	if childrens != "" {
		searchUrl += "&childAges=" + childrens
	}
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		log.Println(err)
		return c.SendString("Internal server error")
	}
	req.Header.Set("x-rapidapi-key", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	var response models.ResponseRoom
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	roomListingResponse := models.ExtractRoomListingResponse(response)
	return c.JSON(roomListingResponse)
}

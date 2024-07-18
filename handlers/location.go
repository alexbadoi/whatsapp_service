package handlers

import (
	"encoding/json"
	"hotel/models"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

const url = "https://agoda-com.p.rapidapi.com/hotels/"

var key = getEnv("AGODA_KEY")

func GetLocations(c *fiber.Ctx) error {
	var query = c.Query("name")
	if query == "" {
		log.Println("Query missing")
		return c.Status(fiber.StatusBadRequest).SendString("Bad request. Query missing")
	}

	req, err := http.NewRequest(http.MethodGet, url+"auto-complete?query="+query, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	req.Header.Set("x-rapidapi-key", key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	var locationResponse models.LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.JSON(locationResponse)
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}
	return value
}

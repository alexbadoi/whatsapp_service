package handlers

import (
	"hotel/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveProposal(c *fiber.Ctx) error {
	var proposal []models.HotelProposal
	if err := c.BodyParser(&proposal); err != nil {
		log.Println(err)
		return err
	}
	id := uuid.New()
	for i := range proposal {
		proposal[i].ProposalID = id.String()
	}
	ok := models.Save(proposal)
	if ok != "ok" {
		log.Println(ok)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusCreated)
}

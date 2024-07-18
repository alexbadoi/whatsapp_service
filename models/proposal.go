package models

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HotelProposal struct {
	ProposalHotelID        int     `gorm:"primaryKey;autoIncrement"`
	ProposalID             string  `gorm:"index"`
	DayStepCount           int     `gorm:"check:day_step_count >= 1 AND day_step_count <= 5" json:"stepCount"`
	BookingDT              string  `json:"bookingDt"`
	HotelName              string  `json:"hotelName"`
	AgodaProprietyID       string  `json:"agodaProprietyId"`
	CheckinDT              string  `json:"checkIn"`
	CheckoutDT             string  `json:"checkOut"`
	NumberAdults           int     `json:"numberAdults"`
	NumberChildren         int     `json:"numberChildren"`
	ChildrenAges           string  `json:"childrenAge"`
	RoomName               string  `json:"roomName"`
	RoomNumber             int     `json:"roomNumber"`
	CommonSpace            string  `json:"commonSpace"`
	RoomSize               string  `json:"roomSize"`
	CancelationPolicyTitle string  `json:"cancelationPolicyTitle"`
	CancelationDeadlineDT  string  `json:"cancelationDeadlineDt"`
	TotalPricePerNight     float64 `json:"totalPricePerNight"`
	Currency               string  `json:"currency"`
}

func Save(p []HotelProposal) string {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	db.AutoMigrate(&HotelProposal{})
	tx := db.Create(&p)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error.Error()
	}
	t, _ := db.DB()
	t.Close()
	return "ok"
}

package models

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ActiveFlg *int           `gorm:"default:1"`
}

type AgodaData struct {
	ProposalDetailID            int        `gorm:"primaryKey"`
	ProposalID                  string     `gorm:"not null"`
	DayCount                    int        `gorm:"not null" json:"dayCount"`
	AgodaHotelLocationID        string     `json:"agodaHotelLocationID"`
	AgodaLocationName           string     `gorm:"type:varchar(255)" json:"agodaLocationName"`
	StartDt                     time.Time  `gorm:"not null" json:"startDt"`
	EndDt                       time.Time  `gorm:"not null" json:"endDt"`
	NumberOfRooms               int        `gorm:"not null" json:"numberOfRooms"`
	NumberOfAdults              int        `gorm:"not null" json:"numberOfAdults"`
	NumberOfChildren            int        `gorm:"not null" json:"numberOfChildren"`
	ChildAgesPipeDelimited      string     `gorm:"type:varchar(50)" json:"childAgesPipeDelimited"`
	AgodaHotelID                int        `json:"agodaHotelID"`
	AgodaHotelName              string     `gorm:"type:varchar(255)" json:"agodaHotelName"`
	AgodaHotelStars             float64    `json:"agodaHotelStars"`
	AgodaHotelRating            float64    `json:"agodaHotelRating"`
	AgodaHotelRatingsCount      int        `json:"agodaHotelRatingsCount"`
	AgodaHotelPicsPipeDelimited string     `gorm:"type:text" json:"agodaHotelPicsPipeDelimited"`
	AgodaRoomPicsPipeDelimited  string     `gorm:"type:text" json:"agodaRoomPicsPipeDelimited"`
	SleepsCount                 int        `gorm:"not null" json:"sleepsCount"`
	RoomPricePerNight           float64    `gorm:"type:decimal(10,2)" json:"roomPricePerNight"`
	RiskFreeBookingBinary       int8       `gorm:"type:tinyint" json:"riskFreeBookingBinary"`
	CancelationDeadlineDt       *time.Time `json:"cancelationDeadlineDt"`
	IncludesBreakfastBinary     int8       `gorm:"type:tinyint" json:"includesBreakfastBinary"`
	IncludesLunchBinary         int8       `gorm:"type:tinyint" json:"includesLunchBinary"`
	IncludesDinnerBinary        int8       `gorm:"type:tinyint" json:"includesDinnerBinary"`
	BedType                     string     `gorm:"type:varchar(255)" json:"bedType"`
	SizeSqm                     string     `json:"sizeSqm"`
	BaseModel
}

func (AgodaData) TableName() string {
	return "agoda_data"
}

func Save(p []AgodaData) string {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	tx := db.Create(&p)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error.Error()
	}
	t, _ := db.DB()
	t.Close()
	return "ok"
}

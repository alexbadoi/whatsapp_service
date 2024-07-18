package models

import (
	"strconv"
	"strings"
	"time"
)

type RoomListingResponse struct {
	RoomGroup []RoomGroupListing `json:"roomGroup"`
}

type RoomGroupListing struct {
	Rooms    []RoomListing `json:"rooms"`
	Name     string        `json:"name"`
	ID       string        `json:"id"`
	Images   []string      `json:"images"`
	Features []string      `json:"features"`
	BedType  string        `json:"bedType"`
	SizeInfo string        `json:"sizeInfo"`
}

type RoomListing struct {
	ID                       string   `json:"id"`
	Price                    float64  `json:"price"`
	Benefits                 []string `json:"benefits"`
	MaxOccupancy             int      `json:"maxOccupancy"`
	CancelationPolicy        string   `json:"cancelationPolicy"`
	CancelationPolicyMaxDate string   `json:"cancelationPolicyMaxDate"`
}

func ExtractRoomListingResponse(response ResponseRoom) RoomListingResponse {
	result := RoomListingResponse{
		RoomGroup: make([]RoomGroupListing, 0, len(response.Data.RoomGroups)),
	}

	for _, roomGroup := range response.Data.RoomGroups {
		groupListing := RoomGroupListing{
			Name:     roomGroup.MasterRoomTypeName,
			ID:       strconv.Itoa(roomGroup.MasterRoomTypeID),
			BedType:  roomGroup.BedType,
			SizeInfo: roomGroup.SizeInfo.FullDescription,
			Rooms:    make([]RoomListing, 0, len(roomGroup.Rooms)),
			Images:   make([]string, 0, len(roomGroup.Images)),
			Features: make([]string, 0, len(roomGroup.Features)),
		}

		for _, image := range roomGroup.Images {
			groupListing.Images = append(groupListing.Images, image.URL)
		}

		for _, feature := range roomGroup.Features {
			if !strings.Contains(feature.Name, "size") {
				groupListing.Features = append(groupListing.Features, feature.Text)
			}
		}

		for _, room := range roomGroup.Rooms {
			roomListing := RoomListing{
				ID:                       room.UID,
				Price:                    room.PricingDisplaySummary.PerRoomPerBook.ChargeTotal.AllInclusive,
				MaxOccupancy:             room.MaxRoomOccupancy,
				CancelationPolicy:        room.CancellationPolicy,
				CancelationPolicyMaxDate: parseAndFormatDate(room.CancellationInfo.FreeCancellationDate),
				Benefits:                 make([]string, 0, len(room.Benefits)),
			}

			for _, benefit := range room.Benefits {
				roomListing.Benefits = append(roomListing.Benefits, benefit.DisplayText)
			}
			if room.ProviderText.Payment.Title != "" {
				roomListing.Benefits = append(roomListing.Benefits, room.ProviderText.Payment.Title)
			}
			groupListing.Rooms = append(groupListing.Rooms, roomListing)
		}

		result.RoomGroup = append(result.RoomGroup, groupListing)
	}

	return result
}

func parseAndFormatDate(dateString string) string {
	dateString = strings.TrimPrefix(dateString, "/Date(")
	dateString = strings.TrimSuffix(dateString, ")/")
	parts := strings.Split(dateString, "+")
	if len(parts) != 2 {
		return "Invalid date format"
	}
	msec, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return "Invalid date format"
	}
	t := time.Unix(msec/1000, 0)
	offset, err := strconv.Atoi(parts[1])
	if err != nil {
		return "Invalid date format"
	}
	location := time.FixedZone("", offset*3600)
	t = t.In(location)
	return t.Format("2006-01-02")
}

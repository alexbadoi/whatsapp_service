package models

type LocationResponse struct {
	Data []Location `json:"data"`
}
type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

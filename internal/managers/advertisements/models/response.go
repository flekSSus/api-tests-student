package models

type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

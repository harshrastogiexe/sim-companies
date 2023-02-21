package core

import "time"

type MarketItem struct {
	ID       int     `json:"id"`
	Kind     int     `json:"kind"`
	Quantity int     `json:"quantity"`
	Quality  int     `json:"quality"`
	Price    float64 `json:"price"`
	Seller   struct {
		ID           int         `json:"id"`
		Company      string      `json:"company"`
		RealmID      int         `json:"realmId"`
		Logo         string      `json:"logo"`
		Certificates int         `json:"certificates"`
		ContestWins  int         `json:"contest_wins"`
		Npc          bool        `json:"npc"`
		CourseID     interface{} `json:"courseId"`
		IP           string      `json:"ip"`
	} `json:"seller"`
	Posted time.Time `json:"posted"`
	Fees   int       `json:"fees"`
}

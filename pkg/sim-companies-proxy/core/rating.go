package core

type Rating struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Insurance   int64  `json:"insurance"`
}

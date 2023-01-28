package core

type Level struct {
	Start               int64  `json:"start"`
	Name                string `json:"name"`
	MaxBuildings        int64  `json:"maxBuildings"`
	CanScrape           bool   `json:"canScrape"`
	CanResearch         bool   `json:"canResearch"`
	CanContracts        bool   `json:"canContracts"`
	CanBonds            bool   `json:"canBonds"`
	CanExecutives       bool   `json:"canExecutives"`
	CanGovernmentOrders bool   `json:"canGovernmentOrders"`
	TimeLimitS          int64  `json:"timeLimitS"`
}

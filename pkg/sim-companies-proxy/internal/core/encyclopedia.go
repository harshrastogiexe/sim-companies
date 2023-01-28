package core

type Encyclopedia interface {
	// fetches resource from https://www.simcompanies.com api service
	GetResource(id string) (*Resource, error)

	// fetches building from https://www.simcompanies.com api service
	GetBuilding(id string) (*Building, error)

	// fetches levels from https://www.simcompanies.com api service
	GetLevels() ([]Level, error)

	// fetches rating from https://www.simcompanies.com api service
	GetRatings() ([]Rating, error)
}

package models

type Building struct {
	BuildingBase
	DoesProduce []ResourceBase
	DoesSell    []ResourceBase
	CostUnits   float64
	Hours       int64
	Wages       float64
}

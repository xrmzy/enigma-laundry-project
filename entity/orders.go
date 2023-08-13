package entity

type Orders struct {
	OrderId      string
	CustomerId   string
	CustomerName string
	Service      string
	Unit         string
	TotalCost    float64
	OutletName   string
	OrderDate    string
	Status       string
}

package entity

type OrderDetails struct {
	OrderDetailsId string
	OrderId        string
	CustomerId     string
	CustomerName   string
	ServiceId      string
	ServiceName    string
	Unit           string
	Price          float64
	TotalCost      float64
	OutletId       string
	OutletName     string
	OutletAddress  string
	OrderDate      string
	WorkDate       string
	Status         string
	Note           string
}

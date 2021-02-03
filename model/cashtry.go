package model

type Cashtry struct {
	TotalCash   float64
	TotalOrder  int
	TVisa       float64
	TVoid       float64
	MonthNo     int
	AverageCash float64
	NoOfCashTry int
	AvgBasket   float64
}

type CashtryReq struct {
	Store uint `json:"store"`
	Year  uint `json:"year"`
}
type CashtryStores struct {
	StoreCode int    `json:"store_code"`
	StoreName string `json:"store_name"`
}

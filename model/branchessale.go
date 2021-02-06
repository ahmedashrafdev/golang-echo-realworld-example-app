package model

type BranchesSale struct {
	StoreCode   int
	Totalamount float64
	StoreName   string
}

type BranchesSaleReq struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
}

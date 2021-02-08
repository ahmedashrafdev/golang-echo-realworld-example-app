package model

type MonthlySales struct {
	Totalamount float64
	DocMonth    int
}

type MonthlySalesReq struct {
	Year uint `json:"Year"`
}

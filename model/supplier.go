package model

type Supplier struct {
	AccountCode    int
	AccountName    string
	DBT            float32
	CRDT           float32
	ReturnBuy      float32
	Buy            float32
	Paid           float32
	CHEQUE         float32
	CHQUnderCollec float32
	Discount       float32
}

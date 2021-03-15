package model

type CashFlow struct {
	DocDate  string
	Income   float32
	Supplier float32
	Expensis float32
	Others   float32
	Bankin   float32
	Cheqout  float32
	Cheqin   float32
}

type CashFlowYearReq struct {
	Year      string
	AccSerial int
}

type CashFlowReq struct {
	DateFrom  string
	DateTo    string
	AccSerial int
}

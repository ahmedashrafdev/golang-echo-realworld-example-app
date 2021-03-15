package model

type Doc struct {
	DocNo int
}

type DocReq struct {
	DevNo     int
	TrSerial  int
	StoreCode int
}
type OpenDoc struct {
	DocNo        int
	StoreCode    int
	AccontSerial int
	TransSerial  int
	AccountName  string
	AccountCode  int
}
type OpenDocReq struct {
	DevNo    int
	TrSerial int
}

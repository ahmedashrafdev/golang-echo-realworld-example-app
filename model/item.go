package model

type GetItemRequest struct {
	BCode     string `json:"BCode" validate:"required"`
	StoreCode int    `json:"StoreCode" validate:"required"`
}

type Item struct {
	Serial        int
	ItemName      string
	MinorPerMajor int
	POSPP         float64
	POSTP         float64
	ByWeight      bool
}
type InsertItemReq struct {
	DNo        int
	TrS        int
	AccS       int
	ItmS       int
	Qnt        float32
	StCode     int
	StCode2    int
	InvNo      int
	ItmBarCode string
	DevNo      int
}

type DocItemsReq struct {
	DocNo     int
	TrSerial  int
	StoreCode int
	DevNo     int
}

type DocItem struct {
	Serial        int
	Qnt           float32
	Item_BarCode  int
	MinorPerMajor int
	ItemName      string
	ByWeight      bool
}

type DeleteItemReq struct {
	Serial int
}

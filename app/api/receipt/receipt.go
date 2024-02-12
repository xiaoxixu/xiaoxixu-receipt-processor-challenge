package receipt

type ReceiptProcessReq struct {
	Retailer     string           `form:"retailer" json:"retailer" binding:"required"`
	PurchaseDate string           `form:"purchaseDate" json:"purchaseDate" binding:"required"`
	PurchaseTime string           `form:"purchaseTime" json:"purchaseTime" binding:"required"`
	Items        []ReceiptItemReq `form:"items" json:"items" binding:"required"`
	Total        string           `form:"total" json:"total" binding:"required"`
}

type ReceiptItemReq struct {
	ShortDescription string `form:"shortDescription"`
	Price            string `form:"price"`
}

type ReceiptProcessResp struct {
	ID string `json:"id"`
}

type ReceiptPointResp struct {
	Points int64 `json:"points"`
}

package models

type Food struct {
	CreateDate int64  `xml:"create_date" json:"create_date"`
	UpdateDate int64  `xml:"update_date" json:"update_date"`
	ID         int    `xml:"id" json:"id"`
	Name       string `xml:"name" json:"name"`
	InitPrice  string `xml:"init_price" json:"init_price"`
	SalePrice  string `xml:"sale_price" json:"sale_price"`
	Title      string `xml:"title" json:"title"`
	Content    string `xml:"content" json:"content"`
	StoreID    int    `xml:"store_id" json:"store_id"`
	MenuID     int    `xml:"menu_id" json:"menu_id"`
	CategoryID int    `xml:"category_id" json:"category_id"`
	Status     int    `xml:"status" json:"status"`
	Image      string `xml:"image" json:"image"`
}

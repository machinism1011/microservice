package model

import "time"

type Order struct {
	ID          int64         `gorm:"primary_key;not_null;auto_increment" json:"id"`
	OrderCode   string        `gorm:"unique_index;not_null" json:"order_code"`
	PayStatus   int32         `json:"pay_status"`
	ShipStatus  int64         `json:"ship_status"`
	Price       float64       `json:"price"`
	OrderDetail []OrderDetail `gorm:"ForeignKey:OrderID" json:"order_detail"`
	CreateAt    time.Time
	UpdateAt    time.Time
}

type OrderDetail struct {
	ID            int64   `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductID     int64   `json:"product_id"`
	ProductNum    int64   `json:"product_num"`
	ProductSizeID int64   `json:"product_size_id"`
	ProductPrice  float64 `json:"product_price"`
	OrderID       int64   `json:"order_id"`
}

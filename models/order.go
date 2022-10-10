package models

import "time"

type Order struct {
	OrderID      int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null;type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"not null;type:timestamp" json:"orderedAt"`
	Items        []Item    `gorm:"constraint:onDelete:CASCADE" json:"items"`
}

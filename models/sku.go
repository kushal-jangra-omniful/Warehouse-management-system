package models

import (
	"time"

	"gorm.io/gorm"
)

type SKU struct {
	ID             string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ProductID      string         `gorm:"type:uuid;not null" json:"product_id"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Fragile        string         `gorm:"type:varchar(10);not null" json:"fragile"`
	Specifications string         `gorm:"type:jsonb" json:"specifications"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	gorm.Model
}

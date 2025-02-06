package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	ID                string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	HubID             string         `gorm:"type:uuid;not null" json:"hub_id"`
	TenantID          string         `gorm:"type:uuid;not null" json:"tenant_id"`
	SellerID          string         `gorm:"type:uuid;not null" json:"seller_id"`
	SKU_ID            string         `gorm:"type:uuid;not null" json:"sku_id"`
	AvailableQuantity int            `gorm:"type:int;default:0" json:"available_quantity"`
	LastUpdated       string         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"last_updated"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	gorm.Model
}

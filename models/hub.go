package models

import (
	"time"

	"gorm.io/gorm"
)

type Hub struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(60);not null" json:"name"`
	TenantID  string         `gorm:"type:uuid;not null" json:"tenant_id"`
	Address   string         `gorm:"type:varchar(200);not null" json:"address"`
	City      string         `gorm:"type:varchar(100);not null" json:"city"`
	State     string         `gorm:"type:varchar(100);not null" json:"state"`
	Country   string         `gorm:"type:varchar(100);not null" json:"country"`
	Pincode   string         `gorm:"type:varchar(10);not null" json:"pincode"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	gorm.Model
}

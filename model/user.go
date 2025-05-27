package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      Name      `gorm:"embedded" json:"name"`
	Address   Address   `gorm:"embedded" json:"address"`
	Active    bool      `json:"active"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdateBy  string    `json:"update_by"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteBy  string    `json:"delete_by"`
	DeleteAt  time.Time `json:"delete_at"`
}

type Name struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Address struct {
	Lane     string `json:"lane"`
	Village  string `json:"village"`
	City     string `json:"city"`
	District string `json:"district"`
	Pincode  int    `json:"pincode"`

	State string `json:"state"`
}

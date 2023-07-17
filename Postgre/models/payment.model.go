package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserId        uuid.UUID `gorm:"type:uuid;references:User.UserId" json:"userid" validate:"required"`
	PaymentMode   string    `json:"paymentmode" validate:"required"`
	SuccessStatus bool      `json:"isSuccess" validate:"required"`
	Created       time.Time `gorm:"autoUpdateTime" json:"createdAt"`
}

package models

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"userId" validate:"required"`
	UserName string    `json:"userName" validate:"required"`
}

//validate:"dive" for SLICE Validation

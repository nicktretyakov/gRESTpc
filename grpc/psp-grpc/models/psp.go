package models

import (
	"time"

	"github.com/google/uuid"
)

type Psp struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     int64
	BankID     int64
	Code       string `gorm:"type:varchar(250);not null"`
	Status     string `gorm:"type:varchar(250);not null"`
	BookedDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

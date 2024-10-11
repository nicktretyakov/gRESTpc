package requests

import (
	"time"

	"github.com/google/uuid"
)

type CreateBankRequest struct {
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}

type UpdateBankRequest struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}

type ListBankRequest struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}

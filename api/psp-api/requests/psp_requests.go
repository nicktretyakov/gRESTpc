package requests

import "time"

type CreatePspRequest struct {
	UserID     string `json:"user_id"`
	BankID     string `json:"bank_id"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	BookedDate time.Time
}

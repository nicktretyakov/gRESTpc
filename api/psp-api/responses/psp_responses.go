package responses

import "time"

type PspResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	BankID     string `json:"bank_id"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	BookedDate time.Time
}

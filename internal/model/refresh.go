package model

import (
	"time"

	"github.com/google/uuid"
)

// Always valid user request struct.
type ReqRFT struct {
	ID     uuid.UUID `json:"id" validate:"uuid"`
	UserID uuid.UUID `json:"user_id" validate:"uuid"`
	Hash   string    `json:"hash" validate:"sha512"`
}

type RefreshToken struct {
	ID         uuid.UUID //json:"id" validate:"uuid
	UserID     uuid.UUID //json:"user_id" validate:"uuid
	Hash       string    //json:"hash" validate:"sha512 // bcrypt hash
	Expiration time.Time //json "expiration"
}

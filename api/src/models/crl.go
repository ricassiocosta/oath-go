package models

import (
	"errors"
	"time"
)

var (
	errMissingSerial = errors.New("certificate serial is missing")
)

// CRL defines the certificate's data stored in the database
type CRL struct {
	ID        uint64    `json:"id,omitempty"`
	Serial    string    `json:"serial,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will validate and format the received user
func (u *CRL) Prepare() error {
	if u.Serial == "" {
		return errMissingSerial
	}

	return nil
}

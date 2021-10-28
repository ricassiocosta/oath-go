package models

import (
	"errors"
	"time"
)

var (
	errMissingSerial = errors.New("certificate serial is missing")
)

// CRLCert defines the certificate's data stored in the CRLCert
type CRLCert struct {
	ID        uint64    `json:"id,omitempty"`
	Serial    string    `json:"serial,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will validate and format the received user
func (u *CRLCert) Prepare() error {
	if u.Serial == "" {
		return errMissingSerial
	}

	return nil
}

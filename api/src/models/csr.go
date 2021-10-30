package models

import (
	"errors"
)

var (
	errMissingCSR = errors.New("certificate signing request is missing")
)

// CSR defines the certificate's CSR
type CSR struct {
	CSR string `json:"csr,omitempty"`
}

// Prepare will validate and format the data received from user
func (u *CSR) Prepare() error {
	if u.CSR == "" {
		return errMissingCSR
	}

	return nil
}

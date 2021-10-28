package repositories

import (
	"database/sql"
	"oath-go/src/models"
)

// CRL represents an CRL repository
type CRL struct {
	db *sql.DB
}

// NewCertRepository creates a CRL repository
func NewCertRepository(db *sql.DB) *CRL {
	return &CRL{db}
}

// GetCRL all certificate's serial in the CRL
func (u CRL) GetCRL() ([]models.CRLCert, error) {
	lines, err := u.db.Query(
		"SELECT * FROM cert_revocation_list",
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var crl []models.CRLCert

	for lines.Next() {
		var user models.CRLCert
		if err = lines.Scan(
			&user.ID,
			&user.Serial,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		crl = append(crl, user)
	}

	return crl, nil
}

// GetCRL all certificate's serial in the CRL
func (u CRL) AddCertToCRL(serial string) error {
	_, err := u.db.Query(
		`INSERT INTO users (serial) VALUES ($1)`,
		serial,
	)
	if err != nil {
		return err
	}
	defer u.db.Close()

	return nil
}
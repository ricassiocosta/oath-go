package repositories

import (
	"database/sql"
	"oath-go/src/models"
)

// CRL represents an CRL repository
type CRL struct {
	db *sql.DB
}

// NewCRLRepository creates a CRL repository
func NewCRLRepository(db *sql.DB) *CRL {
	return &CRL{db}
}

// GetCRL all certificate's serial in the CRL
func (u CRL) GetCRL() ([]models.CRL, error) {
	lines, err := u.db.Query(
		"SELECT * FROM cert_revocation_list",
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var crl []models.CRL

	for lines.Next() {
		var user models.CRL
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

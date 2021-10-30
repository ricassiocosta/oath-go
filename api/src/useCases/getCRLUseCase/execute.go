package getCRLUseCase

import (
	"net/http"
	"oath-go/src/models"
	"oath-go/src/repositories"
	"oath-go/src/utils/database"
	"oath-go/src/utils/responses"
)

// Execute perform all the steps to get the CRL
func Execute(w http.ResponseWriter) ([]models.CRLCert, error) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return nil, err
	}
	defer db.Close()

	repository := repositories.NewCertRepository(db)
	crl, err := repository.GetCRL()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return nil, err
	}

	return crl, nil
}
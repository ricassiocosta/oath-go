package addCertToCRLUseCase

import (
	"encoding/json"
	"net/http"
	"oath-go/src/models"
	"oath-go/src/repositories"
	"oath-go/src/utils/database"
	"oath-go/src/utils/responses"
)

// Execute perform all the steps to add a certificate into CRL
func Execute(w http.ResponseWriter, body []byte) error {
	var certificate models.CRLCert
	if err := json.Unmarshal(body, &certificate); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return err
	}

	if err := certificate.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return err
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return err
	}
	defer db.Close()

	repository := repositories.NewCertRepository(db)
	err = repository.AddCertToCRL(certificate.Serial)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return err
	}

	return nil
}
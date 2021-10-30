package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"oath-go/src/models"
	"oath-go/src/repositories"
	"oath-go/src/utils/database"
	"oath-go/src/utils/responses"
)

// GetCRL is responsible for returning the CRL list
func GetCRL(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCertRepository(db)
	crl, err := repository.GetCRL()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, crl)
}

// AddCertToCRL is responsible add a certificate serial into CRL
func AddCertToCRL(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var certificate models.CRLCert
	if err = json.Unmarshal(body, &certificate); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = certificate.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCertRepository(db)
	err = repository.AddCertToCRL(certificate.Serial)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

// SignCertificate is responsible to return a signed certificate
// for a given CSR
func SignCertificate(w http.ResponseWriter, r *http.Request) {
	
}
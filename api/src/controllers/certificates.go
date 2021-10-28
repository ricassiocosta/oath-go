package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"oath-go/src/database"
	"oath-go/src/models"
	"oath-go/src/repositories"
	"oath-go/src/responses"
)

// GetCRL is responsible for returning the CRL list
func GetCRL(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.CRL
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCRLRepository(db)
	crl, err := repository.GetCRL()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, crl)
}

// AddCertToCRL is responsible add a certificate serial into CRL
func AddCertToCRL(w http.ResponseWriter, r *http.Request) {
	
}

// GetTrustBundle is responsible return a signed certificate 
func GetTrustBundle(w http.ResponseWriter, r *http.Request) {
	
}
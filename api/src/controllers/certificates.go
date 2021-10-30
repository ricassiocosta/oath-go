package controllers

import (
	"io/ioutil"
	"net/http"
	"oath-go/src/useCases/addCertToCRLUseCase"
	"oath-go/src/useCases/getCRLUseCase"
	"oath-go/src/useCases/signCertificateUseCase"
	"oath-go/src/utils/responses"
)

// GetCRL is responsible for returning the CRL list
func GetCRL(w http.ResponseWriter, r *http.Request) {
	crl, err := getCRLUseCase.Execute(w)
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

	err = addCertToCRLUseCase.Execute(w, body)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

// SignCertificate is responsible to return a signed certificate
// for a given CSR
func SignCertificate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	signedCert, err:= signCertificateUseCase.Execute(w, body)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, string(signedCert))
}
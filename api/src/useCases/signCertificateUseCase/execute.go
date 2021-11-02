package signCertificateUseCase

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"oath-go/src/config"
	"oath-go/src/models"
	"oath-go/src/utils/responses"
	"strings"
	"time"
)

var (
	errDecodingPEMFile = errors.New("error while decoding a PEM file")
	errGettingCSRPubKey = errors.New("error while getting CSR Public Key")
)

type ca struct {
	cert *x509.Certificate
	key interface{}
}

// Execute perform all the steps to sign a certificate
func Execute(w http.ResponseWriter, body []byte) ([]byte, error) {
	var csr models.CSR
	if err := json.Unmarshal(body, &csr); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return nil, err
	}

	csr.CSR = strings.Replace(csr.CSR, `\n`, "\n", -1)

	if err := csr.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return nil, err
	}

	ca, err := setupCA()
	if err != nil {
		return nil, err
	}

	parsedCSR, err := setupCSR([]byte(csr.CSR))	
	if err != nil {
		return nil, err
	}

	template, err := setupCertTemplate(parsedCSR)
	if err != nil {
		return nil, err
	}

	csrKey, ok := parsedCSR.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errGettingCSRPubKey
	}

	signedCert, err := ca.signCert(template, csrKey)
	if err != nil {
		return nil, err
	}

	encodedCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: signedCert})

	return encodedCert, nil
}

func setupCertTemplate(csr *x509.CertificateRequest) (*x509.Certificate, error){
	template := &x509.Certificate{
		Subject: csr.Subject,
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(time.Hour * 24 * 365),
	}

	csrKey, err := x509.MarshalPKIXPublicKey(csr.PublicKey)
	if err != nil {
		return nil, err
	}

	serial := sha256.Sum256(csrKey)

	template.SerialNumber = new(big.Int).SetBytes(serial[:])

	return template, nil
}

func setupCSR(csr []byte) (*x509.CertificateRequest, error){
	decodedCSR, _ := pem.Decode([]byte(csr))
	if decodedCSR == nil {
		return nil, errDecodingPEMFile
	}

	parsedCSR, err := x509.ParseCertificateRequest(decodedCSR.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedCSR, nil
}

func setupCA() (*ca, error) {
	caCert, err := ioutil.ReadFile(config.RootCACertPath)
	if err != nil {
		return nil, err
	}

	caKey, err := ioutil.ReadFile(config.RootCAKeyPath)
	if err != nil {
		return nil, err
	}

	decodedCACert, _ := pem.Decode(caCert)
	if decodedCACert == nil {
		return nil, errDecodingPEMFile
	}

	parsedCaCert, err := x509.ParseCertificate(decodedCACert.Bytes)
	if err != nil {
		return nil, err
	}

	decodedCAKey, _ := pem.Decode(caKey)
	if decodedCAKey == nil {
		return nil, errDecodingPEMFile
	}

	parsedCaKey, err := x509.ParsePKCS8PrivateKey(decodedCAKey.Bytes)
	if err != nil {
		return nil, err
	}

	return &ca{
		cert: parsedCaCert,
		key: parsedCaKey,
	}, nil
}

func (ca *ca) signCert(template *x509.Certificate, csrKey *rsa.PublicKey) ([]byte, error) {
	return x509.CreateCertificate(rand.Reader, template, ca.cert, csrKey, ca.key)
}
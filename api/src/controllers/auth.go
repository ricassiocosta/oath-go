package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"oath-go/src/config"
	"oath-go/src/utils/responses"
	"strings"
)

var (
	errInvalidVerificationToken = errors.New("invalid verification token")
)

const (
	getAccessTokenURL = "https://github.com/login/oauth/access_token?client_id={id}&client_secret={secret}&code={code}"
	badVerificationToken = "bad_verification_code"
)

type Auth struct {
	Token string `json:"token,omitempty"`
}

// AuthGithub for get the token from github
func AuthGithub(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	requestToken := query["code"]

	url := strings.Replace(getAccessTokenURL, "{id}", config.GithubClientID, -1)
	url = strings.Replace(url, "{secret}", config.GithubClientSecret, -1)
	url = strings.Replace(url, "{code}", requestToken[0], -1)

	resp, err := http.Get(url)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	token := strings.SplitAfter(string(body), "=")
	token = strings.Split(token[1], "&")

	if token[0] == badVerificationToken {
		responses.Error(w, http.StatusUnauthorized, errInvalidVerificationToken)
		return
	}

	var auth Auth
	auth.Token = token[0]

	responses.JSON(w, http.StatusCreated, auth)
}
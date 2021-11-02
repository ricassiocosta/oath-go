package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"oath-go/src/config"
	"oath-go/src/utils/authentication"
	"oath-go/src/utils/responses"
	"strings"
	"time"
)

var (
	errInvalidVerificationToken = errors.New("invalid verification token")
)

const (
	getAccessTokenURL    = "https://github.com/login/oauth/access_token?client_id={id}&client_secret={secret}&code={code}"
	getUserURL           = "https://api.github.com/user"
	badVerificationToken = "bad_verification_code"
)

type AuthData struct {
	Token string `json:"token,omitempty"`
}

// GithubCallback for get the token from github
func GithubCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	var auth AuthData
	auth.Token = token[0]

	responses.JSON(w, http.StatusAccepted, auth)
}

type User struct {
	ID int `json:"id,omitempty"`
}

// Auth is responsible to handle the user login
func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := r.URL.Query()
	githubToken := query["githubToken"]

	fmt.Println(string(githubToken[0]))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, getUserURL, nil)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	req.Header.Set("Authorization", "token " + githubToken[0])
	resp, err := client.Do(req)
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

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	token, err := authentication.GenerateToken(uint64(user.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var auth AuthData
	auth.Token = token

	responses.JSON(w, http.StatusAccepted, auth)
}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	htmlIndex = `<html><body>
<a href="/GoogleLogin">Log in with Google</a>
</body></html>
`
	clientID     = `107767092716-p6288snfat8ui9o9jnmd7oi0kcmiml3i.apps.googleusercontent.com`
	clientSecret = `BVxnHWnBrTyhzQck091J2Hey`

	oauthStateString = "random" //state
)

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  "http://localhost:8000/GoogleCallback",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: endpoint,
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlIndex)
	w.Write([]byte("Welcome to myOauth2~~!!!"))
	log.Println("welcome myOauth2")
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	log.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println("state is=======>", state)

	code := r.FormValue("code")
	log.Println("code is=======>", code)
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	log.Println("token is=======>", token)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Get token failed with %#v\n", err)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content========> %s\n", contents)
}

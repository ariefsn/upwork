package env

import (
	"net/http"
	u "net/url"
)

type OAuthGithub struct {
	ClientId     string
	ClientSecret string
}

func (o *OAuthGithub) IsDisabled() bool {
	return o.ClientId == "" || o.ClientSecret == ""
}

func (o *OAuthGithub) UrlIdentity() (method, url string) {
	urlParsed, _ := u.Parse("https://github.com/login/oauth/authorize")
	q := urlParsed.Query()
	q.Add("client_id", o.ClientId)
	q.Add("scope", "user")
	urlParsed.RawQuery = q.Encode()
	method = http.MethodGet
	url = urlParsed.String()
	return
}

func (o *OAuthGithub) UrlAuthorize(code string) (method, url string) {
	urlParsed, _ := u.Parse("https://github.com/login/oauth/access_token")
	q := urlParsed.Query()
	q.Add("client_id", o.ClientId)
	q.Add("client_secret", o.ClientSecret)
	q.Add("code", code)
	urlParsed.RawQuery = q.Encode()
	method = http.MethodPost
	url = urlParsed.String()
	return
}

func (o *OAuthGithub) UrlProfile() (method, url string) {
	urlParsed, _ := u.Parse("https://api.github.com/user")
	method = http.MethodGet
	url = urlParsed.String()
	return
}

type OAuth struct {
	Github OAuthGithub
}

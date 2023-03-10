package pkg

import "time"

type Account struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
	AccountID    int       `json:"account_id"`
	Integration  Integration
}

type Integration struct {
	SecretKey        string `json:"secret_key"`
	ClientID         int    `json:"client_id"`
	RedirectURL      string `json:"redirect_url"`
	AuthorizationURL string `json:"authorization_url"`
}

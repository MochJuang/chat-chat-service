package model

import "time"

type AuthenticationResponse struct {
	AccessToken      string    `json:"access_token"`
	RefreshToken     string    `json:"refresh_token"`
	ExpiresIn        time.Time `json:"expires_in"`
	RefreshExpiresIn time.Time `json:"refresh_expires_in"`
}

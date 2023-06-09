package authentication

import (
	"fmt"
	"time"
)

type Passport struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func GeneratePassport(atDuration time.Duration, rtDuration time.Duration, userId string) (*Passport, error) {
	if userId == "" {
		return nil, fmt.Errorf("userId cannot be empty")
	}
	accessToken, err := GenerateJwtToken(atDuration, userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateJwtToken(rtDuration, userId)
	if err != nil {
		return nil, err
	}
	pp := Passport{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &pp, nil
}

func VerifyPassport(pp Passport) (*JwtCustomClaims, *JwtCustomClaims, error) {
	aToken, err := VerifyJwtToken(pp.AccessToken)
	if err != nil {
		return nil, nil, err
	}
	rToken, err := VerifyJwtToken(pp.RefreshToken)
	if err != nil {
		return nil, nil, err
	}
	return aToken, rToken, nil
}

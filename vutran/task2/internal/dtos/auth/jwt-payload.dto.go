package dto

type JwtPayloadDto struct {
	Sub     int64
	Options interface{}
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenClaims struct {
	UserID   string
	Email    string
	ClientID string
}

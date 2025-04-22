package dto

type JwtPayloadDto struct {
	Sub     int64
	Options interface{}
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

package dto

type RegisterClientRequestDto struct {
	Name         string   `json:"name" binding:"required"`
	RedirectURIs []string `json:"redirect_uris" binding:"required"`
}

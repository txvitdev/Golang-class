package dto

type RegisterClientRequestDto struct {
	Name          string   `json:"name" binding:"required"`
	RedirectURIs  []string `json:"redirect_uris" binding:"required"`
	AdminEmail    string   `json:"admin_email" binding:"required,email"`
	AdminPassword string   `json:"admin_password" binding:"required,min=6"`
	AdminFullName string   `json:"admin_full_name"`
}

package dataoject

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

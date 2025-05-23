package responses

type TokenResponse struct {
	LoginResponse
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

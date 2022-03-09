package auth

// AuthenticationRequest represents the request body structure for the authentication endpoint.
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthenticationResponse represents the response body structure for the authentication endpoint.
type AuthenticationResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

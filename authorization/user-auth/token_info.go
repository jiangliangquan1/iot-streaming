package userauth

type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int32  `json:"expired_in"`
}

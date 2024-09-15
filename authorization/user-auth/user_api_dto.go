package userauth

type UserSignUpRequest struct {
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignUpResult struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Nick  string `json:"nick"`
	Email string `json:"email"`
}

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int32  `json:"expired_in"`
}

type UserRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

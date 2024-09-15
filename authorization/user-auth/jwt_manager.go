package userauth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const id = "iot-streaming"
const accessTokenExpiredIn = 3000000
const refreshTokenExpiredIn = 3600

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

type JwtManager struct {
}

func (j *JwtManager) GenerateToken(userid int64, username string, createRefreshToken bool) (*TokenInfo, error) {
	accessToken, err := j.generate(userid, username, accessTokenExpiredIn)
	if err != nil {
		return nil, err
	}

	refreshToken, err1 := j.generate(userid, username, refreshTokenExpiredIn)
	if err1 != nil {
		return nil, err1
	}

	return &TokenInfo{AccessToken: accessToken, RefreshToken: refreshToken, ExpiredIn: accessTokenExpiredIn}, nil

}

func (*JwtManager) generate(userid int64, username string, expiredIn int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expiredIn) * time.Second)
	claims := Claims{
		ID:       userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    id,
			IssuedAt:  nowTime.Unix(),
			Subject:   string(userid),
			Id:        id,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(id))
	return token, err
}

func (*JwtManager) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(id), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func NewJwtManager() *JwtManager {
	return &JwtManager{}
}

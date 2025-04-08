package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinho-yoo-jack/gin-tonic-tester/config"
	"time"
)

type JwtUtils struct {
	secretKey       string
	accessTokenExp  string
	refreshTokenExp string
}

func NewJwtUtils(cfg *config.Config) *JwtUtils {
	return &JwtUtils{cfg.SecretKey, cfg.AccessTokenExp, cfg.RefreshTokenExp}
}

func (j *JwtUtils) GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(j.getAccessTokenExp()).Unix(), // 만료 시간
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtUtils) getAccessTokenExp() time.Duration {
	dur, err := j.convertExpStringToTime(j.accessTokenExp)
	if err != nil {
		return dur
	}
	return time.Minute * 15
}

func (j *JwtUtils) convertExpStringToTime(exp string) (time.Duration, error) {
	return time.ParseDuration(exp)
}

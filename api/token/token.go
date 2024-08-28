package token

import (
	"errors"
	"fmt"
	"time"

	pb "auth-athlevo/genproto/auth"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const signingKey = "secret_key"

func GenerateJWTToken(user *pb.User) (*pb.LoginRes, string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(180 * time.Minute).Unix()

	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, "", fmt.Errorf("error generating access token: %w", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	rftClaims["email"] = user.Email
	rftClaims["role"] = user.Role
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(48 * time.Hour).Unix()

	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, "", fmt.Errorf("error generating refresh token: %w", err)
	}

	res := &pb.LoginRes{
		Token:     access,
		ExpiresAt: time.Now().Add(180 * time.Minute).Format("2006-01-02 15:04:05"),
	}

	return res, refresh, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

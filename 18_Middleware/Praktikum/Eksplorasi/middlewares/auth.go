package middlewares

import (
	"errors"
	"net/http"
	"rest/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			if err != nil {
				return c.JSON(http.StatusUnauthorized, models.BaseResponse{
					Status:  false,
					Message: err.Error(),
				})
			}

			return nil
		},
	}
}

func (jwtConfig *JWTConfig) GenerateToken(userID, userRole string) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuration))))

	claims := &JWTCustomClaims{
		userID,
		userRole,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(c echo.Context) (*JWTCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, errors.New("invalid token")
	}

	claims := user.Claims.(*JWTCustomClaims)

	return claims, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := GetUser(c)

		isInvalid := data == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, models.BaseResponse{
				Status:  false,
				Message: "invalid token",
			})
		}

		if data.UserRole != "admin" {
			return c.JSON(http.StatusForbidden, models.BaseResponse{
				Status:  false,
				Message: "you have no permission to access this endpoint",
			})
		}

		return next(c)
	}
}

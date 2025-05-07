package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	gommonlog "github.com/labstack/gommon/log"
)

var (
	ErrHttpGenericMessage = echo.NewHTTPError(http.StatusInternalServerError, "something went wrong, please try again later")
	ErrWrongCredentials   = echo.NewHTTPError(http.StatusUnauthorized, "username or password is invalid")
	jwtSecret             = "foo"
)

func main() {
	hostport := ":" + os.Getenv("AUTH_API_PORT")
	userAPIAddress := os.Getenv("USERS_API_ADDRESS")

	envJwtSecret := os.Getenv("JWT_SECRET")
	if len(envJwtSecret) != 0 {
		jwtSecret = envJwtSecret
	}

	userService := UserService{
		Client:         http.DefaultClient,
		UserAPIAddress: userAPIAddress,
		AllowedUserHashes: map[string]interface{}{
			"admin_admin": nil,
			"johnd_foo":   nil,
			"janed_ddd":   nil,
		},
	}

	e := echo.New()
	e.Logger.SetLevel(gommonlog.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "Auth API, written in Go\n")
	})

	e.POST("/login", getLoginHandler(userService))

	e.Logger.Fatal(e.Start(hostport))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getLoginHandler(userService UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestData := LoginRequest{}
		if err := json.NewDecoder(c.Request().Body).Decode(&requestData); err != nil {
			log.Printf("could not read credentials from POST body: %s", err.Error())
			return ErrHttpGenericMessage
		}

		ctx := c.Request().Context()
		user, err := userService.Login(ctx, requestData.Username, requestData.Password)
		if err != nil {
			if err != ErrWrongCredentials {
				log.Printf("could not authorize user '%s': %s", requestData.Username, err.Error())
				return ErrHttpGenericMessage
			}
			return ErrWrongCredentials
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["firstname"] = user.FirstName
		claims["lastname"] = user.LastName
		claims["role"] = user.Role
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			log.Printf("could not generate a JWT token: %s", err.Error())
			return ErrHttpGenericMessage
		}

		return c.JSON(http.StatusOK, map[string]string{
			"accessToken": t,
		})
	}
}

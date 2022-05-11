package middlewares

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Config struct{
	// when it returns true, our middleware is skipped
	Filter func(c *fiber.Ctx) bool // required

	// function to  run when there's error in decoding jwt
	Unauthorized fiber.Handler // middleware specific

	// function to decode our jwt token
	Decode func (c *fiber.Ctx) (*jwt.MapClaims, error) // middleware specific

	// set jwt secret
	Secret string // middleware specific

	// set jwt expiry in seconds
	Expiry int64 // middleware specific
}

var ConfigDefault = Config{
	Filter: nil,
	Decode: nil,
	Unauthorized: nil,
	Secret: "secret",
	Expiry: 60,
} 

/*
Function for generating default Config
*/
func configDefault(config ...Config) Config {
	// returns default config, if not provided with one,
	if len(config) < 1 {
		return ConfigDefault
	}

	// overrride default config
	cfg := config[0]

	// Set default values if not passed
	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	// sef default secret if not passed
	if cfg.Secret == "" {
		cfg.Secret = ConfigDefault.Secret
	}

	// Set default expiry if not passed
	if cfg.Expiry == 0 {
		cfg.Expiry = ConfigDefault.Expiry
	}

	if cfg.Decode == nil {
		// Set default decode function if not passed
		cfg.Decode =  func(c *fiber.Ctx) (*jwt.MapClaims, error) {
			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("authorization header is required")
			}

			// parse our jwt token and check for validity
			token, err := jwt.Parse(
				authHeader[7:],
				func(token *jwt.Token) (interface{}, error) {
					// verifying our algo
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["algo"])
					}

					return []byte(cfg.Secret), nil
				},
			)

			if err != nil {
				return nil, errors.New("error parsing token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !(ok && token.Valid){
				return nil, errors.New("invalid Token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, errors.New("jwt token expired")
			}

			return &claims, nil
		}
	}

	// set default unauthorized, if not passed
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

/*
Function to generate jwt token
*/
func Encode(claims *jwt.MapClaims, expiryAfter int64) (string, error){
	// setting default expiryAfter
	if expiryAfter == 0 {
		expiryAfter = ConfigDefault.Expiry
	}

	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter
	token  := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// our signed jwt token string
	signedToken, err := token.SignedString([]byte(ConfigDefault.Secret))
	if err != nil {
		return "", errors.New("error in creating token")
	}
	return signedToken, nil
}

/*
Our main middleware function used to initialize the our middleware
*/
func New(config Config) fiber.Handler {
	// For setting default config
	cfg := configDefault(config)

	return func (c *fiber.Ctx) error {
		// Don't execute middleware if filter returns true
		if cfg.Filter != nil && cfg.Filter(c) {
			fmt.Println("Middle was skipped")
			return c.Next()
		}
		fmt.Println("Middle was run")

		claims, err := cfg.Decode(c)
		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
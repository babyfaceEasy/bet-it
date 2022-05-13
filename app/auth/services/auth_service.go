package services

import (
	"elivate9ja-go/app/admin/services"
	"elivate9ja-go/middlewares"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	adminService services.IAdminService
}

func NewAuthService(adminService services.IAdminService) *AuthService {
	return &AuthService{adminService}
}

func (as *AuthService) LogCustomerIn(email, password string) error {
	return nil
}

func (as *AuthService) LogAdminIn(email, password string) error {

	response, err := as.adminService.VerifyAdmin(email, password)
	if err != nil {
		return errors.New("this user does not exist")
	}
	if !response {
		return errors.New("this user does not exist")
	}

	adminUser, err := as.adminService.GetAdminByEmail(email)
	if err != nil {
		// Read the type of error and know what to do next
		return errors.New("admin with this email does not exist")
	}

	adminClaim := &jwt.MapClaims{"email": adminUser.Email, "isAdmin": true}
	token, err := middlewares.Encode(adminClaim, 1000)
	if err != nil {
		// TODO: can log erorr here
		// TODO: depending on the kind of error returned
		return errors.New("an error occurred pleas try again later.")
	}

	fmt.Println(token)
	return nil
}

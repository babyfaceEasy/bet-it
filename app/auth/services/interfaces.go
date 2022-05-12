package services

type IAuthService interface {
	LogCustomerIn(email, password string) error
	LogAdminIn(email, password string) error
}
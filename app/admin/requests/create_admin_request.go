package requests

type CreateAdminRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	SuperAdmin bool `json:"super_admin"`
}
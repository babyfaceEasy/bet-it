package requests

type LogAdminInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

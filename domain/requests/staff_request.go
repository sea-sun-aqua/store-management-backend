package requests

type StaffRegisterRequest struct {
	Name     string `json:"staff_name"`
	Email    string `json:"staff_email"`
	Password string `json:"staff_password"`
}

type StaffLoginRequest struct {
	Email    string `json:"staff_email"`
	Password string `json:"staff_password"`
}
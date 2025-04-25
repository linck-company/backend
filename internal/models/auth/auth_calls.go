package authmodels

type UserAuthLoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
	Ip         string
}

type UserAuthLogoutRequest struct {
	Token string
}

type UserAuthUpdateMobileNumberRequest struct {
	Username    string
	PhoneNumber string `json:"contact_number"`
}

type UserAuthChangePasswordRequest struct {
	Username           string
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type GetUserRequest struct {
	UserId string
}

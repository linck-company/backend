package authmodels

import (
	genricresponses "backendv1/internal/models/generic_responses"
	"net/http"
)

type UserAuthLoginResponse struct {
	StatusCode int    `json:"status_code"`
	JwtToken   string `json:"jwt_token,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	Message    string `json:"message,omitempty"`
}

type UserAuthChangePasswordResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GetUserDetailsResponse struct {
	StatusCode      int     `json:"status_code"`
	FirstName       *string `json:"first_name"`
	LastName        *string `json:"last_name"`
	Email           *string `json:"email"`
	Mobile_number   *string `json:"mobile_number"`
	ProfilePhotoUrl *string `json:"profile_photo_url"`
}

var LogoutSuccessfulResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusOK,
	Message:    "Logout Successful",
}

var JwtValidResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusOK,
	Message:    "Valid",
}

var JwtInvalidResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Not valid",
}

var ChangePasswordOldPasswordFailure = genricresponses.GenericResponse{
	StatusCode: http.StatusFailedDependency,
	Message:    "Incorrect Old Password",
}

var ConfirmPasswordDoesNotMatchResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusFailedDependency,
	Message:    "New Password and Confirm Password do not match",
}

var AuthFailureResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Invalid Username or Password",
}

var MaxPasswordLimitAccountBlockedResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusForbidden,
	Message:    "Max Password Retries Limit Reached. Your account has been blocked. Contact your faculty advisor or Club Heads to activate your account",
}

var InactiveOrBlockedResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusUnauthorized,
	Message:    "Your account is inactive/blocked. Contact your faculty advisor or Club Heads to activate your account",
}

var UpdateMobileNumberSuccessResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusOK,
	Message:    "Mobile Number updates successfully",
}

var UpdateMobileNumberFailureResponse = genricresponses.GenericResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "Couldn't update phone number",
}

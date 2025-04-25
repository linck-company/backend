package db

import (
	authmodels "backendv1/internal/models/auth"
	"context"
)

type AuthDB interface {
	Database
	InitAuthDbSchema()
	LoginUser(ctx context.Context, user *authmodels.UserAuthLoginRequest) interface{}
	ValidateJWT(ctx context.Context, token string) interface{}
	LogoutUser(ctx context.Context, user *authmodels.UserAuthLogoutRequest) interface{}
	ChangePassword(ctx context.Context, user *authmodels.UserAuthChangePasswordRequest) interface{}
	UpdateMobileNumber(ctx context.Context, user *authmodels.UserAuthUpdateMobileNumberRequest) interface{}
	GetUserDetails(ctx context.Context, user *authmodels.GetUserRequest) interface{}
	//ResetPassword(ctx context.Context, user *callmodels.ResetPassword) interface{}
	//UpdatePassword(ctx context.Context, user *callmodels.UpdatePassword) interface{}
	//CreateUser(ctx context.Context, user *callmodels.UserAuthCreateUser) interface{}
	//UpdateUserDetails(ctx context.Context, user *callmodels.UpdateUserDetails) interface{}
	//CheckIFUserExists(ctx context.Context, user *callmodels.CheckIfUserExists) interface{}
}

package authpg

import (
	authmodels "backendv1/internal/models/auth"
	genricresponses "backendv1/internal/models/generic_responses"
	"backendv1/pkg/errcheck"
	"backendv1/pkg/utils"
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func (adb *AuthPostgres) incrementPasswordAttempts(username string) {
	query := "UPDATE users SET password_tries = password_tries + 1 WHERE username = $1"
	_, err := adb.Pool.Exec(context.Background(), query, username)
	errcheck.LogIfError(err, fmt.Sprintf("Error incrementing password attempts for user %s", username))
}

func (adb *AuthPostgres) resetPasswordAttempts(username string) {
	query := "UPDATE users SET password_tries = 0 WHERE username = $1"
	_, err := adb.Pool.Exec(context.Background(), query, username)
	errcheck.LogIfError(err, fmt.Sprintf("Error resetting password attempts for user %s", username))
}

func (adb *AuthPostgres) changeAccountStatus(status, username string) {
	query := "UPDATE users SET account_status = $1 WHERE username = $2"
	_, err := adb.Pool.Exec(context.Background(), query, status, username)
	errcheck.LogIfError(err, fmt.Sprintf("Error changing account status for user %s to %s", username, status))
}

func (adb *AuthPostgres) updateUserPassword(ctx context.Context, username, password string) error {
	// query := "UPDATE users SET password = $1 WHERE id = $2 AND username = $3"
	query := "UPDATE users SET password = $1 WHERE username = $2"
	_, err := adb.Pool.Exec(ctx, query, utils.BcryptEncode(utils.Base64Decode(password)), username)
	go errcheck.LogIfError(err, fmt.Sprintf("Failed to update password for user %s", username))
	return err
}

func (adb *AuthPostgres) setJwtTokenAsInactive(ctx context.Context, token string) error {
	query := "UPDATE auth_tokens SET is_active = $1 WHERE token = $2"
	_, err := adb.Pool.Exec(ctx, query, false, token)
	return err
}

func (adb *AuthPostgres) addNewJWT(ctx context.Context, jid, token, uid, ip string) error {
	query := "INSERT INTO auth_tokens (id, token, user_id, ip) VALUES ($1, $2, $3, $4)"
	_, err := adb.Pool.Exec(ctx, query, utils.GetNewULID(), token, uid, ip)
	return err
}

func (adb *AuthPostgres) generateLoginUserResponse(qr *authmodels.UserAuthLogin, u *authmodels.UserAuthLoginRequest) interface{} {
	if qr.AccountStatus == adb.BlockedStatus || qr.AccountStatus == adb.InactiveStatus {
		log.Println("Account status is blocked")
		return authmodels.InactiveOrBlockedResponse
	}
	if !utils.BcryptCompare(utils.Base64Decode(u.Password), qr.Password) {
		go adb.incrementPasswordAttempts(u.Username)
		if qr.PasswordRetries >= 5 {
			go adb.changeAccountStatus(adb.BlockedStatus, u.Username)
			return authmodels.MaxPasswordLimitAccountBlockedResponse
		}
		return authmodels.AuthFailureResponse
	}

	if qr.PasswordRetries != 0 {
		go adb.resetPasswordAttempts(u.Username)
	}
	jid := utils.GetNewULID()
	jwt := adb.Jwt.GenerateJWT(qr.Id, qr.Username, qr.UserRole, jid, u.RememberMe)
	adb.addNewJWT(context.Background(), jid, jwt, qr.Id, u.Ip)
	return authmodels.UserAuthLoginResponse{
		UserId:     qr.Id,
		StatusCode: http.StatusOK,
		JwtToken:   jwt,
		Message:    "Login Successful",
	}
}

func (adb *AuthPostgres) validateAndUpdatePassword(ctx context.Context, q *authmodels.UserAuthChangePassword, u *authmodels.UserAuthChangePasswordRequest) interface{} {
	if q.AccountStatus == adb.BlockedStatus || q.AccountStatus == adb.InactiveStatus {
		log.Println("Account status is blocked")
		return authmodels.InactiveOrBlockedResponse
	}

	if !utils.BcryptCompare(utils.Base64Decode(u.OldPassword), q.Password) {
		return authmodels.ChangePasswordOldPasswordFailure
	}

	// if adb.updateUserPassword(ctx, u.Id, u.Username, u.NewPassword) != nil {
	if adb.updateUserPassword(ctx, u.Username, u.NewPassword) != nil {
		return genricresponses.GenericInternalServerErrorResponse
	}

	return authmodels.UserAuthChangePasswordResponse{
		StatusCode: http.StatusOK,
		Message:    "Password Updates Successfully",
	}
}

package authpg

import (
	"backendv1/internal/jwt"
	authmodels "backendv1/internal/models/auth"
	genricresponses "backendv1/internal/models/generic_responses"
	"backendv1/pkg/errcheck"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type AuthPostgres struct {
	Pool *pgxpool.Pool

	ActiveStatus   string
	InvitedStatus  string
	BlockedStatus  string
	InactiveStatus string

	Jwt *jwt.JWT
}

func (adb *AuthPostgres) InitAuthDbSchema() {
	pwd, err := os.Getwd()
	errcheck.FatalIfError(err, "Error getting working directory")
	authSchemaPath := filepath.Join(pwd, "migrations", "schema-v1", "auth.sql")
	authSchemaFile, err := os.ReadFile(authSchemaPath)
	errcheck.FatalIfError(err, "Error reading auth schema file")
	authSchema := string(authSchemaFile)
	authSchemaExecutables := strings.Split(authSchema, ";")
	log.Println("Executing auth schema initialization")
	for _, stmt := range authSchemaExecutables {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		log.Printf("Executing statement: %s\n", strings.Split(stmt, "\n")[0])
		_, err := adb.Pool.Exec(context.Background(), stmt)
		errcheck.LogIfError(err, "Error executing statement")
	}
	log.Println("Auth database schema initialized successfully")
}

func (adb *AuthPostgres) Close() error {
	log.Println("Closing auth database connection")
	adb.Pool.Close()
	return nil
}

func (adb *AuthPostgres) Ping() error {
	return adb.Pool.Ping(context.Background())
}

func (adb *AuthPostgres) LoginUser(ctx context.Context, user *authmodels.UserAuthLoginRequest) interface{} {
	if user == nil || user.Username == "" || user.Password == "" {
		return genricresponses.GenericBadRequestResponse
	}
	query := "SELECT id, password, password_tries, user_type, account_status FROM users WHERE username = $1"
	qr := authmodels.UserAuthLogin{}
	err := adb.Pool.QueryRow(ctx, query, user.Username).Scan(&qr.Id, &qr.Password, &qr.PasswordRetries, &qr.UserRole, &qr.AccountStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authmodels.AuthFailureResponse
		}
		log.Printf("Database query error in LoginUser for user %s: %v", user.Username, err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	qr.Username = user.Username
	return adb.generateLoginUserResponse(&qr, user)
}

func (adb *AuthPostgres) ValidateJWT(ctx context.Context, token string) interface{} {
	query := "SELECT is_active FROM auth_tokens WHERE token = $1"
	var isActive bool
	err := adb.Pool.QueryRow(ctx, query, token).Scan(&isActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		log.Printf("Database query error in ValidateJWT for token %s: %v", token, err)
		return false
	}
	c, err := adb.Jwt.ParseJWT(token)
	if err != nil {
		return false
	}
	if adb.Jwt.IsJWTValid(c) && isActive {
		return authmodels.JwtValidResponse
	}
	return authmodels.JwtInvalidResponse
}

func (adb *AuthPostgres) LogoutUser(ctx context.Context, user *authmodels.UserAuthLogoutRequest) interface{} {
	err := adb.setJwtTokenAsInactive(ctx, user.Token)
	if err != nil {
		log.Printf("Database query error in LogoutUser for user: %v", err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return authmodels.LogoutSuccessfulResponse
}

func (adb *AuthPostgres) ChangePassword(ctx context.Context, user *authmodels.UserAuthChangePasswordRequest) interface{} {
	if user == nil || user.Username == "" || user.OldPassword == "" || user.NewPassword == "" {
		fmt.Println("User not found")
		return genricresponses.GenericBadRequestResponse
	}
	query := "SELECT password, account_status FROM users WHERE username = $1"
	qr := authmodels.UserAuthChangePassword{}

	err := adb.Pool.QueryRow(ctx, query, user.Username).Scan(&qr.Password, &qr.AccountStatus)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows")
			return genricresponses.GenericBadRequestResponse
		}
		log.Printf("Database query error in ChangePassword for user %s: %v", user.Username, err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	return adb.validateAndUpdatePassword(ctx, &qr, user)
}

func (adb *AuthPostgres) UpdateMobileNumber(ctx context.Context, user *authmodels.UserAuthUpdateMobileNumberRequest) interface{} {
	query := "UPDATE users SET contact_number = $1 WHERE username = $2"

	if _, err := adb.Pool.Exec(ctx, query, user.PhoneNumber, user.Username); err != nil {
		log.Printf("Database query error in UpdateMobileNumber for user %s: %v", user.Username, err)
		return authmodels.UpdateMobileNumberFailureResponse
	}
	return authmodels.UpdateMobileNumberSuccessResponse
}

func (adb *AuthPostgres) GetUserDetails(ctx context.Context, user *authmodels.GetUserRequest) interface{} {
	query := `SELECT first_name, last_name, email, contact_number FROM users WHERE id = $1`
	var response authmodels.GetUserDetailsResponse

	err := adb.Pool.QueryRow(ctx, query, user.UserId).Scan(
		&response.FirstName, &response.LastName, &response.Email, &response.Mobile_number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return genricresponses.GenericInternalServerErrorResponse
		}
		log.Printf("Database query error in GetUserDetails for user %s: %v", user.UserId, err)
		return genricresponses.GenericInternalServerErrorResponse
	}
	response.StatusCode = http.StatusOK
	return response
}

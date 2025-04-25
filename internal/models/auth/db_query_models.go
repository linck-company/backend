package authmodels

type UserAuthLogin struct {
	Id              string
	Username        string
	Password        string
	AccountStatus   string
	UserRole        string
	PasswordRetries int
}

type UserAuthChangePassword struct {
	Username      string
	Password      string
	AccountStatus string
}

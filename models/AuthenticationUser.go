package models

type AuthenticationUser struct {
	Username string
	Password string
}

func CreateAuthUserObject(username string, password string) AuthenticationUser {
	newAuthUser := AuthenticationUser{
		Username: username,
		Password: password,
	}

	return newAuthUser
}

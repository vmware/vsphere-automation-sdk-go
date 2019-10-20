package refreshtoken

type Error struct {
	ErrorDetails error
}

func (basicAuthError *Error) Error() string {
	return "Basic Auth Error: UserName and Password are missing."
}

type UsernameError struct{}

func (userNameError *UsernameError) Error() string {
	return "Basic Auth Error: Username is not available."
}

type PasswordError struct{}

func (passwordError *PasswordError) Error() string {
	return "Basic Auth Error: Password is not available."
}

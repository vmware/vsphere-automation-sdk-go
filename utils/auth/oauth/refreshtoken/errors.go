package refreshtoken

type Error struct {
	ErrorDetails error
}

func (basicAuthError *Error) Error() string {
	return Name + " Error: CSPURL and RefreshToken are missing."
}

type CSPURLError struct{}

func (userNameError *CSPURLError) Error() string {
	return Name + " Error: CSPURL is not available."
}

type RefreshTokenError struct{}

func (passwordError *RefreshTokenError) Error() string {
	return Name + " Error: RefreshToken is not available."
}

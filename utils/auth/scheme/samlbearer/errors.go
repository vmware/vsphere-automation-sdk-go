package samlbearer

type Error struct {
}

func (samlBearerAuthError *Error) Error() string {
	return "SAML Bearer Auth Error: saml-bearer-token is missing."
}

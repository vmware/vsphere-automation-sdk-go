package scheme

type Scheme string

const (
	NoAuth            Scheme = "NoAuth"
	BasicAuth         Scheme = "BasicAuth"
	SAMLBearer        Scheme = "SAMLBearer"
	OAuthRefreshToken Scheme = "OAuth.RefreshToken"
)

var All []Scheme = []Scheme{
	NoAuth,
	BasicAuth,
	SAMLBearer,
	OAuthRefreshToken,
}

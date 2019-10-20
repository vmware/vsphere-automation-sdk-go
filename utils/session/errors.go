package session

type ArchTypeError struct {
	archType string
}

func (err *ArchTypeError) Error() string {
	return "Session Manager Error: Invalid Session Architecture Type " + err.archType
}

type URLError struct{}

func (err *URLError) Error() string {
	return "Session Manager Error: Invalid URL"
}

type AuthDetailsError struct{}

func (err *AuthDetailsError) Error() string {
	return "Session Manager Error: Invalid Authentication Details"
}

package host

type ArchTypeError struct {
	archType string
}

func (err *ArchTypeError) Error() string {
	return "Session Manager Error: Invalid Session Architecture Type " + err.archType
}

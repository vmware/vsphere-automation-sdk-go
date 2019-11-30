package arch

type archType string

func (at archType) isType() archType {
	return at
}

func (at archType) String() string {
	return string(at)
}

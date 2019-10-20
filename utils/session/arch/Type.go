package arch

type Type interface {
	isType() archType
	String() string
}

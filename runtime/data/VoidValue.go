package data

type VoidValue struct {
}

func NewVoidValue() *VoidValue {
	return &VoidValue{}
}

func (voidValue *VoidValue) Value() DataValue {
	return nil
}

func (voidValue *VoidValue) Type() DataType {
	return VOID
}

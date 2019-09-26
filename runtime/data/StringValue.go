package data

type StringValue struct {
	value string
}

func NewStringValue(value string) *StringValue {
	var stringValue = &StringValue{value: value}
	return stringValue
}

func (stringValue *StringValue) Type() DataType {
	return STRING
}

func (stringValue *StringValue) Value() string {
	return (stringValue.value)
}

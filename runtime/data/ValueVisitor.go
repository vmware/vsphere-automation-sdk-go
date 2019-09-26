package data

type ValueVisitor interface {
	VisitString(value *StringValue)
	VisitStructure(value *StructValue)
	VisitOptional(value *OptionalValue)
	VisitError(value *ErrorValue)
	VisitList(value *ListValue)
	VisitInteger(value *IntegerValue)
	VisitDouble(value *DoubleValue)
	VisitVoid(value *VoidValue)
	VisitBoolean(value *BooleanValue)
	VisitSecret(value *SecretValue)
}

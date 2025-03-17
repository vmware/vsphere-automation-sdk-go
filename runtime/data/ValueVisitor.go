// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

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

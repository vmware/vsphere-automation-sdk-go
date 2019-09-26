package info

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type UnionCase struct {
	caseValue  string
	fieldNames []string
}

func NewUnionCase() *UnionCase {
	fieldNames := []string{}
	unionCase := UnionCase{fieldNames: fieldNames}
	return &unionCase
}

func (u *UnionCase) CaseValue() string {
	return u.caseValue
}

func (u *UnionCase) SetCaseValue(cval string) {
	u.caseValue = cval
}

func (u *UnionCase) FieldNames() []string {
	return u.fieldNames
}

func (u *UnionCase) SetFieldNames(fieldNames []string) {
	u.fieldNames = fieldNames
}

func (u *UnionCase) AddToFieldNames(fieldName string) {
	u.fieldNames = append(u.fieldNames, fieldName)
}

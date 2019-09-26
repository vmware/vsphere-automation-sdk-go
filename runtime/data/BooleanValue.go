/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data

type BooleanValue struct {
	value bool
}

func NewBooleanValue(value bool) *BooleanValue {
	return &BooleanValue{value: value}
}

func (b *BooleanValue) Type() DataType {
	return BOOLEAN
}

func (b *BooleanValue) Value() bool {
	return (b.value)
}

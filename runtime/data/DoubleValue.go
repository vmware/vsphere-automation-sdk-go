/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data

import "fmt"

type DoubleValue struct {
	value float64
}

func NewDoubleValue(value float64) *DoubleValue {
	return &DoubleValue{value: value}
}
func (doubleValue *DoubleValue) Type() DataType {
	return DOUBLE
}

func (doubleValue *DoubleValue) Value() float64 {
	return doubleValue.value
}

func (doubleValue *DoubleValue) String() string {
	return fmt.Sprintf("%g", doubleValue.value)
}

package info

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type PrimitiveTypeInfo struct {
	name          string
	primitiveType PrimitiveType
}

func NewPrimitiveTypeInfo(name string, primitiveType PrimitiveType) *PrimitiveTypeInfo {
	primitiveTypeInfo := PrimitiveTypeInfo{name: name, primitiveType: primitiveType}
	return &primitiveTypeInfo
}

// Name
func (pt *PrimitiveTypeInfo) Name() string {
	return pt.name
}

func (pt *PrimitiveTypeInfo) SetName(name string) {
	pt.name = name
}

// Primitive Type
func (pt *PrimitiveTypeInfo) PrimitiveType() PrimitiveType {
	return pt.primitiveType
}

func (pt *PrimitiveTypeInfo) SetPrimitiveType(primitiveType PrimitiveType) {
	pt.primitiveType = primitiveType
}

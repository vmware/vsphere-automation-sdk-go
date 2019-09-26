package info

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type UserDefinedTypeInfo struct {
	name          string
	structureInfo *StructureInfo
}

func NewUserDefinedTypeInfo(name string, structureInfo *StructureInfo) *UserDefinedTypeInfo {
	return &UserDefinedTypeInfo{name: name, structureInfo: structureInfo}
}

// Name
func (udt *UserDefinedTypeInfo) Name() string {
	return udt.name
}

// StructureInfo
func (udt *UserDefinedTypeInfo) StructureInfo() *StructureInfo {
	return udt.structureInfo
}

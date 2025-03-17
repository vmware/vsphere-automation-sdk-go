// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package info

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

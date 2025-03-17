// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package info

type EnumerationInfo struct {
	identifier    string
	documentation string
	enumValues    []string
}

func NewEnumerationInfo() *EnumerationInfo {
	enumvalues := []string{}
	enumInfo := EnumerationInfo{enumValues: enumvalues}
	return &enumInfo
}

func (en *EnumerationInfo) EnumValues() []string {
	return en.enumValues
}

func (en *EnumerationInfo) SetEnumValues(enumValues []string) {
	en.enumValues = enumValues
}

func (en *EnumerationInfo) Identifier() string {
	return en.identifier
}

func (en *EnumerationInfo) SetIdentifier(identifier string) {
	en.identifier = identifier
}

func (en *EnumerationInfo) Documentation() string {
	return en.documentation
}

func (en *EnumerationInfo) SetDocumentation(documentation string) {
	en.documentation = documentation
}

// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

type ProviderDefinition struct {
	id       string
	checkSum string
}

func (pd *ProviderDefinition) Identifier() string {
	return pd.id
}

func (pd *ProviderDefinition) CheckSum() string {
	return pd.checkSum
}

func NewProviderDefinition(id string, checkSum string) *ProviderDefinition {
	return &ProviderDefinition{id: id, checkSum: checkSum}
}

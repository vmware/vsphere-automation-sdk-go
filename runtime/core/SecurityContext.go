// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

type SecurityContext interface {
	Property(key string) interface{}
	GetAllProperties() map[string]interface{}
	SetProperty(key string, value interface{})
}

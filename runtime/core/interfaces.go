// Copyright (c) 2022-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

type JSONRPCRequestPreProcessor interface {
	Process(requestBody map[string]interface{}) error
}

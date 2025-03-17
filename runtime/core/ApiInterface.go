// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
)

// ApiInterface responsibility is to convert to native the DataValue and invoke the implementation
// It also provides methods used by introspection services. Implemented by API providers.
type ApiInterface interface {
	Identifier() InterfaceIdentifier
	Definition() InterfaceDefinition
	MethodDefinition(MethodIdentifier) *MethodDefinition
	Invoke(ctx *ExecutionContext, methodId MethodIdentifier, input data.DataValue) MethodResult
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Greeter
// Used by service-side to provide implementations.

package greeter

import (
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
)

// The ``Greeter`` interface provides methods to greet the client.
type GreeterProvider interface {

	// Gives a greeting.
	// @return a string with the greeting.
	// @throws Unauthenticated if the caller is not authenticated.
	Greet(ctx *vapiCore_.ExecutionContext) (string, error)
}

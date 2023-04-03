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

	// Provides a stream of greetings
	// @return a greeting string.
	//
	// @throws InternalServerError
	Greet(ctx *vapiCore_.ExecutionContext) (chan string, chan error)
}

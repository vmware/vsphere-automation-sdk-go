// Copyright © 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause
package impl

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
)

type GreeterImpl struct {
}

func NewGreeterServiceImpl() *GreeterImpl {
	return &GreeterImpl{}
}

func (GreeterImpl GreeterImpl) InformalyGreet(ctx *core.ExecutionContext) (string, error) {
	var informalGreeting = "Hi :)"
	return informalGreeting, nil
}
func (GreeterImpl GreeterImpl) FormalyGreet(ctx *core.ExecutionContext) (string, error) {
	var formalGreeting = "Greetings!"
	return formalGreeting, nil
}

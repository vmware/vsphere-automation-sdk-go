// Copyright © 2022 VMware, Inc. All Rights Reserved.
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

func (GreeterImpl GreeterImpl) Greet(ctx *core.ExecutionContext) (string, error) {
	return "Hi :)", nil
}

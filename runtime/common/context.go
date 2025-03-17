// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package common

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/google/uuid"
)

func NewDefaultApplicationContext() *core.ApplicationContext {
	appContext := core.NewApplicationContext(nil)
	InsertOperationId(appContext)
	return appContext
}

func NewUUID() string {
	return uuid.NewString()
}

func InsertOperationId(appContext *core.ApplicationContext) {
	opId := NewUUID()
	appContext.SetProperty(lib.OPID, &opId)
}

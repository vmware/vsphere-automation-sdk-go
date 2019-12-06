/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type info struct {
	userName string
	password string
}

func (basicAuthInfo *info) UserName() string {
	return basicAuthInfo.userName
}

func (basicAuthInfo *info) Password() string {
	return basicAuthInfo.password
}

func (basicAuthInfo *info) Name() string {
	return Name
}

func (basicAuthInfo *info) AuthInterface() interface{} {
	return Info(basicAuthInfo)
}

func (basicAuthInfo *info) SecurityContext() (core.SecurityContext, error) {
	return security.NewUserPasswordSecurityContext(basicAuthInfo.UserName(), basicAuthInfo.Password()), nil
}

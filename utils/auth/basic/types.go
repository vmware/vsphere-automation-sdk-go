/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package basic

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

// Info represents the details of Basic Authentication.
type Info interface {
	// UserName returns the User Name for Basic Authentication Scheme.
	UserName() string
	// Password returns the Password for Basic Authentication Scheme.
	Password() string
	model.AuthInfo
}

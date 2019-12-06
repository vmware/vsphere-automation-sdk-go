/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package refreshtoken

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

// Info represents the details of OAuth Authentication Scheme by Refresh Token.
type Info interface {
	// CSPURL returns CSP URL for OAuth Authentication Scheme by Refresh Token.
	CSPURL() string
	// RefreshToken returns Refresh Token for OAuth Authentication Scheme by Refresh Token.
	RefreshToken() string
	model.AuthInfo
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/model"
)

// Info represents the details of SAML Bearer Authentication Scheme.
type Info interface {
	// SAMLBearerToken returns the Bearer Token for SAML Bearer Authentication Scheme.
	SAMLBearerToken() string
	model.AuthInfo
}

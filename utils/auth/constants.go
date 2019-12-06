/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package auth

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/basic"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/oauth/refreshtoken"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth/samlbearer"
)

const (
	// NoAuth defines no authentication.
	NoAuth = Scheme("NoAuth")
	// BasicAuth defines Basic Authentication Scheme.
	BasicAuth = Scheme(basic.Name)
	// SAMLBearer defines SAML Bearer Authentication Scheme.
	SAMLBearer = Scheme(samlbearer.Name)
	// OAuthRefreshToken defines OAuth Authentication Scheme by Refresh Token.
	OAuthRefreshToken = Scheme(refreshtoken.Name)
)

// Schemes defines array of all the Authentication Schemes.
var Schemes []Scheme = []Scheme{
	NoAuth,
	BasicAuth,
	SAMLBearer,
	OAuthRefreshToken,
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package refreshtoken

const (
	// CSPDefaultURL defines the default CSP URL for OAuth Authentication
	// by Refresh Token.
	CSPDefaultURL string = "https://console.cloud.vmware.com"
	// CSPRefreshURLSuffix defines the Refresh URL suffix for OAuth
	// Authentication by Refresh Token.
	CSPRefreshURLSuffix string = "/csp/gateway/am/api/auth/api-tokens/authorize"
	// Name defines name of the Authentication Scheme OAuth by Refresh Token.
	Name string = "OAuth.RefreshToken"
)

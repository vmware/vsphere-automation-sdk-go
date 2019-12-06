/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package samlbearer

// Error represents the error for SAML Bearer Authentication Scheme.
type Error struct {
}

func (samlBearerAuthError *Error) Error() string {
	return Name + " Auth Error: saml-bearer-token is missing."
}

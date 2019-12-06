/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package samlbearer represents SAML Bearer Authentication Scheme
// that uses Bearer Token for authentication.
package samlbearer

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/args/keys"
)

// CreateSAMLBearerAuthFromProperties defines and returns new Info for
// SAML Bearer Authentication from map of string to interface.
func CreateSAMLBearerAuthFromProperties(authMap map[string]interface{}) (Info, error) {
	samlToken := authMap[keys.AuthSchemeSAMLBearerTokenKey]
	if samlToken == nil {
		return nil, &Error{}
	}

	return &info{Token: samlToken.(string)}, nil
}

// Create defines and returns new Info for
// SAML Bearer Authentication.
func Create(samlToken string) (Info, error) {
	if len(samlToken) <= 0 {
		return nil, &Error{}
	}

	return &info{Token: samlToken}, nil
}

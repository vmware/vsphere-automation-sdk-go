/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package session represents Session Management for the Hosts.
package session

import (
	"net/http"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

// NewSessionManager creates and returns new Session Manager that manages API session with the Host server.
func NewSessionManager(url string, archType arch.Type, authDetails auth.Info, httpClient http.Client) (Manager, error) {
	if len(url) == 0 {
		return nil, &URLError{}
	}
	if archType == nil {
		return nil, &ArchTypeError{archType: "nil"}
	}
	if authDetails == nil {
		return nil, &AuthDetailsError{}
	}

	sessionManager := manager{URL: url, ArchType: archType, AuthDetails: authDetails, HTTPClient: httpClient}
	if sessionManager.ArchType == arch.JSONRPC {
		sessionManager.connector = client.NewJsonRpcConnector(sessionManager.URL, sessionManager.HTTPClient)
	} else if sessionManager.ArchType == arch.REST {
		sessionManager.connector = client.NewRestConnector(sessionManager.URL, sessionManager.HTTPClient)
	} else {
		return nil, &ArchTypeError{archType: sessionManager.ArchType.String()}
	}
	return &sessionManager, nil
}

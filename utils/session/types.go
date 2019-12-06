/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package session

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

// Manager represents Session Manager.
type Manager interface {
	// Login logs in on the server and creates new session.
	Login() error
	// Logout logs out of the server and deletes the session.
	Logout() error
	// Connector creates and returns new client connector.
	Connector() client.Connector
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

// Info represents details of the HostTypes.
type Info interface {
	// isHostType checks if Info is a valid HostType.
	isHostType() Info
	// String returns the string value of Info of HostType.
	String() string
	// Protocol returns the protocol supported by the HostType.
	Protocol() arch.Type
	// SupportedAuthTypes returns all Authentication Schemes supported by HostType.
	SupportedAuthTypes() []auth.Scheme
	// APIEndPointPrefix returns API Endpoint Prefix for HostType.
	APIEndPointPrefix() string
}

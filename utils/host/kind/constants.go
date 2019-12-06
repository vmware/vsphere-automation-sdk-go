/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

// VCENTER defines Host Type Info for the VMware product vCenter.
var VCENTER Info = info{
	name:              "vCenter",
	protocol:          arch.JSONRPC,
	apiEndPointPrefix: "/api",
	supportedAuthTypes: []auth.Scheme{
		auth.BasicAuth,
		auth.SAMLBearer,
	},
}

// VMC defines Host Type Info for the VMware product VMC.
var VMC Info = info{
	name:              "VMC",
	protocol:          arch.REST,
	apiEndPointPrefix: "",
	supportedAuthTypes: []auth.Scheme{
		auth.OAuthRefreshToken,
	},
}

// NSXT defines Host Type Info for the VMware product NSXT.
var NSXT Info = info{
	name:              "NSXT",
	protocol:          arch.REST,
	apiEndPointPrefix: "",
	supportedAuthTypes: []auth.Scheme{
		auth.OAuthRefreshToken,
	},
}

// All defines array of Info for all the Host Types.
var All = []Info{
	VCENTER,
	VMC,
	NSXT,
}

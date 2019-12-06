/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package kind

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/auth"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session/arch"
)

type info struct {
	name               string
	protocol           arch.Type
	supportedAuthTypes []auth.Scheme
	apiEndPointPrefix  string
}

func (ht info) isHostType() Info {
	return ht
}

func (ht info) String() string {
	return ht.name
}

func (ht info) Protocol() arch.Type {
	return ht.protocol
}

func (ht info) APIEndPointPrefix() string {
	return ht.apiEndPointPrefix
}

func (ht info) SupportedAuthTypes() []auth.Scheme {
	return ht.supportedAuthTypes
}

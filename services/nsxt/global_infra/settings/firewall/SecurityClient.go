/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Security
 * Used by client-side stubs.
 */

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SecurityClient interface {

    // Get the current dfw firewall configurations.
    // @return com.vmware.nsx_policy.model.DfwFirewallConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.DfwFirewallConfiguration, error)

    // Update dfw firewall related configurations.
    //
    // @param dfwFirewallConfigurationParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(dfwFirewallConfigurationParam model.DfwFirewallConfiguration) error

    // Update dfw firewall related configurations.
    //
    // @param dfwFirewallConfigurationParam (required)
    // @return com.vmware.nsx_policy.model.DfwFirewallConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(dfwFirewallConfigurationParam model.DfwFirewallConfiguration) (model.DfwFirewallConfiguration, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Dns
 * Used by client-side stubs.
 */

package management_vms

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

type DnsClient interface {

    // Update the DNS records of management VMs to use public or private IP addresses
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param managementVmIdParam the management VM ID (required)
    // @param ipTypeParam the ip type to associate with FQDN in DNS (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates or invalid body
    // @throws Unauthorized  Forbidden
	Update(orgParam string, sddcParam string, managementVmIdParam string, ipTypeParam string) (model.Task, error)
}

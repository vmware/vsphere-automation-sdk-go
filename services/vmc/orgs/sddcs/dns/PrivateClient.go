/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Private
 * Used by client-side stubs.
 */

package dns

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type PrivateClient interface {

    // Update the DNS records of management VMs to use private IP addresses
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates or invalid body
    // @throws Unauthorized  Forbidden
	Update(orgParam string, sddcParam string) (model.Task, error)
}

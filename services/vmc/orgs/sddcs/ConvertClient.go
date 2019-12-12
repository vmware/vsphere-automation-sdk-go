/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Convert
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ConvertClient interface {

    // This API converts a one host SDDC to a four node DEFAULT SDDC. It takes care of configuring and upgrading the vCenter configurations on the SDDC for high availability and data redundancy.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates, Method not allowed
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the SDDC with given identifier
	Create(orgParam string, sddcParam string) (model.Task, error)
}

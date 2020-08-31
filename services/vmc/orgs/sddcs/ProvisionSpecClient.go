/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ProvisionSpec
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ProvisionSpecClient interface {

    // Get sddc provision spec for an org
    //
    // @param orgParam Organization identifier (required)
    // @return com.vmware.vmc.model.ProvisionSpec
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  BadRequest
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Internal server error.
	Get(orgParam string) (model.ProvisionSpec, error)
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ClusterProvisionSpec
 * Used by client-side stubs.
 */

package sddcs

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type ClusterProvisionSpecClient interface {

    // Get cluster provision spec for a new cluster in an sddc
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.ClusterProvisionSpec
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws InternalServerError  Internal server error.
	Get(orgParam string, sddcParam string) (model.ClusterProvisionSpec, error)
}

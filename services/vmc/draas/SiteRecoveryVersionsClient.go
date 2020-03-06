/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SiteRecoveryVersions
 * Used by client-side stubs.
 */

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

type SiteRecoveryVersionsClient interface {

    // Query site recovery versions for the specified sddc
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param versionSourceParam Represents the source for getting the version from. (optional)
    // @return com.vmware.vmc.draas.model.SiteRecoveryVersions
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find site recovery versions for sddc identifier
	Get(orgParam string, sddcParam string, versionSourceParam *string) (model.SiteRecoveryVersions, error)
}

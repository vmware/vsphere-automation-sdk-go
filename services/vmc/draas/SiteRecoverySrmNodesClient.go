/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SiteRecoverySrmNodes
 * Used by client-side stubs.
 */

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

type SiteRecoverySrmNodesClient interface {

    // Delete a SRM server.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param srmNodeParam SRM node identifier (required)
    // @return com.vmware.vmc.draas.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find SDDC or SRM node
	Delete(orgParam string, sddcParam string, srmNodeParam string) (model.Task, error)

    // Provision an additional SRM server.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param provisionSrmConfigParam Customization, for example can specify custom extension key suffix for SRM. (optional)
    // @return com.vmware.vmc.draas.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Post(orgParam string, sddcParam string, provisionSrmConfigParam *model.ProvisionSrmConfig) (model.Task, error)
}

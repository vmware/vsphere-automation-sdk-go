/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SiteRecovery
 * Used by client-side stubs.
 */

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

type SiteRecoveryClient interface {

    // Deactivate site recovery for the specified sddc
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param forceParam If = 'true', will deactivate site recovery forcefully. (optional)
    // @param deactivateHcxParam If = 'true', will deactivate HCX. (optional)
    // @return com.vmware.vmc.draas.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Delete(orgParam string, sddcParam string, forceParam *bool, deactivateHcxParam *bool) (model.Task, error)

    // Query site recovery configuration for the specified sddc
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @return com.vmware.vmc.draas.model.SiteRecovery
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string) (model.SiteRecovery, error)

    // Activate site recovery for the specified sddc
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param activateSiteRecoveryConfigParam Customization, for example can specify custom extension key suffix for SRM. (optional)
    // @return com.vmware.vmc.draas.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Post(orgParam string, sddcParam string, activateSiteRecoveryConfigParam *model.ActivateSiteRecoveryConfig) (model.Task, error)
}

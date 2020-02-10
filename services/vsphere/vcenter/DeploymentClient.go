/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Deployment
 * Used by client-side stubs.
 */

package vcenter


// The ``Deployment`` interface provides methods to get the status of the vCenter appliance deployment. This interface was added in vSphere API 6.7.
type DeploymentClient interface {

    // Get the current status of the appliance deployment. This method was added in vSphere API 6.7.
    // @return Info structure containing the status information about the appliance.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotFound if appliance state cannot be determined.
	Get() (DeploymentInfo, error)

    // Rollback a failed appliance so it can be configured once again. This method was added in vSphere API 6.7.
    // @throws Unsupported if the appliance is not in FAILED state.
    // @throws Unauthenticated if the caller is not authenticated.
	Rollback() error
}

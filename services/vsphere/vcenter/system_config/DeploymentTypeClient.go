/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DeploymentType
 * Used by client-side stubs.
 */

package system_config


// The ``DeploymentType`` interface provides methods to get/set the type of the appliance. This interface was added in vSphere API 6.7.
type DeploymentTypeClient interface {

    // Get the type of the vCenter appliance. This method was added in vSphere API 6.7.
    // @return The type of the vCenter appliance.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in CONFIGURED state.
	Get() (DeploymentTypeInfo, error)

    // Reconfigure the type of the vCenter appliance. This method was added in vSphere API 6.7.
    //
    // @param specParam ReconfigureSpec to set the appliance type.
    // @throws Unsupported if the appliance is in CONFIGURED state and if not changing the type from VCSA_EMBEDDED to VCSA_EXTERNAL.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if external PSC credentials are not valid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED or CONFIGURED state.
	Reconfigure(specParam DeploymentTypeReconfigureSpec) error
}

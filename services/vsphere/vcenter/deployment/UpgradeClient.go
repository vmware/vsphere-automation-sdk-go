/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Upgrade
 * Used by client-side stubs.
 */

package deployment


// The ``Upgrade`` interface provides methods to configure the upgrade of this appliance from an existing vCenter appliance. This interface was added in vSphere API 6.7.
type UpgradeClient interface {

    // Get the UpgradeSpec parameters used to configure the ongoing appliance upgrade. This method was added in vSphere API 6.7.
    // @return UpgradeSpec parameters being used to configure appliance upgrade.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in UPGRADE_PROGRESS state.
	Get() (UpgradeUpgradeSpec, error)

    // Run sanity checks using the UpgradeSpec parameters passed. This method was added in vSphere API 6.7.
    //
    // @param specParam UpgradeSpec parameters to run sanity check on.
    // @return CheckInfo containing the check results.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if source credentials are not valid.
    // @throws Unauthenticated if source container credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
	Check(specParam UpgradeUpgradeSpec) (CheckInfo, error)

    // Start the appliance installation. This method was added in vSphere API 6.7.
    //
    // @param specParam UpgradeSpec parameters to configure the appliance upgrade.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if source credentials are not valid.
    // @throws Unauthenticated if source container credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
	Start(specParam UpgradeUpgradeSpec) error

    // Cancel the appliance upgrade that is in progress. This method was added in vSphere API 6.7.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in CONFIG_IN_PROGRESS state and if the operation is not INSTALL.
	Cancel() error
}

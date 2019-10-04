/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Upgrade
 * Used by client-side stubs.
 */

package upgrade

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/deployment"
)

// The ``Upgrade`` interface provides methods to configure the upgrade of this appliance from an existing vCenter appliance.
type UpgradeClient interface {


    // Get the UpgradeSpec parameters used to configure the ongoing appliance upgrade.
    // @return UpgradeSpec parameters being used to configure appliance upgrade.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in UPGRADE_PROGRESS state.
    Get() (UpgradeSpec, error) 


    // Run sanity checks using the UpgradeSpec parameters passed.
    //
    // @param specParam UpgradeSpec parameters to run sanity check on.
    // @return CheckInfo containing the check results.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if source credentials are not valid.
    // @throws Unauthenticated if source container credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam UpgradeSpec) (deployment.CheckInfo, error) 


    // Start the appliance installation.
    //
    // @param specParam UpgradeSpec parameters to configure the appliance upgrade.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if source credentials are not valid.
    // @throws Unauthenticated if source container credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Start(specParam UpgradeSpec) error 


    // Cancel the appliance upgrade that is in progress.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in CONFIG_IN_PROGRESS state and if the operation is not INSTALL.
    Cancel() error 

}

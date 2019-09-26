/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Migrate
 * Used by client-side stubs.
 */

package migrate

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
)

// The ``Migrate`` interface provides methods to configure the migration of this appliance from an existing vCenter for Windows. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MigrateClient interface {


    // Get the MigrateSpec parameters used to configure the ongoing appliance migration. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return MigrateSpec parameters being used to configure appliance migration.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if appliance is not in MIGRATE_PROGRESS state.
    Get() (MigrateSpec, error) 


    // Run sanity checks using the MigrateSpec parameters passed. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam MigrateSpec parameters to run sanity check on.
    // @return CheckInfo containing the check results.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if migration assistant credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam MigrateSpec) (deployment.CheckInfo, error) 


    // Start the appliance migration. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam MigrateSpec parameters to configure the appliance migration.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if migration assistant credentials are not valid.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Start(specParam MigrateSpec) error 


    // Cancel the appliance migration that is in progress. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is not in CONFIG_IN_PROGRESS state and if the operation is not INSTALL.
    Cancel() error 

}

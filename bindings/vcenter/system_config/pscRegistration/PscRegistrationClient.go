/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: PscRegistration
 * Used by client-side stubs.
 */

package pscRegistration

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
)

// The ``PscRegistration`` interface provides methods to get and set the PSC_EXTERNAL appliance a VCSA_EXTERNAL appliance is registered with.
type PscRegistrationClient interface {


    // Get information of the PSC that this appliance is registered with.
    // @return Info structure containing information about the external PSC node this appliance is registered with.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws NotAllowedInCurrentState if the appliance is in NOT_INITIALIZED state.
    Get() (Info, error) 


    // Repoint this vCenter Server appliance to a different external PSC.
    //
    // @param specParam RemotePscSpec structure containing information about the external PSC node to repoint this vCenter Server appliance to.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if the passed external PSC credentials is invalid.
    // @throws InvalidArgument if the passed external PSC is not a replicating with the current PSC this appliance is registered with.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Unsupported if the current appliance is not of the type VCSA_EXTERNAL.
    // @throws NotAllowedInCurrentState if the appliance is NOT in CONFIGURED state.
    Repoint(specParam deployment.RemotePscSpec) error 

}

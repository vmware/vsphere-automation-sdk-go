/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: RemotePsc
 * Used by client-side stubs.
 */

package remotePsc

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/deployment"
)

// The ``RemotePsc`` interface provides methods to check if the deployed vCenter Server can register with the remote PSC.
type RemotePscClient interface {


    // Checks whether the remote PSC is reachable and the deployed vCenter Server can be registered with the remote PSC.
    //
    // @param specParam Information to connect to the remote PSC.
    // @return Information about the success or failure of the checks that were performed.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if external PSC credentials are not valid when configuring a VCSA_EXTERNAL appliance.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam deployment.RemotePscSpec) (deployment.CheckInfo, error) 

}

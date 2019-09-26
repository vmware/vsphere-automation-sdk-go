/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Standalone
 * Used by client-side stubs.
 */

package standalone

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
)

// The ``Standalone`` interface provides methods to check if the values provided for the standalone PSC satisfies the requirements.
type StandaloneClient interface {


    // Checks that the information to configure a non-replicated PSC satisfies the requirements.
    //
    // @param specParam Information to configure a non-replicated PSC.
    // @return Information about the success or failure of the checks that were performed.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam deployment.StandalonePscSpec) (deployment.CheckInfo, error) 

}

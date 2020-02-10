/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Standalone
 * Used by client-side stubs.
 */

package psc

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/deployment"
)

// The ``Standalone`` interface provides methods to check if the values provided for the standalone PSC satisfies the requirements. This interface was added in vSphere API 6.7.
type StandaloneClient interface {

    // Checks that the information to configure a non-replicated PSC satisfies the requirements. This method was added in vSphere API 6.7.
    //
    // @param specParam Information to configure a non-replicated PSC.
    // @return Information about the success or failure of the checks that were performed.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
	Check(specParam deployment.StandalonePscSpec) (deployment.CheckInfo, error)
}

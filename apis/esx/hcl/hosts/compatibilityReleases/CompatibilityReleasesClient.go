/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CompatibilityReleases
 * Used by client-side stubs.
 */

package compatibilityReleases

import (
)

// This interface provides methods to list available releases for generating compatibility report for a specific ESXi host. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReleasesClient interface {


    // Lists the locally available ESXi releases for a given host that can be used to generate a compatiblity report. Each host has its own list of supported releases depending on its current release. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Contains the MoID identifying the ESXi host.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @return Available releases for compatibility for a specified host.
    // @throws InternalServerError If there is some internal error. The accompanying error message will give more details about the failure.
    // @throws NotAllowedInCurrentState if there is no compatibility data on the vCenter executing the operation.
    // @throws NotFound if no host with the given MoID can be found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unsupported if the provided host is not supported.
    // @throws ResourceInaccessible if the vCenter this API is executed on is not part of the Customer Experience Improvement Program (CEIP).
    // @throws Error If there is some unknown error. The accompanying error message will give more details about the failure.
    List(hostParam string) (EsxiCompatibilityReleases, error) 

}

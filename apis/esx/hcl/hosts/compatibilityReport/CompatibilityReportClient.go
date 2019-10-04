/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CompatibilityReport
 * Used by client-side stubs.
 */

package compatibilityReport

import (
)

// This interface provides methods to generate hardware compatibility report for a given ESXi host against a specific ESXi release. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportClient interface {


    // Generates hardware compatibility report for a specified ESXi host against specific ESXi release.
    //
    //  The result of this operation can be queried by calling the cis/tasks/{task-id} where the task-id is the response of this operation.. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Contains the MoID identifying the ESXi host.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @param specParam Specifies the input parameters for generating compatibility report.
    // If null host compatibility will be checked against the current release of the ESXi.
    // @return ``Result`` class that contains the requested report and the identifier of the report.
    // @throws NotFound if no host with the given MoID can be found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unsupported if the provided host is not supported.
    // @throws ResourceInaccessible if the vCenter this API is executed on is not part of the Customer Experience Improvement Program (CEIP).
    // @throws NotAllowedInCurrentState if there is no compatibility data on the vCenter executing the operation.
    // @throws Error If there is some unknown error. The accompanying error message will give more details about the failure.
    Create(hostParam string, specParam *Spec) (Result, error) 

}

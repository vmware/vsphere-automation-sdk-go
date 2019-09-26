/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Policy
 * Used by client-side stubs.
 */

package policy

import (
)

// The ``Policy`` interface provides methods to set/get background check for the new updates.
type PolicyClient interface {


    // Gets the automatic update checking and staging policy.
    // @return Structure containing the policy for the appliance update.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    Get() (Info, error) 


    // Sets the automatic update checking and staging policy.
    //
    // @param policyParam Info structure containing the policy for the appliance update.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    Set(policyParam Config) error 

}

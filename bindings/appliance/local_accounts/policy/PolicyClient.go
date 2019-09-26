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

// The ``Policy`` interface provides methods to manage local user accounts
type PolicyClient interface {


    // Get the global password policy.
    // @return Global password policy
    // @throws Error Generic error
    Get() (Info, error) 


    // Set the global password policy.
    //
    // @param policyParam Global password policy
    // @throws InvalidArgument if passed policy values are < -1 or > 99999
    // @throws Error Generic error
    Set(policyParam Info) error 

}

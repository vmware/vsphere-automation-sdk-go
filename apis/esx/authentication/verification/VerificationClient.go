/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Verification
 * Used by client-side stubs.
 */

package verification

import (
)

// ESX Authentication interface 
//
//  This ``Verification`` interface is providing a verification tool for tokens issued by Esx Authentication service.
type VerificationClient interface {


    // Verifies the provided token
    // @return username from the token
    Verify() (string, error) 

}

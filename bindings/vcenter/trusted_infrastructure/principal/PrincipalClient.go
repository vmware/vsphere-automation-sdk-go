/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Principal
 * Used by client-side stubs.
 */

package principal

import (
)

// The ``Principal`` interface contains information about the certificates which sign the tokens used by vCenter for authentication.
type PrincipalClient interface {


    // Returns information about the STS used by this vCenter instance.
    // @return \\\\@{link Info} a summary containing the certificates used to sign tokens and the solution user used to retrieve them.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadStsInfo``.
    Get() (Info, error) 

}

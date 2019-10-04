/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Thumbprint
 * Used by client-side stubs.
 */

package thumbprint

import (
)

// The ``Thumbprint`` interface provides methods to get the thumbprint of the remote PSC.
type ThumbprintClient interface {


    // Gets the SHA1 thumbprint of the remote PSC.
    //
    // @param specParam Information used to connect to the remote PSC.
    // @return The thumbprint of the specified remote PSC
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Error on exception.
    Get(specParam RemoteSpec) (string, error) 

}

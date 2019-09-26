/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Token
 * Used by client-side stubs.
 */

package token

import (
)

// ESXi Authentication interface 
//
//  This ``Token`` interface is providing tokens for accessing control plane services on the host. 
//
//  It is reading the user's credentials from the call context and creates a signed API access token. The token can be presented as authentication to the other host APIs. API clients should treat the token as opaque value. Future ESXi releases may utilize different token technology. There are no backwards compatibility guarantees about the token format and technology. The current implementation uses JSON Web Tokens (JWT) as per RFC 7519.
type TokenClient interface {


    // Creates API access token using API user credentials in the call context. The returned token can be used to authenticate access to other host APIs 
    //
    //  The ESXi Authentication interface accepts the following types of user credentials: 
    //
    // * Username and password
    // * vCenter SAML tokens
    //
    //  
    // @throws Unauthenticated if the credentials provided in the API call cannot be validated or are insufficient to generate API access token
    Create() (TokenInfo, error) 

}

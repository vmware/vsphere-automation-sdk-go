/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TokenExchange
 * Used by client-side stubs.
 */

package tokenExchange

import (
)

// The ``TokenExchange`` interface provides possibility to exchange between different tokens types. Implementation of "OAuth 2.0 Token Exchange" standard (https://tools.ietf.org/html/draft-ietf-oauth-token-exchange-12). **Warning:** This interface is available as technical preview. It may be changed in a future release.
type TokenExchangeClient interface {


    // Exchanges incoming token based on the spec and current client authorization data. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param specParam  ``ExchangeSpec`` class contains arguments that define exchange process.
    // @return Info class that contains new token.
    // @throws InvalidGrant  provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client.
    // @throws InvalidScope  If the server is unwilling or unable to issue a token for all the target services indicated by the ExchangeSpec#resource or ExchangeSpec#audience parameters.
    // @throws Unauthorized  if authorization is not given to a caller.
    Exchange(specParam ExchangeSpec) (Info, error) 

}

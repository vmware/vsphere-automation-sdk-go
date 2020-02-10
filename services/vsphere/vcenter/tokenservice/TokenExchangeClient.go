/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TokenExchange
 * Used by client-side stubs.
 */

package tokenservice


// The ``TokenExchange`` interface provides possibility to exchange between different tokens types. Implementation of "OAuth 2.0 Token Exchange" standard (https://tools.ietf.org/html/draft-ietf-oauth-token-exchange-12). **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TokenExchangeClient interface {

    // Exchanges incoming token based on the spec and current client authorization data. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param specParam  ``ExchangeSpec`` class contains arguments that define exchange process.
    // @return TokenExchangeInfo class that contains new token.
    // @throws InvalidGrant  provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client.
    // @throws InvalidScope  If the server is unwilling or unable to issue a token for all the target services indicated by the TokenExchangeExchangeSpec#resource or TokenExchangeExchangeSpec#audience parameters.
    // @throws Unauthorized  if authorization is not given to a caller.
	Exchange(specParam TokenExchangeExchangeSpec) (TokenExchangeInfo, error)
}

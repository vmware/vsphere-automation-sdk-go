/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SecurityTokenIssuers
 * Used by client-side stubs.
 */

package securityTokenIssuers

import (
)

// The ``SecurityTokenIssuers`` interface provides methods to manage ESX trust to Security Token Issuers. For example WS-Trust SSO STS or OAuth Identity Providers, that issue authentication tokens.
//
//  Each record contains a list of X.509 certificate chains that apply for a unique token issuer. The certificates are used to verify security tokens from this issuer.
//
//  Symmetric signing keys are not suported.
//
//  If a security token is received with issuer that has no configured ``SecurityTokenIssuers`` instance, the authentication fails and an error is returned.
type SecurityTokenIssuersClient interface {


    // List the available security token issuer trusts.
    //
    // @param projectionParam The type of the returned summary - brief, normal or full.
    // If {\\\\@term.unset} a normal projection will be used.
    // @return The list of the current security token issuer trusts.
    // @throws Error if there is a problem accessing the stored data.
    // @throws InvalidArgument if the null argument contains invalid data.
    // @throws Unauthenticated if the user can not be authenticated.
    List(projectionParam *SummaryType) ([]Summary, error) 


    // Create a new trust to a Security Token Issuer.
    //
    // @param specParam Settings for the new security token issuer trust.
    // @throws AlreadyExists if the security token issuer or alias already exist. If the issuer already exists, the value of the null property will be a class that contains all the properties defined in IssuerAlreadyExistsInfo where null is the alias where the issuer is already saved.
    // @throws InvalidArgument if the CreateSpec argument contains invalid data.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Create(specParam CreateSpec) error 


    // Get the details of a security token issuer.
    //
    // @param issuerAliasParam The security token issuer trust identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.trust.security-token-issuer``.
    // @return Details of the security token issuer trust.
    // @throws NotFound if there is no security token issuer trust for the alias.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(issuerAliasParam string) (Info, error) 


    // Update an existing security token issuer trust.
    //
    // @param issuerAliasParam The security token issuer trust identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.trust.security-token-issuer``.
    // @param specParam The new settings.
    // @throws NotFound if there is no security token issuer trust for the alias.
    // @throws Error if there is a problem storing the data.
    // @throws InvalidArgument if the UpdateSpec argument contains invalid data.
    // @throws Unauthenticated if the user can not be authenticated.
    Update(issuerAliasParam string, specParam UpdateSpec) error 


    // Delete an existing security token issuer trust.
    //
    // @param issuerAliasParam The security token issuer trust identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.trust.security-token-issuer``.
    // @throws NotFound if there is no security token issuer trust for the alias.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Delete(issuerAliasParam string) error 

}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Session
// Used by service-side to provide implementations.

package cis

import (
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
)

// The ``Session`` interface allows API clients to manage session tokens including creating, deleting and obtaining information about sessions.
//
//
//
// * The Session#create method creates session token in exchange for another authentication token.
// * The Session#delete method invalidates a session token.
// * The Session#get retrieves information about a session token.
//
//
//
//  The call to the Session#create method is part of the overall authentication process for API clients. For example, the sequence of steps for establishing a session with SAML token is:
//
// * Connect to lookup service.
// * Discover the secure token service (STS) endpoint URL.
// * Connect to the secure token service to obtain a SAML token.
// * Authenticate to the lookup service using the obtained SAML token.
// * Discover the API endpoint URL from lookup service.
// * Call the Session#create method. The Session#create call must include the SAML token.
//
//
//
//  See the programming guide and samples for additional information about establishing API sessions.
//
//  **Execution Context and Security Context**
//
//  To use session based authentication a client should supply the session token obtained through the Session#create method. The client should add the session token in the security context when using SDK classes. Clients using the REST API should supply the session token using the ``vmware-api-session-id`` HTTP header field.
//
//  **Session Lifetime**
//
//  A session begins with call to the Session#create method to exchange a SAML token for a API session token. A session ends under the following circumstances:
//
// * Call to the Session#delete method.
// * The session expires. Session expiration may be caused by one of the following situations:
//
//     * Client inactivity - For a particular session identified by client requests that specify the associated session ID, the lapsed time since the last request exceeds the maximum interval between requests.
//     * Unconditional or absolute session expiration time: At the beginning of the session, the session logic uses the SAML token and the system configuration to calculate absolute expiration time.
//
//
//
//  When a session ends, the authentication logic will reject any subsequent client requests that specify that session. Any operations in progress will continue to completion.
//
//  **Error Handling**
//
//  The Session returns the following exceptions:
//
// * vapiStdErrors_.Unauthenticated exception for any exceptions related to the request.
// * vapiStdErrors_.ServiceUnavailable exception for all exceptions caused by internal service failure.
type SessionProvider interface {

	// Creates a session with the API. This is the equivalent of login. This method exchanges user credentials supplied in the security context for a session token that is to be used for authenticating subsequent calls.
	//
	// To authenticate subsequent calls clients are expected to include the session token. For REST API calls the HTTP ``vmware-api-session-id`` header field should be used for this.
	// @return Newly created session token to be used for authenticating further requests.
	//
	// @throws Unauthenticated  if the session creation fails due to request specific issues. Due to the security nature of the API the details of the error are not disclosed.
	//
	//  Please check the following preconditions if using a SAML token to authenticate:
	//
	// * the supplied token is delegate-able.
	// * the time of client and server system are synchronized.
	// * the token supplied is valid.
	// * if bearer tokens are used check that system configuration allows the API endpoint to accept such tokens.
	// @throws ServiceUnavailable  if session creation fails due to server specific issues, for example connection to a back end component is failing. Due to the security nature of this API further details will not be disclosed in the exception. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Create(ctx *vapiCore_.ExecutionContext) (string, error)
	// Terminates the validity of a session token. This is the equivalent of log out.
	//
	//  A session token is expected as part of the request.
	//
	// @throws Unauthenticated  if the session id is missing from the request or the corresponding session object cannot be found.
	// @throws ServiceUnavailable  if session deletion fails due to server specific issues, for example connection to a back end component is failing. Due to the security nature of this API further details will not be disclosed in the exception. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Delete(ctx *vapiCore_.ExecutionContext) error
	// Returns information about the current session. This method expects a valid session token to be supplied.
	//
	//  A side effect of invoking this method may be a change to the session's last accessed time to the current time if this is supported by the session implementation. Invoking any other method in the API will also update the session's last accessed time.
	//
	//  This API is meant to serve the needs of various front end projects that may want to display the name of the user. Examples of this include various web based user interfaces and logging facilities.
	// @return Information about the session.
	//
	// @throws Unauthenticated  if the session id is missing from the request or the corresponding session object cannot be found.
	// @throws ServiceUnavailable  if session retrieval fails due to server specific issues e.g. connection to back end component is failing. Due to the security nature of this API further details will not be disclosed in the error. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Get(ctx *vapiCore_.ExecutionContext) (SessionInfo, error)
}

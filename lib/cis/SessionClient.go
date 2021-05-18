// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Session
// Used by client-side stubs.

package cis

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

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
// * errors.Unauthenticated exception for any exceptions related to the request.
// * errors.ServiceUnavailable exception for all exceptions caused by internal service failure.
type SessionClient interface {

	// Creates a session with the API. This is the equivalent of login. This method exchanges user credentials supplied in the security context for a session token that is to be used for authenticating subsequent calls.
	//
	// To authenticate subsequent calls clients are expected to include the session token. For REST API calls the HTTP ``vmware-api-session-id`` header field should be used for this.
	// @return Newly created session token to be used for authenticating further requests.
	// @throws Unauthenticated  if the session creation fails due to request specific issues. Due to the security nature of the API the details of the error are not disclosed.
	//
	//  Please check the following preconditions if using a SAML token to authenticate:
	//
	// * the supplied token is delegate-able.
	// * the time of client and server system are synchronized.
	// * the token supplied is valid.
	// * if bearer tokens are used check that system configuration allows the API endpoint to accept such tokens.
	// @throws ServiceUnavailable  if session creation fails due to server specific issues, for example connection to a back end component is failing. Due to the security nature of this API further details will not be disclosed in the exception. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Create() (string, error)

	// Terminates the validity of a session token. This is the equivalent of log out.
	//
	//  A session token is expected as part of the request.
	// @throws Unauthenticated  if the session id is missing from the request or the corresponding session object cannot be found.
	// @throws ServiceUnavailable  if session deletion fails due to server specific issues, for example connection to a back end component is failing. Due to the security nature of this API further details will not be disclosed in the exception. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Delete() error

	// Returns information about the current session. This method expects a valid session token to be supplied.
	//
	//  A side effect of invoking this method may be a change to the session's last accessed time to the current time if this is supported by the session implementation. Invoking any other method in the API will also update the session's last accessed time.
	//
	//  This API is meant to serve the needs of various front end projects that may want to display the name of the user. Examples of this include various web based user interfaces and logging facilities.
	// @return Information about the session.
	// @throws Unauthenticated  if the session id is missing from the request or the corresponding session object cannot be found.
	// @throws ServiceUnavailable  if session retrieval fails due to server specific issues e.g. connection to back end component is failing. Due to the security nature of this API further details will not be disclosed in the error. Please refer to component health information, administrative logs and product specific documentation for possible causes.
	Get() (SessionInfo, error)
}

type sessionClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSessionClient(connector client.Connector) *sessionClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.cis.session")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := sessionClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sessionClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sessionClient) Create() (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sessionCreateInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.cis.session", "create", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sessionCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sessionClient) Delete() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sessionDeleteInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.cis.session", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sessionClient) Get() (SessionInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sessionGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput SessionInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.cis.session", "get", inputDataValue, executionContext)
	var emptyOutput SessionInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sessionGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(SessionInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

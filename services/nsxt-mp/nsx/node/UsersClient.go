// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Users
// Used by client-side stubs.

package node

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type UsersClient interface {

	// Activates the account for this user. When an account is successfully activated, the \"status\" field in the response is \"ACTIVE\". This API is not supported for userid 0 and userid 10000.
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPasswordPropertyParam (required)
	// @return com.vmware.nsx.model.NodeUserProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Activate(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) (model.NodeUserProperties, error)

	// Deactivates the account for this user. Deactivating an account is permanent, unlike an account that is temporarily locked because of too many password failures. A deactivated account has to be explicitly activated. When an account is successfully deactivated, the \"status\" field in the response is \"NOT_ACTIVATED\". This API is not supported for userid 0 and userid 10000.
	//
	// @param useridParam User id of the user (required)
	// @return com.vmware.nsx.model.NodeUserProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deactivate(useridParam string) (model.NodeUserProperties, error)

	// Returns information about a specified user who is configured to log in to the NSX appliance. The valid user IDs are: 0, 10000, 10002.
	//
	// @param useridParam User id of the user (required)
	// @return com.vmware.nsx.model.NodeUserProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(useridParam string) (model.NodeUserProperties, error)

	// Returns the list of users configured to log in to the NSX appliance.
	// @return com.vmware.nsx.model.NodeUserPropertiesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.NodeUserPropertiesListResult, error)

	// Enables a user to reset their own password.
	//
	// @param resetNodeUserOwnPasswordPropertiesParam (required)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetownpassword(resetNodeUserOwnPasswordPropertiesParam model.ResetNodeUserOwnPasswordProperties) error

	// Unlike the PUT version of this call (PUT /node/users/<userid>), this API does not require that the current password for the user be provided. The account of the target user must be \"ACTIVE\" for the call to succeed. This API only supports user IDs 10002, 10003, and 10004.
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPasswordPropertyParam (required)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetpassword(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error

	// Updates attributes of an existing NSX appliance user. This method cannot be used to add a new user. Modifiable attributes include the username, full name of the user, and password. If you specify a password in a PUT request, it is not returned in the response. Nor is it returned in a GET request. The specified password does not meet the following complexity requirements: - minimum 12 characters in length - minimum 1 uppercase character - minimum 1 lowercase character - minimum 1 numeric character - minimum 1 special character - minimum 5 unique characters - default password complexity rules as enforced by the Linux PAM module The valid user IDs are: 0, 10000, 10002. Note that invoking this API does not update any user-related properties of existing objects in the system and does not modify the username field in existing audit log entries.
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPropertiesParam (required)
	// @return com.vmware.nsx.model.NodeUserProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(useridParam string, nodeUserPropertiesParam model.NodeUserProperties) (model.NodeUserProperties, error)
}

type usersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUsersClient(connector client.Connector) *usersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.node.users")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"activate":         core.NewMethodIdentifier(interfaceIdentifier, "activate"),
		"deactivate":       core.NewMethodIdentifier(interfaceIdentifier, "deactivate"),
		"get":              core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":             core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"resetownpassword": core.NewMethodIdentifier(interfaceIdentifier, "resetownpassword"),
		"resetpassword":    core.NewMethodIdentifier(interfaceIdentifier, "resetpassword"),
		"update":           core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := usersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *usersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *usersClient) Activate(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) (model.NodeUserProperties, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersActivateInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserPasswordProperty", nodeUserPasswordPropertyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersActivateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "activate", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usersActivateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *usersClient) Deactivate(useridParam string) (model.NodeUserProperties, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersDeactivateInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersDeactivateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "deactivate", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usersDeactivateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *usersClient) Get(useridParam string) (model.NodeUserProperties, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersGetInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "get", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *usersClient) List() (model.NodeUserPropertiesListResult, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserPropertiesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "list", inputDataValue, executionContext)
	var emptyOutput model.NodeUserPropertiesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserPropertiesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *usersClient) Resetownpassword(resetNodeUserOwnPasswordPropertiesParam model.ResetNodeUserOwnPasswordProperties) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersResetownpasswordInputType(), typeConverter)
	sv.AddStructField("ResetNodeUserOwnPasswordProperties", resetNodeUserOwnPasswordPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersResetownpasswordRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "resetownpassword", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *usersClient) Resetpassword(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersResetpasswordInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserPasswordProperty", nodeUserPasswordPropertyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersResetpasswordRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "resetpassword", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *usersClient) Update(useridParam string, nodeUserPropertiesParam model.NodeUserProperties) (model.NodeUserProperties, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usersUpdateInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserProperties", nodeUserPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.users", "update", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

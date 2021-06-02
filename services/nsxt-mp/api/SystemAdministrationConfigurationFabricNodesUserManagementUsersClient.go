// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesUserManagementUsers
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationFabricNodesUserManagementUsersClient interface {

	// Activates the account for this user. When an account is successfully activated, the \"status\" field in the response is \"ACTIVE\". This API is not supported for userid 0 and userid 10000.
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPasswordPropertyParam (required)
	// @return com.vmware.model.NodeUserProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ActivateNodeUserActivate(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) (model.NodeUserProperties, error)

	// Add SSH public key to authorized_keys file for node user
	//
	// @param useridParam User id of the user (required)
	// @param sshKeyPropertiesParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddNodeUserSshKeyAddSshKey(useridParam string, sshKeyPropertiesParam model.SshKeyProperties) error

	// Deactivates the account for this user. Deactivating an account is permanent, unlike an account that is temporarily locked because of too many password failures. A deactivated account has to be explicitly activated. When an account is successfully deactivated, the \"status\" field in the response is \"NOT_ACTIVATED\". This API is not supported for userid 0 and userid 10000.
	//
	// @param useridParam User id of the user (required)
	// @return com.vmware.model.NodeUserProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeactivateNodeUserDeactivate(useridParam string) (model.NodeUserProperties, error)

	// Remove SSH public key from authorized_keys file for node user
	//
	// @param useridParam User id of the user (required)
	// @param sshKeyBasePropertiesParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteNodeUserSshKeyRemoveSshKey(useridParam string, sshKeyBasePropertiesParam model.SshKeyBaseProperties) error

	// Returns a list of all SSH keys from authorized_keys file for node user
	//
	// @param useridParam User id of the user (required)
	// @return com.vmware.model.SshKeyPropertiesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNodeUserSshKeys(useridParam string) (model.SshKeyPropertiesListResult, error)

	// Returns the list of users configued to log in to the NSX appliance.
	// @return com.vmware.model.NodeUserPropertiesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNodeUsers() (model.NodeUserPropertiesListResult, error)

	// Returns information about a specified user who is configued to log in to the NSX appliance. The valid user IDs are: 0, 10000, 10002.
	//
	// @param useridParam User id of the user (required)
	// @return com.vmware.model.NodeUserProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNodeUser(useridParam string) (model.NodeUserProperties, error)

	// Enables a user to reset their own password.
	//
	// @param nodeUserPasswordPropertyParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetNodeUserOwnPasswordResetOwnPassword(nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error

	//
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPasswordPropertyParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetNodeUserPasswordResetPassword(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error

	// Updates attributes of an existing NSX appliance user. This method cannot be used to add a new user. Modifiable attributes include the username, full name of the user, and password. If you specify a password in a PUT request, it is not returned in the response. Nor is it returned in a GET request. The specified password does not meet the following complexity requirements: - minimum 12 characters in length - minimum 1 uppercase character - minimum 1 lowercase character - minimum 1 numeric character - minimum 1 special character - minimum 5 unique characters - default password complexity rules as enforced by the Linux PAM module The valid user IDs are: 0, 10000, 10002. Note that invoking this API does not update any user-related properties of existing objects in the system and does not modify the username field in existing audit log entries.
	//
	// @param useridParam User id of the user (required)
	// @param nodeUserPropertiesParam (required)
	// @return com.vmware.model.NodeUserProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNodeUser(useridParam string, nodeUserPropertiesParam model.NodeUserProperties) (model.NodeUserProperties, error)
}

type systemAdministrationConfigurationFabricNodesUserManagementUsersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesUserManagementUsersClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesUserManagementUsersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"activate_node_user_activate":                     core.NewMethodIdentifier(interfaceIdentifier, "activate_node_user_activate"),
		"add_node_user_ssh_key_add_ssh_key":               core.NewMethodIdentifier(interfaceIdentifier, "add_node_user_ssh_key_add_ssh_key"),
		"deactivate_node_user_deactivate":                 core.NewMethodIdentifier(interfaceIdentifier, "deactivate_node_user_deactivate"),
		"delete_node_user_ssh_key_remove_ssh_key":         core.NewMethodIdentifier(interfaceIdentifier, "delete_node_user_ssh_key_remove_ssh_key"),
		"list_node_user_ssh_keys":                         core.NewMethodIdentifier(interfaceIdentifier, "list_node_user_ssh_keys"),
		"list_node_users":                                 core.NewMethodIdentifier(interfaceIdentifier, "list_node_users"),
		"read_node_user":                                  core.NewMethodIdentifier(interfaceIdentifier, "read_node_user"),
		"reset_node_user_own_password_reset_own_password": core.NewMethodIdentifier(interfaceIdentifier, "reset_node_user_own_password_reset_own_password"),
		"reset_node_user_password_reset_password":         core.NewMethodIdentifier(interfaceIdentifier, "reset_node_user_password_reset_password"),
		"update_node_user":                                core.NewMethodIdentifier(interfaceIdentifier, "update_node_user"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesUserManagementUsersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ActivateNodeUserActivate(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) (model.NodeUserProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserPasswordProperty", nodeUserPasswordPropertyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "activate_node_user_activate", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) AddNodeUserSshKeyAddSshKey(useridParam string, sshKeyPropertiesParam model.SshKeyProperties) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersAddNodeUserSshKeyAddSshKeyInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("SshKeyProperties", sshKeyPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersAddNodeUserSshKeyAddSshKeyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "add_node_user_ssh_key_add_ssh_key", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) DeactivateNodeUserDeactivate(useridParam string) (model.NodeUserProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "deactivate_node_user_deactivate", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) DeleteNodeUserSshKeyRemoveSshKey(useridParam string, sshKeyBasePropertiesParam model.SshKeyBaseProperties) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersDeleteNodeUserSshKeyRemoveSshKeyInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("SshKeyBaseProperties", sshKeyBasePropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersDeleteNodeUserSshKeyRemoveSshKeyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "delete_node_user_ssh_key_remove_ssh_key", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ListNodeUserSshKeys(useridParam string) (model.SshKeyPropertiesListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SshKeyPropertiesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "list_node_user_ssh_keys", inputDataValue, executionContext)
	var emptyOutput model.SshKeyPropertiesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SshKeyPropertiesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ListNodeUsers() (model.NodeUserPropertiesListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserPropertiesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "list_node_users", inputDataValue, executionContext)
	var emptyOutput model.NodeUserPropertiesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserPropertiesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ReadNodeUser(useridParam string) (model.NodeUserProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "read_node_user", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ResetNodeUserOwnPasswordResetOwnPassword(nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserOwnPasswordResetOwnPasswordInputType(), typeConverter)
	sv.AddStructField("NodeUserPasswordProperty", nodeUserPasswordPropertyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserOwnPasswordResetOwnPasswordRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "reset_node_user_own_password_reset_own_password", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) ResetNodeUserPasswordResetPassword(useridParam string, nodeUserPasswordPropertyParam model.NodeUserPasswordProperty) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserPasswordResetPasswordInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserPasswordProperty", nodeUserPasswordPropertyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserPasswordResetPasswordRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "reset_node_user_password_reset_password", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesUserManagementUsersClient) UpdateNodeUser(useridParam string, nodeUserPropertiesParam model.NodeUserProperties) (model.NodeUserProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserInputType(), typeConverter)
	sv.AddStructField("Userid", useridParam)
	sv.AddStructField("NodeUserProperties", nodeUserPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeUserProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_user_management_users", "update_node_user", inputDataValue, executionContext)
	var emptyOutput model.NodeUserProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeUserProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

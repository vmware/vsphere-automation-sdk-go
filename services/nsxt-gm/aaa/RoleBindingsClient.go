// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RoleBindings
// Used by client-side stubs.

package aaa

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RoleBindingsClient interface {

	// This API is used to assign a user/group any role(s) of choice. It is recommended to use the new property roles_for_paths instead of roles. When using the roles_for_paths, set the read_roles_for_paths as true. User has union of all the roles assigned to it on a particular path and its sub-tree. User name is dealt case-insensitively.
	//
	// @param roleBindingParam (required)
	// @return com.vmware.nsx_global_policy.model.RoleBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(roleBindingParam nsx_global_policyModel.RoleBinding) (nsx_global_policyModel.RoleBinding, error)

	// Delete the user/group's role assignment. If the path is provided then deletes only the roles_for_paths that matches the path. If path is provided for the last roles_for_paths then the whole role binding is deleted provided it is not that of a local user.
	//
	// @param bindingIdParam User/Group's id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error

	// Delete all stale role assignments
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deletestalebindings(cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error

	// Get user/group's role information
	//
	// @param bindingIdParam User/Group's id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	// @return com.vmware.nsx_global_policy.model.RoleBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_global_policyModel.RoleBinding, error)

	// Get all users and groups with their roles. If the root_path is provided then only return role bindings that start-with or are sub-trees of the provided root path. Also filter the roles_for_paths such that only those roles_for_paths appear that start-with or are sub-tree of the provided root path.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	// @return com.vmware.nsx_global_policy.model.RoleBindingListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_global_policyModel.RoleBindingListResult, error)

	// This API is used to update a user/group any role(s) of choice. It is recommended to use the new property roles_for_paths instead of roles. When using the roles_for_paths, set the read_roles_for_paths as true. User has union of all the roles assigned to it on a particular path and its sub-tree. User name is dealt case-insensitively. This API will merge the existing roles_for_paths with the newly provided roles_for_paths excluding roles_for_paths those are marked for deletion.
	//
	// @param bindingIdParam User/Group's id (required)
	// @param roleBindingParam (required)
	// @return com.vmware.nsx_global_policy.model.RoleBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(bindingIdParam string, roleBindingParam nsx_global_policyModel.RoleBinding) (nsx_global_policyModel.RoleBinding, error)
}

type roleBindingsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRoleBindingsClient(connector vapiProtocolClient_.Connector) *roleBindingsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.aaa.role_bindings")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"deletestalebindings": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "deletestalebindings"),
		"get":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := roleBindingsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *roleBindingsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *roleBindingsClient) Create(roleBindingParam nsx_global_policyModel.RoleBinding) (nsx_global_policyModel.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsCreateInputType(), typeConverter)
	sv.AddStructField("RoleBinding", roleBindingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.RoleBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "create", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) Delete(bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsDeleteInputType(), typeConverter)
	sv.AddStructField("BindingId", bindingIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *roleBindingsClient) Deletestalebindings(cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsDeletestalebindingsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsDeletestalebindingsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "deletestalebindings", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *roleBindingsClient) Get(bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_global_policyModel.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsGetInputType(), typeConverter)
	sv.AddStructField("BindingId", bindingIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.RoleBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) List(cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_global_policyModel.RoleBindingListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.RoleBindingListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.RoleBindingListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.RoleBindingListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) Update(bindingIdParam string, roleBindingParam nsx_global_policyModel.RoleBinding) (nsx_global_policyModel.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsUpdateInputType(), typeConverter)
	sv.AddStructField("BindingId", bindingIdParam)
	sv.AddStructField("RoleBinding", roleBindingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.RoleBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.role_bindings", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

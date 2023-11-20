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
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RoleBindingsClient interface {

	// Delete the user/group's role assignment on CSP. If the path is provided then deletes only the roles_for_paths that matches the path. If path is provided for the last roles_for_paths then the whole role binding is deleted provided it is not that of a local user.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param bindingIdParam (required)
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
	Delete(orgIdParam string, projectIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error

	// Get user/group's role information from CSP
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param bindingIdParam (required)
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
	// @return com.vmware.nsx_policy.model.RoleBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_policyModel.RoleBinding, error)

	// Get all users and groups with their roles from CSP. If the root_path is provided then only return role bindings that start-with or are sub-trees of the provided root path. Also filter the roles_for_paths such that only those roles_for_paths appear that start-with or are sub-tree of the provided root path.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
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
	// @return com.vmware.nsx_policy.model.RoleBindingListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_policyModel.RoleBindingListResult, error)

	// This API is used to assign a user/group any role(s) of choice on CSP. It is recommended to use the new property roles_for_paths instead of roles. When using the roles_for_paths, set the read_roles_for_paths as true. User has union of all the roles assigned to it on a particular path and its sub-tree. User name is dealt case-insensitively.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param roleBindingParam (required)
	// @return com.vmware.nsx_policy.model.RoleBinding
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, roleBindingParam nsx_policyModel.RoleBinding) (nsx_policyModel.RoleBinding, error)
}

type roleBindingsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRoleBindingsClient(connector vapiProtocolClient_.Connector) *roleBindingsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.aaa.role_bindings")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
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

func (rIface *roleBindingsClient) Delete(orgIdParam string, projectIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
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

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.aaa.role_bindings", "delete", inputDataValue, executionContext)
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

func (rIface *roleBindingsClient) Get(orgIdParam string, projectIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_policyModel.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
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
		var emptyOutput nsx_policyModel.RoleBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.aaa.role_bindings", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) List(orgIdParam string, projectIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsx_policyModel.RoleBindingListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
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
		var emptyOutput nsx_policyModel.RoleBindingListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.aaa.role_bindings", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.RoleBindingListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.RoleBindingListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) Patch(orgIdParam string, projectIdParam string, roleBindingParam nsx_policyModel.RoleBinding) (nsx_policyModel.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := roleBindingsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(roleBindingsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("RoleBinding", roleBindingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.RoleBinding
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.aaa.role_bindings", "patch", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoleBindingsPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

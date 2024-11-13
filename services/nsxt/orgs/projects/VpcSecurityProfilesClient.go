// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VpcSecurityProfiles
// Used by client-side stubs.

package projects

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VpcSecurityProfilesClient interface {

	// Get security profile with the specified ID. Currently only the default security profile at the project level can be queried using this API.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param securityProfileIdParam (required)
	// @return com.vmware.nsx_policy.model.VpcSecurityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, securityProfileIdParam string) (nsx_policyModel.VpcSecurityProfile, error)

	// Paginated list of VpcSecurityProfiles. Note: Currently this API will return only a single default security profile at the project level.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.VpcSecurityProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.VpcSecurityProfileListResult, error)

	// Update existing default security profile. Note: Currently, user will not be able to create new security profiles.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcSecurityProfileIdParam (required)
	// @param vpcSecurityProfileParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, vpcSecurityProfileParam nsx_policyModel.VpcSecurityProfile) error

	// Update existing security profile. Note: Currently, user will not be able to create new security profiles.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcSecurityProfileIdParam (required)
	// @param vpcSecurityProfileParam (required)
	// @return com.vmware.nsx_policy.model.VpcSecurityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, vpcSecurityProfileParam nsx_policyModel.VpcSecurityProfile) (nsx_policyModel.VpcSecurityProfile, error)
}

type vpcSecurityProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVpcSecurityProfilesClient(connector vapiProtocolClient_.Connector) *vpcSecurityProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vpcSecurityProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vpcSecurityProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vpcSecurityProfilesClient) Get(orgIdParam string, projectIdParam string, securityProfileIdParam string) (nsx_policyModel.VpcSecurityProfile, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcSecurityProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcSecurityProfilesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("SecurityProfileId", securityProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VpcSecurityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VpcSecurityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcSecurityProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VpcSecurityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcSecurityProfilesClient) List(orgIdParam string, projectIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.VpcSecurityProfileListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcSecurityProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcSecurityProfilesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VpcSecurityProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VpcSecurityProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcSecurityProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VpcSecurityProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcSecurityProfilesClient) Patch(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, vpcSecurityProfileParam nsx_policyModel.VpcSecurityProfile) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcSecurityProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcSecurityProfilesPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("VpcSecurityProfile", vpcSecurityProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vpcSecurityProfilesClient) Update(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, vpcSecurityProfileParam nsx_policyModel.VpcSecurityProfile) (nsx_policyModel.VpcSecurityProfile, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcSecurityProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcSecurityProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("VpcSecurityProfile", vpcSecurityProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VpcSecurityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VpcSecurityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcSecurityProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VpcSecurityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

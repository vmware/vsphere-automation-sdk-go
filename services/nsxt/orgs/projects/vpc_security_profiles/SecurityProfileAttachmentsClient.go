// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SecurityProfileAttachments
// Used by client-side stubs.

package vpc_security_profiles

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SecurityProfileAttachmentsClient interface {

	// Deletion of the VPC Security profile attachment to VPC result in detaching the security profile to VPC.
	//
	// @param orgIdParam Organisation identifier (required)
	// @param projectIdParam Project identifier (required)
	// @param vpcSecurityProfileIdParam VPC Security Profile identifier (required)
	// @param securityProfileAttachmentIdParam Security Profile attachment identifier (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string) error

	// Paginated list of SecurityProfileAttachments.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcSecurityProfileIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.SecurityProfileAttachmentListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SecurityProfileAttachmentListResult, error)

	// create of update existing security profile attachment of a VPC.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcSecurityProfileIdParam (required)
	// @param securityProfileAttachmentIdParam (required)
	// @param securityProfileAttachmentParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string, securityProfileAttachmentParam nsx_policyModel.SecurityProfileAttachment) error

	// Update existing security profile attachment of a VPC.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcSecurityProfileIdParam (required)
	// @param securityProfileAttachmentIdParam (required)
	// @param securityProfileAttachmentParam (required)
	// @return com.vmware.nsx_policy.model.SecurityProfileAttachment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string, securityProfileAttachmentParam nsx_policyModel.SecurityProfileAttachment) (nsx_policyModel.SecurityProfileAttachment, error)
}

type securityProfileAttachmentsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSecurityProfileAttachmentsClient(connector vapiProtocolClient_.Connector) *securityProfileAttachmentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles.security_profile_attachments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := securityProfileAttachmentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *securityProfileAttachmentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *securityProfileAttachmentsClient) Delete(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := securityProfileAttachmentsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(securityProfileAttachmentsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("SecurityProfileAttachmentId", securityProfileAttachmentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles.security_profile_attachments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *securityProfileAttachmentsClient) List(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SecurityProfileAttachmentListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := securityProfileAttachmentsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(securityProfileAttachmentsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SecurityProfileAttachmentListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles.security_profile_attachments", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SecurityProfileAttachmentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SecurityProfileAttachmentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SecurityProfileAttachmentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *securityProfileAttachmentsClient) Patch(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string, securityProfileAttachmentParam nsx_policyModel.SecurityProfileAttachment) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := securityProfileAttachmentsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(securityProfileAttachmentsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("SecurityProfileAttachmentId", securityProfileAttachmentIdParam)
	sv.AddStructField("SecurityProfileAttachment", securityProfileAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles.security_profile_attachments", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *securityProfileAttachmentsClient) Update(orgIdParam string, projectIdParam string, vpcSecurityProfileIdParam string, securityProfileAttachmentIdParam string, securityProfileAttachmentParam nsx_policyModel.SecurityProfileAttachment) (nsx_policyModel.SecurityProfileAttachment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := securityProfileAttachmentsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(securityProfileAttachmentsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcSecurityProfileId", vpcSecurityProfileIdParam)
	sv.AddStructField("SecurityProfileAttachmentId", securityProfileAttachmentIdParam)
	sv.AddStructField("SecurityProfileAttachment", securityProfileAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SecurityProfileAttachment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpc_security_profiles.security_profile_attachments", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SecurityProfileAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SecurityProfileAttachmentsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SecurityProfileAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

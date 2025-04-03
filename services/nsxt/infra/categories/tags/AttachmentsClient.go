// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Attachments
// Used by client-side stubs.

package tags

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type AttachmentsClient interface {

    // Delete tag attachments
    //
    // @param categoryIdParam (required)
    // @param tagIdParam (required)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(categoryIdParam string, tagIdParam string) error

    // Get tag attachments
    //
    // @param categoryIdParam (required)
    // @param tagIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param resourceTypeParam Tag Attachment Type (optional)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.TagAttachmentResponse
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(categoryIdParam string, tagIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TagAttachmentResponse, error)

    // Create or update vm tag attachment
    //
    // @param categoryIdParam (required)
    // @param tagIdParam (required)
    // @param tagAttachmentParam (required)
    // @return com.vmware.nsx_policy.model.TagAttachmentResponse
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(categoryIdParam string, tagIdParam string, tagAttachmentParam nsx_policyModel.TagAttachment) (nsx_policyModel.TagAttachmentResponse, error)
}


type attachmentsClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewAttachmentsClient(connector vapiProtocolClient_.Connector) *attachmentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.categories.tags.attachments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := attachmentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *attachmentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *attachmentsClient) Delete(categoryIdParam string, tagIdParam string) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := attachmentsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(attachmentsDeleteInputType(), typeConverter)
	sv.AddStructField("CategoryId", categoryIdParam)
	sv.AddStructField("TagId", tagIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.categories.tags.attachments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (aIface *attachmentsClient) Get(categoryIdParam string, tagIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TagAttachmentResponse, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := attachmentsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(attachmentsGetInputType(), typeConverter)
	sv.AddStructField("CategoryId", categoryIdParam)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ResourceType", resourceTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TagAttachmentResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.categories.tags.attachments", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TagAttachmentResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AttachmentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TagAttachmentResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *attachmentsClient) Patch(categoryIdParam string, tagIdParam string, tagAttachmentParam nsx_policyModel.TagAttachment) (nsx_policyModel.TagAttachmentResponse, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := attachmentsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(attachmentsPatchInputType(), typeConverter)
	sv.AddStructField("CategoryId", categoryIdParam)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("TagAttachment", tagAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TagAttachmentResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.categories.tags.attachments", "patch", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TagAttachmentResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AttachmentsPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TagAttachmentResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


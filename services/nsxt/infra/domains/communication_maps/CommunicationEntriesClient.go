// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CommunicationEntries
// Used by client-side stubs.

package communication_maps

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CommunicationEntriesClient interface {

	// Delete CommunicationEntry
	//  This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param communicationEntryIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) error

	// Read CommunicationEntry
	//  This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param communicationEntryIdParam (required)
	// @return com.vmware.nsx_policy.model.CommunicationEntry
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) (nsx_policyModel.CommunicationEntry, error)

	// List CommunicationEntries
	//  This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.CommunicationEntryListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(domainIdParam string, communicationMapIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.CommunicationEntryListResult, error)

	// Patch the CommunicationEntry. If a communication entry for the given communication-entry-id is not present, the object will get created and if it is present it will be updated. This is a full replace
	//  This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param communicationEntryIdParam (required)
	// @param communicationEntryParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry) error

	// This is used to re-order a communictation entry within a communication map.
	//  This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id?action=revise
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param communicationEntryIdParam (required)
	// @param communicationEntryParam (required)
	// @param anchorPathParam The communication map/communication entry path if operation is 'insert_after' or 'insert_before' (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx_policy.model.CommunicationEntry
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry, anchorPathParam *string, operationParam *string) (nsx_policyModel.CommunicationEntry, error)

	// Update the CommunicationEntry. If a CommunicationEntry with the communication-entry-id is not already present, this API fails with a 404. Creation of CommunicationEntries is not allowed using this API.
	//  This API is deprecated. Please use the following API instead PUT /infra/domains/domain-id/security-policies/securit-policy-id/rules/rule-id
	//
	// Deprecated: This API element is deprecated.
	//
	// @param domainIdParam (required)
	// @param communicationMapIdParam (required)
	// @param communicationEntryIdParam (required)
	// @param communicationEntryParam (required)
	// @return com.vmware.nsx_policy.model.CommunicationEntry
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry) (nsx_policyModel.CommunicationEntry, error)
}

type communicationEntriesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCommunicationEntriesClient(connector vapiProtocolClient_.Connector) *communicationEntriesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"revise": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := communicationEntriesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *communicationEntriesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *communicationEntriesClient) Delete(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesDeleteInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("CommunicationEntryId", communicationEntryIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *communicationEntriesClient) Get(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) (nsx_policyModel.CommunicationEntry, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesGetInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("CommunicationEntryId", communicationEntryIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.CommunicationEntry
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.CommunicationEntry
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommunicationEntriesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.CommunicationEntry), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *communicationEntriesClient) List(domainIdParam string, communicationMapIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.CommunicationEntryListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesListInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.CommunicationEntryListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.CommunicationEntryListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommunicationEntriesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.CommunicationEntryListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *communicationEntriesClient) Patch(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesPatchInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("CommunicationEntryId", communicationEntryIdParam)
	sv.AddStructField("CommunicationEntry", communicationEntryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *communicationEntriesClient) Revise(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry, anchorPathParam *string, operationParam *string) (nsx_policyModel.CommunicationEntry, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesReviseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesReviseInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("CommunicationEntryId", communicationEntryIdParam)
	sv.AddStructField("CommunicationEntry", communicationEntryParam)
	sv.AddStructField("AnchorPath", anchorPathParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.CommunicationEntry
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "revise", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.CommunicationEntry
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommunicationEntriesReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.CommunicationEntry), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *communicationEntriesClient) Update(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam nsx_policyModel.CommunicationEntry) (nsx_policyModel.CommunicationEntry, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := communicationEntriesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(communicationEntriesUpdateInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("CommunicationMapId", communicationMapIdParam)
	sv.AddStructField("CommunicationEntryId", communicationEntryIdParam)
	sv.AddStructField("CommunicationEntry", communicationEntryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.CommunicationEntry
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.communication_maps.communication_entries", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.CommunicationEntry
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommunicationEntriesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.CommunicationEntry), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

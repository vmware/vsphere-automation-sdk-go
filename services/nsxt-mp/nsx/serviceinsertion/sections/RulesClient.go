// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Rules
// Used by client-side stubs.

package sections

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RulesClient interface {

	// Adds a new serviceinsertion rule in existing serviceinsertion section.
	//  Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(sectionIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRule, error)

	// Create multiple serviceinsertion rules in existing serviceinsertion section bounded by limit of 1000 serviceinsertion rules per section.
	//  Note- POST service insertion rules API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createmultiple(sectionIdParam string, serviceInsertionRuleListParam nsxModel.ServiceInsertionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRuleList, error)

	// Delete existing serviceinsertion rule in a serviceinsertion section.
	//  Note- DELETE service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(sectionIdParam string, ruleIdParam string) error

	// Return existing serviceinsertion rule information in a serviceinsertion section. Note- GET service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string, ruleIdParam string) (nsxModel.ServiceInsertionRule, error)

	// Return all serviceinsertion rule(s) information for a given serviceinsertion section.
	//  Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @return com.vmware.nsx.model.ServiceInsertionRuleListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(sectionIdParam string, appliedTosParam *string, cursorParam *string, destinationsParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (nsxModel.ServiceInsertionRuleListResult, error)

	// Modifies existing serviceinsertion rule along with relative position among other serviceinsertion rules inside a serviceinsertion section.
	//  Note- POST service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRule, error)

	// Modifies existing serviceinsertion rule in a serviceinsertion section.
	//  Note- PUT service insertion rule API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param serviceInsertionRuleParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule) (nsxModel.ServiceInsertionRule, error)
}

type rulesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRulesClient(connector vapiProtocolClient_.Connector) *rulesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.sections.rules")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createmultiple": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "createmultiple"),
		"delete":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"revise":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"update":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := rulesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *rulesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *rulesClient) Create(sectionIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesCreateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Createmultiple(sectionIdParam string, serviceInsertionRuleListParam nsxModel.ServiceInsertionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRuleList, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesCreatemultipleRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesCreatemultipleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionRuleList", serviceInsertionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "createmultiple", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesCreatemultipleOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Delete(sectionIdParam string, ruleIdParam string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesDeleteInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "delete", inputDataValue, executionContext)
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

func (rIface *rulesClient) Get(sectionIdParam string, ruleIdParam string) (nsxModel.ServiceInsertionRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) List(sectionIdParam string, appliedTosParam *string, cursorParam *string, destinationsParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (nsxModel.ServiceInsertionRuleListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesListInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRuleListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRuleListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRuleListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Revise(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule, idParam *string, operationParam *string) (nsxModel.ServiceInsertionRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesReviseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "revise", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Update(sectionIdParam string, ruleIdParam string, serviceInsertionRuleParam nsxModel.ServiceInsertionRule) (nsxModel.ServiceInsertionRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := rulesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(rulesUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("ServiceInsertionRule", serviceInsertionRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections.rules", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RulesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

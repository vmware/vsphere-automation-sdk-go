// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Rules
// Used by client-side stubs.

package sections

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type RulesClient interface {

	// Adds a new firewall rule in existing firewall section. Adding firewall rule to a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale content to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(sectionIdParam string, firewallRuleParam model.FirewallRule, idParam *string, operationParam *string) (model.FirewallRule, error)

	// Create multiple firewall rules in existing firewall section bounded by limit of 1000 firewall rules per section. Adding multiple firewall rules in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createmultiple(sectionIdParam string, firewallRuleListParam model.FirewallRuleList, idParam *string, operationParam *string) (model.FirewallRuleList, error)

	// Delete existing firewall rule in a firewall section. Deleting firewall rule in a section modifies parent section and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(sectionIdParam string, ruleIdParam string) error

	// Return existing firewall rule information in a firewall section.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @return com.vmware.nsx.model.FirewallRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string, ruleIdParam string) (model.FirewallRule, error)

	// Return all firewall rule(s) information for a given firewall section.
	//
	// @param sectionIdParam (required)
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param contextProfilesParam Limits results to sections having rules with specific Context Profiles. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param deepSearchParam Toggle to search with direct or indirect references. (optional, default to false)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param extendedSourcesParam Limits results to sections having rules with specific Extended Sources. (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param searchInvalidReferencesParam Return invalid references in results. (optional, default to false)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @return com.vmware.nsx.model.FirewallRuleListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(sectionIdParam string, appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, extendedSourcesParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, searchInvalidReferencesParam *bool, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (model.FirewallRuleListResult, error)

	// Modifies existing firewall rule along with relative position among other firewall rules inside a firewall section. Revising firewall rule in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param firewallRuleParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, ruleIdParam string, firewallRuleParam model.FirewallRule, idParam *string, operationParam *string) (model.FirewallRule, error)

	// Modifies existing firewall rule in a firewall section. Updating firewall rule in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param firewallRuleParam (required)
	// @return com.vmware.nsx.model.FirewallRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, ruleIdParam string, firewallRuleParam model.FirewallRule) (model.FirewallRule, error)
}

type rulesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewRulesClient(connector client.Connector) *rulesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.firewall.sections.rules")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":         core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createmultiple": core.NewMethodIdentifier(interfaceIdentifier, "createmultiple"),
		"delete":         core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":            core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":           core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"revise":         core.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"update":         core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := rulesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *rulesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *rulesClient) Create(sectionIdParam string, firewallRuleParam model.FirewallRule, idParam *string, operationParam *string) (model.FirewallRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesCreateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallRule", firewallRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "create", inputDataValue, executionContext)
	var emptyOutput model.FirewallRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Createmultiple(sectionIdParam string, firewallRuleListParam model.FirewallRuleList, idParam *string, operationParam *string) (model.FirewallRuleList, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesCreatemultipleInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallRuleList", firewallRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesCreatemultipleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "createmultiple", inputDataValue, executionContext)
	var emptyOutput model.FirewallRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesCreatemultipleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Delete(sectionIdParam string, ruleIdParam string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesDeleteInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *rulesClient) Get(sectionIdParam string, ruleIdParam string) (model.FirewallRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "get", inputDataValue, executionContext)
	var emptyOutput model.FirewallRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) List(sectionIdParam string, appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, extendedSourcesParam *string, filterTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, searchInvalidReferencesParam *bool, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string) (model.FirewallRuleListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesListInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("ContextProfiles", contextProfilesParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DeepSearch", deepSearchParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("ExtendedSources", extendedSourcesParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SearchInvalidReferences", searchInvalidReferencesParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRuleListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "list", inputDataValue, executionContext)
	var emptyOutput model.FirewallRuleListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRuleListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Revise(sectionIdParam string, ruleIdParam string, firewallRuleParam model.FirewallRule, idParam *string, operationParam *string) (model.FirewallRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("FirewallRule", firewallRuleParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesReviseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "revise", inputDataValue, executionContext)
	var emptyOutput model.FirewallRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *rulesClient) Update(sectionIdParam string, ruleIdParam string, firewallRuleParam model.FirewallRule) (model.FirewallRule, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(rulesUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("FirewallRule", firewallRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := rulesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules", "update", inputDataValue, executionContext)
	var emptyOutput model.FirewallRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), rulesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

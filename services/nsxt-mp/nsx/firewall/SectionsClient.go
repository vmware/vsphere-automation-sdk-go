// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Sections
// Used by client-side stubs.

package firewall

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SectionsClient interface {

	// Creates new empty firewall section in the system.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(firewallSectionParam nsxModel.FirewallSection, idParam *string, operationParam *string) (nsxModel.FirewallSection, error)

	// Creates a new firewall section with rules. The limit on the number of rules is defined by maxItems in collection types for FirewallRule (FirewallRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number of rule references is not supported. Instead, to create sections, use: POST /api/v1/firewall/sections To create rules, use: POST /api/v1/firewall/sections/<section-id>/rules
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(firewallSectionRuleListParam nsxModel.FirewallSectionRuleList, idParam *string, operationParam *string) (nsxModel.FirewallSectionRuleList, error)

	// Removes firewall section from the system. Firewall section with rules can only be deleted by passing \"cascade=true\" parameter.
	//
	//  Use the following Policy API -
	//  DELETE /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param cascadeParam Flag to cascade delete of this object to all it's child objects. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(sectionIdParam string, cascadeParam *bool) error

	// Returns information about firewall section for the identifier.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string) (nsxModel.FirewallSection, error)

	//
	//
	// Deprecated: This API element is deprecated.
	//
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param contextProfilesParam Limits results to sections having rules with specific Context Profiles. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param deepSearchParam Toggle to search with direct or indirect references. (optional, default to false)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param enforcedOnParam Type of attachment for logical port; for query only. (optional)
	// @param excludeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param extendedSourcesParam Limits results to sections having rules with specific Extended Sources. (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param lockedParam Limit results to sections which are locked/unlocked (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param searchInvalidReferencesParam Return invalid references in results. (optional, default to false)
	// @param searchScopeParam Limit result to sections of a specific enforcement point (optional)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @param type_Param Section Type (optional, default to LAYER3)
	// @return com.vmware.nsx.model.FirewallSectionListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, enforcedOnParam *string, excludeAppliedToTypeParam *string, extendedSourcesParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, lockedParam *bool, pageSizeParam *int64, searchInvalidReferencesParam *bool, searchScopeParam *string, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (nsxModel.FirewallSectionListResult, error)

	// Returns firewall section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number rule references is not supported. Instead, to read firewall rules, use: GET /api/v1/firewall/sections/<section-id>/rules with the appropriate page_size.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Listwithrules(sectionIdParam string) (nsxModel.FirewallSectionRuleList, error)

	// Lock a section.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionLockParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws ResourceBusy  Locked
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Lock(sectionIdParam string, firewallSectionLockParam nsxModel.FirewallSectionLock) (nsxModel.FirewallSection, error)

	// Modifies an existing firewall section along with its relative position among other firewall sections in the system. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	//  Use the following Policy API -
	//  POST /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>?action=revise
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, firewallSectionParam nsxModel.FirewallSection, idParam *string, operationParam *string) (nsxModel.FirewallSection, error)

	// Modifies an existing firewall section along with its relative position among other firewall sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to move a section above or below another section, use: POST /api/v1/firewall/sections/<section-id>?action=revise To modify rules, use: PUT /api/v1/firewall/sections/<section-id>/rules/<rule-id> Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	//  Use the following Policy API -
	//  POST /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>?action=revise
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revisewithrules(sectionIdParam string, firewallSectionRuleListParam nsxModel.FirewallSectionRuleList, idParam *string, operationParam *string) (nsxModel.FirewallSectionRuleList, error)

	// Unlock a section.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionLockParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws ResourceBusy  Locked
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Unlock(sectionIdParam string, firewallSectionLockParam nsxModel.FirewallSectionLock) (nsxModel.FirewallSection, error)

	// Modifies the specified section, but does not modify the section's associated rules. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, firewallSectionParam nsxModel.FirewallSection) (nsxModel.FirewallSection, error)

	// Modifies existing firewall section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to update rule content, use: PUT /api/v1/firewall/sections/<section-id>/rules/<rule-id> Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionRuleListParam (required)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(sectionIdParam string, firewallSectionRuleListParam nsxModel.FirewallSectionRuleList) (nsxModel.FirewallSectionRuleList, error)
}

type sectionsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSectionsClient(connector vapiProtocolClient_.Connector) *sectionsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.firewall.sections")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"listwithrules":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "listwithrules"),
		"lock":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "lock"),
		"revise":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"revisewithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revisewithrules"),
		"unlock":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "unlock"),
		"update":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatewithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "updatewithrules"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := sectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sectionsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sectionsClient) Create(firewallSectionParam nsxModel.FirewallSection, idParam *string, operationParam *string) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsCreateInputType(), typeConverter)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Createwithrules(firewallSectionRuleListParam nsxModel.FirewallSectionRuleList, idParam *string, operationParam *string) (nsxModel.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsCreatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "createwithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Delete(sectionIdParam string, cascadeParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsDeleteInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("Cascade", cascadeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "delete", inputDataValue, executionContext)
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

func (sIface *sectionsClient) Get(sectionIdParam string) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) List(appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, enforcedOnParam *string, excludeAppliedToTypeParam *string, extendedSourcesParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, lockedParam *bool, pageSizeParam *int64, searchInvalidReferencesParam *bool, searchScopeParam *string, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (nsxModel.FirewallSectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsListInputType(), typeConverter)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("ContextProfiles", contextProfilesParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DeepSearch", deepSearchParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("EnforcedOn", enforcedOnParam)
	sv.AddStructField("ExcludeAppliedToType", excludeAppliedToTypeParam)
	sv.AddStructField("ExtendedSources", extendedSourcesParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludeAppliedToType", includeAppliedToTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Locked", lockedParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SearchInvalidReferences", searchInvalidReferencesParam)
	sv.AddStructField("SearchScope", searchScopeParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSectionListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Listwithrules(sectionIdParam string) (nsxModel.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsListwithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsListwithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "listwithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsListwithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Lock(sectionIdParam string, firewallSectionLockParam nsxModel.FirewallSectionLock) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsLockRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsLockInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionLock", firewallSectionLockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "lock", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsLockOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revise(sectionIdParam string, firewallSectionParam nsxModel.FirewallSection, idParam *string, operationParam *string) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsReviseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "revise", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revisewithrules(sectionIdParam string, firewallSectionRuleListParam nsxModel.FirewallSectionRuleList, idParam *string, operationParam *string) (nsxModel.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsRevisewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsRevisewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "revisewithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsRevisewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Unlock(sectionIdParam string, firewallSectionLockParam nsxModel.FirewallSectionLock) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsUnlockRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsUnlockInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionLock", firewallSectionLockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "unlock", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsUnlockOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Update(sectionIdParam string, firewallSectionParam nsxModel.FirewallSection) (nsxModel.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Updatewithrules(sectionIdParam string, firewallSectionRuleListParam nsxModel.FirewallSectionRuleList) (nsxModel.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsUpdatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

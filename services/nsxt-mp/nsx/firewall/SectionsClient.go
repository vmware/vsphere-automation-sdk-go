// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Sections
// Used by client-side stubs.

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type SectionsClient interface {

	// Creates new empty firewall section in the system.
	//
	// @param firewallSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(firewallSectionParam model.FirewallSection, idParam *string, operationParam *string) (model.FirewallSection, error)

	// Creates a new firewall section with rules. The limit on the number of rules is defined by maxItems in collection types for FirewallRule (FirewallRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number of rule references is not supported. Instead, to create sections, use: POST /api/v1/firewall/sections To create rules, use: POST /api/v1/firewall/sections/<section-id>/rules
	//
	// @param firewallSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(firewallSectionRuleListParam model.FirewallSectionRuleList, idParam *string, operationParam *string) (model.FirewallSectionRuleList, error)

	// Removes firewall section from the system. Firewall section with rules can only be deleted by passing \"cascade=true\" parameter.
	//
	// @param sectionIdParam (required)
	// @param cascadeParam Flag to cascade delete of this object to all it's child objects. (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(sectionIdParam string, cascadeParam *bool) error

	// Returns information about firewall section for the identifier.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string) (model.FirewallSection, error)

	// List all firewall section in paginated form. A default page size is limited to 1000 firewall sections. By default list of section is filtered by LAYER3 type.
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
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, enforcedOnParam *string, excludeAppliedToTypeParam *string, extendedSourcesParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, lockedParam *bool, pageSizeParam *int64, searchInvalidReferencesParam *bool, searchScopeParam *string, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.FirewallSectionListResult, error)

	// Returns firewall section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number rule references is not supported. Instead, to read firewall rules, use: GET /api/v1/firewall/sections/<section-id>/rules with the appropriate page_size.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Listwithrules(sectionIdParam string) (model.FirewallSectionRuleList, error)

	// Lock a section.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionLockParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws ResourceBusy  Locked
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Lock(sectionIdParam string, firewallSectionLockParam model.FirewallSectionLock) (model.FirewallSection, error)

	// Modifies an existing firewall section along with its relative position among other firewall sections in the system. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, firewallSectionParam model.FirewallSection, idParam *string, operationParam *string) (model.FirewallSection, error)

	// Modifies an existing firewall section along with its relative position among other firewall sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to move a section above or below another section, use: POST /api/v1/firewall/sections/<section-id>?action=revise To modify rules, use: PUT /api/v1/firewall/sections/<section-id>/rules/<rule-id> Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revisewithrules(sectionIdParam string, firewallSectionRuleListParam model.FirewallSectionRuleList, idParam *string, operationParam *string) (model.FirewallSectionRuleList, error)

	// Unlock a section.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionLockParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws ResourceBusy  Locked
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Unlock(sectionIdParam string, firewallSectionLockParam model.FirewallSectionLock) (model.FirewallSection, error)

	// Modifies the specified section, but does not modify the section's associated rules. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionParam (required)
	// @return com.vmware.nsx.model.FirewallSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, firewallSectionParam model.FirewallSection) (model.FirewallSection, error)

	// Modifies existing firewall section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to update rule content, use: PUT /api/v1/firewall/sections/<section-id>/rules/<rule-id> Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update.
	//
	// @param sectionIdParam (required)
	// @param firewallSectionRuleListParam (required)
	// @return com.vmware.nsx.model.FirewallSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(sectionIdParam string, firewallSectionRuleListParam model.FirewallSectionRuleList) (model.FirewallSectionRuleList, error)
}

type sectionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSectionsClient(connector client.Connector) *sectionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.firewall.sections")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":          core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": core.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"listwithrules":   core.NewMethodIdentifier(interfaceIdentifier, "listwithrules"),
		"lock":            core.NewMethodIdentifier(interfaceIdentifier, "lock"),
		"revise":          core.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"revisewithrules": core.NewMethodIdentifier(interfaceIdentifier, "revisewithrules"),
		"unlock":          core.NewMethodIdentifier(interfaceIdentifier, "unlock"),
		"update":          core.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatewithrules": core.NewMethodIdentifier(interfaceIdentifier, "updatewithrules"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := sectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sectionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sectionsClient) Create(firewallSectionParam model.FirewallSection, idParam *string, operationParam *string) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsCreateInputType(), typeConverter)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "create", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Createwithrules(firewallSectionRuleListParam model.FirewallSectionRuleList, idParam *string, operationParam *string) (model.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsCreatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "createwithrules", inputDataValue, executionContext)
	var emptyOutput model.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Delete(sectionIdParam string, cascadeParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsDeleteInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("Cascade", cascadeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "delete", inputDataValue, executionContext)
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

func (sIface *sectionsClient) Get(sectionIdParam string) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "get", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) List(appliedTosParam *string, contextProfilesParam *string, cursorParam *string, deepSearchParam *bool, destinationsParam *string, enforcedOnParam *string, excludeAppliedToTypeParam *string, extendedSourcesParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, lockedParam *bool, pageSizeParam *int64, searchInvalidReferencesParam *bool, searchScopeParam *string, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.FirewallSectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsListInputType(), typeConverter)
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
		var emptyOutput model.FirewallSectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "list", inputDataValue, executionContext)
	var emptyOutput model.FirewallSectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Listwithrules(sectionIdParam string) (model.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsListwithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsListwithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "listwithrules", inputDataValue, executionContext)
	var emptyOutput model.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsListwithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Lock(sectionIdParam string, firewallSectionLockParam model.FirewallSectionLock) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsLockInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionLock", firewallSectionLockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsLockRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "lock", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsLockOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revise(sectionIdParam string, firewallSectionParam model.FirewallSection, idParam *string, operationParam *string) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsReviseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "revise", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revisewithrules(sectionIdParam string, firewallSectionRuleListParam model.FirewallSectionRuleList, idParam *string, operationParam *string) (model.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsRevisewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsRevisewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "revisewithrules", inputDataValue, executionContext)
	var emptyOutput model.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsRevisewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Unlock(sectionIdParam string, firewallSectionLockParam model.FirewallSectionLock) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsUnlockInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionLock", firewallSectionLockParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsUnlockRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "unlock", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsUnlockOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Update(sectionIdParam string, firewallSectionParam model.FirewallSection) (model.FirewallSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSection", firewallSectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "update", inputDataValue, executionContext)
	var emptyOutput model.FirewallSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Updatewithrules(sectionIdParam string, firewallSectionRuleListParam model.FirewallSectionRuleList) (model.FirewallSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("FirewallSectionRuleList", firewallSectionRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FirewallSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsUpdatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput model.FirewallSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FirewallSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

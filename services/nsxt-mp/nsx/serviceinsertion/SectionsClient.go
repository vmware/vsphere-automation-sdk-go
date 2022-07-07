// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Sections
// Used by client-side stubs.

package serviceinsertion

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

	// Creates new empty Service Insertion section in the system. Note- POST service insertion section API is deprecated. Please use the policy redirection-policy API.
	//
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error)

	// Creates a new serviceinsertion section with rules. The limit on the number of rules is defined by maxItems in collection types for ServiceInsertionRule (ServiceInsertionRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to create sections, use: POST /api/v1/serviceinsertion/sections To create rules, use: POST /api/v1/serviceinsertion/sections/<section-id>/rules Note- POST service insertion section creation with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error)

	// Removes serviceinsertion section from the system. ServiceInsertion section with rules can only be deleted by passing \"cascade=true\" parameter. Note- DELETE service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param cascadeParam Flag to cascade delete of this object to all it's child objects. (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(sectionIdParam string, cascadeParam *bool) error

	// Returns information about serviceinsertion section for the identifier. Note- GET service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string) (model.ServiceInsertionSection, error)

	// List all Service Insertion section in paginated form. A default page size is limited to 1000 sections. By default, the list of section is filtered by L3REDIRECT type. Note- GET service insertion sections API is deprecated. Please use the policy redirection-policy API.
	//
	// @param appliedTosParam AppliedTo's referenced by this section or section's Distributed Service Rules . (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param destinationsParam Destinations referenced by this section's Distributed Service Rules . (optional)
	// @param excludeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param filterTypeParam Filter type (optional, default to FILTER)
	// @param includeAppliedToTypeParam Resource type valid for use as AppliedTo filter in section API (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param servicesParam NSService referenced by this section's Distributed Service Rules . (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourcesParam Sources referenced by this section's Distributed Service Rules . (optional)
	// @param type_Param Section Type (optional, default to L3REDIRECT)
	// @return com.vmware.nsx.model.ServiceInsertionSectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.ServiceInsertionSectionListResult, error)

	// Returns serviceinsertion section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to read serviceinsertion rules, use: GET /api/v1/serviceinsertion/sections/<section-id>/rules with the appropriate page_size. Note- GET service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Listwithrules(sectionIdParam string) (model.ServiceInsertionSectionRuleList, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections in the system. Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to move a section above or below another section, use: POST /api/v1/serviceinsertion/sections/<section-id>?action=revise To modify rules, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id> Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revisewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error)

	// Modifies the specified section, but does not modify the section's associated rules. Note- PUT service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection) (model.ServiceInsertionSection, error)

	// Modifies existing serviceinsertion section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to update rule content, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id> Note- POST service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList) (model.ServiceInsertionSectionRuleList, error)
}

type sectionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSectionsClient(connector client.Connector) *sectionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.sections")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":          core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": core.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"listwithrules":   core.NewMethodIdentifier(interfaceIdentifier, "listwithrules"),
		"revise":          core.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"revisewithrules": core.NewMethodIdentifier(interfaceIdentifier, "revisewithrules"),
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

func (sIface *sectionsClient) Create(serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "create", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Createwithrules(serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsCreatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "createwithrules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
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
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "delete", inputDataValue, executionContext)
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

func (sIface *sectionsClient) Get(sectionIdParam string) (model.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "get", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) List(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (model.ServiceInsertionSectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsListInputType(), typeConverter)
	sv.AddStructField("AppliedTos", appliedTosParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Destinations", destinationsParam)
	sv.AddStructField("ExcludeAppliedToType", excludeAppliedToTypeParam)
	sv.AddStructField("FilterType", filterTypeParam)
	sv.AddStructField("IncludeAppliedToType", includeAppliedToTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Services", servicesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Sources", sourcesParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "list", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Listwithrules(sectionIdParam string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsListwithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsListwithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "listwithrules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsListwithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revise(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection, idParam *string, operationParam *string) (model.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsReviseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "revise", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revisewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsRevisewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsRevisewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "revisewithrules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsRevisewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Update(sectionIdParam string, serviceInsertionSectionParam model.ServiceInsertionSection) (model.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "update", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Updatewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam model.ServiceInsertionSectionRuleList) (model.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sectionsUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceInsertionSectionRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sectionsUpdatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput model.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sectionsUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

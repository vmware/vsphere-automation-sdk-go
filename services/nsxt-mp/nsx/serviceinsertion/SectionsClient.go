// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Sections
// Used by client-side stubs.

package serviceinsertion

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SectionsClient interface {

	// Creates new empty Service Insertion section in the system.
	//  Note- POST service insertion section API is deprecated. Please use the policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceInsertionSectionParam nsxModel.ServiceInsertionSection, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSection, error)

	// Creates a new serviceinsertion section with rules. The limit on the number of rules is defined by maxItems in collection types for ServiceInsertionRule (ServiceInsertionRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to create sections, use: POST /api/v1/serviceinsertion/sections To create rules, use: POST /api/v1/serviceinsertion/sections/<section-id>/rules
	//  Note- POST service insertion section creation with rules API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSectionRuleList, error)

	// Removes serviceinsertion section from the system. ServiceInsertion section with rules can only be deleted by passing \"cascade=true\" parameter.
	//  Note- DELETE service insertion section API is deprecated. Please use policy redirection-policy API.
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

	// Returns information about serviceinsertion section for the identifier.
	//  Note- GET service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string) (nsxModel.ServiceInsertionSection, error)

	// List all Service Insertion section in paginated form. A default page size is limited to 1000 sections. By default, the list of section is filtered by L3REDIRECT type.
	//  Note- GET service insertion sections API is deprecated. Please use the policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
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
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (nsxModel.ServiceInsertionSectionListResult, error)

	// Returns serviceinsertion section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported. Instead, to read serviceinsertion rules, use: GET /api/v1/serviceinsertion/sections/<section-id>/rules with the appropriate page_size.
	//  Note- GET service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Listwithrules(sectionIdParam string) (nsxModel.ServiceInsertionSectionRuleList, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections in the system.
	//  Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revise(sectionIdParam string, serviceInsertionSectionParam nsxModel.ServiceInsertionSection, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSection, error)

	// Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to move a section above or below another section, use: POST /api/v1/serviceinsertion/sections/<section-id>?action=revise To modify rules, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id>
	//  Note- POST service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @param idParam Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'. (optional)
	// @param operationParam Operation (optional, default to insert_top)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Revisewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSectionRuleList, error)

	// Modifies the specified section, but does not modify the section's associated rules.
	//  Note- PUT service insertion section API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(sectionIdParam string, serviceInsertionSectionParam nsxModel.ServiceInsertionSection) (nsxModel.ServiceInsertionSection, error)

	// Modifies existing serviceinsertion section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds. Instead, to update rule content, use: PUT /api/v1/serviceinsertion/sections/<section-id>/rules/<rule-id>
	//  The POST service insertion section with rules API is deprecated. Please use policy redirection-policy API.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param serviceInsertionSectionRuleListParam (required)
	// @return com.vmware.nsx.model.ServiceInsertionSectionRuleList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList) (nsxModel.ServiceInsertionSectionRuleList, error)
}

type sectionsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSectionsClient(connector vapiProtocolClient_.Connector) *sectionsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.sections")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"listwithrules":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "listwithrules"),
		"revise":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revise"),
		"revisewithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "revisewithrules"),
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

func (sIface *sectionsClient) Create(serviceInsertionSectionParam nsxModel.ServiceInsertionSection, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Createwithrules(serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsCreatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "createwithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSectionRuleList), nil
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "delete", inputDataValue, executionContext)
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

func (sIface *sectionsClient) Get(sectionIdParam string) (nsxModel.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) List(appliedTosParam *string, cursorParam *string, destinationsParam *string, excludeAppliedToTypeParam *string, filterTypeParam *string, includeAppliedToTypeParam *string, includedFieldsParam *string, pageSizeParam *int64, servicesParam *string, sortAscendingParam *bool, sortByParam *string, sourcesParam *string, type_Param *string) (nsxModel.ServiceInsertionSectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsListInputType(), typeConverter)
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
		var emptyOutput nsxModel.ServiceInsertionSectionListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Listwithrules(sectionIdParam string) (nsxModel.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsListwithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsListwithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "listwithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsListwithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revise(sectionIdParam string, serviceInsertionSectionParam nsxModel.ServiceInsertionSection, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsReviseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsReviseInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "revise", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsReviseOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Revisewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList, idParam *string, operationParam *string) (nsxModel.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsRevisewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsRevisewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "revisewithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsRevisewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Update(sectionIdParam string, serviceInsertionSectionParam nsxModel.ServiceInsertionSection) (nsxModel.ServiceInsertionSection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsUpdateInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSection", serviceInsertionSectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sectionsClient) Updatewithrules(sectionIdParam string, serviceInsertionSectionRuleListParam nsxModel.ServiceInsertionSectionRuleList) (nsxModel.ServiceInsertionSectionRuleList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sectionsUpdatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sectionsUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("ServiceInsertionSectionRuleList", serviceInsertionSectionRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceInsertionSectionRuleList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.sections", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceInsertionSectionRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SectionsUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceInsertionSectionRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

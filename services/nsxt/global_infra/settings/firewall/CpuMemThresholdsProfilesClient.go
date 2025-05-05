// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CpuMemThresholdsProfiles
// Used by client-side stubs.

package firewall

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CpuMemThresholdsProfilesClient interface {

	// Delete CPU and memory thresholds profile.
	//
	// @param profileIdParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(profileIdParam string, overrideParam *bool) error

	// Read the CPU and memory thresholds profile.
	//
	// @param profileIdParam (required)
	// @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(profileIdParam string) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, error)

	// List all CPU and memory thresholds profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfileListResult, error)

	// Create or update CPU and memory thresholds profile.
	//
	// @param profileIdParam (required)
	// @param policyFirewallCpuMemThresholdsProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) error

	// Create or update CPU and memory thresholds profile.
	//
	// @param profileIdParam (required)
	// @param policyFirewallCpuMemThresholdsProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	// @return com.vmware.nsx_policy.model.PolicyFirewallCpuMemThresholdsProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, error)
}

type cpuMemThresholdsProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCpuMemThresholdsProfilesClient(connector vapiProtocolClient_.Connector) *cpuMemThresholdsProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := cpuMemThresholdsProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cpuMemThresholdsProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cpuMemThresholdsProfilesClient) Delete(profileIdParam string, overrideParam *bool) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles", "delete", inputDataValue, executionContext)
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

func (cIface *cpuMemThresholdsProfilesClient) Get(profileIdParam string) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfilesGetInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CpuMemThresholdsProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cpuMemThresholdsProfilesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfileListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CpuMemThresholdsProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyFirewallCpuMemThresholdsProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cpuMemThresholdsProfilesClient) Patch(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfilesPatchInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("PolicyFirewallCpuMemThresholdsProfile", policyFirewallCpuMemThresholdsProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles", "patch", inputDataValue, executionContext)
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

func (cIface *cpuMemThresholdsProfilesClient) Update(profileIdParam string, policyFirewallCpuMemThresholdsProfileParam nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, overrideParam *bool) (nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("PolicyFirewallCpuMemThresholdsProfile", policyFirewallCpuMemThresholdsProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.cpu_mem_thresholds_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CpuMemThresholdsProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyFirewallCpuMemThresholdsProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

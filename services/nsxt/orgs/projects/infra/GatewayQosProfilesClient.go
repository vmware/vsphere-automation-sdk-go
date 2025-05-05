// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GatewayQosProfiles
// Used by client-side stubs.

package infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type GatewayQosProfilesClient interface {

	// Delete QoS profile
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param qosProfileIdParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, qosProfileIdParam string, overrideParam *bool) error

	// Read gateway QoS profile
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param qosProfileIdParam (required)
	// @return com.vmware.nsx_policy.model.GatewayQosProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, qosProfileIdParam string) (nsx_policyModel.GatewayQosProfile, error)

	// Paginated list of all gateway QoS profle instances
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.GatewayQosProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.GatewayQosProfileListResult, error)

	// If profile with the qos-profile-id is not already present, create a new gateway QoS profile instance. If it already exists, update the gateway QoS profile instance with specified attributes.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param qosProfileIdParam (required)
	// @param gatewayQosProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, qosProfileIdParam string, gatewayQosProfileParam nsx_policyModel.GatewayQosProfile, overrideParam *bool) error

	// If profile with the qos-profile-id is not already present, create a new gateway QoS profile instance. If it already exists, replace the gateway QoS profile instance with this object.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param qosProfileIdParam (required)
	// @param gatewayQosProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	// @return com.vmware.nsx_policy.model.GatewayQosProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, qosProfileIdParam string, gatewayQosProfileParam nsx_policyModel.GatewayQosProfile, overrideParam *bool) (nsx_policyModel.GatewayQosProfile, error)
}

type gatewayQosProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewGatewayQosProfilesClient(connector vapiProtocolClient_.Connector) *gatewayQosProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	gIface := gatewayQosProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *gatewayQosProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *gatewayQosProfilesClient) Delete(orgIdParam string, projectIdParam string, qosProfileIdParam string, overrideParam *bool) error {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayQosProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayQosProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("QosProfileId", qosProfileIdParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *gatewayQosProfilesClient) Get(orgIdParam string, projectIdParam string, qosProfileIdParam string) (nsx_policyModel.GatewayQosProfile, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayQosProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayQosProfilesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("QosProfileId", qosProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayQosProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayQosProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayQosProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayQosProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *gatewayQosProfilesClient) List(orgIdParam string, projectIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.GatewayQosProfileListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayQosProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayQosProfilesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayQosProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayQosProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayQosProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayQosProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *gatewayQosProfilesClient) Patch(orgIdParam string, projectIdParam string, qosProfileIdParam string, gatewayQosProfileParam nsx_policyModel.GatewayQosProfile, overrideParam *bool) error {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayQosProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayQosProfilesPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("QosProfileId", qosProfileIdParam)
	sv.AddStructField("GatewayQosProfile", gatewayQosProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *gatewayQosProfilesClient) Update(orgIdParam string, projectIdParam string, qosProfileIdParam string, gatewayQosProfileParam nsx_policyModel.GatewayQosProfile, overrideParam *bool) (nsx_policyModel.GatewayQosProfile, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := gatewayQosProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(gatewayQosProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("QosProfileId", qosProfileIdParam)
	sv.AddStructField("GatewayQosProfile", gatewayQosProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GatewayQosProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.gateway_qos_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GatewayQosProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GatewayQosProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GatewayQosProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

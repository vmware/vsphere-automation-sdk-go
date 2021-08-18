// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TlsInspectionConfigProfiles
// Used by client-side stubs.

package security

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type TlsInspectionConfigProfilesClient interface {

	// API will delete TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(tlsInspectionConfigProfileParam string, overrideParam *bool) error

	// API will get TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @return com.vmware.nsx_policy.model.PolicyTlsConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tlsInspectionConfigProfileParam string) (model.PolicyTlsConfigProfile, error)

	// API will list all TLS Config Profiles
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyTlsConfigProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyTlsConfigProfileListResult, error)

	// API will create/update TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param policyTlsConfigProfileParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam model.PolicyTlsConfigProfile, overrideParam *bool) error

	// API will update TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param policyTlsConfigProfileParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	// @return com.vmware.nsx_policy.model.PolicyTlsConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam model.PolicyTlsConfigProfile, overrideParam *bool) (model.PolicyTlsConfigProfile, error)
}

type tlsInspectionConfigProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTlsInspectionConfigProfilesClient(connector client.Connector) *tlsInspectionConfigProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := tlsInspectionConfigProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tlsInspectionConfigProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tlsInspectionConfigProfilesClient) Delete(tlsInspectionConfigProfileParam string, overrideParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tlsInspectionConfigProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tlsInspectionConfigProfilesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Get(tlsInspectionConfigProfileParam string) (model.PolicyTlsConfigProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tlsInspectionConfigProfilesGetInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyTlsConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tlsInspectionConfigProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyTlsConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tlsInspectionConfigProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyTlsConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyTlsConfigProfileListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tlsInspectionConfigProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyTlsConfigProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tlsInspectionConfigProfilesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles", "list", inputDataValue, executionContext)
	var emptyOutput model.PolicyTlsConfigProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tlsInspectionConfigProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyTlsConfigProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Patch(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam model.PolicyTlsConfigProfile, overrideParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tlsInspectionConfigProfilesPatchInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("PolicyTlsConfigProfile", policyTlsConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tlsInspectionConfigProfilesPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Update(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam model.PolicyTlsConfigProfile, overrideParam *bool) (model.PolicyTlsConfigProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tlsInspectionConfigProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("PolicyTlsConfigProfile", policyTlsConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyTlsConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tlsInspectionConfigProfilesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.security.tls_inspection_config_profiles", "update", inputDataValue, executionContext)
	var emptyOutput model.PolicyTlsConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tlsInspectionConfigProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyTlsConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

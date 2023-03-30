// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TlsInspectionConfigProfiles
// Used by client-side stubs.

package security

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TlsInspectionConfigProfilesClient interface {

	// API will delete TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	//
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
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tlsInspectionConfigProfileParam string) (nsx_policyModel.PolicyTlsConfigProfile, error)

	// API will list all TLS Config Profiles
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyTlsConfigProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyTlsConfigProfileListResult, error)

	// API will create/update TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param policyTlsConfigProfileParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam nsx_policyModel.PolicyTlsConfigProfile, overrideParam *bool) error

	// API will update TLS Config Profile
	//
	// @param tlsInspectionConfigProfileParam TLS Config Profile ID (required)
	// @param policyTlsConfigProfileParam (required)
	// @param overrideParam Locally override the global object (optional, default to false)
	// @return com.vmware.nsx_policy.model.PolicyTlsConfigProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam nsx_policyModel.PolicyTlsConfigProfile, overrideParam *bool) (nsx_policyModel.PolicyTlsConfigProfile, error)
}

type tlsInspectionConfigProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTlsInspectionConfigProfilesClient(connector vapiProtocolClient_.Connector) *tlsInspectionConfigProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := tlsInspectionConfigProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tlsInspectionConfigProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tlsInspectionConfigProfilesClient) Delete(tlsInspectionConfigProfileParam string, overrideParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Get(tlsInspectionConfigProfileParam string) (nsx_policyModel.PolicyTlsConfigProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfilesGetInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyTlsConfigProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyTlsConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TlsInspectionConfigProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyTlsConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyTlsConfigProfileListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyTlsConfigProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyTlsConfigProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TlsInspectionConfigProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyTlsConfigProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Patch(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam nsx_policyModel.PolicyTlsConfigProfile, overrideParam *bool) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfilesPatchInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("PolicyTlsConfigProfile", policyTlsConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tlsInspectionConfigProfilesClient) Update(tlsInspectionConfigProfileParam string, policyTlsConfigProfileParam nsx_policyModel.PolicyTlsConfigProfile, overrideParam *bool) (nsx_policyModel.PolicyTlsConfigProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("TlsInspectionConfigProfile", tlsInspectionConfigProfileParam)
	sv.AddStructField("PolicyTlsConfigProfile", policyTlsConfigProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyTlsConfigProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.tls_inspection_config_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyTlsConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TlsInspectionConfigProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyTlsConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

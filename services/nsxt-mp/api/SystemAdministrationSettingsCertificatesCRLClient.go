// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationSettingsCertificatesCRL
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationSettingsCertificatesCRLClient interface {

	// Create an entity that will represent a Crl Distribution Point
	//
	// @param crlDistributionPointParam (required)
	// @return com.vmware.model.CrlDistributionPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateCrlDistributionPoint(crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error)

	// Delete a CrlDistributionPoint. It does not delete the actual CRL.
	//
	// @param crlDistributionPointIdParam Unique id of the CrlDistributionPoint to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteCrlDistributionPoint(crlDistributionPointIdParam string) error

	//
	//
	// @param crlDistributionPointIdParam (required)
	// @return com.vmware.model.CrlDistributionPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetCrlDistributionPoint(crlDistributionPointIdParam string) (model.CrlDistributionPoint, error)

	// Return stored CRL in PEM format
	//
	// @param crlPemRequestTypeParam (required)
	// @return String
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetCrlDistributionPointPem(crlPemRequestTypeParam model.CrlPemRequestType) (string, error)

	// Return the status of the CrlDistributionPoint
	//
	// @param crlDistributionPointIdParam (required)
	// @return com.vmware.model.CrlDistributionPointStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetCrlDistributionPointStatus(crlDistributionPointIdParam string) (model.CrlDistributionPointStatus, error)

	// Return the list of CrlDistributionPoints
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.CrlDistributionPointList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListCrlDistributionPoints(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.CrlDistributionPointList, error)

	//
	//
	// @param crlDistributionPointIdParam (required)
	// @param crlDistributionPointParam (required)
	// @return com.vmware.model.CrlDistributionPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateCrlDistributionPoint(crlDistributionPointIdParam string, crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error)
}

type systemAdministrationSettingsCertificatesCRLClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationSettingsCertificatesCRLClient(connector client.Connector) *systemAdministrationSettingsCertificatesCRLClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_settings_certificates_CRL")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_crl_distribution_point":     core.NewMethodIdentifier(interfaceIdentifier, "create_crl_distribution_point"),
		"delete_crl_distribution_point":     core.NewMethodIdentifier(interfaceIdentifier, "delete_crl_distribution_point"),
		"get_crl_distribution_point":        core.NewMethodIdentifier(interfaceIdentifier, "get_crl_distribution_point"),
		"get_crl_distribution_point_pem":    core.NewMethodIdentifier(interfaceIdentifier, "get_crl_distribution_point_pem"),
		"get_crl_distribution_point_status": core.NewMethodIdentifier(interfaceIdentifier, "get_crl_distribution_point_status"),
		"list_crl_distribution_points":      core.NewMethodIdentifier(interfaceIdentifier, "list_crl_distribution_points"),
		"update_crl_distribution_point":     core.NewMethodIdentifier(interfaceIdentifier, "update_crl_distribution_point"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationSettingsCertificatesCRLClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) CreateCrlDistributionPoint(crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointInputType(), typeConverter)
	sv.AddStructField("CrlDistributionPoint", crlDistributionPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlDistributionPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "create_crl_distribution_point", inputDataValue, executionContext)
	var emptyOutput model.CrlDistributionPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlDistributionPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) DeleteCrlDistributionPoint(crlDistributionPointIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLDeleteCrlDistributionPointInputType(), typeConverter)
	sv.AddStructField("CrlDistributionPointId", crlDistributionPointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLDeleteCrlDistributionPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "delete_crl_distribution_point", inputDataValue, executionContext)
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

func (sIface *systemAdministrationSettingsCertificatesCRLClient) GetCrlDistributionPoint(crlDistributionPointIdParam string) (model.CrlDistributionPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointInputType(), typeConverter)
	sv.AddStructField("CrlDistributionPointId", crlDistributionPointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlDistributionPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "get_crl_distribution_point", inputDataValue, executionContext)
	var emptyOutput model.CrlDistributionPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlDistributionPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) GetCrlDistributionPointPem(crlPemRequestTypeParam model.CrlPemRequestType) (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemInputType(), typeConverter)
	sv.AddStructField("CrlPemRequestType", crlPemRequestTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "get_crl_distribution_point_pem", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) GetCrlDistributionPointStatus(crlDistributionPointIdParam string) (model.CrlDistributionPointStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusInputType(), typeConverter)
	sv.AddStructField("CrlDistributionPointId", crlDistributionPointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlDistributionPointStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "get_crl_distribution_point_status", inputDataValue, executionContext)
	var emptyOutput model.CrlDistributionPointStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlDistributionPointStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) ListCrlDistributionPoints(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.CrlDistributionPointList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlDistributionPointList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "list_crl_distribution_points", inputDataValue, executionContext)
	var emptyOutput model.CrlDistributionPointList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlDistributionPointList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCertificatesCRLClient) UpdateCrlDistributionPoint(crlDistributionPointIdParam string, crlDistributionPointParam model.CrlDistributionPoint) (model.CrlDistributionPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointInputType(), typeConverter)
	sv.AddStructField("CrlDistributionPointId", crlDistributionPointIdParam)
	sv.AddStructField("CrlDistributionPoint", crlDistributionPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlDistributionPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_certificates_CRL", "update_crl_distribution_point", inputDataValue, executionContext)
	var emptyOutput model.CrlDistributionPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlDistributionPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

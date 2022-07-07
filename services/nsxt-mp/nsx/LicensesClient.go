// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Licenses
// Used by client-side stubs.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type LicensesClient interface {

	// This will add a license key to the system. The API supports adding only one license key for each license edition type - Standard, Advanced or Enterprise. If a new license key is tried to add for an edition for which the license key already exists, then this API will return an error.
	//
	// @param licenseParam (required)
	// @return com.vmware.nsx.model.License
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(licenseParam model.License) (model.License, error)

	// Deprecated. Use POST /licenses?action=delete API instead.
	//
	// @param licenseKeyParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(licenseKeyParam string) error

	// This will delete the license key identified in the request body by \"license_key\" and its properties from the system. Attempting to delete the last license key will result in an error.
	//
	// @param licenseParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete0(licenseParam model.License) error

	// Deprecated. Use the GET /licenses API instead.
	// @return com.vmware.nsx.model.License
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.License, error)

	// Deprecated. Use GET /licenses API instead.
	//
	// @param licenseKeyParam (required)
	// @return com.vmware.nsx.model.License
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Getlicensebykey(licenseKeyParam string) (model.License, error)

	// Returns all licenses.
	// @return com.vmware.nsx.model.LicensesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.LicensesListResult, error)

	// Deprecated. Use the POST /licenses API instead
	//
	// @param licenseParam (required)
	// @return com.vmware.nsx.model.License
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(licenseParam model.License) (model.License, error)
}

type licensesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLicensesClient(connector client.Connector) *licensesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.licenses")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":          core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"delete_0":        core.NewMethodIdentifier(interfaceIdentifier, "delete_0"),
		"get":             core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"getlicensebykey": core.NewMethodIdentifier(interfaceIdentifier, "getlicensebykey"),
		"list":            core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":          core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := licensesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *licensesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *licensesClient) Create(licenseParam model.License) (model.License, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesCreateInputType(), typeConverter)
	sv.AddStructField("License", licenseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.License
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "create", inputDataValue, executionContext)
	var emptyOutput model.License
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), licensesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.License), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *licensesClient) Delete(licenseKeyParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesDeleteInputType(), typeConverter)
	sv.AddStructField("LicenseKey", licenseKeyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *licensesClient) Delete0(licenseParam model.License) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesDelete0InputType(), typeConverter)
	sv.AddStructField("License", licenseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesDelete0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "delete_0", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *licensesClient) Get() (model.License, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.License
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "get", inputDataValue, executionContext)
	var emptyOutput model.License
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), licensesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.License), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *licensesClient) Getlicensebykey(licenseKeyParam string) (model.License, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesGetlicensebykeyInputType(), typeConverter)
	sv.AddStructField("LicenseKey", licenseKeyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.License
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesGetlicensebykeyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "getlicensebykey", inputDataValue, executionContext)
	var emptyOutput model.License
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), licensesGetlicensebykeyOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.License), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *licensesClient) List() (model.LicensesListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LicensesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "list", inputDataValue, executionContext)
	var emptyOutput model.LicensesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), licensesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LicensesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *licensesClient) Update(licenseParam model.License) (model.License, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(licensesUpdateInputType(), typeConverter)
	sv.AddStructField("License", licenseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.License
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := licensesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.licenses", "update", inputDataValue, executionContext)
	var emptyOutput model.License
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), licensesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.License), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

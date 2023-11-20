// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Sddcs
// Used by client-side stubs.

package orgs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SddcsClient interface {

	// Provision an SDDC in target cloud
	//
	// @param orgParam Organization identifier (required)
	// @param sddcConfigParam sddcConfig (required)
	// @param validateOnlyParam When true, only validates the given sddc configuration without provisioning. (optional)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Create(orgParam string, sddcConfigParam vmcModel.AwsSddcConfig, validateOnlyParam *bool) (vmcModel.Task, error)

	// Delete SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param retainConfigurationParam If = 'true', the SDDC's configuration is retained as a template for later use. This flag is applicable only to SDDCs in ACTIVE state. (optional)
	// @param templateNameParam Only applicable when retainConfiguration is also set to 'true'. When set, this value will be used as the name of the SDDC configuration template generated. (optional)
	// @param forceParam If = true, will delete forcefully. Beware: do not use the force flag if there is a chance an active provisioning or deleting task is running against this SDDC. This option is restricted. (optional)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  The sddc is not in a state that's valid for deletion
	// @throws Unauthorized  Access not allowed to the operation for the current user
	// @throws NotFound  Cannot find the SDDC with given identifier
	Delete(orgParam string, sddcParam string, retainConfigurationParam *bool, templateNameParam *string, forceParam *bool) (vmcModel.Task, error)

	// Get SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @return com.vmware.vmc.model.Sddc
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the SDDC with given identifier
	Get(orgParam string, sddcParam string) (vmcModel.Sddc, error)

	// List all the SDDCs of an organization
	//
	// @param orgParam Organization identifier (required)
	// @param includeDeletedParam When true, forces the result to also include deleted SDDCs. (optional)
	// @param filterParam Filter expression Binary Operators: 'eq', 'ne', 'lt', 'gt', 'le', 'ge', 'mul', 'div', 'mod', 'sub', 'add' Unary Operators: 'not', '-' (minus) String Operators: 'startswith', 'endswith', 'length', 'contains', 'tolower', 'toupper', Nested attributes are composed using '.' Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z Strings should enclosed in single quotes, escape single quote with two single quotes The special literal 'created' will be mapped to the time the resource was first created. Examples: - $filter=(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e) - $filter=(created eq 2016-08-09) - $filter=(created gt 2016-08-09) and (sddc.status eq 'READY') (optional)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	List(orgParam string, includeDeletedParam *bool, filterParam *string) ([]vmcModel.Sddc, error)

	// Patch SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param sddcPatchRequestParam Patch request for the SDDC (required)
	// @return com.vmware.vmc.model.Sddc
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  SDDC cannot be patched
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the SDDC with given identifier
	Patch(orgParam string, sddcParam string, sddcPatchRequestParam vmcModel.SddcPatchRequest) (vmcModel.Sddc, error)
}

type sddcsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSddcsClient(connector vapiProtocolClient_.Connector) *sddcsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := sddcsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sddcsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sddcsClient) Create(orgParam string, sddcConfigParam vmcModel.AwsSddcConfig, validateOnlyParam *bool) (vmcModel.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sddcsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sddcsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcConfig", sddcConfigParam)
	sv.AddStructField("ValidateOnly", validateOnlyParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs", "create", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SddcsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sddcsClient) Delete(orgParam string, sddcParam string, retainConfigurationParam *bool, templateNameParam *string, forceParam *bool) (vmcModel.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sddcsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sddcsDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("RetainConfiguration", retainConfigurationParam)
	sv.AddStructField("TemplateName", templateNameParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs", "delete", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SddcsDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sddcsClient) Get(orgParam string, sddcParam string) (vmcModel.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sddcsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sddcsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Sddc
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.Sddc
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SddcsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sddcsClient) List(orgParam string, includeDeletedParam *bool, filterParam *string) ([]vmcModel.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sddcsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sddcsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("IncludeDeleted", includeDeletedParam)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.Sddc
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs", "list", inputDataValue, executionContext)
	var emptyOutput []vmcModel.Sddc
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SddcsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sddcsClient) Patch(orgParam string, sddcParam string, sddcPatchRequestParam vmcModel.SddcPatchRequest) (vmcModel.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sddcsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sddcsPatchInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("SddcPatchRequest", sddcPatchRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Sddc
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs", "patch", inputDataValue, executionContext)
	var emptyOutput vmcModel.Sddc
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SddcsPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

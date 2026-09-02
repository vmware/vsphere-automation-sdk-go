// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IpBlockConstraints
// Used by client-side stubs.

package ip_blocks

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type IpBlockConstraintsClient interface {

	// Delete the singleton IpAddressBlockConstraint attached to the given IpAddressBlock. After deletion, the block reverts to system defaults: READ_ONLY sharing permission and unrestricted use cases.
	//
	// @param ipBlockIdParam IP Block ID (required)
	// @param constraintIdParam IP Address Block Constraint ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ipBlockIdParam string, constraintIdParam string) error

	// Read the singleton IpAddressBlockConstraint attached to the given IpAddressBlock. Returns 404 if no constraint has been configured.
	//
	// @param ipBlockIdParam IP Block ID (required)
	// @param constraintIdParam IP Address Block Constraint ID (required)
	// @return com.vmware.nsx_policy.model.IpAddressBlockConstraint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipBlockIdParam string, constraintIdParam string) (nsx_policyModel.IpAddressBlockConstraint, error)

	// Create or update the singleton IpAddressBlockConstraint attached to the given IpAddressBlock. Only the fields provided in the request body are modified; omitted fields are left unchanged.
	//
	// @param ipBlockIdParam IP Block ID (required)
	// @param constraintIdParam IP Address Block Constraint ID (required)
	// @param ipAddressBlockConstraintParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(ipBlockIdParam string, constraintIdParam string, ipAddressBlockConstraintParam nsx_policyModel.IpAddressBlockConstraint) error

	// Create or fully replace the singleton IpAddressBlockConstraint attached to the given IpAddressBlock. This is a full replace; fields not present in the request body are reset to their defaults.
	//
	// @param ipBlockIdParam IP Block ID (required)
	// @param constraintIdParam IP Address Block Constraint ID (required)
	// @param ipAddressBlockConstraintParam (required)
	// @return com.vmware.nsx_policy.model.IpAddressBlockConstraint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ipBlockIdParam string, constraintIdParam string, ipAddressBlockConstraintParam nsx_policyModel.IpAddressBlockConstraint) (nsx_policyModel.IpAddressBlockConstraint, error)
}

type ipBlockConstraintsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIpBlockConstraintsClient(connector vapiProtocolClient_.Connector) *ipBlockConstraintsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.ip_blocks.ip_block_constraints")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := ipBlockConstraintsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *ipBlockConstraintsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *ipBlockConstraintsClient) Delete(ipBlockIdParam string, constraintIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipBlockConstraintsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipBlockConstraintsDeleteInputType(), typeConverter)
	sv.AddStructField("IpBlockId", ipBlockIdParam)
	sv.AddStructField("ConstraintId", constraintIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.ip_blocks.ip_block_constraints", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *ipBlockConstraintsClient) Get(ipBlockIdParam string, constraintIdParam string) (nsx_policyModel.IpAddressBlockConstraint, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipBlockConstraintsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipBlockConstraintsGetInputType(), typeConverter)
	sv.AddStructField("IpBlockId", ipBlockIdParam)
	sv.AddStructField("ConstraintId", constraintIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.IpAddressBlockConstraint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.ip_blocks.ip_block_constraints", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.IpAddressBlockConstraint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpBlockConstraintsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.IpAddressBlockConstraint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipBlockConstraintsClient) Patch(ipBlockIdParam string, constraintIdParam string, ipAddressBlockConstraintParam nsx_policyModel.IpAddressBlockConstraint) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipBlockConstraintsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipBlockConstraintsPatchInputType(), typeConverter)
	sv.AddStructField("IpBlockId", ipBlockIdParam)
	sv.AddStructField("ConstraintId", constraintIdParam)
	sv.AddStructField("IpAddressBlockConstraint", ipAddressBlockConstraintParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.ip_blocks.ip_block_constraints", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *ipBlockConstraintsClient) Update(ipBlockIdParam string, constraintIdParam string, ipAddressBlockConstraintParam nsx_policyModel.IpAddressBlockConstraint) (nsx_policyModel.IpAddressBlockConstraint, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipBlockConstraintsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipBlockConstraintsUpdateInputType(), typeConverter)
	sv.AddStructField("IpBlockId", ipBlockIdParam)
	sv.AddStructField("ConstraintId", constraintIdParam)
	sv.AddStructField("IpAddressBlockConstraint", ipAddressBlockConstraintParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.IpAddressBlockConstraint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.ip_blocks.ip_block_constraints", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.IpAddressBlockConstraint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpBlockConstraintsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.IpAddressBlockConstraint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ApiVirtualIp
// Used by client-side stubs.

package cluster

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ApiVirtualIpClient interface {

	// Clears the cluster virtual IPv4 or IPv6 address.
	// Note, query parameter **?action=clear_virtual_ip** clears virtual IPv4 address and **?action=clear_virtual_ip6** clears virtual IPv6 address.
	// @return com.vmware.nsx.model.ClusterVirtualIpProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Clearvirtualip() (nsxModel.ClusterVirtualIpProperties, error)

	// Clears the cluster virtual IPv4 or IPv6 address.
	// Note, query parameter **?action=clear_virtual_ip** clears virtual IPv4 address and **?action=clear_virtual_ip6** clears virtual IPv6 address.
	// @return com.vmware.nsx.model.ClusterVirtualIpProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Clearvirtualip6() (nsxModel.ClusterVirtualIpProperties, error)

	// Returns the configured cluster virtual IPv4 and IPv6 address or null if not configured.
	// @return com.vmware.nsx.model.ClusterVirtualIpProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsxModel.ClusterVirtualIpProperties, error)

	// Sets the cluster virtual IPv4 and IPv6 address. Note, all nodes in the management cluster must be in the same subnet. If not, a 409 CONFLICT status is returned.
	// Query parameter **ip_address** sets virtual IPv4 address and **ip6_address** sets virtual IPv6 address; either or both of the parameters can be specified to set virtual IP address(es).
	// When updating either of any one parameter value this does not changes the value of other unspecified parameter.
	//
	// @param forceParam On enable it ignores duplicate address detection and DNS lookup validation check (optional, default to false)
	// @param ip6AddressParam Virtual IPv6 address, :: if not configured (optional)
	// @param ipAddressParam Virtual IP address, 0.0.0.0 if not configured (optional)
	// @return com.vmware.nsx.model.ClusterVirtualIpProperties
	//
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setvirtualip(forceParam *string, ip6AddressParam *string, ipAddressParam *string) (nsxModel.ClusterVirtualIpProperties, error)
}

type apiVirtualIpClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewApiVirtualIpClient(connector vapiProtocolClient_.Connector) *apiVirtualIpClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.cluster.api_virtual_ip")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"clearvirtualip":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "clearvirtualip"),
		"clearvirtualip6": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "clearvirtualip6"),
		"get":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"setvirtualip":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "setvirtualip"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := apiVirtualIpClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *apiVirtualIpClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *apiVirtualIpClient) Clearvirtualip() (nsxModel.ClusterVirtualIpProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := apiVirtualIpClearvirtualipRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(apiVirtualIpClearvirtualipInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ClusterVirtualIpProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_virtual_ip", "clearvirtualip", inputDataValue, executionContext)
	var emptyOutput nsxModel.ClusterVirtualIpProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ApiVirtualIpClearvirtualipOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ClusterVirtualIpProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *apiVirtualIpClient) Clearvirtualip6() (nsxModel.ClusterVirtualIpProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := apiVirtualIpClearvirtualip6RestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(apiVirtualIpClearvirtualip6InputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ClusterVirtualIpProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_virtual_ip", "clearvirtualip6", inputDataValue, executionContext)
	var emptyOutput nsxModel.ClusterVirtualIpProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ApiVirtualIpClearvirtualip6OutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ClusterVirtualIpProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *apiVirtualIpClient) Get() (nsxModel.ClusterVirtualIpProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := apiVirtualIpGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(apiVirtualIpGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ClusterVirtualIpProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_virtual_ip", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ClusterVirtualIpProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ApiVirtualIpGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ClusterVirtualIpProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *apiVirtualIpClient) Setvirtualip(forceParam *string, ip6AddressParam *string, ipAddressParam *string) (nsxModel.ClusterVirtualIpProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := apiVirtualIpSetvirtualipRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(apiVirtualIpSetvirtualipInputType(), typeConverter)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("Ip6Address", ip6AddressParam)
	sv.AddStructField("IpAddress", ipAddressParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ClusterVirtualIpProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_virtual_ip", "setvirtualip", inputDataValue, executionContext)
	var emptyOutput nsxModel.ClusterVirtualIpProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ApiVirtualIpSetvirtualipOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ClusterVirtualIpProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

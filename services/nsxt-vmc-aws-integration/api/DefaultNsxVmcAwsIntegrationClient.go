
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: NsxVmcAwsIntegration
 * Functions that implement the generated NsxVmcAwsIntegrationClient interface
 */


package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultNsxVmcAwsIntegrationClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultNsxVmcAwsIntegrationClient(connector client.Connector) *DefaultNsxVmcAwsIntegrationClient {
	interfaceName := "com.vmware.api.nsx_vmc_aws_integration"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "attachVif"),
		core.NewMethodIdentifier(interfaceIdentifier, "createPublicIp"),
		core.NewMethodIdentifier(interfaceIdentifier, "deletePublicIp"),
		core.NewMethodIdentifier(interfaceIdentifier, "deleteVifById"),
		core.NewMethodIdentifier(interfaceIdentifier, "getDxBgpInfo"),
		core.NewMethodIdentifier(interfaceIdentifier, "getLinkedVpc"),
		core.NewMethodIdentifier(interfaceIdentifier, "getMgmtVmInfo"),
		core.NewMethodIdentifier(interfaceIdentifier, "getPublicIp"),
		core.NewMethodIdentifier(interfaceIdentifier, "getRealizedStateStatus"),
		core.NewMethodIdentifier(interfaceIdentifier, "getSddcUserConfig"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAccounts"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAdvertisedRoutes"),
		core.NewMethodIdentifier(interfaceIdentifier, "listConnectedServices"),
		core.NewMethodIdentifier(interfaceIdentifier, "listLearnedRoutes"),
		core.NewMethodIdentifier(interfaceIdentifier, "listLinkedVpcs"),
		core.NewMethodIdentifier(interfaceIdentifier, "listMgmtVms"),
		core.NewMethodIdentifier(interfaceIdentifier, "listPublicIps"),
		core.NewMethodIdentifier(interfaceIdentifier, "listVifs"),
		core.NewMethodIdentifier(interfaceIdentifier, "updateConnectedService"),
		core.NewMethodIdentifier(interfaceIdentifier, "updateDxBgpInfo"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errorBindingMap[errors.Canceled{}.Error()] = errors.CanceledBindingType()
	errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errorBindingMap[errors.InvalidElementType{}.Error()] = errors.InvalidElementTypeBindingType()
	errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errorBindingMap[errors.ResourceInUse{}.Error()] = errors.ResourceInUseBindingType()
	errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()


	nIface := DefaultNsxVmcAwsIntegrationClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	nIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	nIface.methodNameToDefMap["attach_vif"] = nIface.attachVifMethodDefinition()
	nIface.methodNameToDefMap["create_public_ip"] = nIface.createPublicIpMethodDefinition()
	nIface.methodNameToDefMap["delete_public_ip"] = nIface.deletePublicIpMethodDefinition()
	nIface.methodNameToDefMap["delete_vif_by_id"] = nIface.deleteVifByIdMethodDefinition()
	nIface.methodNameToDefMap["get_dx_bgp_info"] = nIface.getDxBgpInfoMethodDefinition()
	nIface.methodNameToDefMap["get_linked_vpc"] = nIface.getLinkedVpcMethodDefinition()
	nIface.methodNameToDefMap["get_mgmt_vm_info"] = nIface.getMgmtVmInfoMethodDefinition()
	nIface.methodNameToDefMap["get_public_ip"] = nIface.getPublicIpMethodDefinition()
	nIface.methodNameToDefMap["get_realized_state_status"] = nIface.getRealizedStateStatusMethodDefinition()
	nIface.methodNameToDefMap["get_sddc_user_config"] = nIface.getSddcUserConfigMethodDefinition()
	nIface.methodNameToDefMap["list_accounts"] = nIface.listAccountsMethodDefinition()
	nIface.methodNameToDefMap["list_advertised_routes"] = nIface.listAdvertisedRoutesMethodDefinition()
	nIface.methodNameToDefMap["list_connected_services"] = nIface.listConnectedServicesMethodDefinition()
	nIface.methodNameToDefMap["list_learned_routes"] = nIface.listLearnedRoutesMethodDefinition()
	nIface.methodNameToDefMap["list_linked_vpcs"] = nIface.listLinkedVpcsMethodDefinition()
	nIface.methodNameToDefMap["list_mgmt_vms"] = nIface.listMgmtVmsMethodDefinition()
	nIface.methodNameToDefMap["list_public_ips"] = nIface.listPublicIpsMethodDefinition()
	nIface.methodNameToDefMap["list_vifs"] = nIface.listVifsMethodDefinition()
	nIface.methodNameToDefMap["update_connected_service"] = nIface.updateConnectedServiceMethodDefinition()
	nIface.methodNameToDefMap["update_dx_bgp_info"] = nIface.updateDxBgpInfoMethodDefinition()
	return &nIface
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) AttachVif(vifIdParam string, actionParam string) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "attach_vif")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationAttachVifInputType(), typeConverter)
	sv.AddStructField("VifId", vifIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationAttachVifRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) CreatePublicIp(publicIpIdParam string, publicIpParam model.PublicIp) (model.PublicIp, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "create_public_ip")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationCreatePublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	sv.AddStructField("PublicIp", publicIpParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationCreatePublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.PublicIp
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationCreatePublicIpOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) DeletePublicIp(publicIpIdParam string, forceParam *bool) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "delete_public_ip")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationDeletePublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationDeletePublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) DeleteVifById(vifIdParam string) error {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "delete_vif_by_id")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationDeleteVifByIdInputType(), typeConverter)
	sv.AddStructField("VifId", vifIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationDeleteVifByIdRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetDxBgpInfo() (model.DirectConnectBgpInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_dx_bgp_info")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetDxBgpInfoInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectConnectBgpInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetDxBgpInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.DirectConnectBgpInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetDxBgpInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectConnectBgpInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetLinkedVpc(linkedVpcIdParam string) (model.LinkedVpcInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_linked_vpc")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetLinkedVpcInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LinkedVpcInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetLinkedVpcRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.LinkedVpcInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetLinkedVpcOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LinkedVpcInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetMgmtVmInfo(vmIdParam string) (model.MgmtVmInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_mgmt_vm_info")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetMgmtVmInfoInputType(), typeConverter)
	sv.AddStructField("VmId", vmIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MgmtVmInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetMgmtVmInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.MgmtVmInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetMgmtVmInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MgmtVmInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetPublicIp(publicIpIdParam string) (model.PublicIp, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_public_ip")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetPublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetPublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.PublicIp
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetPublicIpOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetRealizedStateStatus(intentPathParam string) (model.VmcConsolidatedRealizedStatus, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_realized_state_status")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetRealizedStateStatusInputType(), typeConverter)
	sv.AddStructField("IntentPath", intentPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VmcConsolidatedRealizedStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetRealizedStateStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.VmcConsolidatedRealizedStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetRealizedStateStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VmcConsolidatedRealizedStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) GetSddcUserConfig() (model.SddcUserConfiguration, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "get_sddc_user_config")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationGetSddcUserConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SddcUserConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationGetSddcUserConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.SddcUserConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationGetSddcUserConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SddcUserConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListAccounts() (model.VMCAccounts, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_accounts")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListAccountsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VMCAccounts
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListAccountsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.VMCAccounts
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListAccountsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VMCAccounts), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListAdvertisedRoutes() (model.BGPAdvertisedRoutes, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_advertised_routes")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListAdvertisedRoutesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BGPAdvertisedRoutes
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListAdvertisedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.BGPAdvertisedRoutes
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListAdvertisedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BGPAdvertisedRoutes), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListConnectedServices(linkedVpcIdParam string) (model.ConnectedServiceListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_connected_services")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListConnectedServicesInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConnectedServiceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListConnectedServicesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ConnectedServiceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListConnectedServicesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConnectedServiceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListLearnedRoutes() (model.BGPLearnedRoutes, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_learned_routes")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListLearnedRoutesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BGPLearnedRoutes
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListLearnedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.BGPLearnedRoutes
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListLearnedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BGPLearnedRoutes), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListLinkedVpcs() (model.LinkedVpcsListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_linked_vpcs")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListLinkedVpcsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LinkedVpcsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListLinkedVpcsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.LinkedVpcsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListLinkedVpcsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LinkedVpcsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListMgmtVms() (model.MgmtVmsListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_mgmt_vms")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListMgmtVmsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MgmtVmsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListMgmtVmsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.MgmtVmsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListMgmtVmsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MgmtVmsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListPublicIps() (model.PublicIpsListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_public_ips")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListPublicIpsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIpsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListPublicIpsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.PublicIpsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListPublicIpsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIpsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) ListVifs() (model.VifsListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "list_vifs")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationListVifsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VifsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationListVifsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.VifsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationListVifsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VifsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) UpdateConnectedService(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam model.ConnectedServiceStatus) (model.ConnectedServiceStatus, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "update_connected_service")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationUpdateConnectedServiceInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	sv.AddStructField("ServiceName", serviceNameParam)
	sv.AddStructField("ConnectedServiceStatus", connectedServiceStatusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConnectedServiceStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationUpdateConnectedServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ConnectedServiceStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationUpdateConnectedServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConnectedServiceStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) UpdateDxBgpInfo(directConnectBgpInfoParam model.DirectConnectBgpInfo) (model.DirectConnectBgpInfo, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "update_dx_bgp_info")
	sv := bindings.NewStructValueBuilder(nsxVmcAwsIntegrationUpdateDxBgpInfoInputType(), typeConverter)
	sv.AddStructField("DirectConnectBgpInfo", directConnectBgpInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectConnectBgpInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxVmcAwsIntegrationUpdateDxBgpInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := nIface.connector.NewExecutionContext()
	methodResult := nIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.DirectConnectBgpInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxVmcAwsIntegrationUpdateDxBgpInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectConnectBgpInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (nIface *DefaultNsxVmcAwsIntegrationClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := nIface.connector.GetApiProvider().Invoke(nIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (nIface *DefaultNsxVmcAwsIntegrationClient) attachVifMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationAttachVifInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationAttachVifOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "attachVif")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.attachVif method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) createPublicIpMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationCreatePublicIpInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationCreatePublicIpOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "createPublicIp")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.createPublicIp method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) deletePublicIpMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationDeletePublicIpInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationDeletePublicIpOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "deletePublicIp")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deletePublicIp method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) deleteVifByIdMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationDeleteVifByIdInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationDeleteVifByIdOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "deleteVifById")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.deleteVifById method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getDxBgpInfoMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetDxBgpInfoInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetDxBgpInfoOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getDxBgpInfo")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getDxBgpInfo method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getLinkedVpcMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetLinkedVpcInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetLinkedVpcOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getLinkedVpc")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getLinkedVpc method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getMgmtVmInfoMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetMgmtVmInfoInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetMgmtVmInfoOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getMgmtVmInfo")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getMgmtVmInfo method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getPublicIpMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetPublicIpInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetPublicIpOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getPublicIp")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getPublicIp method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getRealizedStateStatusMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetRealizedStateStatusInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetRealizedStateStatusOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getRealizedStateStatus")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getRealizedStateStatus method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) getSddcUserConfigMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetSddcUserConfigInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationGetSddcUserConfigOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getSddcUserConfig")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.getSddcUserConfig method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listAccountsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListAccountsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListAccountsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAccounts")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAccounts method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listAdvertisedRoutesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListAdvertisedRoutesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListAdvertisedRoutesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAdvertisedRoutes")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listAdvertisedRoutes method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listConnectedServicesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListConnectedServicesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListConnectedServicesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listConnectedServices")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listConnectedServices method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listLearnedRoutesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListLearnedRoutesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListLearnedRoutesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listLearnedRoutes")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLearnedRoutes method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listLinkedVpcsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListLinkedVpcsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListLinkedVpcsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listLinkedVpcs")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listLinkedVpcs method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listMgmtVmsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListMgmtVmsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListMgmtVmsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listMgmtVms")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listMgmtVms method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listPublicIpsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListPublicIpsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListPublicIpsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listPublicIps")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listPublicIps method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) listVifsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListVifsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationListVifsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listVifs")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.listVifs method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) updateConnectedServiceMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationUpdateConnectedServiceInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationUpdateConnectedServiceOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "updateConnectedService")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateConnectedService method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *DefaultNsxVmcAwsIntegrationClient) updateDxBgpInfoMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationUpdateDxBgpInfoInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nsxVmcAwsIntegrationUpdateDxBgpInfoOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "updateDxBgpInfo")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultNsxVmcAwsIntegrationClient.updateDxBgpInfo method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PrecheckByClusters
// Used by client-side stubs.

package nvds_urt

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type PrecheckByClustersClient interface {

	// Start precheck for N-VDS to VDS migration by clusters
	//
	// @param precheckParametersParam (required)
	// @param tolerateDifferentConfigurationsParam tolerate differnet configurations (optional, default to true)
	// @return com.vmware.nsx.model.NvdsUpgradePrecheckId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(precheckParametersParam model.PrecheckParameters, tolerateDifferentConfigurationsParam *bool) (model.NvdsUpgradePrecheckId, error)
}

type precheckByClustersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPrecheckByClustersClient(connector client.Connector) *precheckByClustersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.nvds_urt.precheck_by_clusters")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := precheckByClustersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *precheckByClustersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *precheckByClustersClient) Create(precheckParametersParam model.PrecheckParameters, tolerateDifferentConfigurationsParam *bool) (model.NvdsUpgradePrecheckId, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(precheckByClustersCreateInputType(), typeConverter)
	sv.AddStructField("PrecheckParameters", precheckParametersParam)
	sv.AddStructField("TolerateDifferentConfigurations", tolerateDifferentConfigurationsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NvdsUpgradePrecheckId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := precheckByClustersCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.precheck_by_clusters", "create", inputDataValue, executionContext)
	var emptyOutput model.NvdsUpgradePrecheckId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), precheckByClustersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NvdsUpgradePrecheckId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SiteRecoverySrmNodes
// Used by client-side stubs.

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

const _ = core.SupportedByRuntimeVersion1

type SiteRecoverySrmNodesClient interface {

	// Delete a SRM server.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param srmNodeParam SRM node identifier (required)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find SDDC or SRM node
	Delete(orgParam string, sddcParam string, srmNodeParam string) (model.Task, error)

	// Provision an additional SRM server.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param provisionSrmConfigParam Customization, for example can specify custom extension key suffix for SRM. (optional)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Post(orgParam string, sddcParam string, provisionSrmConfigParam *model.ProvisionSrmConfig) (model.Task, error)
}

type siteRecoverySrmNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSiteRecoverySrmNodesClient(connector client.Connector) *siteRecoverySrmNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.site_recovery_srm_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"post":   core.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := siteRecoverySrmNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *siteRecoverySrmNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *siteRecoverySrmNodesClient) Delete(orgParam string, sddcParam string, srmNodeParam string) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoverySrmNodesDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("SrmNode", srmNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoverySrmNodesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery_srm_nodes", "delete", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoverySrmNodesDeleteOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *siteRecoverySrmNodesClient) Post(orgParam string, sddcParam string, provisionSrmConfigParam *model.ProvisionSrmConfig) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoverySrmNodesPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("ProvisionSrmConfig", provisionSrmConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoverySrmNodesPostRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery_srm_nodes", "post", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoverySrmNodesPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

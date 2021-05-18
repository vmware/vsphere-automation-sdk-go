// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SiteRecovery
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

type SiteRecoveryClient interface {

	// Deactivate site recovery for the specified sddc
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param forceParam If = 'true', will deactivate site recovery forcefully. (optional)
	// @param deactivateHcxParam If = 'true', will deactivate HCX. (optional)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Delete(orgParam string, sddcParam string, forceParam *bool, deactivateHcxParam *bool) (model.Task, error)

	// Query site recovery configuration for the specified sddc
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @return com.vmware.vmc.draas.model.SiteRecovery
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string) (model.SiteRecovery, error)

	// Activate site recovery for the specified sddc
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param activateSiteRecoveryConfigParam Customization, for example can specify custom extension key suffix for SRM. (optional)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find site recovery configuration for sddc identifier
	Post(orgParam string, sddcParam string, activateSiteRecoveryConfigParam *model.ActivateSiteRecoveryConfig) (model.Task, error)
}

type siteRecoveryClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSiteRecoveryClient(connector client.Connector) *siteRecoveryClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.site_recovery")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"post":   core.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := siteRecoveryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *siteRecoveryClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *siteRecoveryClient) Delete(orgParam string, sddcParam string, forceParam *bool, deactivateHcxParam *bool) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoveryDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("DeactivateHcx", deactivateHcxParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoveryDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery", "delete", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoveryDeleteOutputType())
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

func (sIface *siteRecoveryClient) Get(orgParam string, sddcParam string) (model.SiteRecovery, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoveryGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SiteRecovery
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoveryGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery", "get", inputDataValue, executionContext)
	var emptyOutput model.SiteRecovery
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoveryGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SiteRecovery), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *siteRecoveryClient) Post(orgParam string, sddcParam string, activateSiteRecoveryConfigParam *model.ActivateSiteRecoveryConfig) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoveryPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("ActivateSiteRecoveryConfig", activateSiteRecoveryConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoveryPostRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery", "post", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoveryPostOutputType())
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

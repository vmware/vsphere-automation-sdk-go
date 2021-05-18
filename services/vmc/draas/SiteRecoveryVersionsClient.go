// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SiteRecoveryVersions
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

type SiteRecoveryVersionsClient interface {

	// Query site recovery versions for the specified sddc
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param versionSourceParam Represents the source for getting the version from. (optional)
	// @return com.vmware.vmc.draas.model.SiteRecoveryVersions
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find site recovery versions for sddc identifier
	Get(orgParam string, sddcParam string, versionSourceParam *string) (model.SiteRecoveryVersions, error)
}

type siteRecoveryVersionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSiteRecoveryVersionsClient(connector client.Connector) *siteRecoveryVersionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.site_recovery_versions")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := siteRecoveryVersionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *siteRecoveryVersionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *siteRecoveryVersionsClient) Get(orgParam string, sddcParam string, versionSourceParam *string) (model.SiteRecoveryVersions, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(siteRecoveryVersionsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("VersionSource", versionSourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SiteRecoveryVersions
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := siteRecoveryVersionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.site_recovery_versions", "get", inputDataValue, executionContext)
	var emptyOutput model.SiteRecoveryVersions
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), siteRecoveryVersionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SiteRecoveryVersions), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

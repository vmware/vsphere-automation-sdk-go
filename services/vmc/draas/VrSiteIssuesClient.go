// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VrSiteIssues
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

type VrSiteIssuesClient interface {

	// Query VR site issues
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Not found
	Get(orgParam string, sddcParam string) ([]model.HmsSiteIssueInfo, error)
}

type vrSiteIssuesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewVrSiteIssuesClient(connector client.Connector) *vrSiteIssuesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.vr_site_issues")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	vIface := vrSiteIssuesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vrSiteIssuesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vrSiteIssuesClient) Get(orgParam string, sddcParam string) ([]model.HmsSiteIssueInfo, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vrSiteIssuesGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.HmsSiteIssueInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vrSiteIssuesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.vr_site_issues", "get", inputDataValue, executionContext)
	var emptyOutput []model.HmsSiteIssueInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vrSiteIssuesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.HmsSiteIssueInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

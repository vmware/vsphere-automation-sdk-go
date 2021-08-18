// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: OrgRoot
// Used by client-side stubs.

package nsx_policy

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type OrgRootClient interface {

	// Read Org Root. Returns only the org root related properties. Inner object are not populated.
	//
	// @param basePathParam Base Path for retrieving hierarchical intent (optional)
	// @param filterParam Filter string as java regex (optional)
	// @param typeFilterParam Filter string to retrieve hierarchy. (optional)
	// @return com.vmware.nsx_policy.model.OrgRoot
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(basePathParam *string, filterParam *string, typeFilterParam *string) (model.OrgRoot, error)

	// Patch API at org-root to create/update/delete entire or part of intent hierarchy. Hierarchical API: Provides users a way to create entire or part of intent in single API invocation. Input is expressed in a tree format. Each node in tree can have multiple children of different types. System will resolve the dependecies of nodes within the intent tree and will create the model. Children for any node can be specified using ChildResourceReference or ChildPolicyConfigResource. If a resource is specified using ChildResourceReference then it will not be updated only its children will be updated. If Object is specified using ChildPolicyConfigResource, object along with its children will be updated. Hierarchical API can also be used to delete any sub-branch of entire tree.
	//
	// @param orgRootParam (required)
	// @param enforceRevisionCheckParam Force revision check (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgRootParam model.OrgRoot, enforceRevisionCheckParam *bool) error
}

type orgRootClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewOrgRootClient(connector client.Connector) *orgRootClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": core.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	oIface := orgRootClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *orgRootClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *orgRootClient) Get(basePathParam *string, filterParam *string, typeFilterParam *string) (model.OrgRoot, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgRootGetInputType(), typeConverter)
	sv.AddStructField("BasePath", basePathParam)
	sv.AddStructField("Filter", filterParam)
	sv.AddStructField("TypeFilter", typeFilterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OrgRoot
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgRootGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root", "get", inputDataValue, executionContext)
	var emptyOutput model.OrgRoot
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), orgRootGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OrgRoot), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *orgRootClient) Patch(orgRootParam model.OrgRoot, enforceRevisionCheckParam *bool) error {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgRootPatchInputType(), typeConverter)
	sv.AddStructField("OrgRoot", orgRootParam)
	sv.AddStructField("EnforceRevisionCheck", enforceRevisionCheckParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgRootPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

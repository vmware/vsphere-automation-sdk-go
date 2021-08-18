// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DeploymentGroups
// Used by client-side stubs.

package orgs

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type DeploymentGroupsClient interface {

	// Delete DeploymentGroup.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, deploymentGroupIdParam string) error

	// Get DeploymentGroup
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @return com.vmware.nsx_policy.model.DeploymentGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, deploymentGroupIdParam string) (model.DeploymentGroup, error)

	// Paginated list of DeploymentGroup.
	//
	// @param orgIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.DeploymentGroupListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DeploymentGroupListResult, error)

	// If a DeploymentGroup with the deployment-group-id is not already present, create a new DeploymentGroup. If it already exists, update the DeploymentGroup.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param deploymentGroupParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, deploymentGroupIdParam string, deploymentGroupParam model.DeploymentGroup) error

	// Update the DeploymentGroup.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param deploymentGroupParam (required)
	// @return com.vmware.nsx_policy.model.DeploymentGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, deploymentGroupIdParam string, deploymentGroupParam model.DeploymentGroup) (model.DeploymentGroup, error)
}

type deploymentGroupsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDeploymentGroupsClient(connector client.Connector) *deploymentGroupsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs.deployment_groups")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := deploymentGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentGroupsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentGroupsClient) Delete(orgIdParam string, deploymentGroupIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentGroupsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentGroupsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *deploymentGroupsClient) Get(orgIdParam string, deploymentGroupIdParam string) (model.DeploymentGroup, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentGroupsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DeploymentGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentGroupsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups", "get", inputDataValue, executionContext)
	var emptyOutput model.DeploymentGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentGroupsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DeploymentGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentGroupsClient) List(orgIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DeploymentGroupListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentGroupsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DeploymentGroupListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentGroupsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups", "list", inputDataValue, executionContext)
	var emptyOutput model.DeploymentGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentGroupsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DeploymentGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentGroupsClient) Patch(orgIdParam string, deploymentGroupIdParam string, deploymentGroupParam model.DeploymentGroup) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentGroupsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("DeploymentGroup", deploymentGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentGroupsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *deploymentGroupsClient) Update(orgIdParam string, deploymentGroupIdParam string, deploymentGroupParam model.DeploymentGroup) (model.DeploymentGroup, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(deploymentGroupsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("DeploymentGroup", deploymentGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DeploymentGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deploymentGroupsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups", "update", inputDataValue, executionContext)
	var emptyOutput model.DeploymentGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), deploymentGroupsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DeploymentGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

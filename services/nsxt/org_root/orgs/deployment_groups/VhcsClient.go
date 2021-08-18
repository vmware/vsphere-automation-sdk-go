// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Vhcs
// Used by client-side stubs.

package deployment_groups

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type VhcsClient interface {

	// Delete VHC.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string) error

	// Get VHC
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @return com.vmware.nsx_policy.model.Vhc
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string) (model.Vhc, error)

	// Paginated list of VHC.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.VhcListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, deploymentGroupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.VhcListResult, error)

	// If a VHC with the vhc-id is not already present, create a new VHC. If it already exists, update the VHC.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param vhcParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, vhcParam model.Vhc) error

	// Update the VHC. For local VHC, local transport zone and edge clusters needs to be added.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param vhcParam (required)
	// @return com.vmware.nsx_policy.model.Vhc
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, vhcParam model.Vhc) (model.Vhc, error)
}

type vhcsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewVhcsClient(connector client.Connector) *vhcsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	vIface := vhcsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vhcsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vhcsClient) Delete(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vhcsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vhcsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vhcsClient) Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string) (model.Vhc, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vhcsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Vhc
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vhcsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs", "get", inputDataValue, executionContext)
	var emptyOutput model.Vhc
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vhcsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Vhc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vhcsClient) List(orgIdParam string, deploymentGroupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.VhcListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vhcsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VhcListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vhcsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs", "list", inputDataValue, executionContext)
	var emptyOutput model.VhcListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vhcsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VhcListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vhcsClient) Patch(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, vhcParam model.Vhc) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vhcsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Vhc", vhcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vhcsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vhcsClient) Update(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, vhcParam model.Vhc) (model.Vhc, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vhcsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Vhc", vhcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Vhc
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vhcsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs", "update", inputDataValue, executionContext)
	var emptyOutput model.Vhc
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vhcsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Vhc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

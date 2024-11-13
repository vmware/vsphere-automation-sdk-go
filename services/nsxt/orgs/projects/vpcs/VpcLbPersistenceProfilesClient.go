// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VpcLbPersistenceProfiles
// Used by client-side stubs.

package vpcs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VpcLbPersistenceProfilesClient interface {

	// Delete the LBPersistenceProfile along with all the entities contained by this LBPersistenceProfile. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param vpcLbPersistenceProfileIdParam LBPersistenceProfile ID (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, forceParam *bool) error

	// Read a LBPersistenceProfile. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param vpcLbPersistenceProfileIdParam LBPersistenceProfile ID (required)
	// @return com.vmware.nsx_policy.model.LBPersistenceProfile
	// The return value will contain all the properties defined in nsx_policyModel.LBPersistenceProfile.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string) (*vapiData_.StructValue, error)

	// Paginated list of all LBPersistenceProfiles specified Vpc. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.LBPersistenceProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.LBPersistenceProfileListResult, error)

	// If a LBPersistenceProfile with the vpc-lb-persistence-profile-id is not already present, create a new LBPersistenceProfile. If it already exists, update the LBPersistenceProfile. This is a full replace. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param vpcLbPersistenceProfileIdParam LBPersistenceProfile ID (required)
	// @param lbPersistenceProfileParam (required)
	// The parameter must contain all the properties defined in nsx_policyModel.LBPersistenceProfile.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, lbPersistenceProfileParam *vapiData_.StructValue) error

	// If a LBPersistenceProfile with the vpc-lb-persistence-profile-id is not already present, create a new LBPersistenceProfile. If it already exists, update the LBPersistenceProfile. This is a full replace. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param vpcLbPersistenceProfileIdParam LBPersistenceProfile ID (required)
	// @param lbPersistenceProfileParam (required)
	// The parameter must contain all the properties defined in nsx_policyModel.LBPersistenceProfile.
	// @return com.vmware.nsx_policy.model.LBPersistenceProfile
	// The return value will contain all the properties defined in nsx_policyModel.LBPersistenceProfile.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, lbPersistenceProfileParam *vapiData_.StructValue) (*vapiData_.StructValue, error)
}

type vpcLbPersistenceProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVpcLbPersistenceProfilesClient(connector vapiProtocolClient_.Connector) *vpcLbPersistenceProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vpcLbPersistenceProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vpcLbPersistenceProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vpcLbPersistenceProfilesClient) Delete(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, forceParam *bool) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbPersistenceProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbPersistenceProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("VpcLbPersistenceProfileId", vpcLbPersistenceProfileIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vpcLbPersistenceProfilesClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string) (*vapiData_.StructValue, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbPersistenceProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbPersistenceProfilesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("VpcLbPersistenceProfileId", vpcLbPersistenceProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles", "get", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcLbPersistenceProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcLbPersistenceProfilesClient) List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.LBPersistenceProfileListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbPersistenceProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbPersistenceProfilesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.LBPersistenceProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.LBPersistenceProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcLbPersistenceProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.LBPersistenceProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vpcLbPersistenceProfilesClient) Patch(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, lbPersistenceProfileParam *vapiData_.StructValue) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbPersistenceProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbPersistenceProfilesPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("VpcLbPersistenceProfileId", vpcLbPersistenceProfileIdParam)
	sv.AddStructField("LbPersistenceProfile", lbPersistenceProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vpcLbPersistenceProfilesClient) Update(orgIdParam string, projectIdParam string, vpcIdParam string, vpcLbPersistenceProfileIdParam string, lbPersistenceProfileParam *vapiData_.StructValue) (*vapiData_.StructValue, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbPersistenceProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbPersistenceProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("VpcLbPersistenceProfileId", vpcLbPersistenceProfileIdParam)
	sv.AddStructField("LbPersistenceProfile", lbPersistenceProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_persistence_profiles", "update", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcLbPersistenceProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

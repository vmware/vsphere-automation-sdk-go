// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VpcLbNodeCapacityStatus
// Used by client-side stubs.

package vpcs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VpcLbNodeCapacityStatusClient interface {

	// API is used to retrieve available load balancer capacity status for all service clusters shared with VPC. Currently Edge cluster is supported as service cluster. If the parameter ?service_cluster_path=<service-cluster-path> exists, then API only shows the available load balancer capacity status for this specific service cluster. NSX Load Balancer availability in terms of use-cases and editions is specified in NSX Feature and Edition Guide. Please review before consuming those APIs.
	//
	// @param orgIdParam Org ID (required)
	// @param projectIdParam Project ID (required)
	// @param vpcIdParam VPC ID (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param serviceClusterPathParam The service cluster path (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx_policy.model.AggregateLBNodeCapacitySummary
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, enforcementPointPathParam *string, serviceClusterPathParam *string, sourceParam *string) (nsx_policyModel.AggregateLBNodeCapacitySummary, error)
}

type vpcLbNodeCapacityStatusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVpcLbNodeCapacityStatusClient(connector vapiProtocolClient_.Connector) *vpcLbNodeCapacityStatusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_node_capacity_status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vpcLbNodeCapacityStatusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vpcLbNodeCapacityStatusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vpcLbNodeCapacityStatusClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, enforcementPointPathParam *string, serviceClusterPathParam *string, sourceParam *string) (nsx_policyModel.AggregateLBNodeCapacitySummary, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vpcLbNodeCapacityStatusGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vpcLbNodeCapacityStatusGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("ServiceClusterPath", serviceClusterPathParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AggregateLBNodeCapacitySummary
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.vpc_lb_node_capacity_status", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AggregateLBNodeCapacitySummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VpcLbNodeCapacityStatusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AggregateLBNodeCapacitySummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

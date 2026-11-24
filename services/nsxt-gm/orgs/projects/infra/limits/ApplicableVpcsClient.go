// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ApplicableVpcs
// Used by client-side stubs.

package limits

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ApplicableVpcsClient interface {

	// This API returns, for each requested limit, the VPCs on which the limit is applicable. A limit is applicable on a VPC if it is explicitly applied to the VPC, or if it is applied to the project of the VPC and at least one of the IP blocks in the limit is used by the VPC. A maximum of 25 limit paths can be requested in a single call. A limit path that does not exist is omitted from the response. At /infra, the request represents the default project where limits are not supported, so no applicable VPCs are returned unless all_projects is set to true, in which case VPCs from all projects are considered. In the context of a project, only VPCs of that project are considered.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param limitPathsParam Comma separated policy paths of the limits for which applicable VPCs are requested. A maximum of 25 paths is allowed. Limits can be defined at /infra or within a project. (required)
	// @param allProjectsParam Applicable only when the request is made at /infra. If true, VPCs from all projects on which a limit is applicable are returned. If false or not set, the request at /infra represents the default project, where limits are not supported, and the list of applicable VPCs is empty for every limit. When the request is made in the context of a project, VPCs are always restricted to that project and setting this field to true results in an error. (optional)
	// @return com.vmware.nsx_global_policy.model.LimitApplicableVpcsListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, limitPathsParam string, allProjectsParam *bool) (nsx_global_policyModel.LimitApplicableVpcsListResult, error)
}

type applicableVpcsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewApplicableVpcsClient(connector vapiProtocolClient_.Connector) *applicableVpcsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.orgs.projects.infra.limits.applicable_vpcs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := applicableVpcsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *applicableVpcsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *applicableVpcsClient) Get(orgIdParam string, projectIdParam string, limitPathsParam string, allProjectsParam *bool) (nsx_global_policyModel.LimitApplicableVpcsListResult, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := applicableVpcsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(applicableVpcsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("LimitPaths", limitPathsParam)
	sv.AddStructField("AllProjects", allProjectsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.LimitApplicableVpcsListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.orgs.projects.infra.limits.applicable_vpcs", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.LimitApplicableVpcsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ApplicableVpcsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.LimitApplicableVpcsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

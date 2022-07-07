// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: MemberTypes
// Used by client-side stubs.

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type MemberTypesClient interface {

	// It retrieves member types for a given group. In case of nested groups, it calculates member types of child groups as well. Considers member type for members added via static members and dynamic membership criteria.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param domainIdParam Domain ID (required)
	// @param groupIdParam Group ID (required)
	// @return com.vmware.nsx_policy.model.GroupMemberTypeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, domainIdParam string, groupIdParam string) (model.GroupMemberTypeListResult, error)
}

type memberTypesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMemberTypesClient(connector client.Connector) *memberTypesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.infra.domains.groups.member_types")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := memberTypesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *memberTypesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *memberTypesClient) Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, domainIdParam string, groupIdParam string) (model.GroupMemberTypeListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(memberTypesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMemberTypeListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := memberTypesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.infra.domains.groups.member_types", "get", inputDataValue, executionContext)
	var emptyOutput model.GroupMemberTypeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), memberTypesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMemberTypeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

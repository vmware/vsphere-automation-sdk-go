// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Get_vm_group_execution_details
// Used by client-side stubs.

package actions

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type Get_vm_group_execution_detailsClient interface {

	// The result includes a \"logical_switch_id_to_vm_instance_id_and_vnics_map\" list and an optional \"failedVmInstanceIds\" list which includes the uuids of VMs that are not found in the source VC. Construct a map of vmInstanceUuid to (vnic, ls_id) from the \"logical_switch_id_to_vm_instance_id_and_vnics_map\", then use the map to populate the relocate spec of each VM and migrate the VM as below: The VM object can be found in source VC by the vmInstanceUuid using VC API serviceinstance.content.searchIndex.FindByUuid(uuid=vmInstanceUuid, vmSearch=True, instanceUuid=True) For each VM vNIC, if the device key \"vnic\" is not found in the map of vmInstanceUuid to (vnic, ls_id), then skip the vNIC. Otherwise form a VIF-id for the vNIC by vmInstanceUuid + ':' + str(vnic), e.g. \"52630e5d-ce6f-fac0-424c-4aa4bdf6bd56:4001\", and use it to setup the vNIC's network backing. For VDS6.x migration, opaque-network backing in NVDS will be used; setup the vNIC device as below: - vdevice.backing = vim.vm.device.VirtualEthernetCard.OpaqueNetworkBackingInfo() - vdevice.backing.opaqueNetworkId = ls_id - vdevice.backing.opaqueNetworkType = 'nsx.LogicalSwitch' - vdevice.externalId = VIF-id For VDS 7.0 and later versions, \"nsx\" LogicalSwitch backing in VDS will be used. Go through all VDS DVPGs in the target VC to find each DVPG whose config.backingType is \"nsx\" and config.logicalSwitchUuid is not blank, build a map of dvpg.config.logicalSwitchUuid to [dvpg.key, dvpg.config.distributedVirtualSwitch.uuid] once. For each vNIC device, get the dvpg value from the map by the ls_id and then set it up as below: - vdsPgConn = vim.dvs.PortConnection() - vdsPgConn.portgroupKey = dvpg.key - vdsPgConn.switchUuid = dvpg.config.distributedVirtualSwitch.uuid - vdevice.backing = vim.vm.device.VirtualEthernetCard.DistributedVirtualPortBackingInfo() - vdevice.backing.port = vdsPgConn - vdevice.externalId = VIF-id
	//
	// @param groupIdParam User defined VM group ID (required)
	// @param federationSiteIdParam ID of the site in NSX-T Federation (optional)
	// @return com.vmware.nsx.model.VmGroupExecutionDetails
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(groupIdParam string, federationSiteIdParam *string) (nsxModel.VmGroupExecutionDetails, error)
}

type get_vm_group_execution_detailsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewGet_vm_group_execution_detailsClient(connector vapiProtocolClient_.Connector) *get_vm_group_execution_detailsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.vmgroup.actions.get_vm_group_execution_details")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	gIface := get_vm_group_execution_detailsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *get_vm_group_execution_detailsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *get_vm_group_execution_detailsClient) List(groupIdParam string, federationSiteIdParam *string) (nsxModel.VmGroupExecutionDetails, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := getVmGroupExecutionDetailsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(getVmGroupExecutionDetailsListInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("FederationSiteId", federationSiteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.VmGroupExecutionDetails
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.vmgroup.actions.get_vm_group_execution_details", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.VmGroupExecutionDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GetVmGroupExecutionDetailsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.VmGroupExecutionDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

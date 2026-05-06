// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodeProfiles
// Used by client-side stubs.

package nsx

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TransportNodeProfilesClient interface {

	// Transport node profile captures the configuration needed to create a transport node. A transport node profile can be attached to compute collections for automatic TN creation of member hosts.
	//  This api is now deprecated. Please use new api - /policy/api/v1/infra/host-transport-node-profiles/<host-transport-node-profile-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeProfileParam (required)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @return com.vmware.nsx.model.TransportNodeProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(transportNodeProfileParam nsxModel.TransportNodeProfile, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeProfile, error)

	// Deletes the specified transport node profile. A transport node profile can be deleted only when it is not attached to any compute collection.
	//  This api is now deprecated. Please use new api - /policy/api/v1/infra/host-transport-node-profiles/<host-transport-node-profile-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeProfileIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(transportNodeProfileIdParam string) error

	// Returns information about a specified transport node profile.
	//  This api is now deprecated. Please use new api - /policy/api/v1/infra/host-transport-node-profiles/<host-transport-node-profile-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeProfileIdParam (required)
	// @return com.vmware.nsx.model.TransportNodeProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(transportNodeProfileIdParam string) (nsxModel.TransportNodeProfile, error)

	// Returns information about all transport node profiles.
	//  This api is now deprecated. Please use new api - /policy/api/v1/infra/host-transport-node-profiles
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.TransportNodeProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.TransportNodeProfileListResult, error)

	// When configurations of a transport node profile(TNP) is updated, all the transport nodes in all the compute collections to which this TNP is attached are updated to reflect the updated configuration.
	//  This api is now deprecated. Please use new api - /policy/api/v1/infra/host-transport-node-profiles/<host-transport-node-profile-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param transportNodeProfileIdParam (required)
	// @param transportNodeProfileParam (required)
	// @param esxMgmtIfMigrationDestParam A comma separated list of network ids. When migrating vmks into logical switches, the ids are the logical switches's ids. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). This property can only used together with 'if_id'. (optional)
	// @param ifIdParam A comma separated list of vmk interfaces (for example, vmk0,vmk1). This property can only used along with 'esx_mgmt_if_migration_dest'. If all vmk interfaces will be migrated into the same logical switch or DV portgroup, the 'esx_mgmt_if_migration_dest' can be just one logical switch id or DV portgroup name. Otherwise the number of vmks in this list must equal the number of ids in 'esx_mgmt_if_migration_dest' list, and the orders of the two lists are important because the vmks match the network ids one by one in the same order. (optional)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @param pingIpParam A comma separated list of IP addresses that match the vmk interfaces given in property 'if_id\" or 'vnic' one-by-one in the same order. '0.0.0.0' is a special IP that indicates the pre-migration gateway of the vmk will be pinged post-migration. If a VMK does not need the ping ip or a VM NIC is given inside 'vnic', the ping ip must be skipped but the comma has to stay. For example, '0.0.0.0,,10.1.1.1' indicates the vmk or VM NIC at the 2nd position does not need ping post-migration. Right after all ESX vmk interfaces are migrated, ping packets will be sent through each vmk to its given ping_ip to check if the migraton will break the network connectivity or not. If any vmk_ping fails, the whole migration of all vmks will be rolled back and transport-node will be in failed state. (optional)
	// @param skipValidationParam If this property is set true, all front-end validation for vmk, vnic, and/or pnic migration will be skipped. This is useful when the remote host becomes unreachable as a result of a migration; in which case the front-end validation will always fail because data from the remote host is no longer available. Skipping the validation will allow user to undo the migration by updating the transport node first and then restoring the host network connectivity. (optional, default to false)
	// @param vnicParam A comma separated list of vmk interfaces and/or one VM NIC. Only one VM NIC is allowed in the list; the format must be vmInstanceUuid:DeviceId like '50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. An example list is 'vmk0,vmk1,50ca5f2d-1fa2-432d-991e-f01e0e16d182:4000'. The property can only be used along with 'vnic_migration_dest'. (optional)
	// @param vnicMigrationDestParam A comma separated list of vif ids, or port group names. When migrating into logical switches, the ids are vif ids in the logical ports created in the logical switches. When migrating out of logical switches, the ids are vSphere Standard Switch portgroup names in a single vSphere Standard Switch, or distributed virtual portgroup names in a single distributed virtual switch (DVS). The property can only be used in combination with property 'vnic'. The number of vnic interfaces in 'vnic' must equal the number of vif ids or port-group names in this list. The items in the two lists match by the the order. (optional)
	// @return com.vmware.nsx.model.TransportNodeProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(transportNodeProfileIdParam string, transportNodeProfileParam nsxModel.TransportNodeProfile, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsxModel.TransportNodeProfile, error)
}

type transportNodeProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTransportNodeProfilesClient(connector vapiProtocolClient_.Connector) *transportNodeProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_node_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := transportNodeProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodeProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodeProfilesClient) Create(transportNodeProfileParam nsxModel.TransportNodeProfile, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeProfilesCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeProfilesCreateInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfile", transportNodeProfileParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_profiles", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeProfilesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeProfilesClient) Delete(transportNodeProfileIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodeProfilesClient) Get(transportNodeProfileIdParam string) (nsxModel.TransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeProfilesGetInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeProfilesClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.TransportNodeProfileListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeProfilesClient) Update(transportNodeProfileIdParam string, transportNodeProfileParam nsxModel.TransportNodeProfile, esxMgmtIfMigrationDestParam *string, ifIdParam *string, overrideNsxOwnershipParam *bool, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (nsxModel.TransportNodeProfile, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	sv.AddStructField("TransportNodeProfile", transportNodeProfileParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

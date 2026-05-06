// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodeCollections
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

type TransportNodeCollectionsClient interface {

	// When transport node collection is created the hosts which are part of compute collection will be prepared automatically i.e. NSX Manager attempts to install the NSX components on hosts. Transport nodes for these hosts are created using the configuration specified in transport node profile.
	//
	// @param transportNodeCollectionParam (required)
	// @param applyProfileParam This flag should be used when the configuration specified by the transport_node_profile_id should not be applied to existing hosts referred to by the compute_collection_id during transport node collection creation. If this flag is set to false, the TNP configuration will not be applied to any of the hosts in the cluster during creation. Any transport node that exists in the cluster that has a different configuration than the TNP configuration will have the is_overridden flag set to true. This will result in the transport node collection creation completing with a status of PROFILE_MISMATCH. If this flag is set to true, the default value, the TNP configuration will be applied to all hosts in the cluster during transport node collection creation. (optional, default to true)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @return com.vmware.nsx.model.TransportNodeCollection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(transportNodeCollectionParam nsxModel.TransportNodeCollection, applyProfileParam *bool, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeCollection, error)

	// By deleting transport node collection, we are detaching the transport node profile(TNP) from the compute collection. It has no effect on existing transport nodes. However, new hosts added to the compute collection will no longer be automatically converted to NSX transport node. Detaching TNP from compute collection does not delete TNP.
	//
	// @param transportNodeCollectionIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(transportNodeCollectionIdParam string) error

	// Returns transport node collection by id
	//
	// @param transportNodeCollectionIdParam (required)
	// @return com.vmware.nsx.model.TransportNodeCollection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(transportNodeCollectionIdParam string) (nsxModel.TransportNodeCollection, error)

	// Returns all Transport Node collections
	//
	// @param clusterMoidParam Managed object ID of cluster in VC. vc_instance_uuid has to be provided along with this parameter otherwise it will return empty list. (optional)
	// @param computeCollectionIdParam Compute collection id against which the serach will be done. If this parameter is provided then other parameters will be ignored. (optional)
	// @param vcInstanceUuidParam This is UUID of VC deployment as seen in managed objects of VC as \"instanceUuid\". cluster_moid has to be provided along with this parameter otherwise it will return empty list. (optional)
	// @return com.vmware.nsx.model.TransportNodeCollectionListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(clusterMoidParam *string, computeCollectionIdParam *string, vcInstanceUuidParam *string) (nsxModel.TransportNodeCollectionListResult, error)

	// This API is relevant for compute collection on which vLCM is enabled. This API shpuld be invoked to retry the realization of transport node profile on the compute collection. This is useful when profile realization had failed because of error in vLCM. This API has no effect if vLCM is not enabled on the computer collection.
	//
	// @param transportNodeCollectionIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Retryprofilerealization(transportNodeCollectionIdParam string) error

	// Attach different transport node profile to compute collection by updating transport node collection.
	//
	// @param transportNodeCollectionIdParam (required)
	// @param transportNodeCollectionParam (required)
	// @param overrideNsxOwnershipParam Flag indicating whether the NSX ownership constraints (on Managed Objects like Host/Cluster/DVS) should be overridden/bypassed. Note: Overriding/bypassing NSX ownership constraints is not recommended at all. This indicates, you want to use/configure/own certain Managed Objects (like Cluster, Host or DVS) which seem to be already in use/configured/owned by some other NSX instance. This option should be used with caution. It should only be used to come out of situations where: a. The other NSX instance no longer intends to use the Managed Objects (and has already unconfigured NSX configurations) but the ownership still lies with it (incorrectly) and you want those Managed Objects to be used/configured/owned by this NSX instance. b. The other NSX instance has crashed or decommisioned but the ownership still lies with it and you want those Managed Objects to be used/configured/owned by this NSX instance. Enabling this option, while the Managed Objects affected by this operation are actively used by other NSX, can lead to problematic states on both the NSX instances. For example, if a TN is forcefully reconfigured by this NSX instance (using override_nsx_ownership=true), while it was already configured and in use by the other NSX instance, it could corrupt the HostSwitch configurations pushed down by the other NSX instance. (optional, default to false)
	// @return com.vmware.nsx.model.TransportNodeCollection
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(transportNodeCollectionIdParam string, transportNodeCollectionParam nsxModel.TransportNodeCollection, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeCollection, error)
}

type transportNodeCollectionsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTransportNodeCollectionsClient(connector vapiProtocolClient_.Connector) *transportNodeCollectionsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_node_collections")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":                  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":                  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":                     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"retryprofilerealization": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "retryprofilerealization"),
		"update":                  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := transportNodeCollectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodeCollectionsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodeCollectionsClient) Create(transportNodeCollectionParam nsxModel.TransportNodeCollection, applyProfileParam *bool, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeCollection, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsCreateInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollection", transportNodeCollectionParam)
	sv.AddStructField("ApplyProfile", applyProfileParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeCollection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeCollectionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeCollectionsClient) Delete(transportNodeCollectionIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsDeleteInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "delete", inputDataValue, executionContext)
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

func (tIface *transportNodeCollectionsClient) Get(transportNodeCollectionIdParam string) (nsxModel.TransportNodeCollection, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsGetInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeCollection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeCollectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeCollectionsClient) List(clusterMoidParam *string, computeCollectionIdParam *string, vcInstanceUuidParam *string) (nsxModel.TransportNodeCollectionListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsListInputType(), typeConverter)
	sv.AddStructField("ClusterMoid", clusterMoidParam)
	sv.AddStructField("ComputeCollectionId", computeCollectionIdParam)
	sv.AddStructField("VcInstanceUuid", vcInstanceUuidParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeCollectionListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeCollectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeCollectionsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeCollectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeCollectionsClient) Retryprofilerealization(transportNodeCollectionIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsRetryprofilerealizationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsRetryprofilerealizationInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "retryprofilerealization", inputDataValue, executionContext)
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

func (tIface *transportNodeCollectionsClient) Update(transportNodeCollectionIdParam string, transportNodeCollectionParam nsxModel.TransportNodeCollection, overrideNsxOwnershipParam *bool) (nsxModel.TransportNodeCollection, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeCollectionsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeCollectionsUpdateInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	sv.AddStructField("TransportNodeCollection", transportNodeCollectionParam)
	sv.AddStructField("OverrideNsxOwnership", overrideNsxOwnershipParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeCollection
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_node_collections", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeCollectionsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

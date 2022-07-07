// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServerProfiles
// Used by client-side stubs.

package dhcp

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type ServerProfilesClient interface {

	// Create a DHCP server profile. If no edge member is specified, edge members to run the dhcp servers will be auto-allocated from the edge cluster.
	//
	// @param dhcpProfileParam (required)
	// @return com.vmware.nsx.model.DhcpProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error)

	// Delete a DHCP server profile specified by the profile id.
	//
	// @param profileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(profileIdParam string) error

	// Return the DHCP profile specified by the profile id.
	//
	// @param profileIdParam (required)
	// @return com.vmware.nsx.model.DhcpProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(profileIdParam string) (model.DhcpProfile, error)

	// Get a paginated list of DHCP server profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.DhcpProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DhcpProfileListResult, error)

	// As changing edge-cluster-id of a DhcpProfile by a PUT is disallowed, this re-allocate API is used to modify the edge-cluster-id and members of a given DhcpProfile. Only the edge-cluster-id and the edge-cluster-member-indexes fields will be picked up by this re-allication API. The othere fields in the payload will be ignored. If the edge-cluster-id in the payload DhcpProfile is different from the current edge-cluster-id of the profile, the referencing DHCP server(s) will be re-allocated to the new edge cluster. If the edge-cluster-id is not changed, the referencing DHCP server(s) will be re-allocated to the given edge members in the edge cluster. In this case, this REST API will act same as that of updating a DhcpProfile. If the edge cluster member indexes are provided, they should exist in the given edge cluster. If the indexes are not specified in the DhcpProfile, edge members will be auto-allocated from the given edge cluster. Please note that re-allocating edge-cluster will cause lose of all exisitng DHCP lease information. This API is used only when loosing DHCP leases is not a real problem, e.g. cross-site migration or failover and all client hosts will be reboot and get new IP addresses.
	//
	// @param serverProfileIdParam (required)
	// @param dhcpProfileParam (required)
	// @return com.vmware.nsx.model.DhcpProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reallocate(serverProfileIdParam string, dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error)

	// If both the edge_cluster_member_indexes in the DhcpProfile are changed in a same PUT API, e.g. change from [a,b] to [x,y], the current DHCP server leases will be lost, which could cause the network crash due to ip conflicts. Hence the suggestion is to change only one member index in one single update, e.g. from [a, b] to [a,y]. Please note, the edge_cluster_id in DhcpProfile can NOT be changed by this PUT operation because all existing DHCP leases will lost. If losing leases is not a problem, a dedicated re-allocation API is suggested to modify the edge-cluster-id, i.e. \"POST /api/v1/dhcp/dhcp-profiles/<profileiid>?action=reallocate\". Meanwhile, if the edge_cluster_member_indexes was specified currently but now is changed to none (not specified) via a PUT operation, the edge nodes will not be auto-selected from edge cluster. Instead, the previously-allocated edge nodes will continue to be used by the DHCP server. This is because changing both edge nodes of a DHCP server will lose all existing leases. In case re-allocation is required and leases lost is not a problem (or can be recovered), please invoke the reallocate API mentioned above with new DhcpProfile to accomplish the intent.
	//
	// @param profileIdParam (required)
	// @param dhcpProfileParam (required)
	// @return com.vmware.nsx.model.DhcpProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(profileIdParam string, dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error)
}

type serverProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewServerProfilesClient(connector client.Connector) *serverProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.dhcp.server_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":     core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":     core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":        core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":       core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reallocate": core.NewMethodIdentifier(interfaceIdentifier, "reallocate"),
		"update":     core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := serverProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serverProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serverProfilesClient) Create(dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesCreateInputType(), typeConverter)
	sv.AddStructField("DhcpProfile", dhcpProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "create", inputDataValue, executionContext)
	var emptyOutput model.DhcpProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serverProfilesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serverProfilesClient) Delete(profileIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serverProfilesClient) Get(profileIdParam string) (model.DhcpProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesGetInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "get", inputDataValue, executionContext)
	var emptyOutput model.DhcpProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serverProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serverProfilesClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DhcpProfileListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "list", inputDataValue, executionContext)
	var emptyOutput model.DhcpProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serverProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serverProfilesClient) Reallocate(serverProfileIdParam string, dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesReallocateInputType(), typeConverter)
	sv.AddStructField("ServerProfileId", serverProfileIdParam)
	sv.AddStructField("DhcpProfile", dhcpProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesReallocateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "reallocate", inputDataValue, executionContext)
	var emptyOutput model.DhcpProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serverProfilesReallocateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serverProfilesClient) Update(profileIdParam string, dhcpProfileParam model.DhcpProfile) (model.DhcpProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serverProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("DhcpProfile", dhcpProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serverProfilesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.server_profiles", "update", inputDataValue, executionContext)
	var emptyOutput model.DhcpProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serverProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

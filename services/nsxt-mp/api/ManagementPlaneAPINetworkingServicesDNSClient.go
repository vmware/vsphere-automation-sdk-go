// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingServicesDNS
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPINetworkingServicesDNSClient interface {

	// Clear the current cache of the DNS forwarder.
	//
	// @param forwarderIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ClearDnsForwarderCacheClearCache(forwarderIdParam string) error

	// Create a DNS forwader upon a logical router. There is only one DNS forwarder can be created upon a given logical router.
	//
	// @param dnsForwarderParam (required)
	// @return com.vmware.model.DnsForwarder
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateDnsForwader(dnsForwarderParam model.DnsForwarder) (model.DnsForwarder, error)

	// Delete a specific DNS forwarder.
	//
	// @param forwarderIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteDnsForwarder(forwarderIdParam string) error

	// Disable the DNS forwarder if the forwarder is currently enbled. If the DNS forwarder is already disabled, the forwarder will not be re-disabled. Please note, once a DNS forwarder is disabled then enabled, the previous DNS forwarder statistics counters will be reset.
	//
	// @param forwarderIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DisableDnsForwarderDisable(forwarderIdParam string) error

	// Enable the DNS forwarder if the forwarder is currently disabled. If the DNS forwarder is already enabled, the forwarder will not be re-enabled. Please note, once a DNS forwarder is disabled then enabled, the previous DNS forwarder statistics counters will be reset.
	//
	// @param forwarderIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	EnableDnsForwarderEnable(forwarderIdParam string) error

	// Return the realized state information of a DNS forwarder. After a DNS forwarder was created or updated, you can invoke this API to check the realization state of the forwarder.
	//
	// @param forwarderIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.model.ConfigurationState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetDnsForwarderState(forwarderIdParam string, barrierIdParam *int64, requestIdParam *string) (model.ConfigurationState, error)

	// Returns the current status of the given DNS forwarder.
	//
	// @param forwarderIdParam (required)
	// @return com.vmware.model.DnsForwarderStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetDnsForwarderStatus(forwarderIdParam string) (model.DnsForwarderStatus, error)

	// Return the given count of recent failed DNS queries from DNS forwarder. Since the DNS forwarder is running in Acitve/Standby HA mode on transport nodes, the given count of queries will be returned from each nodes. Hence the total queries returned could be doubled. If no count is specified, 100 recent failed queries are returned. If the recent failures is less than the given count, all the failures will be returned. The maximum count is 1,000.
	//
	// @param forwarderIdParam (required)
	// @param countParam The count of the failed DNS queries (optional, default to 100)
	// @return com.vmware.model.DnsFailedQueries
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetFailedDnsQueries(forwarderIdParam string, countParam *int64) (model.DnsFailedQueries, error)

	// Get a paginated list of DNS forwarders.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.DnsForwarderListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListDnsForwaders(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DnsForwarderListResult, error)

	// Query the nameserver for an ip-address or a FQDN of the given an address optionally using an specified DNS server. If the address is a fqdn, nslookup will resolve ip-address with it. If the address is an ip-address, do a reverse lookup and answer fqdn(s).
	//
	// @param forwarderIdParam (required)
	// @param addressParam IP address or FQDN for nslookup (optional)
	// @param serverIpParam IPv4 address (optional)
	// @param sourceIpParam IPv4 address (optional)
	// @return com.vmware.model.DnsAnswer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	LookupAddress(forwarderIdParam string, addressParam *string, serverIpParam *string, sourceIpParam *string) (model.DnsAnswer, error)

	// Retrieve a DNS forwarder.
	//
	// @param forwarderIdParam (required)
	// @return com.vmware.model.DnsForwarder
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadDnsForwader(forwarderIdParam string) (model.DnsForwarder, error)

	// Update a specific DNS forwarder.
	//
	// @param forwarderIdParam (required)
	// @param dnsForwarderParam (required)
	// @return com.vmware.model.DnsForwarder
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateDnsForwarder(forwarderIdParam string, dnsForwarderParam model.DnsForwarder) (model.DnsForwarder, error)
}

type managementPlaneAPINetworkingServicesDNSClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingServicesDNSClient(connector client.Connector) *managementPlaneAPINetworkingServicesDNSClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_services_DNS")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"clear_dns_forwarder_cache_clear_cache": core.NewMethodIdentifier(interfaceIdentifier, "clear_dns_forwarder_cache_clear_cache"),
		"create_dns_forwader":                   core.NewMethodIdentifier(interfaceIdentifier, "create_dns_forwader"),
		"delete_dns_forwarder":                  core.NewMethodIdentifier(interfaceIdentifier, "delete_dns_forwarder"),
		"disable_dns_forwarder_disable":         core.NewMethodIdentifier(interfaceIdentifier, "disable_dns_forwarder_disable"),
		"enable_dns_forwarder_enable":           core.NewMethodIdentifier(interfaceIdentifier, "enable_dns_forwarder_enable"),
		"get_dns_forwarder_state":               core.NewMethodIdentifier(interfaceIdentifier, "get_dns_forwarder_state"),
		"get_dns_forwarder_status":              core.NewMethodIdentifier(interfaceIdentifier, "get_dns_forwarder_status"),
		"get_failed_dns_queries":                core.NewMethodIdentifier(interfaceIdentifier, "get_failed_dns_queries"),
		"list_dns_forwaders":                    core.NewMethodIdentifier(interfaceIdentifier, "list_dns_forwaders"),
		"lookup_address":                        core.NewMethodIdentifier(interfaceIdentifier, "lookup_address"),
		"read_dns_forwader":                     core.NewMethodIdentifier(interfaceIdentifier, "read_dns_forwader"),
		"update_dns_forwarder":                  core.NewMethodIdentifier(interfaceIdentifier, "update_dns_forwarder"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingServicesDNSClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) ClearDnsForwarderCacheClearCache(forwarderIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSClearDnsForwarderCacheClearCacheInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSClearDnsForwarderCacheClearCacheRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "clear_dns_forwarder_cache_clear_cache", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) CreateDnsForwader(dnsForwarderParam model.DnsForwarder) (model.DnsForwarder, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSCreateDnsForwaderInputType(), typeConverter)
	sv.AddStructField("DnsForwarder", dnsForwarderParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsForwarder
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSCreateDnsForwaderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "create_dns_forwader", inputDataValue, executionContext)
	var emptyOutput model.DnsForwarder
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSCreateDnsForwaderOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsForwarder), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) DeleteDnsForwarder(forwarderIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSDeleteDnsForwarderInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSDeleteDnsForwarderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "delete_dns_forwarder", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) DisableDnsForwarderDisable(forwarderIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSDisableDnsForwarderDisableInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSDisableDnsForwarderDisableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "disable_dns_forwarder_disable", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) EnableDnsForwarderEnable(forwarderIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSEnableDnsForwarderEnableInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSEnableDnsForwarderEnableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "enable_dns_forwarder_enable", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) GetDnsForwarderState(forwarderIdParam string, barrierIdParam *int64, requestIdParam *string) (model.ConfigurationState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSGetDnsForwarderStateInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConfigurationState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSGetDnsForwarderStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "get_dns_forwarder_state", inputDataValue, executionContext)
	var emptyOutput model.ConfigurationState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSGetDnsForwarderStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConfigurationState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) GetDnsForwarderStatus(forwarderIdParam string) (model.DnsForwarderStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSGetDnsForwarderStatusInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsForwarderStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSGetDnsForwarderStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "get_dns_forwarder_status", inputDataValue, executionContext)
	var emptyOutput model.DnsForwarderStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSGetDnsForwarderStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsForwarderStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) GetFailedDnsQueries(forwarderIdParam string, countParam *int64) (model.DnsFailedQueries, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSGetFailedDnsQueriesInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	sv.AddStructField("Count", countParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsFailedQueries
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSGetFailedDnsQueriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "get_failed_dns_queries", inputDataValue, executionContext)
	var emptyOutput model.DnsFailedQueries
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSGetFailedDnsQueriesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsFailedQueries), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) ListDnsForwaders(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DnsForwarderListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSListDnsForwadersInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsForwarderListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSListDnsForwadersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "list_dns_forwaders", inputDataValue, executionContext)
	var emptyOutput model.DnsForwarderListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSListDnsForwadersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsForwarderListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) LookupAddress(forwarderIdParam string, addressParam *string, serverIpParam *string, sourceIpParam *string) (model.DnsAnswer, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSLookupAddressInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	sv.AddStructField("Address", addressParam)
	sv.AddStructField("ServerIp", serverIpParam)
	sv.AddStructField("SourceIp", sourceIpParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsAnswer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSLookupAddressRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "lookup_address", inputDataValue, executionContext)
	var emptyOutput model.DnsAnswer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSLookupAddressOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsAnswer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) ReadDnsForwader(forwarderIdParam string) (model.DnsForwarder, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSReadDnsForwaderInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsForwarder
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSReadDnsForwaderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "read_dns_forwader", inputDataValue, executionContext)
	var emptyOutput model.DnsForwarder
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSReadDnsForwaderOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsForwarder), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingServicesDNSClient) UpdateDnsForwarder(forwarderIdParam string, dnsForwarderParam model.DnsForwarder) (model.DnsForwarder, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingServicesDNSUpdateDnsForwarderInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	sv.AddStructField("DnsForwarder", dnsForwarderParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsForwarder
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingServicesDNSUpdateDnsForwarderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_services_DNS", "update_dns_forwarder", inputDataValue, executionContext)
	var emptyOutput model.DnsForwarder
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingServicesDNSUpdateDnsForwarderOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsForwarder), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

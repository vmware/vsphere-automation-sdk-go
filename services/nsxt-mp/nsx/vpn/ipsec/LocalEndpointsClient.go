// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LocalEndpoints
// Used by client-side stubs.

package ipsec

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type LocalEndpointsClient interface {

	// Create custom IPSec local endpoint.
	//
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error)

	// Delete custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error

	// Get custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipsecVpnLocalEndpointIdParam string) (model.IPSecVPNLocalEndpoint, error)

	// Get paginated list of all local endpoints.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipsecVpnServiceIdParam Id of the IPSec VPN service (optional)
	// @param logicalRouterIdParam Id of logical router (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNLocalEndpointListResult, error)

	// Edit custom IPSec local endpoint.
	//
	// @param ipsecVpnLocalEndpointIdParam (required)
	// @param ipSecVPNLocalEndpointParam (required)
	// @return com.vmware.nsx.model.IPSecVPNLocalEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error)
}

type localEndpointsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLocalEndpointsClient(connector client.Connector) *localEndpointsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.vpn.ipsec.local_endpoints")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := localEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *localEndpointsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *localEndpointsClient) Create(ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(localEndpointsCreateInputType(), typeConverter)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := localEndpointsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "create", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), localEndpointsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) Delete(ipsecVpnLocalEndpointIdParam string, forceParam *bool) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(localEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := localEndpointsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *localEndpointsClient) Get(ipsecVpnLocalEndpointIdParam string) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(localEndpointsGetInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := localEndpointsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), localEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) List(cursorParam *string, includedFieldsParam *string, ipsecVpnServiceIdParam *string, logicalRouterIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPSecVPNLocalEndpointListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(localEndpointsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpsecVpnServiceId", ipsecVpnServiceIdParam)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := localEndpointsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), localEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *localEndpointsClient) Update(ipsecVpnLocalEndpointIdParam string, ipSecVPNLocalEndpointParam model.IPSecVPNLocalEndpoint) (model.IPSecVPNLocalEndpoint, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(localEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("IpsecVpnLocalEndpointId", ipsecVpnLocalEndpointIdParam)
	sv.AddStructField("IpSecVPNLocalEndpoint", ipSecVPNLocalEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IPSecVPNLocalEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := localEndpointsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.local_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput model.IPSecVPNLocalEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), localEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IPSecVPNLocalEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Leases
// Used by client-side stubs.

package servers

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type LeasesClient interface {

	// Delete a single DHCP lease entry specified by ip and mac. The DHCP server matches the DHCP lease with the given ip address and the mac address. The matched lease entry will be deleted. If no lease matches, the request is ignored. The DHCP lease to be deleted will be removed by the system from both active and standby node. The system will report error if the DHCP lease could not be removed from both nodes. If the DHCP lease could not be removed on either node, please check the DHCP server status. Once the DHCP server status is UP, please invoke the deletion API again to ensure the lease gets deleted from both nodes.
	//
	// @param serverIdParam (required)
	// @param ipParam IPv4 or IPv6 address (required)
	// @param macParam MAC Address (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serverIdParam string, ipParam string, macParam string) error

	// Get specific leases of a given dhcp server. As a dhcp server could manage millions of leases, the API has to limit the number of the returned leases via two mutually-excluded request parameters, i.e. \"pool_id\" and \"address\". Either a \"pool_id\" or an \"address\" can be provided, but not both in a same call. If a \"pool_id\" is specified, the leases of the specific pool are returned. If an \"address\" is specified, only the lease(s) represented y this address is(are) returned. The \"address\" can be a single IP, an ip-range, or a mac address.
	//
	// @param serverIdParam (required)
	// @param addressParam can be an ip address, or an ip range, or a mac address (optional)
	// @param poolIdParam The uuid of dhcp ip pool (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx.model.DhcpLeases
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serverIdParam string, addressParam *string, poolIdParam *string, sourceParam *string) (model.DhcpLeases, error)
}

type leasesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLeasesClient(connector client.Connector) *leasesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.dhcp.servers.leases")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := leasesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *leasesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *leasesClient) Delete(serverIdParam string, ipParam string, macParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(leasesDeleteInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("Ip", ipParam)
	sv.AddStructField("Mac", macParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := leasesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.leases", "delete", inputDataValue, executionContext)
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

func (lIface *leasesClient) List(serverIdParam string, addressParam *string, poolIdParam *string, sourceParam *string) (model.DhcpLeases, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(leasesListInputType(), typeConverter)
	sv.AddStructField("ServerId", serverIdParam)
	sv.AddStructField("Address", addressParam)
	sv.AddStructField("PoolId", poolIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DhcpLeases
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := leasesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dhcp.servers.leases", "list", inputDataValue, executionContext)
	var emptyOutput model.DhcpLeases
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), leasesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DhcpLeases), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

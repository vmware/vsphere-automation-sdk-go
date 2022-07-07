// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FailedQueries
// Used by client-side stubs.

package forwarders

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type FailedQueriesClient interface {

	// Return the given count of recent failed DNS queries from DNS forwarder. Since the DNS forwarder is running in Acitve/Standby HA mode on transport nodes, the given count of queries will be returned from each nodes. Hence the total queries returned could be doubled. If no count is specified, 100 recent failed queries are returned. If the recent failures is less than the given count, all the failures will be returned. The maximum count is 1,000.
	//
	// @param forwarderIdParam (required)
	// @param countParam The count of the failed DNS queries (optional, default to 100)
	// @return com.vmware.nsx.model.DnsFailedQueries
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(forwarderIdParam string, countParam *int64) (model.DnsFailedQueries, error)
}

type failedQueriesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewFailedQueriesClient(connector client.Connector) *failedQueriesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.dns.forwarders.failed_queries")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	fIface := failedQueriesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *failedQueriesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *failedQueriesClient) Get(forwarderIdParam string, countParam *int64) (model.DnsFailedQueries, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(failedQueriesGetInputType(), typeConverter)
	sv.AddStructField("ForwarderId", forwarderIdParam)
	sv.AddStructField("Count", countParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DnsFailedQueries
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := failedQueriesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.dns.forwarders.failed_queries", "get", inputDataValue, executionContext)
	var emptyOutput model.DnsFailedQueries
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), failedQueriesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DnsFailedQueries), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

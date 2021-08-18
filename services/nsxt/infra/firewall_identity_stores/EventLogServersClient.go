// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EventLogServers
// Used by client-side stubs.

package firewall_identity_stores

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type EventLogServersClient interface {

	// The API tests a Event Log server connection for an already configured domain. If the connection is successful, the response will be HTTP status 200. Otherwise the response will be HTTP status 500 and corresponding error message will be returned.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param actionParam event log server test requested (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(firewallIdentityStoreIdParam string, eventLogServerIdParam string, actionParam string, enforcementPointPathParam *string) error

	// Delete a Event Log server for Firewall Identity store
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error

	// Get a specific Event Log server for a given Firewall Identity store
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.DirectoryEventLogServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (model.DirectoryEventLogServer, error)

	// Get a specific Event Log server for a given Firewall Identity store
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.DirectoryEventLogServerListResults
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(firewallIdentityStoreIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DirectoryEventLogServerListResults, error)

	// More than one Event Log server can be created and only one event log server is used to synchronize directory objects. If more than one Event Log server is configured, NSX will try all the servers until it is able to successfully connect to one.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param directoryEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam model.DirectoryEventLogServer, enforcementPointPathParam *string) error

	// Update a event log server for Firewall Identity store
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log Server identifier (required)
	// @param directoryEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.DirectoryEventLogServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam model.DirectoryEventLogServer, enforcementPointPathParam *string) (model.DirectoryEventLogServer, error)
}

type eventLogServersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewEventLogServersClient(connector client.Connector) *eventLogServersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	eIface := eventLogServersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *eventLogServersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *eventLogServersClient) Create(firewallIdentityStoreIdParam string, eventLogServerIdParam string, actionParam string, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersCreateInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Delete(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersDeleteInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Get(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (model.DirectoryEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersGetInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectoryEventLogServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "get", inputDataValue, executionContext)
	var emptyOutput model.DirectoryEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), eventLogServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectoryEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *eventLogServersClient) List(firewallIdentityStoreIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DirectoryEventLogServerListResults, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersListInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectoryEventLogServerListResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "list", inputDataValue, executionContext)
	var emptyOutput model.DirectoryEventLogServerListResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), eventLogServersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectoryEventLogServerListResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *eventLogServersClient) Patch(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam model.DirectoryEventLogServer, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersPatchInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("DirectoryEventLogServer", directoryEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Update(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam model.DirectoryEventLogServer, enforcementPointPathParam *string) (model.DirectoryEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersUpdateInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("DirectoryEventLogServer", directoryEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectoryEventLogServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "update", inputDataValue, executionContext)
	var emptyOutput model.DirectoryEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), eventLogServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectoryEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VirtualServers
// Used by client-side stubs.

package loadbalancer

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type VirtualServersClient interface {

	// Create a load balancer virtual server.
	//
	// @param lbVirtualServerParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error)

	// It is used to create virtual servers, the associated rules and bind the rules to the virtual server. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. If the virtual server need to consume some existed rules without change, those rules should not be specified in this array, otherwise, the rules are updated.
	//
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.nsx.model.LbVirtualServerWithRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error)

	// Delete a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @param deleteAssociatedRulesParam Delete associated rules (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error

	// Retrieve a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(virtualServerIdParam string) (model.LbVirtualServer, error)

	// Retrieve a paginated list of load balancer virtual servers.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.LbVirtualServerListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbVirtualServerListResult, error)

	// Update a load balancer virtual server.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(virtualServerIdParam string, lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error)

	// It is used to update virtual servers, the associated rules and update the binding of virtual server and rules. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. To delete old rules, the rules should not be configured in new action, the UUID of deleted rules should be also removed from rule_ids. To update rules, the rules should be specified with new change and configured with identifier. If there are some rules which are not modified, those rule should not be specified in the rules list, the UUID list of rules should be specified in rule_ids of LbVirtualServer.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.nsx.model.LbVirtualServerWithRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(virtualServerIdParam string, lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error)
}

type virtualServersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewVirtualServersClient(connector client.Connector) *virtualServersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.loadbalancer.virtual_servers")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":          core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": core.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":          core.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatewithrules": core.NewMethodIdentifier(interfaceIdentifier, "updatewithrules"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	vIface := virtualServersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *virtualServersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *virtualServersClient) Create(lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersCreateInputType(), typeConverter)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "create", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Createwithrules(lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerWithRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersCreatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "createwithrules", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Delete(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersDeleteInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("DeleteAssociatedRules", deleteAssociatedRulesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *virtualServersClient) Get(virtualServerIdParam string) (model.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersGetInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "get", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LbVirtualServerListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "list", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Update(virtualServerIdParam string, lbVirtualServerParam model.LbVirtualServer) (model.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersUpdateInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "update", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Updatewithrules(virtualServerIdParam string, lbVirtualServerWithRuleParam model.LbVirtualServerWithRule) (model.LbVirtualServerWithRule, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualServersUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LbVirtualServerWithRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualServersUpdatewithrulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput model.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualServersUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

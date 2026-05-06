// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VirtualServers
// Used by client-side stubs.

package loadbalancer

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VirtualServersClient interface {

	// Create a load balancer virtual server.
	//
	//  Use the following Policy API -
	//  PATCH /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param lbVirtualServerParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(lbVirtualServerParam nsxModel.LbVirtualServer) (nsxModel.LbVirtualServer, error)

	// It is used to create virtual servers, the associated rules and bind the rules to the virtual server. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. If the virtual server need to consume some existed rules without change, those rules should not be specified in this array, otherwise, the rules are updated.
	//
	//  Use the following Policy API -
	//  PATCH /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.nsx.model.LbVirtualServerWithRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createwithrules(lbVirtualServerWithRuleParam nsxModel.LbVirtualServerWithRule) (nsxModel.LbVirtualServerWithRule, error)

	// Delete a load balancer virtual server.
	//
	//  Use the following Policy API -
	//  DELETE /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param virtualServerIdParam (required)
	// @param deleteAssociatedRulesParam If this is set to true, the associated rules are also deleted when virtual server is deleted if the rules are not used by other virtual servers. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error

	// Retrieve a load balancer virtual server.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param virtualServerIdParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(virtualServerIdParam string) (nsxModel.LbVirtualServer, error)

	// Retrieve a paginated list of load balancer virtual servers.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/lb-virtual-servers
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.LbVirtualServerListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.LbVirtualServerListResult, error)

	// Update a load balancer virtual server.
	//
	//  Use the following Policy API -
	//  PUT/PATCH /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerParam (required)
	// @return com.vmware.nsx.model.LbVirtualServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(virtualServerIdParam string, lbVirtualServerParam nsxModel.LbVirtualServer) (nsxModel.LbVirtualServer, error)

	// It is used to update virtual servers, the associated rules and update the binding of virtual server and rules. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. To delete old rules, the rules should not be configured in new action, the UUID of deleted rules should be also removed from rule_ids. To update rules, the rules should be specified with new change and configured with identifier. If there are some rules which are not modified, those rule should not be specified in the rules list, the UUID list of rules should be specified in rule_ids of LbVirtualServer.
	//
	//  Use the following Policy API -
	//  PUT/PATCH /policy/api/v1/infra/lb-virtual-servers/<lb-virtual-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param virtualServerIdParam (required)
	// @param lbVirtualServerWithRuleParam (required)
	// @return com.vmware.nsx.model.LbVirtualServerWithRule
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatewithrules(virtualServerIdParam string, lbVirtualServerWithRuleParam nsxModel.LbVirtualServerWithRule) (nsxModel.LbVirtualServerWithRule, error)
}

type virtualServersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVirtualServersClient(connector vapiProtocolClient_.Connector) *virtualServersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.loadbalancer.virtual_servers")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"createwithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "createwithrules"),
		"delete":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatewithrules": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "updatewithrules"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := virtualServersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *virtualServersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *virtualServersClient) Create(lbVirtualServerParam nsxModel.LbVirtualServer) (nsxModel.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersCreateInputType(), typeConverter)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Createwithrules(lbVirtualServerWithRuleParam nsxModel.LbVirtualServerWithRule) (nsxModel.LbVirtualServerWithRule, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersCreatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersCreatewithrulesInputType(), typeConverter)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServerWithRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "createwithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersCreatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Delete(virtualServerIdParam string, deleteAssociatedRulesParam *bool) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersDeleteInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("DeleteAssociatedRules", deleteAssociatedRulesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *virtualServersClient) Get(virtualServerIdParam string) (nsxModel.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersGetInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.LbVirtualServerListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServerListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Update(virtualServerIdParam string, lbVirtualServerParam nsxModel.LbVirtualServer) (nsxModel.LbVirtualServer, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersUpdateInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServer", lbVirtualServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualServersClient) Updatewithrules(virtualServerIdParam string, lbVirtualServerWithRuleParam nsxModel.LbVirtualServerWithRule) (nsxModel.LbVirtualServerWithRule, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualServersUpdatewithrulesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualServersUpdatewithrulesInputType(), typeConverter)
	sv.AddStructField("VirtualServerId", virtualServerIdParam)
	sv.AddStructField("LbVirtualServerWithRule", lbVirtualServerWithRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LbVirtualServerWithRule
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.loadbalancer.virtual_servers", "updatewithrules", inputDataValue, executionContext)
	var emptyOutput nsxModel.LbVirtualServerWithRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualServersUpdatewithrulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LbVirtualServerWithRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

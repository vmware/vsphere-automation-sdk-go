// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesDNS
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

type SystemAdministrationConfigurationFabricNodesDNSClient interface {

	// Returns the list of servers that the node uses to look up IP addresses associated with given domain names.
	// @return com.vmware.model.NodeNameServersProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNodeNameServers() (model.NodeNameServersProperties, error)

	// Returns the domain list that the node uses to complete unqualified host names. When a host name does not include a fully qualified domain name (FQDN), the NSX Management node appends the first-listed domain name to the host name before the host name is looked up. The NSX Management node continues this for each entry in the domain list until it finds a match.
	// @return com.vmware.model.NodeSearchDomainsProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNodeSearchDomains() (model.NodeSearchDomainsProperties, error)

	// Modifies the list of servers that the node uses to look up IP addresses associated with given domain names. If DHCP is configured, this method returns a 409 CONFLICT error, because DHCP manages the list of name servers.
	//
	// @param nodeNameServersPropertiesParam (required)
	// @return com.vmware.model.NodeNameServersProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNodeNameServers(nodeNameServersPropertiesParam model.NodeNameServersProperties) (model.NodeNameServersProperties, error)

	// Modifies the list of domain names that the node uses to complete unqualified host names. If DHCP is configured, this method returns a 409 CONFLICT error, because DHCP manages the list of name servers.
	//
	// @param nodeSearchDomainsPropertiesParam (required)
	// @return com.vmware.model.NodeSearchDomainsProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNodeSearchDomains(nodeSearchDomainsPropertiesParam model.NodeSearchDomainsProperties) (model.NodeSearchDomainsProperties, error)
}

type systemAdministrationConfigurationFabricNodesDNSClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesDNSClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesDNSClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_DNS")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"read_node_name_servers":     core.NewMethodIdentifier(interfaceIdentifier, "read_node_name_servers"),
		"read_node_search_domains":   core.NewMethodIdentifier(interfaceIdentifier, "read_node_search_domains"),
		"update_node_name_servers":   core.NewMethodIdentifier(interfaceIdentifier, "update_node_name_servers"),
		"update_node_search_domains": core.NewMethodIdentifier(interfaceIdentifier, "update_node_search_domains"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesDNSClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesDNSClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesDNSClient) ReadNodeNameServers() (model.NodeNameServersProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDNSReadNodeNameServersInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeNameServersProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDNSReadNodeNameServersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_DNS", "read_node_name_servers", inputDataValue, executionContext)
	var emptyOutput model.NodeNameServersProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDNSReadNodeNameServersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeNameServersProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDNSClient) ReadNodeSearchDomains() (model.NodeSearchDomainsProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDNSReadNodeSearchDomainsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeSearchDomainsProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDNSReadNodeSearchDomainsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_DNS", "read_node_search_domains", inputDataValue, executionContext)
	var emptyOutput model.NodeSearchDomainsProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDNSReadNodeSearchDomainsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeSearchDomainsProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDNSClient) UpdateNodeNameServers(nodeNameServersPropertiesParam model.NodeNameServersProperties) (model.NodeNameServersProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDNSUpdateNodeNameServersInputType(), typeConverter)
	sv.AddStructField("NodeNameServersProperties", nodeNameServersPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeNameServersProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDNSUpdateNodeNameServersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_DNS", "update_node_name_servers", inputDataValue, executionContext)
	var emptyOutput model.NodeNameServersProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDNSUpdateNodeNameServersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeNameServersProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDNSClient) UpdateNodeSearchDomains(nodeSearchDomainsPropertiesParam model.NodeSearchDomainsProperties) (model.NodeSearchDomainsProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDNSUpdateNodeSearchDomainsInputType(), typeConverter)
	sv.AddStructField("NodeSearchDomainsProperties", nodeSearchDomainsPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeSearchDomainsProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDNSUpdateNodeSearchDomainsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_DNS", "update_node_search_domains", inputDataValue, executionContext)
	var emptyOutput model.NodeSearchDomainsProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDNSUpdateNodeSearchDomainsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeSearchDomainsProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

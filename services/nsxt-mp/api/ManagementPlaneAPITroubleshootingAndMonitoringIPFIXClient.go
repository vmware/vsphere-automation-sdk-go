// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPITroubleshootingAndMonitoringIPFIX
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPITroubleshootingAndMonitoringIPFIXClient interface {

	// Create a new IPFIX collector configuration
	//
	// @param ipfixCollectorConfigParam (required)
	// @return com.vmware.model.IpfixCollectorConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIpfixCollectorConfig(ipfixCollectorConfigParam model.IpfixCollectorConfig) (model.IpfixCollectorConfig, error)

	// Create a new IPFIX collector profile with essential properties.
	//
	// @param ipfixCollectorUpmProfileParam (required)
	// @return com.vmware.model.IpfixCollectorUpmProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIpfixCollectorUpmProfile(ipfixCollectorUpmProfileParam model.IpfixCollectorUpmProfile) (model.IpfixCollectorUpmProfile, error)

	// Create a new IPFIX configuration
	//
	// @param ipfixConfigParam (required)
	// The parameter must contain all the properties defined in model.IpfixConfig.
	// @return com.vmware.model.IpfixConfig
	// The return value will contain all the properties defined in model.IpfixConfig.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIpfixConfig(ipfixConfigParam *data.StructValue) (*data.StructValue, error)

	// Create a new IPFIX profile with essential properties.
	//
	// @param ipfixUpmProfileParam (required)
	// The parameter must contain all the properties defined in model.IpfixUpmProfile.
	// @return com.vmware.model.IpfixUpmProfile
	// The return value will contain all the properties defined in model.IpfixUpmProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateIpfixUpmProfile(ipfixUpmProfileParam *data.StructValue) (*data.StructValue, error)

	// Delete an existing IPFIX collector configuration
	//
	// @param collectorConfigIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIpfixCollectorConfig(collectorConfigIdParam string) error

	// Delete an existing IPFIX collector profile by ID.
	//
	// @param ipfixCollectorProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string) error

	// Delete an existing IPFIX configuration
	//
	// @param configIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIpfixConfig(configIdParam string) error

	// Delete an existing IPFIX profile by ID.
	//
	// @param ipfixProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteIpfixUpmProfile(ipfixProfileIdParam string) error

	// Get an existing IPFIX collector configuration
	//
	// @param collectorConfigIdParam (required)
	// @return com.vmware.model.IpfixCollectorConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIpfixCollectorConfig(collectorConfigIdParam string) (model.IpfixCollectorConfig, error)

	// Get an existing IPFIX collector profile by profile ID.
	//
	// @param ipfixCollectorProfileIdParam (required)
	// @return com.vmware.model.IpfixCollectorUpmProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string) (model.IpfixCollectorUpmProfile, error)

	// Get an existing IPFIX configuration
	//
	// @param configIdParam (required)
	// @return com.vmware.model.IpfixConfig
	// The return value will contain all the properties defined in model.IpfixConfig.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIpfixConfig(configIdParam string) (*data.StructValue, error)

	// Deprecated - Please use /ipfix-profiles for switch IPFIX profile and /ipfix-collector-profiles for IPFIX collector profile.
	// @return com.vmware.model.IpfixObsPointsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIpfixObsPoints() (model.IpfixObsPointsListResult, error)

	// Get an existing IPFIX profile by profile ID.
	//
	// @param ipfixProfileIdParam (required)
	// @return com.vmware.model.IpfixUpmProfile
	// The return value will contain all the properties defined in model.IpfixUpmProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIpfixUpmProfile(ipfixProfileIdParam string) (*data.StructValue, error)

	//
	// @return com.vmware.model.IpfixObsPointConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSwitchIpfixConfig() (model.IpfixObsPointConfig, error)

	// List IPFIX collector configurations
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IpfixCollectorConfigListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIpfixCollectorConfig(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpfixCollectorConfigListResult, error)

	// Query IPFIX collector profiles with list parameters. List result can be filtered by profile type defined by IpfixCollectorUpmProfileType.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param profileTypesParam IPFIX Collector Profile Type List (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IpfixCollectorUpmProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIpfixCollectorUpmProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypesParam *string, sortAscendingParam *bool, sortByParam *string) (model.IpfixCollectorUpmProfileListResult, error)

	// List IPFIX configuration
	//
	// @param appliedToParam Applied To (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipfixConfigTypeParam Supported IPFIX Config Types. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IpfixConfigListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIpfixConfig(appliedToParam *string, cursorParam *string, includedFieldsParam *string, ipfixConfigTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpfixConfigListResult, error)

	// Query IPFIX profiles with list parameters. List result can be filtered by profile type defined by IpfixUpmProfileType.
	//
	// @param appliedToEntityIdParam ID of Entity Applied with Profile (optional)
	// @param appliedToEntityTypeParam Supported Entity Types (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param profileTypesParam IPFIX Profile Type List (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IpfixUpmProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListIpfixUpmProfiles(appliedToEntityIdParam *string, appliedToEntityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypesParam *string, sortAscendingParam *bool, sortByParam *string) (model.IpfixUpmProfileListResult, error)

	// Update an existing IPFIX collector configuration
	//
	// @param collectorConfigIdParam (required)
	// @param ipfixCollectorConfigParam (required)
	// @return com.vmware.model.IpfixCollectorConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIpfixCollectorConfig(collectorConfigIdParam string, ipfixCollectorConfigParam model.IpfixCollectorConfig) (model.IpfixCollectorConfig, error)

	// Update an existing IPFIX collector profile with profile ID and modified properties.
	//
	// @param ipfixCollectorProfileIdParam (required)
	// @param ipfixCollectorUpmProfileParam (required)
	// @return com.vmware.model.IpfixCollectorUpmProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string, ipfixCollectorUpmProfileParam model.IpfixCollectorUpmProfile) (model.IpfixCollectorUpmProfile, error)

	// Update an existing IPFIX configuration
	//
	// @param configIdParam (required)
	// @param ipfixConfigParam (required)
	// The parameter must contain all the properties defined in model.IpfixConfig.
	// @return com.vmware.model.IpfixConfig
	// The return value will contain all the properties defined in model.IpfixConfig.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIpfixConfig(configIdParam string, ipfixConfigParam *data.StructValue) (*data.StructValue, error)

	// Update an existing IPFIX profile with profile ID and modified properties.
	//
	// @param ipfixProfileIdParam (required)
	// @param ipfixUpmProfileParam (required)
	// The parameter must contain all the properties defined in model.IpfixUpmProfile.
	// @return com.vmware.model.IpfixUpmProfile
	// The return value will contain all the properties defined in model.IpfixUpmProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIpfixUpmProfile(ipfixProfileIdParam string, ipfixUpmProfileParam *data.StructValue) (*data.StructValue, error)

	//
	//
	// @param ipfixObsPointConfigParam (required)
	// @return com.vmware.model.IpfixObsPointConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateSwitchIpfixConfig(ipfixObsPointConfigParam model.IpfixObsPointConfig) (model.IpfixObsPointConfig, error)
}

type managementPlaneAPITroubleshootingAndMonitoringIPFIXClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPITroubleshootingAndMonitoringIPFIXClient(connector client.Connector) *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_ipfix_collector_config":      core.NewMethodIdentifier(interfaceIdentifier, "create_ipfix_collector_config"),
		"create_ipfix_collector_upm_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_ipfix_collector_upm_profile"),
		"create_ipfix_config":                core.NewMethodIdentifier(interfaceIdentifier, "create_ipfix_config"),
		"create_ipfix_upm_profile":           core.NewMethodIdentifier(interfaceIdentifier, "create_ipfix_upm_profile"),
		"delete_ipfix_collector_config":      core.NewMethodIdentifier(interfaceIdentifier, "delete_ipfix_collector_config"),
		"delete_ipfix_collector_upm_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_ipfix_collector_upm_profile"),
		"delete_ipfix_config":                core.NewMethodIdentifier(interfaceIdentifier, "delete_ipfix_config"),
		"delete_ipfix_upm_profile":           core.NewMethodIdentifier(interfaceIdentifier, "delete_ipfix_upm_profile"),
		"get_ipfix_collector_config":         core.NewMethodIdentifier(interfaceIdentifier, "get_ipfix_collector_config"),
		"get_ipfix_collector_upm_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_ipfix_collector_upm_profile"),
		"get_ipfix_config":                   core.NewMethodIdentifier(interfaceIdentifier, "get_ipfix_config"),
		"get_ipfix_obs_points":               core.NewMethodIdentifier(interfaceIdentifier, "get_ipfix_obs_points"),
		"get_ipfix_upm_profile":              core.NewMethodIdentifier(interfaceIdentifier, "get_ipfix_upm_profile"),
		"get_switch_ipfix_config":            core.NewMethodIdentifier(interfaceIdentifier, "get_switch_ipfix_config"),
		"list_ipfix_collector_config":        core.NewMethodIdentifier(interfaceIdentifier, "list_ipfix_collector_config"),
		"list_ipfix_collector_upm_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_ipfix_collector_upm_profiles"),
		"list_ipfix_config":                  core.NewMethodIdentifier(interfaceIdentifier, "list_ipfix_config"),
		"list_ipfix_upm_profiles":            core.NewMethodIdentifier(interfaceIdentifier, "list_ipfix_upm_profiles"),
		"update_ipfix_collector_config":      core.NewMethodIdentifier(interfaceIdentifier, "update_ipfix_collector_config"),
		"update_ipfix_collector_upm_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_ipfix_collector_upm_profile"),
		"update_ipfix_config":                core.NewMethodIdentifier(interfaceIdentifier, "update_ipfix_config"),
		"update_ipfix_upm_profile":           core.NewMethodIdentifier(interfaceIdentifier, "update_ipfix_upm_profile"),
		"update_switch_ipfix_config":         core.NewMethodIdentifier(interfaceIdentifier, "update_switch_ipfix_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPITroubleshootingAndMonitoringIPFIXClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) CreateIpfixCollectorConfig(ipfixCollectorConfigParam model.IpfixCollectorConfig) (model.IpfixCollectorConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigInputType(), typeConverter)
	sv.AddStructField("IpfixCollectorConfig", ipfixCollectorConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "create_ipfix_collector_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) CreateIpfixCollectorUpmProfile(ipfixCollectorUpmProfileParam model.IpfixCollectorUpmProfile) (model.IpfixCollectorUpmProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixCollectorUpmProfile", ipfixCollectorUpmProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorUpmProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "create_ipfix_collector_upm_profile", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorUpmProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixCollectorUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorUpmProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) CreateIpfixConfig(ipfixConfigParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigInputType(), typeConverter)
	sv.AddStructField("IpfixConfig", ipfixConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "create_ipfix_config", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) CreateIpfixUpmProfile(ipfixUpmProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixUpmProfile", ipfixUpmProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "create_ipfix_upm_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXCreateIpfixUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) DeleteIpfixCollectorConfig(collectorConfigIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorConfigInputType(), typeConverter)
	sv.AddStructField("CollectorConfigId", collectorConfigIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "delete_ipfix_collector_config", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) DeleteIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixCollectorProfileId", ipfixCollectorProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixCollectorUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "delete_ipfix_collector_upm_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) DeleteIpfixConfig(configIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixConfigInputType(), typeConverter)
	sv.AddStructField("ConfigId", configIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "delete_ipfix_config", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) DeleteIpfixUpmProfile(ipfixProfileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixProfileId", ipfixProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXDeleteIpfixUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "delete_ipfix_upm_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetIpfixCollectorConfig(collectorConfigIdParam string) (model.IpfixCollectorConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigInputType(), typeConverter)
	sv.AddStructField("CollectorConfigId", collectorConfigIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_ipfix_collector_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string) (model.IpfixCollectorUpmProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixCollectorProfileId", ipfixCollectorProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorUpmProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_ipfix_collector_upm_profile", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorUpmProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixCollectorUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorUpmProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetIpfixConfig(configIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigInputType(), typeConverter)
	sv.AddStructField("ConfigId", configIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_ipfix_config", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetIpfixObsPoints() (model.IpfixObsPointsListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixObsPointsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_ipfix_obs_points", inputDataValue, executionContext)
	var emptyOutput model.IpfixObsPointsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixObsPointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixObsPointsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetIpfixUpmProfile(ipfixProfileIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixProfileId", ipfixProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_ipfix_upm_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetIpfixUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) GetSwitchIpfixConfig() (model.IpfixObsPointConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixObsPointConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "get_switch_ipfix_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixObsPointConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXGetSwitchIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixObsPointConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) ListIpfixCollectorConfig(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpfixCollectorConfigListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorConfigListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "list_ipfix_collector_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) ListIpfixCollectorUpmProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypesParam *string, sortAscendingParam *bool, sortByParam *string) (model.IpfixCollectorUpmProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ProfileTypes", profileTypesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorUpmProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "list_ipfix_collector_upm_profiles", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorUpmProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixCollectorUpmProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorUpmProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) ListIpfixConfig(appliedToParam *string, cursorParam *string, includedFieldsParam *string, ipfixConfigTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpfixConfigListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigInputType(), typeConverter)
	sv.AddStructField("AppliedTo", appliedToParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpfixConfigType", ipfixConfigTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixConfigListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "list_ipfix_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) ListIpfixUpmProfiles(appliedToEntityIdParam *string, appliedToEntityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypesParam *string, sortAscendingParam *bool, sortByParam *string) (model.IpfixUpmProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesInputType(), typeConverter)
	sv.AddStructField("AppliedToEntityId", appliedToEntityIdParam)
	sv.AddStructField("AppliedToEntityType", appliedToEntityTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ProfileTypes", profileTypesParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixUpmProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "list_ipfix_upm_profiles", inputDataValue, executionContext)
	var emptyOutput model.IpfixUpmProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXListIpfixUpmProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixUpmProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) UpdateIpfixCollectorConfig(collectorConfigIdParam string, ipfixCollectorConfigParam model.IpfixCollectorConfig) (model.IpfixCollectorConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigInputType(), typeConverter)
	sv.AddStructField("CollectorConfigId", collectorConfigIdParam)
	sv.AddStructField("IpfixCollectorConfig", ipfixCollectorConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "update_ipfix_collector_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) UpdateIpfixCollectorUpmProfile(ipfixCollectorProfileIdParam string, ipfixCollectorUpmProfileParam model.IpfixCollectorUpmProfile) (model.IpfixCollectorUpmProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixCollectorProfileId", ipfixCollectorProfileIdParam)
	sv.AddStructField("IpfixCollectorUpmProfile", ipfixCollectorUpmProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixCollectorUpmProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "update_ipfix_collector_upm_profile", inputDataValue, executionContext)
	var emptyOutput model.IpfixCollectorUpmProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixCollectorUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixCollectorUpmProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) UpdateIpfixConfig(configIdParam string, ipfixConfigParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigInputType(), typeConverter)
	sv.AddStructField("ConfigId", configIdParam)
	sv.AddStructField("IpfixConfig", ipfixConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "update_ipfix_config", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) UpdateIpfixUpmProfile(ipfixProfileIdParam string, ipfixUpmProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileInputType(), typeConverter)
	sv.AddStructField("IpfixProfileId", ipfixProfileIdParam)
	sv.AddStructField("IpfixUpmProfile", ipfixUpmProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "update_ipfix_upm_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateIpfixUpmProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPITroubleshootingAndMonitoringIPFIXClient) UpdateSwitchIpfixConfig(ipfixObsPointConfigParam model.IpfixObsPointConfig) (model.IpfixObsPointConfig, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigInputType(), typeConverter)
	sv.AddStructField("IpfixObsPointConfig", ipfixObsPointConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpfixObsPointConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_troubleshooting_and_monitoring_IPFIX", "update_switch_ipfix_config", inputDataValue, executionContext)
	var emptyOutput model.IpfixObsPointConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPITroubleshootingAndMonitoringIPFIXUpdateSwitchIpfixConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpfixObsPointConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

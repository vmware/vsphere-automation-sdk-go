// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceVMCOnAWSExternalConnectivity
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

const _ = core.SupportedByRuntimeVersion1

type CloudServiceVMCOnAWSExternalConnectivityClient interface {

	// Attach the specified VIF (Virtual Interface) to the SDDC.
	//
	// @param vifIdParam (required)
	// @param actionParam Direct Connect VIF (Virtual Interface) action (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AttachVif(vifIdParam string, actionParam string) error

	// Create or update the associated Group connection info for reflecting the SDDC’s association state with SDDC group.
	//
	// @param sddcGroupIdParam (required)
	// @param associatedBaseGroupConnectionInfoParam (required)
	// The parameter must contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
	// @return com.vmware.model.AssociatedBaseGroupConnectionInfo
	// The return value will contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateAssociatedGroupConnectionInfo(sddcGroupIdParam string, associatedBaseGroupConnectionInfoParam *data.StructValue) (*data.StructValue, error)

	// This API is used to create DX BGP related information.
	//
	// @param directConnectBgpInfoParam (required)
	// @return com.vmware.model.DirectConnectBgpInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateDxBgp(directConnectBgpInfoParam model.DirectConnectBgpInfo) (model.DirectConnectBgpInfo, error)

	// Delete the associated Group connection info for the local SDDC before removing the SDDC from the Group.
	//
	// @param sddcGroupIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteAssociatedGroupConnectionInfo(sddcGroupIdParam string) error

	// Detach and delete a Direct Connect VIF (Virtual Interface) from the SDDC. Detach and delete are one operation in VMC provider, a Direct Connect VIF will be deleted after detached.
	//
	// @param vifIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteVifById(vifIdParam string) error

	// The API is used to get the associated Group connection info for the local SDDC. It returns the run time attachment state for the local SDDC, others infomations are from DB.
	//
	// @param sddcGroupIdParam (required)
	// @return com.vmware.model.AssociatedBaseGroupConnectionInfo
	// The return value will contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAssociatedGroupConnectionInfo(sddcGroupIdParam string) (*data.StructValue, error)

	// Get Direct Connect BGP related information, including current Autonomous System Number of the VGW attached to the VPC
	// @return com.vmware.model.DirectConnectBgpInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetDxBgpInfo() (model.DirectConnectBgpInfo, error)

	// Retrieve routes that are advertised by the SDDC to the various connectivities. This API is a live query to VMC provider.
	//
	// @param connectivityTypeParam SDDC connectivity (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ExternalSddcRoutesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListAdvertisedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error)

	// Retrieve routes that are advertised by the SDDC to the various connectivities in CSV foramt.
	//
	// @param connectivityTypeParam SDDC connectivity (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ExternalSddcRoutesListResultInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListAdvertisedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error)

	// Retrieve BGP routes that are advertised by Direct Connect from VMC provider to on-premise datacenter. This API is a live query to VMC provider.
	// @return com.vmware.model.BGPAdvertisedRoutes
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListAdvertisedRoutes() (model.BGPAdvertisedRoutes, error)

	// This API is used to list all associated Group connection infos for the local SDDC. It returns the run time attachment state for the local SDDC, others infomations are from DB.
	// @return com.vmware.model.AssociatedGroupConnectionInfosListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListAssociatedGroupConnectionInfos() (model.AssociatedGroupConnectionInfosListResult, error)

	// Return all non-connected VIFs (with states \"avalible\", \"down\", \"pending\" and \"confirming\") and connected VIFs that are available to the SDDC.
	// @return com.vmware.model.VifsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListDirectConnectVifs() (model.VifsListResult, error)

	// Retrieve routes that are learned for the SDDDC. This API is a db query.
	//
	// @param connectivityTypeParam SDDC connectivity (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ExternalSddcRoutesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLearnedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error)

	// Get TGW learned routes for the SDDDC in CSV format.
	//
	// @param connectivityTypeParam SDDC connectivity (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ExternalSddcRoutesListResultInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLearnedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error)

	// Retrieve BGP routes that are learned by Direct Connect from on-premise datacenter. This API is a live query to VMC provider.
	// @return com.vmware.model.BGPLearnedRoutes
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLearnedRoutes() (model.BGPLearnedRoutes, error)

	// This API is used to update DX BGP related information. For ASN update, VIFs should be disconnected from the DX VGW before making this API call. The ASN update operation will be synchronous at this point. In the future the user should make use of the Get RealizationStatus call to check update status. While an ASN update call is in progress, any other DX BGP update request will be rejected.
	//
	// @param directConnectBgpInfoParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateDxBgpInfo(directConnectBgpInfoParam model.DirectConnectBgpInfo) error
}

type cloudServiceVMCOnAWSExternalConnectivityClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceVMCOnAWSExternalConnectivityClient(connector client.Connector) *cloudServiceVMCOnAWSExternalConnectivityClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"attach_vif": core.NewMethodIdentifier(interfaceIdentifier, "attach_vif"),
		"create_associated_group_connection_info": core.NewMethodIdentifier(interfaceIdentifier, "create_associated_group_connection_info"),
		"create_dx_bgp": core.NewMethodIdentifier(interfaceIdentifier, "create_dx_bgp"),
		"delete_associated_group_connection_info":           core.NewMethodIdentifier(interfaceIdentifier, "delete_associated_group_connection_info"),
		"delete_vif_by_id":                                  core.NewMethodIdentifier(interfaceIdentifier, "delete_vif_by_id"),
		"get_associated_group_connection_info":              core.NewMethodIdentifier(interfaceIdentifier, "get_associated_group_connection_info"),
		"get_dx_bgp_info":                                   core.NewMethodIdentifier(interfaceIdentifier, "get_dx_bgp_info"),
		"list_advertised_external_routes":                   core.NewMethodIdentifier(interfaceIdentifier, "list_advertised_external_routes"),
		"list_advertised_external_routes_in_csv_format_csv": core.NewMethodIdentifier(interfaceIdentifier, "list_advertised_external_routes_in_csv_format_csv"),
		"list_advertised_routes":                            core.NewMethodIdentifier(interfaceIdentifier, "list_advertised_routes"),
		"list_associated_group_connection_infos":            core.NewMethodIdentifier(interfaceIdentifier, "list_associated_group_connection_infos"),
		"list_direct_connect_vifs":                          core.NewMethodIdentifier(interfaceIdentifier, "list_direct_connect_vifs"),
		"list_learned_external_routes":                      core.NewMethodIdentifier(interfaceIdentifier, "list_learned_external_routes"),
		"list_learned_external_routes_in_csv_format_csv":    core.NewMethodIdentifier(interfaceIdentifier, "list_learned_external_routes_in_csv_format_csv"),
		"list_learned_routes":                               core.NewMethodIdentifier(interfaceIdentifier, "list_learned_routes"),
		"update_dx_bgp_info":                                core.NewMethodIdentifier(interfaceIdentifier, "update_dx_bgp_info"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceVMCOnAWSExternalConnectivityClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) AttachVif(vifIdParam string, actionParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityAttachVifInputType(), typeConverter)
	sv.AddStructField("VifId", vifIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityAttachVifRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "attach_vif", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) CreateAssociatedGroupConnectionInfo(sddcGroupIdParam string, associatedBaseGroupConnectionInfoParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoInputType(), typeConverter)
	sv.AddStructField("SddcGroupId", sddcGroupIdParam)
	sv.AddStructField("AssociatedBaseGroupConnectionInfo", associatedBaseGroupConnectionInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "create_associated_group_connection_info", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) CreateDxBgp(directConnectBgpInfoParam model.DirectConnectBgpInfo) (model.DirectConnectBgpInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpInputType(), typeConverter)
	sv.AddStructField("DirectConnectBgpInfo", directConnectBgpInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectConnectBgpInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "create_dx_bgp", inputDataValue, executionContext)
	var emptyOutput model.DirectConnectBgpInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectConnectBgpInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) DeleteAssociatedGroupConnectionInfo(sddcGroupIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityDeleteAssociatedGroupConnectionInfoInputType(), typeConverter)
	sv.AddStructField("SddcGroupId", sddcGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityDeleteAssociatedGroupConnectionInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "delete_associated_group_connection_info", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) DeleteVifById(vifIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityDeleteVifByIdInputType(), typeConverter)
	sv.AddStructField("VifId", vifIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityDeleteVifByIdRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "delete_vif_by_id", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) GetAssociatedGroupConnectionInfo(sddcGroupIdParam string) (*data.StructValue, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoInputType(), typeConverter)
	sv.AddStructField("SddcGroupId", sddcGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "get_associated_group_connection_info", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) GetDxBgpInfo() (model.DirectConnectBgpInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DirectConnectBgpInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "get_dx_bgp_info", inputDataValue, executionContext)
	var emptyOutput model.DirectConnectBgpInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DirectConnectBgpInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListAdvertisedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInputType(), typeConverter)
	sv.AddStructField("ConnectivityType", connectivityTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalSddcRoutesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_advertised_external_routes", inputDataValue, executionContext)
	var emptyOutput model.ExternalSddcRoutesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalSddcRoutesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListAdvertisedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("ConnectivityType", connectivityTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalSddcRoutesListResultInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_advertised_external_routes_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.ExternalSddcRoutesListResultInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalSddcRoutesListResultInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListAdvertisedRoutes() (model.BGPAdvertisedRoutes, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BGPAdvertisedRoutes
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_advertised_routes", inputDataValue, executionContext)
	var emptyOutput model.BGPAdvertisedRoutes
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BGPAdvertisedRoutes), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListAssociatedGroupConnectionInfos() (model.AssociatedGroupConnectionInfosListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AssociatedGroupConnectionInfosListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_associated_group_connection_infos", inputDataValue, executionContext)
	var emptyOutput model.AssociatedGroupConnectionInfosListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AssociatedGroupConnectionInfosListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListDirectConnectVifs() (model.VifsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VifsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_direct_connect_vifs", inputDataValue, executionContext)
	var emptyOutput model.VifsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VifsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListLearnedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInputType(), typeConverter)
	sv.AddStructField("ConnectivityType", connectivityTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalSddcRoutesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_learned_external_routes", inputDataValue, executionContext)
	var emptyOutput model.ExternalSddcRoutesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalSddcRoutesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListLearnedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("ConnectivityType", connectivityTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalSddcRoutesListResultInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_learned_external_routes_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.ExternalSddcRoutesListResultInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalSddcRoutesListResultInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) ListLearnedRoutes() (model.BGPLearnedRoutes, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BGPLearnedRoutes
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "list_learned_routes", inputDataValue, executionContext)
	var emptyOutput model.BGPLearnedRoutes
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BGPLearnedRoutes), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSExternalConnectivityClient) UpdateDxBgpInfo(directConnectBgpInfoParam model.DirectConnectBgpInfo) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSExternalConnectivityUpdateDxBgpInfoInputType(), typeConverter)
	sv.AddStructField("DirectConnectBgpInfo", directConnectBgpInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSExternalConnectivityUpdateDxBgpInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_external_connectivity", "update_dx_bgp_info", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

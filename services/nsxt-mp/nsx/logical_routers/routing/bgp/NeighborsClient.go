// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Neighbors
// Used by client-side stubs.

package bgp

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type NeighborsClient interface {

	// Add a new BGP Neighbor on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param bgpNeighborParam (required)
	// @return com.vmware.nsx.model.BgpNeighbor
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(logicalRouterIdParam string, bgpNeighborParam model.BgpNeighbor) (model.BgpNeighbor, error)

	// Delete a specific BGP Neighbor on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param idParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(logicalRouterIdParam string, idParam string) error

	// Read a specific BGP Neighbor on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param idParam (required)
	// @return com.vmware.nsx.model.BgpNeighbor
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(logicalRouterIdParam string, idParam string) (model.BgpNeighbor, error)

	// Paginated list of BGP Neighbors on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.BgpNeighborListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BgpNeighborListResult, error)

	// Read a specific BGP Neighbor details with password on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param idParam (required)
	// @return com.vmware.nsx.model.BgpNeighbor
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Showsensitivedata(logicalRouterIdParam string, idParam string) (model.BgpNeighbor, error)

	// Unset/Delete the password property on the specific BGP Neighbor. No other property of the BgpNeighbor can be updated using this API
	//
	// @param logicalRouterIdParam (required)
	// @param idParam (required)
	// @param actionParam (optional)
	// @return com.vmware.nsx.model.BgpNeighbor
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Unsetpassword(logicalRouterIdParam string, idParam string, actionParam *string) (model.BgpNeighbor, error)

	// Update a specific BGP Neighbor on a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param idParam (required)
	// @param bgpNeighborParam (required)
	// @return com.vmware.nsx.model.BgpNeighbor
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(logicalRouterIdParam string, idParam string, bgpNeighborParam model.BgpNeighbor) (model.BgpNeighbor, error)
}

type neighborsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewNeighborsClient(connector client.Connector) *neighborsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.logical_routers.routing.bgp.neighbors")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":            core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":            core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":               core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":              core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"showsensitivedata": core.NewMethodIdentifier(interfaceIdentifier, "showsensitivedata"),
		"unsetpassword":     core.NewMethodIdentifier(interfaceIdentifier, "unsetpassword"),
		"update":            core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	nIface := neighborsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *neighborsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *neighborsClient) Create(logicalRouterIdParam string, bgpNeighborParam model.BgpNeighbor) (model.BgpNeighbor, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsCreateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("BgpNeighbor", bgpNeighborParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighbor
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "create", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighbor
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighbor), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) Delete(logicalRouterIdParam string, idParam string) error {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsDeleteInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *neighborsClient) Get(logicalRouterIdParam string, idParam string) (model.BgpNeighbor, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsGetInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighbor
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "get", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighbor
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighbor), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) List(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BgpNeighborListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsListInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "list", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) Showsensitivedata(logicalRouterIdParam string, idParam string) (model.BgpNeighbor, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsShowsensitivedataInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighbor
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsShowsensitivedataRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "showsensitivedata", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighbor
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsShowsensitivedataOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighbor), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) Unsetpassword(logicalRouterIdParam string, idParam string, actionParam *string) (model.BgpNeighbor, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsUnsetpasswordInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighbor
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsUnsetpasswordRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "unsetpassword", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighbor
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsUnsetpasswordOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighbor), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) Update(logicalRouterIdParam string, idParam string, bgpNeighborParam model.BgpNeighbor) (model.BgpNeighbor, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(neighborsUpdateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("BgpNeighbor", bgpNeighborParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighbor
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := neighborsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.neighbors", "update", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighbor
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), neighborsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighbor), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

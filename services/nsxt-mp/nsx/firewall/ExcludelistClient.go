// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Excludelist
// Used by client-side stubs.

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type ExcludelistClient interface {

	// Add a new object in the exclude list.
	//
	// @param resourceReferenceParam (required)
	// @return com.vmware.nsx.model.ResourceReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addmember(resourceReferenceParam model.ResourceReference) (model.ResourceReference, error)

	// Check if the object a member of the exclude list
	//
	// @param objectIdParam identifier of the object (required)
	// @param deepCheckParam Check all parents (optional, default to false)
	// @param objectTypeParam Object type of an entity (optional)
	// @return com.vmware.nsx.model.ResourceReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Checkifexists(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (model.ResourceReference, error)

	// Get list of entities in exclude list.
	// @return com.vmware.nsx.model.ExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.ExcludeList, error)

	// Remove an existing object from the exclude list.
	//
	// @param objectIdParam identifier of the object (required)
	// @param deepCheckParam Check all parents (optional, default to false)
	// @param objectTypeParam Object type of an entity (optional)
	// @return com.vmware.nsx.model.ResourceReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Removemember(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (model.ResourceReference, error)

	// Modify exclude list.
	//
	// @param excludeListParam (required)
	// @return com.vmware.nsx.model.ExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(excludeListParam model.ExcludeList) (model.ExcludeList, error)
}

type excludelistClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewExcludelistClient(connector client.Connector) *excludelistClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.firewall.excludelist")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"addmember":     core.NewMethodIdentifier(interfaceIdentifier, "addmember"),
		"checkifexists": core.NewMethodIdentifier(interfaceIdentifier, "checkifexists"),
		"get":           core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"removemember":  core.NewMethodIdentifier(interfaceIdentifier, "removemember"),
		"update":        core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	eIface := excludelistClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *excludelistClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *excludelistClient) Addmember(resourceReferenceParam model.ResourceReference) (model.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(excludelistAddmemberInputType(), typeConverter)
	sv.AddStructField("ResourceReference", resourceReferenceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ResourceReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := excludelistAddmemberRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "addmember", inputDataValue, executionContext)
	var emptyOutput model.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), excludelistAddmemberOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Checkifexists(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (model.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(excludelistCheckifexistsInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("DeepCheck", deepCheckParam)
	sv.AddStructField("ObjectType", objectTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ResourceReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := excludelistCheckifexistsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "checkifexists", inputDataValue, executionContext)
	var emptyOutput model.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), excludelistCheckifexistsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Get() (model.ExcludeList, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(excludelistGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := excludelistGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "get", inputDataValue, executionContext)
	var emptyOutput model.ExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), excludelistGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Removemember(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (model.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(excludelistRemovememberInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("DeepCheck", deepCheckParam)
	sv.AddStructField("ObjectType", objectTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ResourceReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := excludelistRemovememberRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "removemember", inputDataValue, executionContext)
	var emptyOutput model.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), excludelistRemovememberOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Update(excludeListParam model.ExcludeList) (model.ExcludeList, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(excludelistUpdateInputType(), typeConverter)
	sv.AddStructField("ExcludeList", excludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := excludelistUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "update", inputDataValue, executionContext)
	var emptyOutput model.ExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), excludelistUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

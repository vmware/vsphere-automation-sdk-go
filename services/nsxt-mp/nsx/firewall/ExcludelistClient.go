// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Excludelist
// Used by client-side stubs.

package firewall

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ExcludelistClient interface {

	// Add a new object in the exclude list.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/settings/firewall/security/exclude-list
	//
	// Deprecated: This API element is deprecated.
	//
	// @param resourceReferenceParam (required)
	// @return com.vmware.nsx.model.ResourceReference
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addmember(resourceReferenceParam nsxModel.ResourceReference) (nsxModel.ResourceReference, error)

	// Check if the object a member of the exclude list
	//
	// @param objectIdParam identifier of the object (required)
	// @param deepCheckParam Check all parents (optional, default to false)
	// @param objectTypeParam Object type of an entity (optional)
	// @return com.vmware.nsx.model.ResourceReference
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Checkifexists(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (nsxModel.ResourceReference, error)

	// Get list of entities in exclude list.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/settings/firewall/security/exclude-list
	//
	// Deprecated: This API element is deprecated.
	// @return com.vmware.nsx.model.ExcludeList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsxModel.ExcludeList, error)

	// Remove an existing object from the exclude list.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/settings/firewall/security/exclude-list
	//
	// Deprecated: This API element is deprecated.
	//
	// @param objectIdParam identifier of the object (required)
	// @param deepCheckParam Check all parents (optional, default to false)
	// @param objectTypeParam Object type of an entity (optional)
	// @return com.vmware.nsx.model.ResourceReference
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Removemember(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (nsxModel.ResourceReference, error)

	// Modify exclude list.
	//
	//  Use the following Policy API -
	//  PUT|PATCH /policy/api/v1/infra/settings/firewall/security/exclude-list
	//
	// Deprecated: This API element is deprecated.
	//
	// @param excludeListParam (required)
	// @return com.vmware.nsx.model.ExcludeList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(excludeListParam nsxModel.ExcludeList) (nsxModel.ExcludeList, error)
}

type excludelistClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewExcludelistClient(connector vapiProtocolClient_.Connector) *excludelistClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.firewall.excludelist")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"addmember":     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "addmember"),
		"checkifexists": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "checkifexists"),
		"get":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"removemember":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "removemember"),
		"update":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := excludelistClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *excludelistClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *excludelistClient) Addmember(resourceReferenceParam nsxModel.ResourceReference) (nsxModel.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := excludelistAddmemberRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(excludelistAddmemberInputType(), typeConverter)
	sv.AddStructField("ResourceReference", resourceReferenceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ResourceReference
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "addmember", inputDataValue, executionContext)
	var emptyOutput nsxModel.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExcludelistAddmemberOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Checkifexists(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (nsxModel.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := excludelistCheckifexistsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(excludelistCheckifexistsInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("DeepCheck", deepCheckParam)
	sv.AddStructField("ObjectType", objectTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ResourceReference
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "checkifexists", inputDataValue, executionContext)
	var emptyOutput nsxModel.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExcludelistCheckifexistsOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Get() (nsxModel.ExcludeList, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := excludelistGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(excludelistGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ExcludeList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExcludelistGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Removemember(objectIdParam string, deepCheckParam *bool, objectTypeParam *string) (nsxModel.ResourceReference, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := excludelistRemovememberRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(excludelistRemovememberInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("DeepCheck", deepCheckParam)
	sv.AddStructField("ObjectType", objectTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ResourceReference
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "removemember", inputDataValue, executionContext)
	var emptyOutput nsxModel.ResourceReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExcludelistRemovememberOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ResourceReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *excludelistClient) Update(excludeListParam nsxModel.ExcludeList) (nsxModel.ExcludeList, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := excludelistUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(excludelistUpdateInputType(), typeConverter)
	sv.AddStructField("ExcludeList", excludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ExcludeList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.excludelist", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ExcludelistUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

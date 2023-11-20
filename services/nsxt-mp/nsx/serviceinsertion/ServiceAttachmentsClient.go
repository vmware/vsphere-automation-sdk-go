// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceAttachments
// Used by client-side stubs.

package serviceinsertion

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ServiceAttachmentsClient interface {

	// Adds a new Service attachment. A service attachment represents a point on NSX entity (Example: Logical Router) to which service instance can be connected through an InstanceEndpoint.
	//  This API has been deprecated, please use below Policy API
	//  For North-South service insertion
	//  PUT /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PATCH /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PUT /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PATCH /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> For East-West service insertion
	//  PUT /policy/api/v1/infra/segments/service-segments/<service-segment-id> PATCH /policy/api/v1/infra/segments/service-segments/<service-segment-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceAttachmentParam (required)
	// @return com.vmware.nsx.model.ServiceAttachment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceAttachmentParam nsxModel.ServiceAttachment) (nsxModel.ServiceAttachment, error)

	// Delete existing service attachment from system. Before deletion, please make sure that, no instance endpoints are connected to this attachment. In turn no appliance should be connected to this attachment.
	//  This API has been deprecated, please use below Policy API
	//  For North-South service insertion
	//  DELETE /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> DELETE /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> For East-West service insertion
	//  DELETE /policy/api/v1/infra/segments/service-segments/>service-segment-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceAttachmentIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serviceAttachmentIdParam string) error

	// Returns detailed Attachment information for a given service attachment.
	//  This API has been deprecated, please use below Policy API
	//  For North-South service insertion
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> GET /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> For East-West service insertion
	//  GET /policy/api/v1/infra/segments/service-segments/<service-segment-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceAttachmentIdParam (required)
	// @return com.vmware.nsx.model.ServiceAttachment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serviceAttachmentIdParam string) (nsxModel.ServiceAttachment, error)

	// Returns all Service-Attachement(s) present in the system.
	//  This API has been deprecated, please use below Policy API
	//  For North-South service insertion
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces GET /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces For East-West service insertion
	//  GET /policy/api/v1/infra/segments/service-segments
	//
	// Deprecated: This API element is deprecated.
	// @return com.vmware.nsx.model.ServiceAttachmentListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (nsxModel.ServiceAttachmentListResult, error)

	// Modifies an existing service attachment. Updates to name, description and Logical Router list only supported.
	//  This API has been deprecated, please use below Policy API
	//  For North-South service insertion
	//  PUT /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PATCH /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PUT /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> PATCH /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-interfaces/<interface-id> For East-West service insertion
	//  PUT /policy/api/v1/infra/segments/service-segments/<service-segment-id> PATCH /policy/api/v1/infra/segments/service-segments/<service-segment-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceAttachmentIdParam (required)
	// @param serviceAttachmentParam (required)
	// @return com.vmware.nsx.model.ServiceAttachment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serviceAttachmentIdParam string, serviceAttachmentParam nsxModel.ServiceAttachment) (nsxModel.ServiceAttachment, error)
}

type serviceAttachmentsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewServiceAttachmentsClient(connector vapiProtocolClient_.Connector) *serviceAttachmentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.service_attachments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := serviceAttachmentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceAttachmentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceAttachmentsClient) Create(serviceAttachmentParam nsxModel.ServiceAttachment) (nsxModel.ServiceAttachment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceAttachmentsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceAttachmentsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceAttachment", serviceAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceAttachment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.service_attachments", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceAttachmentsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceAttachmentsClient) Delete(serviceAttachmentIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceAttachmentsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceAttachmentsDeleteInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.service_attachments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceAttachmentsClient) Get(serviceAttachmentIdParam string) (nsxModel.ServiceAttachment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceAttachmentsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceAttachmentsGetInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceAttachment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.service_attachments", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceAttachmentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceAttachmentsClient) List() (nsxModel.ServiceAttachmentListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceAttachmentsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceAttachmentsListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceAttachmentListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.service_attachments", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceAttachmentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceAttachmentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceAttachmentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceAttachmentsClient) Update(serviceAttachmentIdParam string, serviceAttachmentParam nsxModel.ServiceAttachment) (nsxModel.ServiceAttachment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceAttachmentsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceAttachmentsUpdateInputType(), typeConverter)
	sv.AddStructField("ServiceAttachmentId", serviceAttachmentIdParam)
	sv.AddStructField("ServiceAttachment", serviceAttachmentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceAttachment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.service_attachments", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceAttachment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceAttachmentsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceAttachment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

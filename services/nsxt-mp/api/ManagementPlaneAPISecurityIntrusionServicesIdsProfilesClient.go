// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPISecurityIntrusionServicesIdsProfiles
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

type ManagementPlaneAPISecurityIntrusionServicesIdsProfilesClient interface {

	// Returns information about the specified IDSProfile.
	//
	// @param idsProfileIdParam IDSProfile Id (required)
	// @return com.vmware.model.IDSProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetIDSProfile(idsProfileIdParam string) (model.IDSProfile, error)
}

type managementPlaneAPISecurityIntrusionServicesIdsProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPISecurityIntrusionServicesIdsProfilesClient(connector client.Connector) *managementPlaneAPISecurityIntrusionServicesIdsProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_security_intrusion_services_ids_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_IDS_profile": core.NewMethodIdentifier(interfaceIdentifier, "get_IDS_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPISecurityIntrusionServicesIdsProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPISecurityIntrusionServicesIdsProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPISecurityIntrusionServicesIdsProfilesClient) GetIDSProfile(idsProfileIdParam string) (model.IDSProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIntrusionServicesIdsProfilesGetIDSProfileInputType(), typeConverter)
	sv.AddStructField("IdsProfileId", idsProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IDSProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIntrusionServicesIdsProfilesGetIDSProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_intrusion_services_ids_profiles", "get_IDS_profile", inputDataValue, executionContext)
	var emptyOutput model.IDSProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIntrusionServicesIdsProfilesGetIDSProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IDSProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

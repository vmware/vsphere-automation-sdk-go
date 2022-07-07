// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UserInfo
// Used by client-side stubs.

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type UserInfoClient interface {

	// This API will return the name and role information of the user invoking this API request. This API is available for all NSX users no matter their authentication method (Local account, VIDM, LDAP etc).
	//
	// @param rootPathParam Prefix path of the context (optional)
	// @return com.vmware.nsx_global_policy.model.UserInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(rootPathParam *string) (model.UserInfo, error)
}

type userInfoClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUserInfoClient(connector client.Connector) *userInfoClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.aaa.user_info")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := userInfoClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *userInfoClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *userInfoClient) Get(rootPathParam *string) (model.UserInfo, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(userInfoGetInputType(), typeConverter)
	sv.AddStructField("RootPath", rootPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UserInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := userInfoGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.user_info", "get", inputDataValue, executionContext)
	var emptyOutput model.UserInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), userInfoGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UserInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

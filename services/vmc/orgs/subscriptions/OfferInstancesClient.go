// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: OfferInstances
// Used by client-side stubs.

package subscriptions

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type OfferInstancesClient interface {

	// List all offers available for the specific product type in the specific region
	//
	// @param orgParam Organization identifier (required)
	// @param regionParam Region for the offer (required)
	// @param productTypeParam Type of the product in offers. \*This has been deprecated\*. Please use product & type parameters (required)
	// @param productParam The product that you are trying to purchase, eg. host. This needs to be accompanied by the type parameter (optional)
	// @param type_Param The type/flavor of the product you are trying it purchase,eg. an \\\\`r5.metal\\\\` host. This needs to be accompanied by the product parameter. (optional)
	// @return com.vmware.vmc.model.OfferInstancesHolder
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request. Type of the product not supported.
	// @throws Unauthorized  Forbidden
	List(orgParam string, regionParam string, productTypeParam string, productParam *string, type_Param *string) (vmcModel.OfferInstancesHolder, error)
}

type offerInstancesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOfferInstancesClient(connector vapiProtocolClient_.Connector) *offerInstancesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.subscriptions.offer_instances")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := offerInstancesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *offerInstancesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *offerInstancesClient) List(orgParam string, regionParam string, productTypeParam string, productParam *string, type_Param *string) (vmcModel.OfferInstancesHolder, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := offerInstancesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(offerInstancesListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Region", regionParam)
	sv.AddStructField("ProductType", productTypeParam)
	sv.AddStructField("Product", productParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.OfferInstancesHolder
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.subscriptions.offer_instances", "list", inputDataValue, executionContext)
	var emptyOutput vmcModel.OfferInstancesHolder
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OfferInstancesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.OfferInstancesHolder), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

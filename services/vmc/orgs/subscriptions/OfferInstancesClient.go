// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: OfferInstances
// Used by client-side stubs.

package subscriptions

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type OfferInstancesClient interface {

	// List all offers available for the specific product type in the specific region
	//
	// @param orgParam Organization identifier (required)
	// @param regionParam Region for the offer (required)
	// @param productTypeParam Type of the product in offers. \*This has been deprecated\*. Please use product & type parameters (required)
	// @param productParam The product that you are trying to purchase, eg. host. This needs to be accompanied by the type parameter (optional)
	// @param type_Param The type/flavor of the product you are trying it purchase,eg. an \\\\`r5.metal\\\\` host. This needs to be accompanied by the product parameter. (optional)
	// @return com.vmware.vmc.model.OfferInstancesHolder
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request. Type of the product not supported.
	// @throws Unauthorized  Forbidden
	List(orgParam string, regionParam string, productTypeParam string, productParam *string, type_Param *string) (model.OfferInstancesHolder, error)
}

type offerInstancesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewOfferInstancesClient(connector client.Connector) *offerInstancesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.subscriptions.offer_instances")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	oIface := offerInstancesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *offerInstancesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *offerInstancesClient) List(orgParam string, regionParam string, productTypeParam string, productParam *string, type_Param *string) (model.OfferInstancesHolder, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(offerInstancesListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Region", regionParam)
	sv.AddStructField("ProductType", productTypeParam)
	sv.AddStructField("Product", productParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OfferInstancesHolder
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := offerInstancesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.subscriptions.offer_instances", "list", inputDataValue, executionContext)
	var emptyOutput model.OfferInstancesHolder
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), offerInstancesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OfferInstancesHolder), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Subscriptions
// Used by client-side stubs.

package orgs

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type SubscriptionsClient interface {

	// Initiates the creation of a subscription
	//
	// @param orgParam Organization identifier (required)
	// @param subscriptionRequestParam subscriptionRequest (required)
	// @return com.vmware.vmc.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
	Create(orgParam string, subscriptionRequestParam model.SubscriptionRequest) (model.Task, error)

	// Get subscription details for a given subscription id
	//
	// @param orgParam Organization identifier (required)
	// @param subscriptionParam SubscriptionId for an sddc. (required)
	// @return com.vmware.vmc.model.SubscriptionDetails
	// @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
	// @throws NotFound  Not Found
	Get(orgParam string, subscriptionParam string) (model.SubscriptionDetails, error)

	// Returns all subscriptions for a given org id
	//
	// @param orgParam Organization identifier (required)
	// @param offerTypeParam Offer Type \* \\\\`ON_DEMAND\\\\` - on-demand subscription \* \\\\`TERM\\\\` - term subscription \* All subscriptions if not specified (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws InternalServerError  Server error. Check retryable flag to see if request should be retried.
	// @throws NotFound  Not Found
	Get0(orgParam string, offerTypeParam *string) ([]model.SubscriptionDetails, error)
}

type subscriptionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSubscriptionsClient(connector client.Connector) *subscriptionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.subscriptions")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"get_0":  core.NewMethodIdentifier(interfaceIdentifier, "get_0"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := subscriptionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *subscriptionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *subscriptionsClient) Create(orgParam string, subscriptionRequestParam model.SubscriptionRequest) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(subscriptionsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SubscriptionRequest", subscriptionRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := subscriptionsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.subscriptions", "create", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), subscriptionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *subscriptionsClient) Get(orgParam string, subscriptionParam string) (model.SubscriptionDetails, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(subscriptionsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Subscription", subscriptionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SubscriptionDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := subscriptionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.subscriptions", "get", inputDataValue, executionContext)
	var emptyOutput model.SubscriptionDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), subscriptionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SubscriptionDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *subscriptionsClient) Get0(orgParam string, offerTypeParam *string) ([]model.SubscriptionDetails, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(subscriptionsGet0InputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("OfferType", offerTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.SubscriptionDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := subscriptionsGet0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.subscriptions", "get_0", inputDataValue, executionContext)
	var emptyOutput []model.SubscriptionDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), subscriptionsGet0OutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.SubscriptionDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

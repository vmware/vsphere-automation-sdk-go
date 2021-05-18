/* Copyright © 2019-2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package local

import (
	"errors"
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/provider/introspection"
)

type LocalProvider struct {
	name string
	//key service id
	serviceMap   map[string]core.ApiInterface
	introspector *introspection.LocalProviderIntrospector
	errorDefs    []data.ErrorDefinition
	errorValues  []data.DataValue
}

const LocalProviderDefaultName = "LocalProvider"

func NewLocalProvider(name string, loadIntrospection bool) (*LocalProvider, error) {
	if name == "" {
		name = LocalProviderDefaultName
	}
	var emptyMap = make(map[string]core.ApiInterface)

	//load introspection
	var introspector, err = introspection.NewLocalProviderIntrospector(name)
	if err != nil {
		return nil, err
	}
	var localProvider = &LocalProvider{name: name, serviceMap: emptyMap, introspector: introspector}

	if loadIntrospection {
		log.Info("Loading Introspection Services")
		var services = introspector.GetIntrospectionServices()
		var adderror = localProvider.AddInterfaces(services)
		if adderror != nil {
			return nil, adderror
		}
	}

	//these are the errors that should be augmented
	//when `get` operation is invoked on
	//com.vmware.vapi.std.introspection.Operation
	var errorDefs = []data.ErrorDefinition{
		bindings.INTERNAL_SERVER_ERROR_DEF,
		bindings.INVALID_ARGUMENT_ERROR_DEF,
		bindings.OP_NOT_FOUND_ERROR_DEF,
	}
	localProvider.errorDefs = errorDefs
	var errorValues = make([]data.DataValue, 0)
	for _, errorDef := range errorDefs {
		var dataVal = data.ConvertOperationDataDefinitionToDataValue(errorDef)
		errorValues = append(errorValues, dataVal)
	}
	localProvider.errorValues = errorValues
	return localProvider, nil

}

/**
 * Adds the specified API interface to the provider.
 *
 * @param iface    API interface to add; must not be <code>null</code>
 *
 * @returns error if API interface with the same ID as
 *         <code>iface</code> is already registered
 */
func (localProvider *LocalProvider) AddInterface(apiInterface core.ApiInterface) error {
	if apiInterface == nil {
		return errors.New("interface cannot be nil")
	}
	var serviceID = apiInterface.Identifier().Name()
	if _, ok := localProvider.serviceMap[serviceID]; ok {
		log.Errorf("Duplicate interface %s found. Registration failed ", serviceID)
		return fmt.Errorf("Duplicate interface %s found. Registration failed ", serviceID)
	}
	localProvider.serviceMap[serviceID] = apiInterface
	if localProvider.introspector != nil {
		localProvider.introspector.AddService(serviceID, apiInterface)
	}
	log.Info("Registered service " + serviceID)
	return nil
}

func (localProvider *LocalProvider) AddInterfaces(apiInterfaces []core.ApiInterface) error {
	if apiInterfaces == nil {
		return errors.New("interfaces cannot be nil")
	}
	for _, apiInterface := range apiInterfaces {
		var err = localProvider.AddInterface(apiInterface)
		if err != nil {
			return err
		}
	}
	return nil
}

// RegisterInterfaces adds a set of interfaces to a local provider. Using the
// DualInterface allows to handle REST enabled bindings.
func (localProvider *LocalProvider) RegisterInterfaces(apiInterfaces []DualInterface) error {
	if apiInterfaces == nil {
		return errors.New("interfaces cannot be nil")
	}
	for _, apiInterface := range apiInterfaces {
		var err = localProvider.AddInterface(apiInterface)
		if err != nil {
			return err
		}
	}
	return nil
}

func (localProvider *LocalProvider) ComputeCheckSum() string {
	return localProvider.introspector.GetCheckSum()
}

func (localProvider *LocalProvider) validateError(errorValue *data.ErrorValue, methodDefinition *core.MethodDefinition) []error {
	if methodDefinition.HasErrorDefinition(errorValue.Name()) {
		var errorDef = methodDefinition.ErrorDefinition(errorValue.Name())
		var messages = errorDef.Validate(errorValue)
		return messages
	}
	var methodID = methodDefinition.Identifier()
	args := map[string]string{
		"errorName":  errorValue.Name(),
		"methodName": methodID.Name()}
	var message = l10n.NewRuntimeError("vapi.method.status.errors.invalid", args)
	log.Errorf(message.Error())
	return []error{message}

}

func (localProvider *LocalProvider) Invoke(serviceID string, operationID string, input data.DataValue, ctx *core.ExecutionContext) core.MethodResult {

	var interfaceID = core.NewInterfaceIdentifier(serviceID)
	var methodID = core.NewMethodIdentifier(interfaceID, operationID)
	log.Debugf("Searching for service %s", serviceID)
	var apiInterface = localProvider.serviceMap[serviceID]
	if apiInterface == nil {
		var errorValue = bindings.CreateErrorValueFromMessageId(bindings.OP_NOT_FOUND_ERROR_DEF,
			"vapi.method.input.invalid.interface", map[string]string{"serviceId": serviceID})
		return core.NewMethodResult(nil, errorValue)
	}
	log.Debugf("Searching for operation %s", operationID)
	var methodDef = apiInterface.MethodDefinition(methodID)
	methodIdArgs := map[string]string{"methodId": methodID.Name()}
	if methodDef == nil {
		var errorValue = bindings.CreateErrorValueFromMessageId(bindings.OP_NOT_FOUND_ERROR_DEF,
			"vapi.method.input.invalid.method", methodIdArgs)
		return core.NewMethodResult(nil, errorValue)
	}
	log.Debug("Validating input")
	//Step 1: Verify input
	if input == nil || input.Type() != data.STRUCTURE {
		log.Errorf("Expected input of type STRUCTURE but found nil")
		var errorValue = bindings.CreateErrorValueFromMessageId(bindings.INVALID_ARGUMENT_ERROR_DEF,
			"vapi.method.input.invalid", methodIdArgs)
		return core.NewMethodResult(nil, errorValue)
	}
	var inputDef = methodDef.InputDefinition()
	//This case should not happen
	if inputDef.Type() != data.STRUCTURE {
		log.Error("vapi.method.input.invalid.definition ")
		var errorValue = bindings.CreateErrorValueFromMessageId(bindings.INTERNAL_SERVER_ERROR_DEF,
			"vapi.method.input.invalid.definition", methodIdArgs)
		return core.NewMethodResult(nil, errorValue)
	}

	//Step 2: Execute method
	log.Debug("Invoking operation")
	methodResult := apiInterface.Invoke(ctx, methodID, input)

	//Step 3: Validate output
	log.Debug("Validating output")
	var outputDef = methodDef.OutputDefinition()
	if methodResult.IsResponseStream() {
		// Validation for stream response is done in binding (__Name__ApiInterface.go)
	} else {
		if methodResult.IsSuccess() {
			var messages = outputDef.Validate(methodResult.Output())
			if len(messages) != 0 {
				log.Errorf("Output validation failed for method %s", methodID.Name())
				var errorValue = bindings.CreateErrorValueFromMessages(bindings.INTERNAL_SERVER_ERROR_DEF, messages)
				return core.NewMethodResult(nil, errorValue)
			}
		} else {
			var err = methodResult.Error()
			var messages = localProvider.validateError(err, methodDef)
			if len(messages) != 0 {
				var errorVal = bindings.CreateErrorValueFromErrorValueAndMessages(bindings.INTERNAL_SERVER_ERROR_DEF, err, messages)
				return core.NewMethodResult(nil, errorVal)
			}
			return core.NewMethodResult(nil, err)
		}
		log.Debug("Request processing complete")
	}

	return methodResult

}

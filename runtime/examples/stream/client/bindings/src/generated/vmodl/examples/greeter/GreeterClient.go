// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Greeter
// Used by client-side stubs.

package greeter

import (
	"context"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Greeter`` interface provides methods to greet the client.
type GreeterClient interface {

	// Provides a stream of greetings
	// @return a greeting string.
	//
	// @throws InternalServerError
	Greet() (chan string, context.CancelFunc, chan error)
}

type greeterClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewGreeterClient(connector vapiProtocolClient_.Connector) *greeterClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("vmodl.examples.greeter.greeter")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"greet": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "greet"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	gIface := greeterClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *greeterClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *greeterClient) Greet() (chan string, context.CancelFunc, chan error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := greeterGreetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, true))

	ctx, cancelFunc := context.WithCancel(executionContext.Context())
	executionContext.WithContext(ctx)
	sv := vapiBindings_.NewStructValueBuilder(greeterGreetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return nil, cancelFunc, vapiCore_.GetErrorChan(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("vmodl.examples.greeter.greeter", "greet", inputDataValue, executionContext)
	resChan := make(chan string)
	errChan := make(chan error)
	go func() {
		for mR := range methodResult.ResponseStream() {
			if mR.Error() != nil {
				methodError, converterError := typeConverter.ConvertToGolang(mR.Error(), gIface.GetErrorBindingType(mR.Error().Name()))
				if converterError != nil {
					for _, err := range converterError {
						errChan <- err
					}
				} else {
					errChan <- methodError.(error)
				}
			} else {
				output, converterError := typeConverter.ConvertToGolang(mR.Output(), GreeterGreetOutputType())
				if converterError != nil {
					for _, err := range converterError {
						errChan <- err
					}
				} else {
					resChan <- output.(string)
				}
			}
		}
		close(resChan)
		close(errChan)
	}()
	return resChan, cancelFunc, errChan
}


/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Status
 * Functions that implement the generated StatusClient interface
 */


package status
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type StatusClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewStatusClientImpl(connector client.Connector) *StatusClientImpl {
      interfaceName := "com.vmware.esx.kms.providers.status"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      sIface := StatusClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      sIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      sIface.methodNameToDefMap["get"] = sIface.getMethodDefinition()
      return &sIface
}

func (sIface *StatusClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(sIface.interfaceName)
}

func (sIface *StatusClientImpl) Definition() core.InterfaceDefinition {
    return sIface.interfaceDefinition
}

func (sIface *StatusClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := sIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (sIface *StatusClientImpl) Get(providerParam string) (Info, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *StatusClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := sIface.connector.GetApiProvider().Invoke(sIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (sIface *StatusClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for StatusClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



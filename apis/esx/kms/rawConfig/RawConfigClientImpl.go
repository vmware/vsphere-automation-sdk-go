
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: RawConfig
 * Functions that implement the generated RawConfigClient interface
 */


package rawConfig
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type RawConfigClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewRawConfigClientImpl(connector client.Connector) *RawConfigClientImpl {
      interfaceName := "com.vmware.esx.kms.raw_config"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "set"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      rIface := RawConfigClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      rIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      rIface.methodNameToDefMap["get"] = rIface.getMethodDefinition()
      rIface.methodNameToDefMap["set"] = rIface.setMethodDefinition()
      return &rIface
}

func (rIface *RawConfigClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(rIface.interfaceName)
}

func (rIface *RawConfigClientImpl) Definition() core.InterfaceDefinition {
    return rIface.interfaceDefinition
}

func (rIface *RawConfigClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := rIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (rIface *RawConfigClientImpl) Get(type_Param *InfoType) (Info, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= rIface.Invoke(rIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (rIface *RawConfigClientImpl) Set(specParam SetSpec) error {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "set")
	sv := bindings.NewStructValueBuilder(setInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= rIface.Invoke(rIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *RawConfigClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := rIface.connector.GetApiProvider().Invoke(rIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (rIface *RawConfigClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
      typeConverter := rIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      rIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (rIface *RawConfigClientImpl) setMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
      typeConverter := rIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(setInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(setOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "set")
      errorDefinitions := make([]data.ErrorDefinition,0)
      rIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for RawConfigClientImpl.set method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}




/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: EnhancedSecurity
 * Functions that implement the generated EnhancedSecurityClient interface
 */


package enhancedSecurity
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type EnhancedSecurityClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewEnhancedSecurityClientImpl(connector client.Connector) *EnhancedSecurityClientImpl {
      interfaceName := "com.vmware.appliance.system.security.enhanced_security"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
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
      eIface := EnhancedSecurityClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      eIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      eIface.methodNameToDefMap["set"] = eIface.setMethodDefinition()
      return &eIface
}

func (eIface *EnhancedSecurityClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(eIface.interfaceName)
}

func (eIface *EnhancedSecurityClientImpl) Definition() core.InterfaceDefinition {
    return eIface.interfaceDefinition
}

func (eIface *EnhancedSecurityClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := eIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (eIface *EnhancedSecurityClientImpl) Set(enabledParam bool) error {
	typeConverter := eIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(eIface.interfaceIdentifier, "set")
	sv := bindings.NewStructValueBuilder(setInputType(), typeConverter)
	sv.AddStructField("Enabled", enabledParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= eIface.Invoke(eIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *EnhancedSecurityClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := eIface.connector.GetApiProvider().Invoke(eIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (eIface *EnhancedSecurityClientImpl) setMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(eIface.interfaceName)
      typeConverter := eIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(setInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(setOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for EnhancedSecurityClientImpl.set method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for EnhancedSecurityClientImpl.set method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "set")
      errorDefinitions := make([]data.ErrorDefinition,0)
      eIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for EnhancedSecurityClientImpl.set method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



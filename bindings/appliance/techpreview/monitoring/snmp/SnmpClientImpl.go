
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Snmp
 * Functions that implement the generated SnmpClient interface
 */


package snmp
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type SnmpClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewSnmpClientImpl(connector client.Connector) *SnmpClientImpl {
      interfaceName := "com.vmware.appliance.techpreview.monitoring.snmp"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "reset"),
          core.NewMethodIdentifier(interfaceIdentifier, "enable"),
          core.NewMethodIdentifier(interfaceIdentifier, "hash"),
          core.NewMethodIdentifier(interfaceIdentifier, "limits"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "disable"),
          core.NewMethodIdentifier(interfaceIdentifier, "set"),
          core.NewMethodIdentifier(interfaceIdentifier, "test"),
          core.NewMethodIdentifier(interfaceIdentifier, "stats"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      sIface := SnmpClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      sIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      sIface.methodNameToDefMap["reset"] = sIface.resetMethodDefinition()
      sIface.methodNameToDefMap["enable"] = sIface.enableMethodDefinition()
      sIface.methodNameToDefMap["hash"] = sIface.hashMethodDefinition()
      sIface.methodNameToDefMap["limits"] = sIface.limitsMethodDefinition()
      sIface.methodNameToDefMap["get"] = sIface.getMethodDefinition()
      sIface.methodNameToDefMap["disable"] = sIface.disableMethodDefinition()
      sIface.methodNameToDefMap["set"] = sIface.setMethodDefinition()
      sIface.methodNameToDefMap["test"] = sIface.testMethodDefinition()
      sIface.methodNameToDefMap["stats"] = sIface.statsMethodDefinition()
      return &sIface
}

func (sIface *SnmpClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(sIface.interfaceName)
}

func (sIface *SnmpClientImpl) Definition() core.InterfaceDefinition {
    return sIface.interfaceDefinition
}

func (sIface *SnmpClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := sIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (sIface *SnmpClientImpl) Reset() error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "reset")
	sv := bindings.NewStructValueBuilder(resetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Enable() error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "enable")
	sv := bindings.NewStructValueBuilder(enableInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := enableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Hash(configParam SNMPHashConfig) (SNMPHashResults, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "hash")
	sv := bindings.NewStructValueBuilder(hashInputType(), typeConverter)
	sv.AddStructField("Config", configParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput SNMPHashResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hashRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput SNMPHashResults
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hashOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(SNMPHashResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Limits() (SNMPLimits, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "limits")
	sv := bindings.NewStructValueBuilder(limitsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput SNMPLimits
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := limitsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput SNMPLimits
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), limitsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(SNMPLimits), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Get() (SNMPConfigReadOnly, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput SNMPConfigReadOnly
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput SNMPConfigReadOnly
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(SNMPConfigReadOnly), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Disable() error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "disable")
	sv := bindings.NewStructValueBuilder(disableInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := disableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Set(configParam SNMPConfig) error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "set")
	sv := bindings.NewStructValueBuilder(setInputType(), typeConverter)
	sv.AddStructField("Config", configParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Test() (SNMPTestResults, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "test")
	sv := bindings.NewStructValueBuilder(testInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput SNMPTestResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := testRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput SNMPTestResults
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), testOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(SNMPTestResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SnmpClientImpl) Stats() (SNMPStats, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "stats")
	sv := bindings.NewStructValueBuilder(statsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput SNMPStats
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := statsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput SNMPStats
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), statsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(SNMPStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *SnmpClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := sIface.connector.GetApiProvider().Invoke(sIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (sIface *SnmpClientImpl) resetMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(resetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(resetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.reset method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.reset method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "reset")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.reset method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) enableMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(enableInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(enableOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.enable method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.enable method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "enable")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.enable method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) hashMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(hashInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(hashOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.hash method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.hash method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "hash")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.hash method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) limitsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(limitsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(limitsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.limits method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.limits method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "limits")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.limits method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) disableMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(disableInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(disableOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.disable method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.disable method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "disable")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.disable method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) setMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(setInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(setOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.set method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.set method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "set")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.set method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) testMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(testInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(testOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.test method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.test method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "test")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.test method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SnmpClientImpl) statsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(statsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(statsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.stats method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.stats method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "stats")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SnmpClientImpl.stats method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}




/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Pending
 * Functions that implement the generated PendingClient interface
 */


package pending
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance/update"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type PendingClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewPendingClientImpl(connector client.Connector) *PendingClientImpl {
      interfaceName := "com.vmware.appliance.update.pending"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "precheck"),
          core.NewMethodIdentifier(interfaceIdentifier, "stage"),
          core.NewMethodIdentifier(interfaceIdentifier, "validate"),
          core.NewMethodIdentifier(interfaceIdentifier, "install"),
          core.NewMethodIdentifier(interfaceIdentifier, "stageAndInstall"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      pIface := PendingClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      pIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      pIface.methodNameToDefMap["list"] = pIface.listMethodDefinition()
      pIface.methodNameToDefMap["get"] = pIface.getMethodDefinition()
      pIface.methodNameToDefMap["precheck"] = pIface.precheckMethodDefinition()
      pIface.methodNameToDefMap["stage"] = pIface.stageMethodDefinition()
      pIface.methodNameToDefMap["validate"] = pIface.validateMethodDefinition()
      pIface.methodNameToDefMap["install"] = pIface.installMethodDefinition()
      pIface.methodNameToDefMap["stage_and_install"] = pIface.stageAndInstallMethodDefinition()
      return &pIface
}

func (pIface *PendingClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(pIface.interfaceName)
}

func (pIface *PendingClientImpl) Definition() core.InterfaceDefinition {
    return pIface.interfaceDefinition
}

func (pIface *PendingClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := pIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (pIface *PendingClientImpl) List(sourceTypeParam SourceType, urlParam *string) ([]update.Summary, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(listInputType(), typeConverter)
	sv.AddStructField("SourceType", sourceTypeParam)
	sv.AddStructField("Url", urlParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []update.Summary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []update.Summary
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]update.Summary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PendingClientImpl) Get(versionParam string) (Info, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PendingClientImpl) Precheck(versionParam string) (PrecheckResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "precheck")
	sv := bindings.NewStructValueBuilder(precheckInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput PrecheckResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := precheckRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput PrecheckResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), precheckOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(PrecheckResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PendingClientImpl) Stage(versionParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "stage")
	sv := bindings.NewStructValueBuilder(stageInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := stageRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (pIface *PendingClientImpl) Validate(versionParam string, userDataParam map[string]string) (appliance.Notifications, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "validate")
	sv := bindings.NewStructValueBuilder(validateInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	sv.AddStructField("UserData", userDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput appliance.Notifications
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := validateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput appliance.Notifications
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), validateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(appliance.Notifications), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PendingClientImpl) Install(versionParam string, userDataParam map[string]string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "install")
	sv := bindings.NewStructValueBuilder(installInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	sv.AddStructField("UserData", userDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := installRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (pIface *PendingClientImpl) StageAndInstall(versionParam string, userDataParam map[string]string) error {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "stage_and_install")
	sv := bindings.NewStructValueBuilder(stageAndInstallInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	sv.AddStructField("UserData", userDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := stageAndInstallRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *PendingClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := pIface.connector.GetApiProvider().Invoke(pIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (pIface *PendingClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.get method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) precheckMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(precheckInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(precheckOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "precheck")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.precheck method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) stageMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(stageInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(stageOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "stage")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.AlreadyExists error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stage method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) validateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(validateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(validateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "validate")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.validate method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) installMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(installInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(installOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "install")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.install method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PendingClientImpl) stageAndInstallMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(stageAndInstallInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(stageAndInstallOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "stageAndInstall")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.AlreadyInDesiredState error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PendingClientImpl.stageAndInstall method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



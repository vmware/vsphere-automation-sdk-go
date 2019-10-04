
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Job
 * Functions that implement the generated JobClient interface
 */


package job
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type JobClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewJobClientImpl(connector client.Connector) *JobClientImpl {
      interfaceName := "com.vmware.appliance.recovery.restore.job"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "cancel"),
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
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
      jIface := JobClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      jIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      jIface.methodNameToDefMap["cancel"] = jIface.cancelMethodDefinition()
      jIface.methodNameToDefMap["create"] = jIface.createMethodDefinition()
      jIface.methodNameToDefMap["get"] = jIface.getMethodDefinition()
      return &jIface
}

func (jIface *JobClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(jIface.interfaceName)
}

func (jIface *JobClientImpl) Definition() core.InterfaceDefinition {
    return jIface.interfaceDefinition
}

func (jIface *JobClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := jIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (jIface *JobClientImpl) Cancel() (ReturnResult, error) {
	typeConverter := jIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(jIface.interfaceIdentifier, "cancel")
	sv := bindings.NewStructValueBuilder(cancelInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput ReturnResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cancelRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	jIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= jIface.Invoke(jIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput ReturnResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cancelOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(ReturnResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), jIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (jIface *JobClientImpl) Create(pieceParam RestoreRequest) (RestoreJobStatus, error) {
	typeConverter := jIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(jIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(createInputType(), typeConverter)
	sv.AddStructField("Piece", pieceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput RestoreJobStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := createRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	jIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= jIface.Invoke(jIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput RestoreJobStatus
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), createOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(RestoreJobStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), jIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (jIface *JobClientImpl) Get() (RestoreJobStatus, error) {
	typeConverter := jIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(jIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput RestoreJobStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	jIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= jIface.Invoke(jIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput RestoreJobStatus
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(RestoreJobStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), jIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (jIface *JobClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := jIface.connector.GetApiProvider().Invoke(jIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (jIface *JobClientImpl) cancelMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(jIface.interfaceName)
      typeConverter := jIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(cancelInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(cancelOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.cancel method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.cancel method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "cancel")
      errorDefinitions := make([]data.ErrorDefinition,0)
      jIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.cancel method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (jIface *JobClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(jIface.interfaceName)
      typeConverter := jIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(createInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(createOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      jIface.errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.FeatureInUseBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.create method's errors.FeatureInUse error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      jIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.create method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      jIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.create method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (jIface *JobClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(jIface.interfaceName)
      typeConverter := jIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      jIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for JobClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}




/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Backup
 * Functions that implement the generated BackupClient interface
 */


package backup
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type BackupClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewBackupClientImpl(connector client.Connector) *BackupClientImpl {
      interfaceName := "com.vmware.appliance.recovery.backup"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "validate"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      bIface := BackupClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      bIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      bIface.methodNameToDefMap["validate"] = bIface.validateMethodDefinition()
      return &bIface
}

func (bIface *BackupClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(bIface.interfaceName)
}

func (bIface *BackupClientImpl) Definition() core.InterfaceDefinition {
    return bIface.interfaceDefinition
}

func (bIface *BackupClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := bIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (bIface *BackupClientImpl) Validate(pieceParam BackupRequest) (ReturnResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "validate")
	sv := bindings.NewStructValueBuilder(validateInputType(), typeConverter)
	sv.AddStructField("Piece", pieceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput ReturnResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := validateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput ReturnResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), validateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(ReturnResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BackupClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := bIface.connector.GetApiProvider().Invoke(bIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (bIface *BackupClientImpl) validateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(validateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(validateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BackupClientImpl.validate method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BackupClientImpl.validate method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "validate")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.FeatureInUseBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BackupClientImpl.validate method's errors.FeatureInUse error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      bIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BackupClientImpl.validate method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



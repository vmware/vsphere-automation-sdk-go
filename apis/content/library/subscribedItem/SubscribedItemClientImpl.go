
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: SubscribedItem
 * Functions that implement the generated SubscribedItemClient interface
 */


package subscribedItem
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type SubscribedItemClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewSubscribedItemClientImpl(connector client.Connector) *SubscribedItemClientImpl {
      interfaceName := "com.vmware.content.library.subscribed_item"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "evict"),
          core.NewMethodIdentifier(interfaceIdentifier, "sync"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      sIface := SubscribedItemClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      sIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      sIface.methodNameToDefMap["evict"] = sIface.evictMethodDefinition()
      sIface.methodNameToDefMap["sync"] = sIface.syncMethodDefinition()
      return &sIface
}

func (sIface *SubscribedItemClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(sIface.interfaceName)
}

func (sIface *SubscribedItemClientImpl) Definition() core.InterfaceDefinition {
    return sIface.interfaceDefinition
}

func (sIface *SubscribedItemClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := sIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (sIface *SubscribedItemClientImpl) Evict(libraryItemIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "evict")
	sv := bindings.NewStructValueBuilder(evictInputType(), typeConverter)
	sv.AddStructField("LibraryItemId", libraryItemIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := evictRestMetadata()
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
func (sIface *SubscribedItemClientImpl) Sync(libraryItemIdParam string, forceSyncContentParam bool) error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "sync")
	sv := bindings.NewStructValueBuilder(syncInputType(), typeConverter)
	sv.AddStructField("LibraryItemId", libraryItemIdParam)
	sv.AddStructField("ForceSyncContent", forceSyncContentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := syncRestMetadata()
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

func (sIface *SubscribedItemClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := sIface.connector.GetApiProvider().Invoke(sIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (sIface *SubscribedItemClientImpl) evictMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(evictInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(evictOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "evict")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.InvalidElementType{}.Error()] = errors.InvalidElementTypeBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidElementTypeBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's errors.InvalidElementType error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's errors.InvalidElementConfiguration error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.evict method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SubscribedItemClientImpl) syncMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(syncInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(syncOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.sync method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.sync method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "sync")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.sync method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.InvalidElementType{}.Error()] = errors.InvalidElementTypeBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidElementTypeBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.sync method's errors.InvalidElementType error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SubscribedItemClientImpl.sync method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



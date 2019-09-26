
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: ImportSession
 * Functions that implement the generated ImportSessionClient interface
 */


package importSession
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
    "net/url"
)


type ImportSessionClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewImportSessionClientImpl(connector client.Connector) *ImportSessionClientImpl {
      interfaceName := "com.vmware.vcenter.ovf.import_session"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "createForResourcePool"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "tryInstantiate"),
          core.NewMethodIdentifier(interfaceIdentifier, "instantiate"),
          core.NewMethodIdentifier(interfaceIdentifier, "progress"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
          core.NewMethodIdentifier(interfaceIdentifier, "preview"),
          core.NewMethodIdentifier(interfaceIdentifier, "probe"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      iIface := ImportSessionClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      iIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      iIface.methodNameToDefMap["create_for_resource_pool"] = iIface.createForResourcePoolMethodDefinition()
      iIface.methodNameToDefMap["get"] = iIface.getMethodDefinition()
      iIface.methodNameToDefMap["try_instantiate"] = iIface.tryInstantiateMethodDefinition()
      iIface.methodNameToDefMap["instantiate"] = iIface.instantiateMethodDefinition()
      iIface.methodNameToDefMap["progress"] = iIface.progressMethodDefinition()
      iIface.methodNameToDefMap["delete"] = iIface.deleteMethodDefinition()
      iIface.methodNameToDefMap["preview"] = iIface.previewMethodDefinition()
      iIface.methodNameToDefMap["probe"] = iIface.probeMethodDefinition()
      return &iIface
}

func (iIface *ImportSessionClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(iIface.interfaceName)
}

func (iIface *ImportSessionClientImpl) Definition() core.InterfaceDefinition {
    return iIface.interfaceDefinition
}

func (iIface *ImportSessionClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := iIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (iIface *ImportSessionClientImpl) CreateForResourcePool(clientTokenParam *string, resourcePoolParam string, hostSystemParam *string, folderParam *string, createSpecParam *data.StructValue) (string, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "create_for_resource_pool")
	sv := bindings.NewStructValueBuilder(createForResourcePoolInputType(), typeConverter)
	sv.AddStructField("ClientToken", clientTokenParam)
	sv.AddStructField("ResourcePool", resourcePoolParam)
	sv.AddStructField("HostSystem", hostSystemParam)
	sv.AddStructField("Folder", folderParam)
	sv.AddStructField("CreateSpec", createSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := createForResourcePoolRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), createForResourcePoolOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Get(idParam string) (Info, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(getInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) TryInstantiate(idParam string, instantiationParametersParam []*data.StructValue) (OvfValidationResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "try_instantiate")
	sv := bindings.NewStructValueBuilder(tryInstantiateInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("InstantiationParameters", instantiationParametersParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput OvfValidationResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tryInstantiateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput OvfValidationResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tryInstantiateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(OvfValidationResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Instantiate(idParam string, instantiationParametersParam []*data.StructValue) error {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "instantiate")
	sv := bindings.NewStructValueBuilder(instantiateInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("InstantiationParameters", instantiationParametersParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := instantiateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Progress(idParam string, percentParam int64) error {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "progress")
	sv := bindings.NewStructValueBuilder(progressInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Percent", percentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := progressRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Delete(idParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(deleteInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := deleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Preview(ovfDescriptorParam string) (Preview, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "preview")
	sv := bindings.NewStructValueBuilder(previewInputType(), typeConverter)
	sv.AddStructField("OvfDescriptor", ovfDescriptorParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Preview
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := previewRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Preview
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), previewOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Preview), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (iIface *ImportSessionClientImpl) Probe(uriParam url.URL, sslCertificateThumbprintParam *string) (ProbeResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(iIface.interfaceIdentifier, "probe")
	sv := bindings.NewStructValueBuilder(probeInputType(), typeConverter)
	sv.AddStructField("Uri", uriParam)
	sv.AddStructField("SslCertificateThumbprint", sslCertificateThumbprintParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput ProbeResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := probeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= iIface.Invoke(iIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput ProbeResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), probeOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(ProbeResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ImportSessionClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := iIface.connector.GetApiProvider().Invoke(iIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (iIface *ImportSessionClientImpl) createForResourcePoolMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(createForResourcePoolInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(createForResourcePoolOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.createForResourcePool method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.createForResourcePool method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "createForResourcePool")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.createForResourcePool method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.createForResourcePool method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.createForResourcePool method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) tryInstantiateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tryInstantiateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tryInstantiateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.tryInstantiate method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.tryInstantiate method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "tryInstantiate")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.tryInstantiate method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.tryInstantiate method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.tryInstantiate method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) instantiateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(instantiateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(instantiateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "instantiate")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's errors.Unsupported error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.instantiate method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) progressMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(progressInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(progressOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.progress method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.progress method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "progress")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.progress method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.progress method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      iIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.progress method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(deleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(deleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) previewMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(previewInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(previewOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.preview method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.preview method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "preview")
      errorDefinitions := make([]data.ErrorDefinition,0)
      iIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.preview method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (iIface *ImportSessionClientImpl) probeMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(iIface.interfaceName)
      typeConverter := iIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(probeInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(probeOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.probe method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for ImportSessionClientImpl.probe method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "probe")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



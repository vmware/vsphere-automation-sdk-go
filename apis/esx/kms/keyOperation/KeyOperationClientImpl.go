
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: KeyOperation
 * Functions that implement the generated KeyOperationClient interface
 */


package keyOperation
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type KeyOperationClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewKeyOperationClientImpl(connector client.Connector) *KeyOperationClientImpl {
      interfaceName := "com.vmware.esx.kms.key_operation"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "generateKey"),
          core.NewMethodIdentifier(interfaceIdentifier, "encrypt"),
          core.NewMethodIdentifier(interfaceIdentifier, "decrypt"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      kIface := KeyOperationClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      kIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      kIface.methodNameToDefMap["generate_key"] = kIface.generateKeyMethodDefinition()
      kIface.methodNameToDefMap["encrypt"] = kIface.encryptMethodDefinition()
      kIface.methodNameToDefMap["decrypt"] = kIface.decryptMethodDefinition()
      return &kIface
}

func (kIface *KeyOperationClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(kIface.interfaceName)
}

func (kIface *KeyOperationClientImpl) Definition() core.InterfaceDefinition {
    return kIface.interfaceDefinition
}

func (kIface *KeyOperationClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := kIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (kIface *KeyOperationClientImpl) GenerateKey(providerParam string, numOfBytesParam int64) (GeneratedKey, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "generate_key")
	sv := bindings.NewStructValueBuilder(generateKeyInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("NumOfBytes", numOfBytesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput GeneratedKey
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := generateKeyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput GeneratedKey
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), generateKeyOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(GeneratedKey), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (kIface *KeyOperationClientImpl) Encrypt(providerParam string, plaintextParam string) (EncryptResult, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "encrypt")
	sv := bindings.NewStructValueBuilder(encryptInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("Plaintext", plaintextParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput EncryptResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := encryptRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput EncryptResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), encryptOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(EncryptResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (kIface *KeyOperationClientImpl) Decrypt(providerParam string, ciphertextParam string) (DecryptResult, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "decrypt")
	sv := bindings.NewStructValueBuilder(decryptInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("Ciphertext", ciphertextParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput DecryptResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := decryptRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput DecryptResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), decryptOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(DecryptResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (kIface *KeyOperationClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := kIface.connector.GetApiProvider().Invoke(kIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (kIface *KeyOperationClientImpl) generateKeyMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
      typeConverter := kIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(generateKeyInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(generateKeyOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "generateKey")
      errorDefinitions := make([]data.ErrorDefinition,0)
      kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.generateKey method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (kIface *KeyOperationClientImpl) encryptMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
      typeConverter := kIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(encryptInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(encryptOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "encrypt")
      errorDefinitions := make([]data.ErrorDefinition,0)
      kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.encrypt method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (kIface *KeyOperationClientImpl) decryptMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
      typeConverter := kIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(decryptInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(decryptOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "decrypt")
      errorDefinitions := make([]data.ErrorDefinition,0)
      kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for KeyOperationClientImpl.decrypt method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



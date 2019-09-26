
/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Batch
 * Functions that implement the generated BatchClient interface
 */


package batch
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)


type BatchClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewBatchClientImpl(connector client.Connector) *BatchClientImpl {
      interfaceName := "com.vmware.cis.tagging.batch"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "getCategories"),
          core.NewMethodIdentifier(interfaceIdentifier, "getAllCategories"),
          core.NewMethodIdentifier(interfaceIdentifier, "getTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "getAllTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listTagsForCategories"),
          core.NewMethodIdentifier(interfaceIdentifier, "findTagsByName"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAllAttachedObjectsOnTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      bIface := BatchClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      bIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      bIface.methodNameToDefMap["get_categories"] = bIface.getCategoriesMethodDefinition()
      bIface.methodNameToDefMap["get_all_categories"] = bIface.getAllCategoriesMethodDefinition()
      bIface.methodNameToDefMap["get_tags"] = bIface.getTagsMethodDefinition()
      bIface.methodNameToDefMap["get_all_tags"] = bIface.getAllTagsMethodDefinition()
      bIface.methodNameToDefMap["list_tags_for_categories"] = bIface.listTagsForCategoriesMethodDefinition()
      bIface.methodNameToDefMap["find_tags_by_name"] = bIface.findTagsByNameMethodDefinition()
      bIface.methodNameToDefMap["list_attached_objects"] = bIface.listAttachedObjectsMethodDefinition()
      bIface.methodNameToDefMap["list_attached_objects_on_tags"] = bIface.listAttachedObjectsOnTagsMethodDefinition()
      bIface.methodNameToDefMap["list_all_attached_objects_on_tags"] = bIface.listAllAttachedObjectsOnTagsMethodDefinition()
      bIface.methodNameToDefMap["list_attached_tags"] = bIface.listAttachedTagsMethodDefinition()
      bIface.methodNameToDefMap["list_attached_tags_on_objects"] = bIface.listAttachedTagsOnObjectsMethodDefinition()
      return &bIface
}

func (bIface *BatchClientImpl) Identifier() core.InterfaceIdentifier {
      return core.NewInterfaceIdentifier(bIface.interfaceName)
}

func (bIface *BatchClientImpl) Definition() core.InterfaceDefinition {
    return bIface.interfaceDefinition
}

func (bIface *BatchClientImpl) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
    if val, ok := bIface.methodNameToDefMap[mi.Name()]; ok {
      return val
    }
    return nil
}

func (bIface *BatchClientImpl) GetCategories(categoryIdsParam []string) ([]tagging.CategoryModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_categories")
	sv := bindings.NewStructValueBuilder(getCategoriesInputType(), typeConverter)
	sv.AddStructField("CategoryIds", categoryIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []tagging.CategoryModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []tagging.CategoryModel
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getCategoriesOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]tagging.CategoryModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) GetAllCategories() ([]tagging.CategoryModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_all_categories")
	sv := bindings.NewStructValueBuilder(getAllCategoriesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []tagging.CategoryModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getAllCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []tagging.CategoryModel
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getAllCategoriesOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]tagging.CategoryModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) GetTags(tagIdsParam []string) ([]tagging.TagModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_tags")
	sv := bindings.NewStructValueBuilder(getTagsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []tagging.TagModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []tagging.TagModel
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]tagging.TagModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) GetAllTags() ([]tagging.TagModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_all_tags")
	sv := bindings.NewStructValueBuilder(getAllTagsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []tagging.TagModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := getAllTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []tagging.TagModel
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), getAllTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]tagging.TagModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListTagsForCategories(categoryIdsParam []string) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_tags_for_categories")
	sv := bindings.NewStructValueBuilder(listTagsForCategoriesInputType(), typeConverter)
	sv.AddStructField("CategoryIds", categoryIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listTagsForCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listTagsForCategoriesOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) FindTagsByName(tagNameParam string) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "find_tags_by_name")
	sv := bindings.NewStructValueBuilder(findTagsByNameInputType(), typeConverter)
	sv.AddStructField("TagName", tagNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := findTagsByNameRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), findTagsByNameOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListAttachedObjects(tagIdsParam []string) ([]std.DynamicID, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_objects")
	sv := bindings.NewStructValueBuilder(listAttachedObjectsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []std.DynamicID
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listAttachedObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []std.DynamicID
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listAttachedObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]std.DynamicID), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListAttachedObjectsOnTags(tagIdsParam []string) ([]TagToObjects, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_objects_on_tags")
	sv := bindings.NewStructValueBuilder(listAttachedObjectsOnTagsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []TagToObjects
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listAttachedObjectsOnTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagToObjects
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listAttachedObjectsOnTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]TagToObjects), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListAllAttachedObjectsOnTags() ([]TagToObjects, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_all_attached_objects_on_tags")
	sv := bindings.NewStructValueBuilder(listAllAttachedObjectsOnTagsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []TagToObjects
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listAllAttachedObjectsOnTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagToObjects
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listAllAttachedObjectsOnTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]TagToObjects), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListAttachedTags(objectIdsParam []std.DynamicID) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_tags")
	sv := bindings.NewStructValueBuilder(listAttachedTagsInputType(), typeConverter)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listAttachedTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listAttachedTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (bIface *BatchClientImpl) ListAttachedTagsOnObjects(objectIdsParam []std.DynamicID) ([]ObjectToTags, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_tags_on_objects")
	sv := bindings.NewStructValueBuilder(listAttachedTagsOnObjectsInputType(), typeConverter)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []ObjectToTags
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := listAttachedTagsOnObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []ObjectToTags
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), listAttachedTagsOnObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]ObjectToTags), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := bIface.connector.GetApiProvider().Invoke(bIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (bIface *BatchClientImpl) getCategoriesMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getCategoriesInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getCategoriesOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getCategories method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getCategories method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getCategories")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) getAllCategoriesMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getAllCategoriesInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getAllCategoriesOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllCategories method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllCategories method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getAllCategories")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) getTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getTags")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) getAllTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(getAllTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(getAllTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getAllTags")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listTagsForCategoriesMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listTagsForCategoriesInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listTagsForCategoriesOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listTagsForCategories method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listTagsForCategories method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listTagsForCategories")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) findTagsByNameMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(findTagsByNameInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(findTagsByNameOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.findTagsByName method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.findTagsByName method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "findTagsByName")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listAttachedObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listAttachedObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listAttachedObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listAttachedObjectsOnTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listAttachedObjectsOnTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listAttachedObjectsOnTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listAllAttachedObjectsOnTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listAllAttachedObjectsOnTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listAllAttachedObjectsOnTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAllAttachedObjectsOnTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listAttachedTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listAttachedTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listAttachedTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (bIface *BatchClientImpl) listAttachedTagsOnObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
      typeConverter := bIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(listAttachedTagsOnObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(listAttachedTagsOnObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



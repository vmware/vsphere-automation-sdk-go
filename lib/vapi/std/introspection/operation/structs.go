/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Operation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package operation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The DataDefinition structure describes a vAPI data type.
 type DataDefinition struct {
    Type_ DataDefinition_DataType
    ElementDefinition *DataDefinition
    Name *string
    Fields map[string]DataDefinition
}




    
    // The DataDefinition_DataType enumeration provides values representing the data types supported by the vAPI infrastructure.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DataDefinition_DataType string

    const (
        // Indicates the value is a binary type.
         DataDefinition_DataType_BINARY DataDefinition_DataType = "BINARY"
        // Indicates the value is a boolean type. The possible values are True and False equivalent of the language used to invoke this operation.
         DataDefinition_DataType_BOOLEAN DataDefinition_DataType = "BOOLEAN"
        // Indicates the value is a double type. It is a 64 bit floating point number.
         DataDefinition_DataType_DOUBLE DataDefinition_DataType = "DOUBLE"
        // Indicates the value is a dynamic structure. This means, any data of type DataDefinition_DataType#DataType_STRUCTURE can be used.
         DataDefinition_DataType_DYNAMIC_STRUCTURE DataDefinition_DataType = "DYNAMIC_STRUCTURE"
        // Indicates the value is a specific error type.
         DataDefinition_DataType_ERROR DataDefinition_DataType = "ERROR"
        // Indicates the value is arbitrary error type. This means, any data of type DataDefinition_DataType#DataType_ERROR can be used.
         DataDefinition_DataType_ANY_ERROR DataDefinition_DataType = "ANY_ERROR"
        // Indicates the value is a list data type. Any value of this type can have zero or more elements in the list.
         DataDefinition_DataType_LIST DataDefinition_DataType = "LIST"
        // Indicates the value is a long data type. It is a 64 bit signed integer number.
         DataDefinition_DataType_LONG DataDefinition_DataType = "LONG"
        // Indicates the value is an opaque type. This means, data of any DataDefinition_DataType can be used.
         DataDefinition_DataType_OPAQUE DataDefinition_DataType = "OPAQUE"
        // Indicates the value is an optional data type. Any value of this type can be null.
         DataDefinition_DataType_OPTIONAL DataDefinition_DataType = "OPTIONAL"
        // Indicates the value is a secret data type. This is used for sensitive information. The server will not log any data of this type and if possible wipe the data from the memory after usage.
         DataDefinition_DataType_SECRET DataDefinition_DataType = "SECRET"
        // Indicates the value is a string data type. This is a unicode string.
         DataDefinition_DataType_STRING DataDefinition_DataType = "STRING"
        // Indicates the value is a structure data type. A structure has string identifier and a set of fields with corresponding values.
         DataDefinition_DataType_STRUCTURE DataDefinition_DataType = "STRUCTURE"
        // Indicates the value is a structure reference. This is used to break circular dependencies in the type references. This just has a string identifier of the structure. Clients have to maintain a list of structures already visited and use that to resolve this reference.
         DataDefinition_DataType_STRUCTURE_REF DataDefinition_DataType = "STRUCTURE_REF"
        // Indicates the value is a void data type.
         DataDefinition_DataType_VOID DataDefinition_DataType = "VOID"
    )

    func (d DataDefinition_DataType) DataDefinition_DataType() bool {
        switch d {
            case DataDefinition_DataType_BINARY:
                return true
            case DataDefinition_DataType_BOOLEAN:
                return true
            case DataDefinition_DataType_DOUBLE:
                return true
            case DataDefinition_DataType_DYNAMIC_STRUCTURE:
                return true
            case DataDefinition_DataType_ERROR:
                return true
            case DataDefinition_DataType_ANY_ERROR:
                return true
            case DataDefinition_DataType_LIST:
                return true
            case DataDefinition_DataType_LONG:
                return true
            case DataDefinition_DataType_OPAQUE:
                return true
            case DataDefinition_DataType_OPTIONAL:
                return true
            case DataDefinition_DataType_SECRET:
                return true
            case DataDefinition_DataType_STRING:
                return true
            case DataDefinition_DataType_STRUCTURE:
                return true
            case DataDefinition_DataType_STRUCTURE_REF:
                return true
            case DataDefinition_DataType_VOID:
                return true
            default:
                return false
        }
    }



// Information about a vAPI operation.
 type Info struct {
    InputDefinition DataDefinition
    OutputDefinition DataDefinition
    ErrorDefinitions []DataDefinition
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fields["operation_id"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    fieldNameMap["operation_id"] = "OperationId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}



func DataDefinitionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vapi.std.introspection.operation.data_definition.data_type", reflect.TypeOf(DataDefinition_DataType(DataDefinition_DataType_BINARY)))
    fieldNameMap["type"] = "Type_"
    fields["element_definition"] = bindings.NewOptionalType(bindings.NewReferenceType(DataDefinitionBindingType))
    fieldNameMap["element_definition"] = "ElementDefinition"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["fields"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(DataDefinitionBindingType),reflect.TypeOf(map[string]DataDefinition{})))
    fieldNameMap["fields"] = "Fields"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "OPTIONAL": []bindings.FieldData {
                 bindings.NewFieldData("element_definition", true),
            },
            "LIST": []bindings.FieldData {
                 bindings.NewFieldData("element_definition", true),
            },
            "STRUCTURE": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("fields", true),
            },
            "STRUCTURE_REF": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
            },
            "ERROR": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("fields", true),
            },
            "BINARY": []bindings.FieldData {},
            "BOOLEAN": []bindings.FieldData {},
            "DOUBLE": []bindings.FieldData {},
            "DYNAMIC_STRUCTURE": []bindings.FieldData {},
            "ANY_ERROR": []bindings.FieldData {},
            "LONG": []bindings.FieldData {},
            "OPAQUE": []bindings.FieldData {},
            "SECRET": []bindings.FieldData {},
            "STRING": []bindings.FieldData {},
            "VOID": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.data_definition",fields, reflect.TypeOf(DataDefinition{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["input_definition"] = bindings.NewReferenceType(DataDefinitionBindingType)
    fieldNameMap["input_definition"] = "InputDefinition"
    fields["output_definition"] = bindings.NewReferenceType(DataDefinitionBindingType)
    fieldNameMap["output_definition"] = "OutputDefinition"
    fields["error_definitions"] = bindings.NewListType(bindings.NewReferenceType(DataDefinitionBindingType), reflect.TypeOf([]DataDefinition{}))
    fieldNameMap["error_definitions"] = "ErrorDefinitions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



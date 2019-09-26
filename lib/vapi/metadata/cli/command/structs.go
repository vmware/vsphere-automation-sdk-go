/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Command.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package command

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``FormatterType`` enumeration class defines supported CLI output formatter types. See Info#formatter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type FormatterType string

const (
    // Displays command output as it is.
     FormatterType_SIMPLE FormatterType = "SIMPLE"
    // Displays command output in table format.
     FormatterType_TABLE FormatterType = "TABLE"
    // Displays command output in JSON format.
     FormatterType_JSON FormatterType = "JSON"
    // Displays command output in XML format.
     FormatterType_XML FormatterType = "XML"
    // Displays command output in CSV format.
     FormatterType_CSV FormatterType = "CSV"
    // Displays command output in HTML format.
     FormatterType_HTML FormatterType = "HTML"
)

func (f FormatterType) FormatterType() bool {
    switch f {
        case FormatterType_SIMPLE:
            return true
        case FormatterType_TABLE:
            return true
        case FormatterType_JSON:
            return true
        case FormatterType_XML:
            return true
        case FormatterType_CSV:
            return true
        case FormatterType_HTML:
            return true
        default:
            return false
    }
}




// The ``GenericType`` enumeration class defines generic types supported by ``Command`` interface. See OptionInfo#generic.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type GenericType string

const (
    // Default case.
     GenericType_NONE GenericType = "NONE"
    // Input parameter is an optional.
     GenericType_OPTIONAL GenericType = "OPTIONAL"
    // Input parameter is a list.
     GenericType_LIST GenericType = "LIST"
    // Input parameter is an optional of type list.
     GenericType_OPTIONAL_LIST GenericType = "OPTIONAL_LIST"
    // Input parameter is a list of optionals.
     GenericType_LIST_OPTIONAL GenericType = "LIST_OPTIONAL"
)

func (g GenericType) GenericType() bool {
    switch g {
        case GenericType_NONE:
            return true
        case GenericType_OPTIONAL:
            return true
        case GenericType_LIST:
            return true
        case GenericType_OPTIONAL_LIST:
            return true
        case GenericType_LIST_OPTIONAL:
            return true
        default:
            return false
    }
}





// The ``OutputFieldInfo`` class describes the name used by the CLI to display a single property of a class element in the interface definition language.
 type OutputFieldInfo struct {
    FieldName string
    DisplayName string
}






// The ``OutputInfo`` class describes the names used by the CLI to display the properties of a class element in the interface definition language as well as the order in which the properties will be displayed.
 type OutputInfo struct {
    StructureId string
    OutputFields []OutputFieldInfo
}






// The ``OptionInfo`` class describes information about a specific input option of a command.
 type OptionInfo struct {
    LongOption string
    ShortOption *string
    FieldName string
    Description string
    Type_ string
    Generic GenericType
}






// The ``Identity`` class uniquely identifies a command in the CLI commands tree.
 type Identity struct {
    Path string
    Name string
}






// The ``Info`` class contains information about a command. It includes the identity of the command, a description, information about the interface and method that implement the command, and CLI-specific information for the command.
 type Info struct {
    Identity Identity
    Description string
    ServiceId string
    OperationId string
    Options []OptionInfo
    Formatter *FormatterType
    OutputFieldList []OutputInfo
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["path"] = "Path"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(IdentityBindingType), reflect.TypeOf([]Identity{}))
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
    fields["identity"] = bindings.NewReferenceType(IdentityBindingType)
    fieldNameMap["identity"] = "Identity"
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


func fingerprintInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fingerprintOutputType() bindings.BindingType {
    return bindings.NewStringType()
}

func fingerprintRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func OutputFieldInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["field_name"] = bindings.NewStringType()
    fieldNameMap["field_name"] = "FieldName"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.output_field_info",fields, reflect.TypeOf(OutputFieldInfo{}), fieldNameMap, validators)
}

func OutputInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["structure_id"] = bindings.NewIdType([]string {"com.vmware.vapi.structure"}, "")
    fieldNameMap["structure_id"] = "StructureId"
    fields["output_fields"] = bindings.NewListType(bindings.NewReferenceType(OutputFieldInfoBindingType), reflect.TypeOf([]OutputFieldInfo{}))
    fieldNameMap["output_fields"] = "OutputFields"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.output_info",fields, reflect.TypeOf(OutputInfo{}), fieldNameMap, validators)
}

func OptionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["long_option"] = bindings.NewStringType()
    fieldNameMap["long_option"] = "LongOption"
    fields["short_option"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["short_option"] = "ShortOption"
    fields["field_name"] = bindings.NewStringType()
    fieldNameMap["field_name"] = "FieldName"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["generic"] = bindings.NewEnumType("com.vmware.vapi.metadata.cli.command.generic_type", reflect.TypeOf(GenericType(GenericType_NONE)))
    fieldNameMap["generic"] = "Generic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.option_info",fields, reflect.TypeOf(OptionInfo{}), fieldNameMap, validators)
}

func IdentityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewStringType()
    fieldNameMap["path"] = "Path"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.identity",fields, reflect.TypeOf(Identity{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["identity"] = bindings.NewReferenceType(IdentityBindingType)
    fieldNameMap["identity"] = "Identity"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    fields["operation_id"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation_id"] = "OperationId"
    fields["options"] = bindings.NewListType(bindings.NewReferenceType(OptionInfoBindingType), reflect.TypeOf([]OptionInfo{}))
    fieldNameMap["options"] = "Options"
    fields["formatter"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.metadata.cli.command.formatter_type", reflect.TypeOf(FormatterType(FormatterType_SIMPLE))))
    fieldNameMap["formatter"] = "Formatter"
    fields["output_field_list"] = bindings.NewListType(bindings.NewReferenceType(OutputInfoBindingType), reflect.TypeOf([]OutputInfo{}))
    fieldNameMap["output_field_list"] = "OutputFieldList"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



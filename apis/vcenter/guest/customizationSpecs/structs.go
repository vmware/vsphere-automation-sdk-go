/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CustomizationSpecs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package customizationSpecs

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/guest"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// The resource type for a vCenter guest customization specification
const CustomizationSpecs_RESOURCE_TYPE = "com.vmware.vcenter.guest.CustomizationSpec"


// The ``OsType`` enumeration class defines the types of guest operating systems for which guest customization is supported.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type OsType string

const (
    // A customization specification for a Windows guest operating system
     OsType_WINDOWS OsType = "WINDOWS"
    // A customization specification for a Linux guest operating system
     OsType_LINUX OsType = "LINUX"
)

func (o OsType) OsType() bool {
    switch o {
        case OsType_WINDOWS:
            return true
        case OsType_LINUX:
            return true
        default:
            return false
    }
}




// The ``Format`` enumeration class specifies the formats the customization specification can be exported to. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Format string

const (
    // JSON format. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Format_JSON Format = "JSON"
    // XML format. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Format_XML Format = "XML"
)

func (f Format) Format() bool {
    switch f {
        case Format_JSON:
            return true
        case Format_XML:
            return true
        default:
            return false
    }
}





// The ``Metadata`` class contains metadata i.e. name and description related to a customization specification. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Metadata struct {
    Description string
    Name string
}






// The ``CreateSpec`` class contains specification information and specification object that can be passed to the CustomizationSpecs#create method. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Spec guest.CustomizationSpec
    Description string
    Name string
}






// The ``Spec`` class contains the specification information and specification object. This is passed to the CustomizationSpecs#set method. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Spec struct {
    Fingerprint string
    Spec guest.CustomizationSpec
    Description string
    Name string
}






// The ``Info`` class describes a guest customization specification and the timestamp when it was last modified. This is returned by the CustomizationSpecs#get method. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    LastModified time.Time
    Spec Spec
}






// The ``FilterSpec`` class contains properties used to filter the results when listing guest customization specifications (see CustomizationSpecs#list). If multiple properties are specified, only guest customization specifications matching all of the properties match the filter.
 type FilterSpec struct {
    Names map[string]bool
    OSType *OsType
}






// The ``Summary`` class contains commonly used information about a guest customization specification.
 type Summary struct {
    Name string
    Description string
    OSType OsType
    LastModified time.Time
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.OS_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(OsType(OsType_WINDOWS))))
    paramsTypeMap["filter.names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, ""), reflect.TypeOf(map[string]bool{})))
    queryParams["filter.names"] = "names"
    queryParams["filter.OS_type"] = "os_type"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/guest/customization-specs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/guest/customization-specs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    fieldNameMap["name"] = "Name"
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
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    pathParams["name"] = "name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/guest/customization-specs/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["name"] = "Name"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(SpecBindingType)
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    pathParams["name"] = "name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/guest/customization-specs/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    pathParams["name"] = "name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vcenter/guest/customization-specs/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func exportInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    fields["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(Format(Format_JSON)))
    fieldNameMap["name"] = "Name"
    fieldNameMap["format"] = "Format"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportOutputType() bindings.BindingType {
    return bindings.NewStringType()
}

func exportRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(Format(Format_JSON)))
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    pathParams["name"] = "name"
    queryParams["format"] = "format"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/guest/customization-specs/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func importSpecificationInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["customization_spec"] = bindings.NewStringType()
    fieldNameMap["customization_spec"] = "CustomizationSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importSpecificationOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CreateSpecBindingType)
}

func importSpecificationRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["customization_spec"] = bindings.NewStringType()
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "customization_spec",
      "POST",
      "/vcenter/guest/customization-specs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func MetadataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.metadata",fields, reflect.TypeOf(Metadata{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(guest.CustomizationSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["fingerprint"] = bindings.NewStringType()
    fieldNameMap["fingerprint"] = "Fingerprint"
    fields["spec"] = bindings.NewReferenceType(guest.CustomizationSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["last_modified"] = bindings.NewDateTimeType()
    fieldNameMap["last_modified"] = "LastModified"
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["OS_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(OsType(OsType_WINDOWS))))
    fieldNameMap["OS_type"] = "OSType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.vcenter.guest.CustomizationSpec"}, "")
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["OS_type"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(OsType(OsType_WINDOWS)))
    fieldNameMap["OS_type"] = "OSType"
    fields["last_modified"] = bindings.NewDateTimeType()
    fieldNameMap["last_modified"] = "LastModified"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}



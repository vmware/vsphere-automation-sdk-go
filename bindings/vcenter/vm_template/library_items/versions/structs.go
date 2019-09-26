/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Versions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package versions

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for library item versions.
const Versions_RESOURCE_TYPE = "com.vmware.content.library.item.Version"



// The ``Summary`` class contains commonly used information about a version of a library item containing a virtual machine template.
 type Summary struct {
    Version string
    VmTemplate string
}






// The ``Info`` class contains information about a version of a library item containing a virtual machine template.
 type Info struct {
    VmTemplate string
}






// The ``RollbackSpec`` class defines the information required to rollback a virtual machine template library item to a previous version.
 type RollbackSpec struct {
    Message string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["version"] = "Version"
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
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func rollbackInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(RollbackSpecBindingType))
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["version"] = "Version"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rollbackOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
}

func rollbackRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["version"] = "Version"
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fieldNameMap["version"] = "Version"
    fields["vm_template"] = bindings.NewStringType()
    fieldNameMap["vm_template"] = "VmTemplate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_template"] = bindings.NewStringType()
    fieldNameMap["vm_template"] = "VmTemplate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func RollbackSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.rollback_spec",fields, reflect.TypeOf(RollbackSpec{}), fieldNameMap, validators)
}



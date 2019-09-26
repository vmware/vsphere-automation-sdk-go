/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CheckOuts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package checkOuts

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)




// The ``CheckOutSpec`` class defines the information required to check out a library item containing a virtual machine template.
 type CheckOutSpec struct {
    Name *string
    Placement *PlacementSpec
    PoweredOn *bool
}






// The ``PlacementSpec`` class contains information used to place a checked out virtual machine onto resources within the vCenter inventory. The specified compute resource should have access to the storage associated with the checked out virtual machine.
 type PlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``CheckInSpec`` class defines the information required to check in a virtual machine into a library item.
 type CheckInSpec struct {
    Message string
}






// The ``Summary`` class contains commonly used information about a checked out virtual machine.
 type Summary struct {
    Vm string
}






// The ``Info`` class contains information about a checked out virtual machine.
 type Info struct {
    Time time.Time
    User string
}









func checkOutInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutSpecBindingType))
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func checkOutRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func checkInInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckInSpecBindingType))
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkInOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
}

func checkInRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
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
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["vm"] = "Vm"
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["template_library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["template_library_item"] = "TemplateLibraryItem"
    fieldNameMap["vm"] = "Vm"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func CheckOutSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["powered_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["powered_on"] = "PoweredOn"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.check_out_spec",fields, reflect.TypeOf(CheckOutSpec{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.placement_spec",fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}

func CheckInSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.check_in_spec",fields, reflect.TypeOf(CheckInSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["time"] = bindings.NewDateTimeType()
    fieldNameMap["time"] = "Time"
    fields["user"] = bindings.NewStringType()
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



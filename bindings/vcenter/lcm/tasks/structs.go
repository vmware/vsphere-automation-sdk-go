/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tasks.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tasks

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/lcm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The type of operation to be executed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Operation string

const (
    // The install deployment operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_INSTALL_DEPLOY Operation = "INSTALL_DEPLOY"
    // The install precheck operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_INSTALL_PRECHECK Operation = "INSTALL_PRECHECK"
    // The upgrade deployment operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_UPGRADE_DEPLOY Operation = "UPGRADE_DEPLOY"
    // The upgrade precheck operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Operation_UPGRADE_PRECHECK Operation = "UPGRADE_PRECHECK"
)

func (o Operation) Operation() bool {
    switch o {
        case Operation_INSTALL_DEPLOY:
            return true
        case Operation_INSTALL_PRECHECK:
            return true
        case Operation_UPGRADE_DEPLOY:
            return true
        case Operation_UPGRADE_PRECHECK:
            return true
        default:
            return false
    }
}





// The data container of the task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Id string
    Task *string
    Type_ Operation
    Status lcm.TaskInfo
}









func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.task"}, "")
    fieldNameMap["task"] = "Task"
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.task"}, "")
    fieldNameMap["id"] = "Id"
    fields["task"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["task"] = "Task"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.tasks.operation", reflect.TypeOf(Operation(Operation_INSTALL_DEPLOY)))
    fieldNameMap["type"] = "Type_"
    fields["status"] = bindings.NewReferenceType(lcm.TaskInfoBindingType)
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.tasks.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package VM

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The {\\\\@Status} enumeration class defines he valid compliance status values for a virtual machine or virtual disk.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // The virtual machine or virtual disk is in compliance.
     Status_COMPLIANT Status = "COMPLIANT"
    // The virtual machine or virtual disk is in not in compliance.
     Status_NON_COMPLIANT Status = "NON_COMPLIANT"
    // Compliance status of the virtual machine or virtual disk is not known.
     Status_UNKNOWN_COMPLIANCE Status = "UNKNOWN_COMPLIANCE"
    // Compliance computation is not applicable for this virtual machine or disk because it does not have any storage requirement that apply to the object-based datastore on which the entity is placed.
     Status_NOT_APPLICABLE Status = "NOT_APPLICABLE"
    // Compliance status becomes out of date when the profile associated with the virtual machine or disk is edited and not applied. The compliance status will remain out of date until the latest policy is applied.
     Status_OUT_OF_DATE Status = "OUT_OF_DATE"
)

func (s Status) Status() bool {
    switch s {
        case Status_COMPLIANT:
            return true
        case Status_NON_COMPLIANT:
            return true
        case Status_UNKNOWN_COMPLIANCE:
            return true
        case Status_NOT_APPLICABLE:
            return true
        case Status_OUT_OF_DATE:
            return true
        default:
            return false
    }
}





// Provides the compliance details of a virtual machine and its associated entities which match the given compliance statuses.
 type Info struct {
    VmHome *Status
    Disks map[string]Status
}






// The ``FilterSpec`` class contains Status used to filter the results when listing virtual machines (see VM#list).
 type FilterSpec struct {
    Status map[Status]bool
    Vms map[string]bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewReferenceType(FilterSpecBindingType)
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"VirtualMachine"}, ""), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400,"UnableToAllocateResource": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(Status(Status_COMPLIANT))))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(Status(Status_COMPLIANT))),reflect.TypeOf(map[string]Status{}))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.VM.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.VM.status", reflect.TypeOf(Status(Status_COMPLIANT))), reflect.TypeOf(map[Status]bool{}))
    fieldNameMap["status"] = "Status"
    fields["vms"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.VM.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}



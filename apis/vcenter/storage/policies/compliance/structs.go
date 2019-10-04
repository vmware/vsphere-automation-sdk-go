/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Compliance.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package compliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// This enumeration defines the set of status values for a compliance operation.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // Entity is in compliance.
     Status_COMPLIANT Status = "COMPLIANT"
    // Entity is out of compliance.
     Status_NON_COMPLIANT Status = "NON_COMPLIANT"
    // Compliance status of the entity is not known.
     Status_UNKNOWN Status = "UNKNOWN"
    // Compliance computation is not applicable for this entity because it does not have any storage requirement that apply to the object-based datastore on which the entity is placed.
     Status_NOT_APPLICABLE Status = "NOT_APPLICABLE"
    // Compliance status becomes out of date when the profile associated with the entity is edited and not applied. The compliance status will remain out of date until the latest policy is applied to the entity.
     Status_OUT_OF_DATE Status = "OUT_OF_DATE"
)

func (s Status) Status() bool {
    switch s {
        case Status_COMPLIANT:
            return true
        case Status_NON_COMPLIANT:
            return true
        case Status_UNKNOWN:
            return true
        case Status_NOT_APPLICABLE:
            return true
        case Status_OUT_OF_DATE:
            return true
        default:
            return false
    }
}





// Provides the details of a virtual machine and its associated entities which match the given compliance statuses.
 type Summary struct {
    Vm string
    VmHome *Status
    Disks map[string]Status
}






// The ``FilterSpec`` class contains complianceStatus used to filter the results when listing entities (see Compliance#list).
 type FilterSpec struct {
    Status map[Status]bool
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Status(Status_COMPLIANT))))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Status(Status_COMPLIANT))),reflect.TypeOf(map[string]Status{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.storage.policies.compliance.status", reflect.TypeOf(Status(Status_COMPLIANT))), reflect.TypeOf(map[Status]bool{}))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.storage.policies.compliance.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}



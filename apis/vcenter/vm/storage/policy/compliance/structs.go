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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``Status`` enumeration class defines the storage compliance status of a virtual machine and its applicable entities.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // Entity is in compliance.
     Status_COMPLIANT Status = "COMPLIANT"
    // Entity is out of compliance.
     Status_NON_COMPLIANT Status = "NON_COMPLIANT"
    // Compliance status of the entity is not known.
     Status_UNKNOWN_COMPLIANCE Status = "UNKNOWN_COMPLIANCE"
    // Compliance computation is not applicable for this entity because it does not have any storage requirements that apply to the datastore on which it is placed.
     Status_NOT_APPLICABLE Status = "NOT_APPLICABLE"
    // The Compliance status becomes out-of-date when the profile associated with the entity is edited but not applied. The compliance status remains out-of-date until the edited policy is applied to the entity.
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





// The ``VmComplianceInfo`` class contains information about storage policy compliance associated with a virtual machine.
 type VmComplianceInfo struct {
    Status Status
    CheckTime time.Time
    Policy *string
    FailureCause []std.LocalizableMessage
}






// The ``Info`` class contains information about the storage policy compliance of a virtual machine, including information about it's home directory and/or it's virtual disks.
 type Info struct {
    OverallCompliance Status
    VmHome *VmComplianceInfo
    Disks map[string]VmComplianceInfo
}






// The ``CheckSpec`` class contains properties used to specify the entities on which the storage policy compliance check is to be invoked.
 type CheckSpec struct {
    VmHome bool
    Disks map[string]bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewOptionalType(bindings.NewReferenceType(InfoBindingType))
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
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func checkInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["check_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckSpecBindingType))
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["check_spec"] = "CheckSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutputType() bindings.BindingType {
    return bindings.NewOptionalType(bindings.NewReferenceType(InfoBindingType))
}

func checkRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func VmComplianceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.compliance.status", reflect.TypeOf(Status(Status_COMPLIANT)))
    fieldNameMap["status"] = "Status"
    fields["check_time"] = bindings.NewDateTimeType()
    fieldNameMap["check_time"] = "CheckTime"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    fields["failure_cause"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["failure_cause"] = "FailureCause"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.vm_compliance_info",fields, reflect.TypeOf(VmComplianceInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["overall_compliance"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.compliance.status", reflect.TypeOf(Status(Status_COMPLIANT)))
    fieldNameMap["overall_compliance"] = "OverallCompliance"
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewReferenceType(VmComplianceInfoBindingType))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(VmComplianceInfoBindingType),reflect.TypeOf(map[string]VmComplianceInfo{}))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewBooleanType()
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.compliance.check_spec",fields, reflect.TypeOf(CheckSpec{}), fieldNameMap, validators)
}



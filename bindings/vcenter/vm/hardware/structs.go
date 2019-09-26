/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Hardware.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Version`` enumeration class defines the valid virtual hardware versions for a virtual machine. See https://kb.vmware.com/s/article/1003746 (Virtual machine hardware versions (1003746)).
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Version string

const (
    // Hardware version 3, first supported in ESXi 2.5.
     Version_VMX_03 Version = "VMX_03"
    // Hardware version 4, first supported in ESXi 3.0.
     Version_VMX_04 Version = "VMX_04"
    // Hardware version 6, first supported in WS 6.0.
     Version_VMX_06 Version = "VMX_06"
    // Hardware version 7, first supported in ESXi 4.0.
     Version_VMX_07 Version = "VMX_07"
    // Hardware version 8, first supported in ESXi 5.0.
     Version_VMX_08 Version = "VMX_08"
    // Hardware version 9, first supported in ESXi 5.1.
     Version_VMX_09 Version = "VMX_09"
    // Hardware version 10, first supported in ESXi 5.5.
     Version_VMX_10 Version = "VMX_10"
    // Hardware version 11, first supported in ESXi 6.0.
     Version_VMX_11 Version = "VMX_11"
    // Hardware version 12, first supported in Workstation 12.0.
     Version_VMX_12 Version = "VMX_12"
    // Hardware version 13, first supported in ESXi 6.5.
     Version_VMX_13 Version = "VMX_13"
    // Hardware version 14, first supported in ESXi 6.7.
     Version_VMX_14 Version = "VMX_14"
    // Hardware version 15, first supported in ESXi 6.7.0 Update 2.
     Version_VMX_15 Version = "VMX_15"
    // Hardware version 16, first supported in Workstation 15.0.
     Version_VMX_16 Version = "VMX_16"
    // Hardware version 17, first supported in ESX 7.0.
     Version_VMX_17 Version = "VMX_17"
    // Future hardware version, not supported yet. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Version_VMX_FUTURE Version = "VMX_FUTURE"
)

func (v Version) Version() bool {
    switch v {
        case Version_VMX_03:
            return true
        case Version_VMX_04:
            return true
        case Version_VMX_06:
            return true
        case Version_VMX_07:
            return true
        case Version_VMX_08:
            return true
        case Version_VMX_09:
            return true
        case Version_VMX_10:
            return true
        case Version_VMX_11:
            return true
        case Version_VMX_12:
            return true
        case Version_VMX_13:
            return true
        case Version_VMX_14:
            return true
        case Version_VMX_15:
            return true
        case Version_VMX_16:
            return true
        case Version_VMX_17:
            return true
        case Version_VMX_FUTURE:
            return true
        default:
            return false
    }
}




// The ``UpgradePolicy`` enumeration class defines the valid virtual hardware upgrade policies for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UpgradePolicy string

const (
    // Do not upgrade the virtual machine when it is powered on.
     UpgradePolicy_NEVER UpgradePolicy = "NEVER"
    // Run scheduled upgrade when the virtual machine is powered on after a clean shutdown of the guest operating system.
     UpgradePolicy_AFTER_CLEAN_SHUTDOWN UpgradePolicy = "AFTER_CLEAN_SHUTDOWN"
    // Run scheduled upgrade when the virtual machine is powered on.
     UpgradePolicy_ALWAYS UpgradePolicy = "ALWAYS"
)

func (u UpgradePolicy) UpgradePolicy() bool {
    switch u {
        case UpgradePolicy_NEVER:
            return true
        case UpgradePolicy_AFTER_CLEAN_SHUTDOWN:
            return true
        case UpgradePolicy_ALWAYS:
            return true
        default:
            return false
    }
}




// The ``UpgradeStatus`` enumeration class defines the valid virtual hardware upgrade statuses for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UpgradeStatus string

const (
    // No scheduled upgrade has been attempted.
     UpgradeStatus_NONE UpgradeStatus = "NONE"
    // Upgrade is scheduled but has not yet been run.
     UpgradeStatus_PENDING UpgradeStatus = "PENDING"
    // The most recent scheduled upgrade was successful.
     UpgradeStatus_SUCCESS UpgradeStatus = "SUCCESS"
    // The most recent scheduled upgrade was not successful.
     UpgradeStatus_FAILED UpgradeStatus = "FAILED"
)

func (u UpgradeStatus) UpgradeStatus() bool {
    switch u {
        case UpgradeStatus_NONE:
            return true
        case UpgradeStatus_PENDING:
            return true
        case UpgradeStatus_SUCCESS:
            return true
        case UpgradeStatus_FAILED:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information related to the virtual hardware of a virtual machine.
 type Info struct {
    Version Version
    UpgradePolicy UpgradePolicy
    UpgradeVersion *Version
    UpgradeStatus UpgradeStatus
    UpgradeError *data.ErrorValue
}






// The ``UpdateSpec`` class describes the updates to virtual hardware settings of a virtual machine.
 type UpdateSpec struct {
    UpgradePolicy *UpgradePolicy
    UpgradeVersion *Version
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func upgradeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Version(Version_VMX_03))))
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Version(Version_VMX_03)))
    fieldNameMap["version"] = "Version"
    fields["upgrade_policy"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(UpgradePolicy(UpgradePolicy_NEVER)))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Version(Version_VMX_03))))
    fieldNameMap["upgrade_version"] = "UpgradeVersion"
    fields["upgrade_status"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_status", reflect.TypeOf(UpgradeStatus(UpgradeStatus_NONE)))
    fieldNameMap["upgrade_status"] = "UpgradeStatus"
    fields["upgrade_error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["upgrade_error"] = "UpgradeError"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("upgrade_policy",
        map[string][]bindings.FieldData {
            "AFTER_CLEAN_SHUTDOWN": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", true),
            },
            "ALWAYS": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", true),
            },
            "NEVER": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("upgrade_status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_error", true),
            },
            "NONE": []bindings.FieldData {},
            "PENDING": []bindings.FieldData {},
            "SUCCESS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["upgrade_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(UpgradePolicy(UpgradePolicy_NEVER))))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Version(Version_VMX_03))))
    fieldNameMap["upgrade_version"] = "UpgradeVersion"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("upgrade_policy",
        map[string][]bindings.FieldData {
            "AFTER_CLEAN_SHUTDOWN": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", false),
            },
            "ALWAYS": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", false),
            },
            "NEVER": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}



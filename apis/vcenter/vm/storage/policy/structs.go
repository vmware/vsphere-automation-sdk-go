/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``VmHomePolicySpec`` class provides a specification for the storage policy to be associated with the virtual machine home's directory.
 type VmHomePolicySpec struct {
    Type_ VmHomePolicySpec_PolicyType
    Policy *string
}




    
    // The ``PolicyType`` enumeration class defines the choices for how to specify the policy to be associated with the virtual machine home's directory.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type VmHomePolicySpec_PolicyType string

    const (
        // Use the specified policy (see VmHomePolicySpec#policy).
         VmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY VmHomePolicySpec_PolicyType = "USE_SPECIFIED_POLICY"
        // Use the default storage policy of the datastore.
         VmHomePolicySpec_PolicyType_USE_DEFAULT_POLICY VmHomePolicySpec_PolicyType = "USE_DEFAULT_POLICY"
    )

    func (p VmHomePolicySpec_PolicyType) VmHomePolicySpec_PolicyType() bool {
        switch p {
            case VmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY:
                return true
            case VmHomePolicySpec_PolicyType_USE_DEFAULT_POLICY:
                return true
            default:
                return false
        }
    }



// The ``DiskPolicySpec`` class provides a specification for the storage policy to be associated with the virtual disks.
 type DiskPolicySpec struct {
    Type_ DiskPolicySpec_PolicyType
    Policy *string
}




    
    // The ``DiskPolicySpec`` enumeration class defines the choices for how to specify the policy to be associated with a virtual disk.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DiskPolicySpec_PolicyType string

    const (
        // Use the specified policy (see DiskPolicySpec#policy).
         DiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY DiskPolicySpec_PolicyType = "USE_SPECIFIED_POLICY"
        // Use the default storage policy of the datastore.
         DiskPolicySpec_PolicyType_USE_DEFAULT_POLICY DiskPolicySpec_PolicyType = "USE_DEFAULT_POLICY"
    )

    func (p DiskPolicySpec_PolicyType) DiskPolicySpec_PolicyType() bool {
        switch p {
            case DiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY:
                return true
            case DiskPolicySpec_PolicyType_USE_DEFAULT_POLICY:
                return true
            default:
                return false
        }
    }



// The ``UpdateSpec`` class describes the updates to be made to the storage policies associated with the virtual machine home and/or its virtual disks.
 type UpdateSpec struct {
    VmHome *VmHomePolicySpec
    Disks map[string]DiskPolicySpec
}






// The ``Info`` class contains information about the storage policies associated with virtual machine's home directory and virtual hard disks.
 type Info struct {
    VmHome *string
    Disks map[string]string
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
       map[string]int{"Error": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500})
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
       map[string]int{"Error": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func VmHomePolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec.policy_type", reflect.TypeOf(VmHomePolicySpec_PolicyType(VmHomePolicySpec_PolicyType_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_DEFAULT_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.vm_home_policy_spec",fields, reflect.TypeOf(VmHomePolicySpec{}), fieldNameMap, validators)
}

func DiskPolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec.policy_type", reflect.TypeOf(DiskPolicySpec_PolicyType(DiskPolicySpec_PolicyType_USE_SPECIFIED_POLICY)))
    fieldNameMap["type"] = "Type_"
    fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "USE_SPECIFIED_POLICY": []bindings.FieldData {
                 bindings.NewFieldData("policy", true),
            },
            "USE_DEFAULT_POLICY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.disk_policy_spec",fields, reflect.TypeOf(DiskPolicySpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewReferenceType(VmHomePolicySpecBindingType))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(DiskPolicySpecBindingType),reflect.TypeOf(map[string]DiskPolicySpec{})))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_home"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, ""))
    fieldNameMap["vm_home"] = "VmHome"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["disks"] = "Disks"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.storage.policy.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



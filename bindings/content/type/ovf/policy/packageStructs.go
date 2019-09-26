/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.type.ovf.policy.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// Provide information of the membership of a particular storage policy group. 
//
// It is valid for disk, virtual machine or virtual machine collection.
type StoragePolicy struct {
    GroupId string
}






// Provide information of storage policy for a group of disks, virtual machines and/or virtual machine collections.
type StoragePolicyGroup struct {
    Id string
    Name string
    Description *string
}










func StoragePolicyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["group_id"] = bindings.NewStringType()
    fieldNameMap["group_id"] = "GroupId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.policy.storage_policy",fields, reflect.TypeOf(StoragePolicy{}), fieldNameMap, validators)
}

func StoragePolicyGroupBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.ovf.policy.storage_policy_group",fields, reflect.TypeOf(StoragePolicyGroup{}), fieldNameMap, validators)
}



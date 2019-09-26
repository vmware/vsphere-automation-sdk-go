/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.lcm.discovery.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package discovery

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``Info`` class contains information about a VMware product which is present in the customer Environemnt. The following information about the products are present: 
//
// * Name
// * Version
// * Deployments
// * Automatically Discovered or Manually Added
type Product struct {
    InstalledProduct string
    Name string
    Version string
    TargetVersion *string
    Deployments []string
    Auto bool
}










func ProductBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["installed_product"] = bindings.NewIdType([]string {"PRODUCT"}, "")
    fieldNameMap["installed_product"] = "InstalledProduct"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["target_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["target_version"] = "TargetVersion"
    fields["deployments"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["deployments"] = "Deployments"
    fields["auto"] = bindings.NewBooleanType()
    fieldNameMap["auto"] = "Auto"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.product",fields, reflect.TypeOf(Product{}), fieldNameMap, validators)
}



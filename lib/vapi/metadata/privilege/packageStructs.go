/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.metadata.privilege.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package privilege

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``ComponentData`` class contains the privilege information of the component along with its fingerprint.
type ComponentData struct {
    Info ComponentInfo
    Fingerprint string
}






// The ``ComponentInfo`` class contains the privilege information of a component element. 
//
//  For an explanation of privilege information contained within component elements, see Component.
type ComponentInfo struct {
    Packages map[string]PackageInfo
}






// The ``OperationInfo`` class contains privilege information of an operation element. 
//
//  For an explanation of containment within operation elements, see Operation.
type OperationInfo struct {
    Privileges []string
    PrivilegeInfo []PrivilegeInfo
}






// The ``PackageInfo`` class contains the privilege information of a package element. 
//
//  For an explanation of privilege information contained within package elements, see Package.
type PackageInfo struct {
    Privileges []string
    Services map[string]ServiceInfo
}






// The ``PrivilegeInfo`` class contains the privilege information for a parameter element in an operation element.
type PrivilegeInfo struct {
    PropertyPath string
    Privileges []string
}






// The ``ServiceInfo`` class contains privilege information of a service element. 
//
//  For an explanation of privilege information contained within service elements, see Service.
type ServiceInfo struct {
    Operations map[string]OperationInfo
}










func ComponentDataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewReferenceType(ComponentInfoBindingType)
    fieldNameMap["info"] = "Info"
    fields["fingerprint"] = bindings.NewStringType()
    fieldNameMap["fingerprint"] = "Fingerprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.component_data",fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType),reflect.TypeOf(map[string]PackageInfo{}))
    fieldNameMap["packages"] = "Packages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["privileges"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["privileges"] = "Privileges"
    fields["privilege_info"] = bindings.NewListType(bindings.NewReferenceType(PrivilegeInfoBindingType), reflect.TypeOf([]PrivilegeInfo{}))
    fieldNameMap["privilege_info"] = "PrivilegeInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.operation_info",fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["privileges"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["privileges"] = "Privileges"
    fields["services"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType),reflect.TypeOf(map[string]ServiceInfo{}))
    fieldNameMap["services"] = "Services"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.package_info",fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func PrivilegeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["property_path"] = bindings.NewStringType()
    fieldNameMap["property_path"] = "PropertyPath"
    fields["privileges"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["privileges"] = "Privileges"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.privilege_info",fields, reflect.TypeOf(PrivilegeInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType),reflect.TypeOf(map[string]OperationInfo{}))
    fieldNameMap["operations"] = "Operations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.privilege.service_info",fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}



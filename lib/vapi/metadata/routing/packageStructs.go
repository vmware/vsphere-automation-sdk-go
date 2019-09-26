/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.metadata.routing.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package routing

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// Routing information of the vAPI component along with its checksum
type ComponentData struct {
    Info ComponentInfo
    Fingerprint string
}






// Information about a vAPI component that contains routing information For an explanation of routing information within components, see Component
type ComponentInfo struct {
    Packages map[string]PackageInfo
}






// Information about a vAPI operation that contains routing information. For an explanation of containment within operations, see null
type OperationInfo struct {
    RoutingInfo RoutingInfo
}






// Information about a vAPI package containing routing information. 
//
//  For an explanation of routing information within packages, see Package
type PackageInfo struct {
    RoutingInfo RoutingInfo
    Services map[string]ServiceInfo
}






// Routing information
type RoutingInfo struct {
    RoutingPath string
    RoutingStrategy string
    OperationHints []string
    IdTypes map[string]string
}






// Information about a vAPI service that has routing information A service is said to contain routing information if any of its operations have routing information
type ServiceInfo struct {
    RoutingInfo RoutingInfo
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
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.component_data",fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType),reflect.TypeOf(map[string]PackageInfo{}))
    fieldNameMap["packages"] = "Packages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
    fieldNameMap["routing_info"] = "RoutingInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.operation_info",fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
    fieldNameMap["routing_info"] = "RoutingInfo"
    fields["services"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType),reflect.TypeOf(map[string]ServiceInfo{}))
    fieldNameMap["services"] = "Services"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.package_info",fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func RoutingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["routing_path"] = bindings.NewStringType()
    fieldNameMap["routing_path"] = "RoutingPath"
    fields["routing_strategy"] = bindings.NewStringType()
    fieldNameMap["routing_strategy"] = "RoutingStrategy"
    fields["operation_hints"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["operation_hints"] = "OperationHints"
    fields["id_types"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["id_types"] = "IdTypes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.routing_info",fields, reflect.TypeOf(RoutingInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["routing_info"] = bindings.NewReferenceType(RoutingInfoBindingType)
    fieldNameMap["routing_info"] = "RoutingInfo"
    fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType),reflect.TypeOf(map[string]OperationInfo{}))
    fieldNameMap["operations"] = "Operations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.routing.service_info",fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}



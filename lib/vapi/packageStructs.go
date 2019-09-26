/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vapi

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/authentication"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/cli"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``ComponentInfo`` class holds component metadata of the different metadata types for an API component. The class allows any combination of metadata types to be aggregated into one instance.
type ComponentInfo struct {
    Metamodel metamodel.ComponentInfo
    Cli *cli.ComponentInfo
    Authentication *authentication.ComponentInfo
    Routing *routing.ComponentInfo
    Privilege *privilege.ComponentInfo
}






// The ``MetadataInfo`` is a class which holds a map of the available metadata aggregated in a ComponentInfo class.
type MetadataInfo struct {
    Version string
    Metadata map[string]ComponentInfo
}










func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["metamodel"] = bindings.NewReferenceType(metamodel.ComponentInfoBindingType)
    fieldNameMap["metamodel"] = "Metamodel"
    fields["cli"] = bindings.NewOptionalType(bindings.NewReferenceType(cli.ComponentInfoBindingType))
    fieldNameMap["cli"] = "Cli"
    fields["authentication"] = bindings.NewOptionalType(bindings.NewReferenceType(authentication.ComponentInfoBindingType))
    fieldNameMap["authentication"] = "Authentication"
    fields["routing"] = bindings.NewOptionalType(bindings.NewReferenceType(routing.ComponentInfoBindingType))
    fieldNameMap["routing"] = "Routing"
    fields["privilege"] = bindings.NewOptionalType(bindings.NewReferenceType(privilege.ComponentInfoBindingType))
    fieldNameMap["privilege"] = "Privilege"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func MetadataInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["metadata"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.component"}, ""), bindings.NewReferenceType(ComponentInfoBindingType),reflect.TypeOf(map[string]ComponentInfo{}))
    fieldNameMap["metadata"] = "Metadata"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata_info",fields, reflect.TypeOf(MetadataInfo{}), fieldNameMap, validators)
}



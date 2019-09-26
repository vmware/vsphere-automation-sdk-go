/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.metadata.cli.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cli

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/cli/command"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/metadata/cli/namespace"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``ComponentInfo`` is an aggregated class for CLI commands and namespaces information.
type ComponentInfo struct {
    Namespaces []namespace.Info
    Commands []command.Info
}










func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespaces"] = bindings.NewListType(bindings.NewReferenceType(namespace.InfoBindingType), reflect.TypeOf([]namespace.Info{}))
    fieldNameMap["namespaces"] = "Namespaces"
    fields["commands"] = bindings.NewListType(bindings.NewReferenceType(command.InfoBindingType), reflect.TypeOf([]command.Info{}))
    fieldNameMap["commands"] = "Commands"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}



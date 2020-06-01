/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vapi.metadata.cli.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cli

import (
	"reflect"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``ComponentInfo`` is an aggregated class for CLI commands and namespaces information.
type ComponentInfo struct {
    // Information for all CLI namespaces of a component
	Namespaces []NamespaceInfo
    // Information for all CLI commands of a component
	Commands []CommandInfo
}




func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespaces"] = bindings.NewListType(bindings.NewReferenceType(NamespaceInfoBindingType), reflect.TypeOf([]NamespaceInfo{}))
	fieldNameMap["namespaces"] = "Namespaces"
	fields["commands"] = bindings.NewListType(bindings.NewReferenceType(CommandInfoBindingType), reflect.TypeOf([]CommandInfo{}))
	fieldNameMap["commands"] = "Commands"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}



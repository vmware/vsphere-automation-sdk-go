/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Version.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``VersionStruct`` class Structure representing appliance version information.
type VersionVersionStruct struct {
    // Appliance version.
	Version string
    // Appliance name.
	Product string
    // Appliance build number.
	Build string
    // Type of product. Same product can have different deployment options, which is represented by type.
	Type_ string
    // Summary of patch (empty string, if the appliance has not been patched)
	Summary string
    // Release date of patch (empty string, if the appliance has not been patched)
	Releasedate string
    // Display the date and time when this system was first installed. Value will not change on subsequent updates.
	InstallTime string
}



func versionGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(VersionVersionStructBindingType)
}

func versionGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500})
}


func VersionVersionStructBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	fields["build"] = bindings.NewStringType()
	fieldNameMap["build"] = "Build"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["summary"] = bindings.NewStringType()
	fieldNameMap["summary"] = "Summary"
	fields["releasedate"] = bindings.NewStringType()
	fieldNameMap["releasedate"] = "Releasedate"
	fields["install_time"] = bindings.NewStringType()
	fieldNameMap["install_time"] = "InstallTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.system.version.version_struct", fields, reflect.TypeOf(VersionVersionStruct{}), fieldNameMap, validators)
}


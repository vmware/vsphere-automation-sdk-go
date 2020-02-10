/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Versions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package library_items

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for library item versions. This constant field was added in vSphere API 6.9.1.
const Versions_RESOURCE_TYPE = "com.vmware.content.library.item.Version"


// The ``Summary`` class contains commonly used information about a version of a library item containing a virtual machine template. This class was added in vSphere API 6.9.1.
type VersionsSummary struct {
    // The version of the library item. This property was added in vSphere API 6.9.1.
	Version string
    // Identifier of the virtual machine template associated with the library item version. This property is the managed object identifier used to identify the virtual machine template in the vSphere Management (SOAP) API. This property was added in vSphere API 6.9.1.
	VmTemplate string
}

// The ``Info`` class contains information about a version of a library item containing a virtual machine template. This class was added in vSphere API 6.9.1.
type VersionsInfo struct {
    // Identifier of the virtual machine template associated with the library item version. This property is the managed object identifier used to identify the virtual machine template in the vSphere Management (SOAP) API. This property was added in vSphere API 6.9.1.
	VmTemplate string
}

// The ``RollbackSpec`` class defines the information required to rollback a virtual machine template library item to a previous version. This class was added in vSphere API 6.9.1.
type VersionsRollbackSpec struct {
    // Message describing the reason for the rollback. This property was added in vSphere API 6.9.1.
	Message string
}



func versionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(VersionsSummaryBindingType), reflect.TypeOf([]VersionsSummary{}))
}

func versionsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func versionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(VersionsInfoBindingType)
}

func versionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func versionsRollbackInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(VersionsRollbackSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionsRollbackOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
}

func versionsRollbackRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(VersionsRollbackSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func versionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func versionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["version"] = "Version"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func VersionsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
	fieldNameMap["version"] = "Version"
	fields["vm_template"] = bindings.NewStringType()
	fieldNameMap["vm_template"] = "VmTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.summary", fields, reflect.TypeOf(VersionsSummary{}), fieldNameMap, validators)
}

func VersionsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_template"] = bindings.NewStringType()
	fieldNameMap["vm_template"] = "VmTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.info", fields, reflect.TypeOf(VersionsInfo{}), fieldNameMap, validators)
}

func VersionsRollbackSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewStringType()
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.versions.rollback_spec", fields, reflect.TypeOf(VersionsRollbackSpec{}), fieldNameMap, validators)
}


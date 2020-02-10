/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CustomizationSpecs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)

// The resource type for a vCenter guest customization specification. This constant field was added in vSphere API 6.7.1.
const CustomizationSpecs_RESOURCE_TYPE = "com.vmware.vcenter.guest.CustomizationSpec"


// The ``OsType`` enumeration class defines the types of guest operating systems for which guest customization is supported. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CustomizationSpecsOsType string

const (
    // A customization specification for a Windows guest operating system. This constant field was added in vSphere API 6.7.1.
	CustomizationSpecsOsType_WINDOWS CustomizationSpecsOsType = "WINDOWS"
    // A customization specification for a Linux guest operating system. This constant field was added in vSphere API 6.7.1.
	CustomizationSpecsOsType_LINUX CustomizationSpecsOsType = "LINUX"
)

func (o CustomizationSpecsOsType) CustomizationSpecsOsType() bool {
	switch o {
	case CustomizationSpecsOsType_WINDOWS:
		return true
	case CustomizationSpecsOsType_LINUX:
		return true
	default:
		return false
	}
}


// The ``Format`` enumeration class specifies the formats the customization specification can be exported to. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CustomizationSpecsFormat string

const (
    // JSON format. This constant field was added in vSphere API 7.0.0.
	CustomizationSpecsFormat_JSON CustomizationSpecsFormat = "JSON"
    // XML format. This constant field was added in vSphere API 7.0.0.
	CustomizationSpecsFormat_XML CustomizationSpecsFormat = "XML"
)

func (f CustomizationSpecsFormat) CustomizationSpecsFormat() bool {
	switch f {
	case CustomizationSpecsFormat_JSON:
		return true
	case CustomizationSpecsFormat_XML:
		return true
	default:
		return false
	}
}


// The ``Metadata`` class contains metadata i.e. name and description related to a customization specification. This class was added in vSphere API 7.0.0.
type CustomizationSpecsMetadata struct {
    // Description of the specification. This property was added in vSphere API 7.0.0.
	Description string
    // Name of the specification. This property was added in vSphere API 7.0.0.
	Name string
}

// The ``CreateSpec`` class contains specification information and specification object that can be passed to the CustomizationSpecs#create method. This class was added in vSphere API 7.0.0.
type CustomizationSpecsCreateSpec struct {
    // The specification object. This property was added in vSphere API 7.0.0.
	Spec CustomizationSpec
    // Description of the specification. This property was added in vSphere API 7.0.0.
	Description string
    // Name of the specification. This property was added in vSphere API 7.0.0.
	Name string
}

// The ``Spec`` class contains the specification information and specification object. This is passed to the CustomizationSpecs#set method. This class was added in vSphere API 7.0.0.
type CustomizationSpecsSpec struct {
    // The fingerprint is a unique identifier for a given version of the configuration. Each change to the configuration will update this value. A client cannot change this value. If specified when updating a specification, the changes will only be applied if the current fingerprint matches the specified fingerprint. This field can be used to guard against updates that has happened between the specification content was read and until it is applied. This property was added in vSphere API 7.0.0.
	Fingerprint string
    // The specification object. This property was added in vSphere API 7.0.0.
	Spec CustomizationSpec
    // Description of the specification. This property was added in vSphere API 7.0.0.
	Description string
    // Name of the specification. This property was added in vSphere API 7.0.0.
	Name string
}

// The ``Info`` class describes a guest customization specification and the timestamp when it was last modified. This is returned by the CustomizationSpecs#get method. This class was added in vSphere API 7.0.0.
type CustomizationSpecsInfo struct {
    // Time when the specification was last modified. This property was added in vSphere API 7.0.0.
	LastModified time.Time
    // The Spec object including specification and metadata information. This property was added in vSphere API 7.0.0.
	Spec CustomizationSpecsSpec
}

// The ``FilterSpec`` class contains properties used to filter the results when listing guest customization specifications (see CustomizationSpecs#list). If multiple properties are specified, only guest customization specifications matching all of the properties match the filter. This class was added in vSphere API 6.7.1.
type CustomizationSpecsFilterSpec struct {
    // Names that guest customization specifications must have to match the filter (see CustomizationSpecsSummary#name). This property was added in vSphere API 6.7.1.
	Names map[string]bool
    // Guest operating system type that guest customization specifications must have to match the filter (see CustomizationSpecsSummary#osType). This property was added in vSphere API 6.7.1.
	OSType *CustomizationSpecsOsType
}

// The ``Summary`` class contains commonly used information about a guest customization specification. This class was added in vSphere API 6.7.1.
type CustomizationSpecsSummary struct {
    // Name of the guest customization specification. This property was added in vSphere API 6.7.1.
	Name string
    // Description of the guest customization specification. This property was added in vSphere API 6.7.1.
	Description string
    // Guest operating system type for which that this guest customization specification applies. This property was added in vSphere API 6.7.1.
	OSType CustomizationSpecsOsType
    // Date and tme when this guest customization specification was last modified. This property was added in vSphere API 6.7.1.
	LastModified time.Time
}



func customizationSpecsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CustomizationSpecsFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CustomizationSpecsSummaryBindingType), reflect.TypeOf([]CustomizationSpecsSummary{}))
}

func customizationSpecsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CustomizationSpecsFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
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
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
}

func customizationSpecsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CustomizationSpecsInfoBindingType)
}

func customizationSpecsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
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
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	fieldNameMap["name"] = "Name"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func customizationSpecsSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	fieldNameMap["name"] = "Name"
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
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func customizationSpecsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
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
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsExportInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(CustomizationSpecsFormat(CustomizationSpecsFormat_JSON)))
	fieldNameMap["name"] = "Name"
	fieldNameMap["format"] = "Format"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsExportOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func customizationSpecsExportRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fields["format"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.format", reflect.TypeOf(CustomizationSpecsFormat(CustomizationSpecsFormat_JSON)))
	fieldNameMap["name"] = "Name"
	fieldNameMap["format"] = "Format"
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
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func customizationSpecsImportSpecificationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["customization_spec"] = bindings.NewStringType()
	fieldNameMap["customization_spec"] = "CustomizationSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customizationSpecsImportSpecificationOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CustomizationSpecsCreateSpecBindingType)
}

func customizationSpecsImportSpecificationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["customization_spec"] = bindings.NewStringType()
	fieldNameMap["customization_spec"] = "CustomizationSpec"
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
		map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func CustomizationSpecsMetadataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.metadata", fields, reflect.TypeOf(CustomizationSpecsMetadata{}), fieldNameMap, validators)
}

func CustomizationSpecsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.create_spec", fields, reflect.TypeOf(CustomizationSpecsCreateSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fingerprint"] = bindings.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.spec", fields, reflect.TypeOf(CustomizationSpecsSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_modified"] = bindings.NewDateTimeType()
	fieldNameMap["last_modified"] = "LastModified"
	fields["spec"] = bindings.NewReferenceType(CustomizationSpecsSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.info", fields, reflect.TypeOf(CustomizationSpecsInfo{}), fieldNameMap, validators)
}

func CustomizationSpecsFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["OS_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(CustomizationSpecsOsType(CustomizationSpecsOsType_WINDOWS))))
	fieldNameMap["OS_type"] = "OSType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.filter_spec", fields, reflect.TypeOf(CustomizationSpecsFilterSpec{}), fieldNameMap, validators)
}

func CustomizationSpecsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.guest.CustomizationSpec"}, "")
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["OS_type"] = bindings.NewEnumType("com.vmware.vcenter.guest.customization_specs.os_type", reflect.TypeOf(CustomizationSpecsOsType(CustomizationSpecsOsType_WINDOWS)))
	fieldNameMap["OS_type"] = "OSType"
	fields["last_modified"] = bindings.NewDateTimeType()
	fieldNameMap["last_modified"] = "LastModified"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.guest.customization_specs.summary", fields, reflect.TypeOf(CustomizationSpecsSummary{}), fieldNameMap, validators)
}


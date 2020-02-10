/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Staged.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``Info`` class contains information about the staged update. This class was added in vSphere API 6.7.
type StagedInfo struct {
    // Is staging complete. This property was added in vSphere API 6.7.
	StagingComplete bool
    // Version in form of X.Y.Z.P. e.g. 6.5.1.5400. This property was added in vSphere API 6.7.
	Version string
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5". This property was added in vSphere API 6.7.
	Description std.LocalizableMessage
    // Update priority. This property was added in vSphere API 6.7.
	Priority CommonInfoPriority
    // Update severity. This property was added in vSphere API 6.7.
	Severity CommonInfoSeverity
    // Update category. This property was added in vSphere API 6.7.
	UpdateType CommonInfoCategory
    // Update release date. This property was added in vSphere API 6.7.
	ReleaseDate time.Time
    // Flag indicating whether reboot is required after update. This property was added in vSphere API 6.7.
	RebootRequired bool
    // Download Size of update in Megabytes. This property was added in vSphere API 6.7.
	Size int64
}



func stagedGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stagedGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(StagedInfoBindingType)
}

func stagedGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}

func stagedDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stagedDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func stagedDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func StagedInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["staging_complete"] = bindings.NewBooleanType()
	fieldNameMap["staging_complete"] = "StagingComplete"
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfoPriority(CommonInfoPriority_HIGH)))
	fieldNameMap["priority"] = "Priority"
	fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfoSeverity(CommonInfoSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfoCategory(CommonInfoCategory_SECURITY)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.staged.info", fields, reflect.TypeOf(StagedInfo{}), fieldNameMap, validators)
}


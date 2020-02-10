/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Policies.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package storage

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for vAPI metadata policy. This constant field was added in vSphere API 6.7.
const Policies_RESOURCE_TYPE = "com.vmware.vcenter.StoragePolicy"


// The ``FilterSpec`` class contains properties used to filter the results when listing the storage policies (see Policies#list). This class was added in vSphere API 6.7.
type PoliciesFilterSpec struct {
    // Identifiers of storage policies that can match the filter. This property was added in vSphere API 6.7.
	Policies map[string]bool
}

// The ``Summary`` class contains commonly used information about a storage policy. This class was added in vSphere API 6.7.
type PoliciesSummary struct {
    // Identifier of the storage policy. This property was added in vSphere API 6.7.
	Policy string
    // Name of the storage policy. This property was added in vSphere API 6.7.
	Name string
    // Description of the storage policy. This property was added in vSphere API 6.7.
	Description string
}

// The ``CompatibleDatastoreInfo`` class contains compatible datastore's information. This class was added in vSphere API 6.7.
type PoliciesCompatibleDatastoreInfo struct {
    // Identifier of the datastore. This property was added in vSphere API 6.7.
	Datastore string
}

// The ``CompatibilityInfo`` class contains info about a list of datastores compatible with a specific storage policy. This class was added in vSphere API 6.7.
type PoliciesCompatibilityInfo struct {
    // Info about a list of datastores compatible with a specific storage policy. This property was added in vSphere API 6.7.
	CompatibleDatastores []PoliciesCompatibleDatastoreInfo
}



func policiesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(PoliciesFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policiesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(PoliciesSummaryBindingType), reflect.TypeOf([]PoliciesSummary{}))
}

func policiesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(PoliciesFilterSpecBindingType))
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
		map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}

func policiesCheckCompatibilityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fields["datastores"] = bindings.NewSetType(bindings.NewIdType([]string{"Datastore"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["policy"] = "Policy"
	fieldNameMap["datastores"] = "Datastores"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policiesCheckCompatibilityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PoliciesCompatibilityInfoBindingType)
}

func policiesCheckCompatibilityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fields["datastores"] = bindings.NewSetType(bindings.NewIdType([]string{"Datastore"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["policy"] = "Policy"
	fieldNameMap["datastores"] = "Datastores"
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
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}


func PoliciesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["policies"] = "Policies"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.filter_spec", fields, reflect.TypeOf(PoliciesFilterSpec{}), fieldNameMap, validators)
}

func PoliciesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fieldNameMap["policy"] = "Policy"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.summary", fields, reflect.TypeOf(PoliciesSummary{}), fieldNameMap, validators)
}

func PoliciesCompatibleDatastoreInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewIdType([]string{"Datastore"}, "")
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatible_datastore_info", fields, reflect.TypeOf(PoliciesCompatibleDatastoreInfo{}), fieldNameMap, validators)
}

func PoliciesCompatibilityInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible_datastores"] = bindings.NewListType(bindings.NewReferenceType(PoliciesCompatibleDatastoreInfoBindingType), reflect.TypeOf([]PoliciesCompatibleDatastoreInfo{}))
	fieldNameMap["compatible_datastores"] = "CompatibleDatastores"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatibility_info", fields, reflect.TypeOf(PoliciesCompatibilityInfo{}), fieldNameMap, validators)
}


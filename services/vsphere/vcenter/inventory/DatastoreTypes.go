/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Datastore.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package inventory

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains information about a datastore.
type DatastoreInfo struct {
    // Type of the datastore.
	Type_ string
}



func datastoreFindInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastores"] = bindings.NewListType(bindings.NewIdType([]string{"Datastore"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["datastores"] = "Datastores"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datastoreFindOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"Datastore"}, ""), bindings.NewOptionalType(bindings.NewReferenceType(DatastoreInfoBindingType)),reflect.TypeOf(map[string]*DatastoreInfo{}))
}

func datastoreFindRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["datastores"] = bindings.NewListType(bindings.NewIdType([]string{"Datastore"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["datastores"] = "Datastores"
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
		map[string]int{"NotFound": 404})
}


func DatastoreInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	isv1 := bindings.NewIsOneOfValidator(
		"type",
		[]string{
			"Datastore",
			"StoragePod",
		},
	)
	validators = append(validators, isv1)
	return bindings.NewStructType("com.vmware.vcenter.inventory.datastore.info", fields, reflect.TypeOf(DatastoreInfo{}), fieldNameMap, validators)
}


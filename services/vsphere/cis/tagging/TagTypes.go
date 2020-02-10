/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Tag.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagging

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``CreateSpec`` class describes a tag. 
//
//  Use the Tag#create method to create a tag defined by the create specification.
type TagCreateSpec struct {
    // The display name of the tag. The name must be unique within its category.
	Name string
    // The description of the tag.
	Description string
    // The unique identifier of the parent category in which this tag will be created.
	CategoryId string
    // This property was added in vSphere API 6.7.
	TagId *string
}

// The ``UpdateSpec`` class describes the updates to be made to an existing tag. 
//
//  Use the Tag#update method to modify a tag. When you call the method, you specify the tag identifier. You obtain the tag identifier when you call the Tag#create method. You can also retrieve an identifier by using the Tag#list method.
type TagUpdateSpec struct {
    // The display name of the tag.
	Name *string
    // The description of the tag.
	Description *string
}



func tagCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["create_spec"] = bindings.NewReferenceType(TagCreateSpecBindingType)
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
}

func tagCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["create_spec"] = bindings.NewReferenceType(TagCreateSpecBindingType)
	fieldNameMap["create_spec"] = "CreateSpec"
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403})
}

func tagGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TagModelBindingType)
}

func tagGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}

func tagUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["update_spec"] = bindings.NewReferenceType(TagUpdateSpecBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["update_spec"] = bindings.NewReferenceType(TagUpdateSpecBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["update_spec"] = "UpdateSpec"
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403})
}

func tagDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}

func tagListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func tagListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{})
}

func tagListUsedTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["used_by_entity"] = "UsedByEntity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagListUsedTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func tagListUsedTagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["used_by_entity"] = "UsedByEntity"
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
		map[string]int{})
}

func tagListTagsForCategoryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, "")
	fieldNameMap["category_id"] = "CategoryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagListTagsForCategoryOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func tagListTagsForCategoryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["category_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, "")
	fieldNameMap["category_id"] = "CategoryId"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}

func tagAddToUsedByInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["used_by_entity"] = "UsedByEntity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAddToUsedByOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagAddToUsedByRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["used_by_entity"] = "UsedByEntity"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}

func tagRemoveFromUsedByInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["used_by_entity"] = "UsedByEntity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagRemoveFromUsedByOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagRemoveFromUsedByRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["used_by_entity"] = bindings.NewStringType()
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["used_by_entity"] = "UsedByEntity"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}

func tagRevokePropagatingPermissionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagRevokePropagatingPermissionsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagRevokePropagatingPermissionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func TagCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["category_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, "")
	fieldNameMap["category_id"] = "CategoryId"
	fields["tag_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""))
	fieldNameMap["tag_id"] = "TagId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag.create_spec", fields, reflect.TypeOf(TagCreateSpec{}), fieldNameMap, validators)
}

func TagUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag.update_spec", fields, reflect.TypeOf(TagUpdateSpec{}), fieldNameMap, validators)
}


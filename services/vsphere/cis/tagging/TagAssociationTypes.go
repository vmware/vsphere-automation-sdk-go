/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: TagAssociation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagging

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``BatchResult`` class describes the result of performing the same method on several tags or objects in a single invocation. This class was added in vSphere API 6.5.
type TagAssociationBatchResult struct {
    // This is true if the batch method completed without any errors. Otherwise it is false and all or some methods have failed. This property was added in vSphere API 6.5.
	Success bool
    // The array of error messages. This property was added in vSphere API 6.5.
	ErrorMessages []std.LocalizableMessage
}

// The ``TagToObjects`` class describes a tag and its related objects. Use the TagAssociation#listAttachedObjectsOnTags method to retrieve a array with each element containing a tag and the objects to which it is attached. This class was added in vSphere API 6.5.
type TagAssociationTagToObjects struct {
    // The identifier of the tag. This property was added in vSphere API 6.5.
	TagId string
    // The identifiers of the related objects. This property was added in vSphere API 6.5.
	ObjectIds []std.DynamicID
}

// The ``ObjectToTags`` class describes an object and its related tags. Use the TagAssociation#listAttachedTagsOnObjects method to retrieve a array with each element containing an object and the tags attached to it. This class was added in vSphere API 6.5.
type TagAssociationObjectToTags struct {
    // The identifier of the object. This property was added in vSphere API 6.5.
	ObjectId std.DynamicID
    // The identifiers of the related tags. This property was added in vSphere API 6.5.
	TagIds []string
}



func tagAssociationAttachInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationAttachOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagAssociationAttachRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_id"] = "ObjectId"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthorized": 403,"Unauthenticated": 401})
}

func tagAssociationAttachMultipleTagsToObjectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationAttachMultipleTagsToObjectOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TagAssociationBatchResultBindingType)
}

func tagAssociationAttachMultipleTagsToObjectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["tag_ids"] = "TagIds"
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
		map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}

func tagAssociationAttachTagToMultipleObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationAttachTagToMultipleObjectsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TagAssociationBatchResultBindingType)
}

func tagAssociationAttachTagToMultipleObjectsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_ids"] = "ObjectIds"
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

func tagAssociationDetachInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationDetachOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tagAssociationDetachRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_id"] = "ObjectId"
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

func tagAssociationDetachMultipleTagsFromObjectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationDetachMultipleTagsFromObjectOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TagAssociationBatchResultBindingType)
}

func tagAssociationDetachMultipleTagsFromObjectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["tag_ids"] = "TagIds"
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
		map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}

func tagAssociationDetachTagFromMultipleObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationDetachTagFromMultipleObjectsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TagAssociationBatchResultBindingType)
}

func tagAssociationDetachTagFromMultipleObjectsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["tag_id"] = "TagId"
	fieldNameMap["object_ids"] = "ObjectIds"
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

func tagAssociationListAttachedObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationListAttachedObjectsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
}

func tagAssociationListAttachedObjectsRestMetadata() protocol.OperationRestMetadata {
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

func tagAssociationListAttachedObjectsOnTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationListAttachedObjectsOnTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TagAssociationTagToObjectsBindingType), reflect.TypeOf([]TagAssociationTagToObjects{}))
}

func tagAssociationListAttachedObjectsOnTagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
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
		map[string]int{"Unauthenticated": 401})
}

func tagAssociationListAttachedTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationListAttachedTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func tagAssociationListAttachedTagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
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
		map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}

func tagAssociationListAttachedTagsOnObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationListAttachedTagsOnObjectsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TagAssociationObjectToTagsBindingType), reflect.TypeOf([]TagAssociationObjectToTags{}))
}

func tagAssociationListAttachedTagsOnObjectsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
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
		map[string]int{"Unauthenticated": 401})
}

func tagAssociationListAttachableTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tagAssociationListAttachableTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func tagAssociationListAttachableTagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
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
		map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}


func TagAssociationBatchResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["success"] = bindings.NewBooleanType()
	fieldNameMap["success"] = "Success"
	fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag_association.batch_result", fields, reflect.TypeOf(TagAssociationBatchResult{}), fieldNameMap, validators)
}

func TagAssociationTagToObjectsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag_association.tag_to_objects", fields, reflect.TypeOf(TagAssociationTagToObjects{}), fieldNameMap, validators)
}

func TagAssociationObjectToTagsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag_association.object_to_tags", fields, reflect.TypeOf(TagAssociationObjectToTags{}), fieldNameMap, validators)
}


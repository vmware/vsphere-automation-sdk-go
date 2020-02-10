/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.cis.tagging.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagging

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``CategoryModel`` class defines a category that is used to group one or more tags.
type CategoryModel struct {
    // The unique identifier of the category.
	Id string
    // The display name of the category.
	Name string
    // The description of the category.
	Description string
    // The associated cardinality (SINGLE, MULTIPLE) of the category.
	Cardinality CategoryModelCardinality
    // The types of objects that the tags in this category can be attached to. If the map with bool value is empty, then tags can be attached to all types of objects. This field works only for objects that reside in Inventory Service (IS). For non IS objects, this check is not performed today and hence a tag can be attached to any non IS object.
	AssociableTypes map[string]bool
    // The map with bool value of users that can use this category. To add users to this, you need to have the edit privilege on the category. Similarly, to unsubscribe from this category, you need the edit privilege on the category. You should not modify other users subscription from this map with bool value.
	UsedBy map[string]bool
}

// The ``Cardinality`` enumeration class defines the number of tags in a category that can be assigned to an object.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CategoryModelCardinality string

const (
    // An object can only be assigned one of the tags in this category. For example, if a category is "Operating System", then different tags of this category would be "Windows", "Linux", and so on. In this case a VM object can be assigned only one of these tags and hence the cardinality of the associated category here is single.
	CategoryModelCardinality_SINGLE CategoryModelCardinality = "SINGLE"
    // An object can be assigned several of the tags in this category. For example, if a category is "Server", then different tags of this category would be "AppServer", "DatabaseServer" and so on. In this case a VM object can be assigned more than one of the above tags and hence the cardinality of the associated category here is multiple.
	CategoryModelCardinality_MULTIPLE CategoryModelCardinality = "MULTIPLE"
)

func (c CategoryModelCardinality) CategoryModelCardinality() bool {
	switch c {
	case CategoryModelCardinality_SINGLE:
		return true
	case CategoryModelCardinality_MULTIPLE:
		return true
	default:
		return false
	}
}


// The ``TagModel`` class defines a tag that can be attached to vSphere objects.
type TagModel struct {
    // The unique identifier of the tag.
	Id string
    // The identifier of the parent category in which this tag will be created.
	CategoryId string
    // The display name of the tag.
	Name string
    // The description of the tag.
	Description string
    // The map with bool value of users that can use this tag. To add users to this, you need to have the edit privilege on the tag. Similarly, to unsubscribe from this tag, you need the edit privilege on the tag. You should not modify other users subscription from this map with bool value.
	UsedBy map[string]bool
}




func CategoryModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, "")
	fieldNameMap["id"] = "Id"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["cardinality"] = bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(CategoryModelCardinality(CategoryModelCardinality_SINGLE)))
	fieldNameMap["cardinality"] = "Cardinality"
	fields["associable_types"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["associable_types"] = "AssociableTypes"
	fields["used_by"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["used_by"] = "UsedBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.category_model", fields, reflect.TypeOf(CategoryModel{}), fieldNameMap, validators)
}

func TagModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["id"] = "Id"
	fields["category_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, "")
	fieldNameMap["category_id"] = "CategoryId"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["used_by"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["used_by"] = "UsedBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.tag_model", fields, reflect.TypeOf(TagModel{}), fieldNameMap, validators)
}



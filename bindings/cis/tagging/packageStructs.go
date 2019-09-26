/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
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
    Id string
    Name string
    Description string
    Cardinality CategoryModel_Cardinality
    AssociableTypes map[string]bool
    UsedBy map[string]bool
}




    
    // The ``Cardinality`` enumeration class defines the number of tags in a category that can be assigned to an object.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CategoryModel_Cardinality string

    const (
        // An object can only be assigned one of the tags in this category. For example, if a category is "Operating System", then different tags of this category would be "Windows", "Linux", and so on. In this case a VM object can be assigned only one of these tags and hence the cardinality of the associated category here is single.
         CategoryModel_Cardinality_SINGLE CategoryModel_Cardinality = "SINGLE"
        // An object can be assigned several of the tags in this category. For example, if a category is "Server", then different tags of this category would be "AppServer", "DatabaseServer" and so on. In this case a VM object can be assigned more than one of the above tags and hence the cardinality of the associated category here is multiple.
         CategoryModel_Cardinality_MULTIPLE CategoryModel_Cardinality = "MULTIPLE"
    )

    func (c CategoryModel_Cardinality) CategoryModel_Cardinality() bool {
        switch c {
            case CategoryModel_Cardinality_SINGLE:
                return true
            case CategoryModel_Cardinality_MULTIPLE:
                return true
            default:
                return false
        }
    }



// The ``TagModel`` class defines a tag that can be attached to vSphere objects.
type TagModel struct {
    Id string
    CategoryId string
    Name string
    Description string
    UsedBy map[string]bool
}










func CategoryModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["cardinality"] = bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(CategoryModel_Cardinality(CategoryModel_Cardinality_SINGLE)))
    fieldNameMap["cardinality"] = "Cardinality"
    fields["associable_types"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["associable_types"] = "AssociableTypes"
    fields["used_by"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["used_by"] = "UsedBy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.category_model",fields, reflect.TypeOf(CategoryModel{}), fieldNameMap, validators)
}

func TagModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fieldNameMap["id"] = "Id"
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["used_by"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["used_by"] = "UsedBy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.tag_model",fields, reflect.TypeOf(TagModel{}), fieldNameMap, validators)
}



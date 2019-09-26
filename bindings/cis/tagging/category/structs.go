/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Category.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package category

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``CreateSpec`` class is used to create a category. 
//
//  Use the Category#create method to create a category defined by the create specification.
 type CreateSpec struct {
    Name string
    Description string
    Cardinality tagging.CategoryModel_Cardinality
    AssociableTypes map[string]bool
    CategoryId *string
}






// The ``UpdateSpec`` class describes the updates to be made to an existing category. 
//
//  Use the Category#update method to modify a category. When you call the method, specify the category identifier. You obtain the category identifier when you call the Category#create method. You can also retrieve an identifier by using the Category#list method.
 type UpdateSpec struct {
    Name *string
    Description *string
    Cardinality *tagging.CategoryModel_Cardinality
    AssociableTypes map[string]bool
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["create_spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(tagging.CategoryModelBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["update_spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["update_spec"] = "UpdateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{})
}


func listUsedCategoriesInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listUsedCategoriesOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
}

func listUsedCategoriesRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{})
}


func addToUsedByInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addToUsedByOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func addToUsedByRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func removeFromUsedByInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func removeFromUsedByOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func removeFromUsedByRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func revokePropagatingPermissionsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func revokePropagatingPermissionsOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func revokePropagatingPermissionsRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["cardinality"] = bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(tagging.CategoryModel_Cardinality(tagging.CategoryModel_Cardinality_SINGLE)))
    fieldNameMap["cardinality"] = "Cardinality"
    fields["associable_types"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["associable_types"] = "AssociableTypes"
    fields["category_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""))
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.category.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["cardinality"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(tagging.CategoryModel_Cardinality(tagging.CategoryModel_Cardinality_SINGLE))))
    fieldNameMap["cardinality"] = "Cardinality"
    fields["associable_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["associable_types"] = "AssociableTypes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.category.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}



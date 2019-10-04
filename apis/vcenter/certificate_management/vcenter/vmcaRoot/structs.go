/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VmcaRoot.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vmcaRoot

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``CreateSpec`` contains information. to generate a Private Key and CSR.
 type CreateSpec struct {
    KeySize *int64
    CommonName *string
    Organization *string
    OrganizationUnit *string
    Locality *string
    StateOrProvince *string
    Country *string
    EmailAddress *string
    SubjectAltName []string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecBindingType))
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewVoidType()
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
       map[string]int{"Error": 500})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["key_size"] = "KeySize"
    fields["common_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["common_name"] = "CommonName"
    fields["organization"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["organization"] = "Organization"
    fields["organization_unit"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["organization_unit"] = "OrganizationUnit"
    fields["locality"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["locality"] = "Locality"
    fields["state_or_province"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["state_or_province"] = "StateOrProvince"
    fields["country"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["country"] = "Country"
    fields["email_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["email_address"] = "EmailAddress"
    fields["subject_alt_name"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["subject_alt_name"] = "SubjectAltName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.vmca_root.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}



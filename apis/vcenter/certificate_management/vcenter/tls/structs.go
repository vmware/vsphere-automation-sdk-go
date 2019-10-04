/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tls.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tls

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)




// The ``Info`` class contains information from a TLS certificate.
 type Info struct {
    Version int64
    SerialNumber string
    SignatureAlgorithm string
    IssuerDn string
    ValidFrom time.Time
    ValidTo time.Time
    SubjectDn string
    Thumbprint string
    IsCA bool
    PathLengthConstraint int64
    KeyUsage []string
    ExtendedKeyUsage []string
    SubjectAlternativeName []string
    AuthorityInformationAccessUri []string
    Cert string
}






// The ``Spec`` class contains information for a Certificate and Private Key.
 type Spec struct {
    Cert string
    Key *string
    RootCert *string
}






// The ``ReplaceSpec`` class contains information to generate a Private Key , CSR and hence VMCA signed machine SSL.
 type ReplaceSpec struct {
    KeySize *int64
    CommonName *string
    Organization string
    OrganizationUnit string
    Locality string
    StateOrProvince string
    Country string
    EmailAddress string
    SubjectAltName []string
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"AlreadyExists": 400,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
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
       map[string]int{"NotFound": 404,"Error": 500})
}


func renewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["duration"] = "Duration"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func renewOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func renewRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Error": 500})
}


func replaceVmcaSignedInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ReplaceSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func replaceVmcaSignedOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func replaceVmcaSignedRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIntegerType()
    fieldNameMap["version"] = "Version"
    fields["serial_number"] = bindings.NewStringType()
    fieldNameMap["serial_number"] = "SerialNumber"
    fields["signature_algorithm"] = bindings.NewStringType()
    fieldNameMap["signature_algorithm"] = "SignatureAlgorithm"
    fields["issuer_dn"] = bindings.NewStringType()
    fieldNameMap["issuer_dn"] = "IssuerDn"
    fields["valid_from"] = bindings.NewDateTimeType()
    fieldNameMap["valid_from"] = "ValidFrom"
    fields["valid_to"] = bindings.NewDateTimeType()
    fieldNameMap["valid_to"] = "ValidTo"
    fields["subject_dn"] = bindings.NewStringType()
    fieldNameMap["subject_dn"] = "SubjectDn"
    fields["thumbprint"] = bindings.NewStringType()
    fieldNameMap["thumbprint"] = "Thumbprint"
    fields["is_CA"] = bindings.NewBooleanType()
    fieldNameMap["is_CA"] = "IsCA"
    fields["path_length_constraint"] = bindings.NewIntegerType()
    fieldNameMap["path_length_constraint"] = "PathLengthConstraint"
    fields["key_usage"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["key_usage"] = "KeyUsage"
    fields["extended_key_usage"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["extended_key_usage"] = "ExtendedKeyUsage"
    fields["subject_alternative_name"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["subject_alternative_name"] = "SubjectAlternativeName"
    fields["authority_information_access_uri"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["authority_information_access_uri"] = "AuthorityInformationAccessUri"
    fields["cert"] = bindings.NewStringType()
    fieldNameMap["cert"] = "Cert"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cert"] = bindings.NewStringType()
    fieldNameMap["cert"] = "Cert"
    fields["key"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["key"] = "Key"
    fields["root_cert"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["root_cert"] = "RootCert"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}

func ReplaceSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["key_size"] = "KeySize"
    fields["common_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["common_name"] = "CommonName"
    fields["organization"] = bindings.NewStringType()
    fieldNameMap["organization"] = "Organization"
    fields["organization_unit"] = bindings.NewStringType()
    fieldNameMap["organization_unit"] = "OrganizationUnit"
    fields["locality"] = bindings.NewStringType()
    fieldNameMap["locality"] = "Locality"
    fields["state_or_province"] = bindings.NewStringType()
    fieldNameMap["state_or_province"] = "StateOrProvince"
    fields["country"] = bindings.NewStringType()
    fieldNameMap["country"] = "Country"
    fields["email_address"] = bindings.NewStringType()
    fieldNameMap["email_address"] = "EmailAddress"
    fields["subject_alt_name"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["subject_alt_name"] = "SubjectAltName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.replace_spec",fields, reflect.TypeOf(ReplaceSpec{}), fieldNameMap, validators)
}



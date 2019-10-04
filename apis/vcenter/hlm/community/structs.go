/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Community.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package community

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``AddSpec`` class is the specification used for connecting to another node and adding it to the local community. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type AddSpec struct {
    Hostname string
    Port *string
    Token string
    SslThumbprint *string
    SslVerify *bool
    AdminGroups map[string]bool
}






// The ``AddCheckSpec`` class is the specification used for specifying how to perform network checks within a community before adding a new peer. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type AddCheckSpec struct {
    AddSpec AddSpec
    CheckAll bool
}






// The ``CertificateInfo`` class contains information about the SSL certificate for a target node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CertificateInfo struct {
    SslThumbprint string
}






// The ``Info`` class has information about the community. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Members []Member
}






// The ``Member`` class has information about a community member. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Member struct {
    Hostname string
}






// The ``CheckInfo`` class contains information on the {\\\\@operation} check. Connectivity checks are validated pairwise. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CheckInfo struct {
    StatusFromDestination Status
    StatusFromSource Status
    Failed bool
}






// The ``Status`` class contains properties that are used to describe the outcome of a single direction connectivity check. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Status struct {
    Source string
    Destination string
    Latency int64
    Failed bool
    Errors []std.LocalizableMessage
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
       map[string]int{"Unauthorized": 403,"Error": 500})
}


func addInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(AddSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func addRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func removeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func removeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func removeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"Error": 500})
}


func checkInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(AddCheckSpecBindingType))
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(CheckInfoBindingType), reflect.TypeOf([]CheckInfo{}))
}

func checkRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}



func AddSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["port"] = "Port"
    fields["token"] = bindings.NewStringType()
    fieldNameMap["token"] = "Token"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["admin_groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["admin_groups"] = "AdminGroups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.add_spec",fields, reflect.TypeOf(AddSpec{}), fieldNameMap, validators)
}

func AddCheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["add_spec"] = bindings.NewReferenceType(AddSpecBindingType)
    fieldNameMap["add_spec"] = "AddSpec"
    fields["check_all"] = bindings.NewBooleanType()
    fieldNameMap["check_all"] = "CheckAll"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.add_check_spec",fields, reflect.TypeOf(AddCheckSpec{}), fieldNameMap, validators)
}

func CertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ssl_thumbprint"] = bindings.NewStringType()
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.certificate_info",fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["members"] = bindings.NewListType(bindings.NewReferenceType(MemberBindingType), reflect.TypeOf([]Member{}))
    fieldNameMap["members"] = "Members"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func MemberBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.member",fields, reflect.TypeOf(Member{}), fieldNameMap, validators)
}

func CheckInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status_from_destination"] = bindings.NewReferenceType(StatusBindingType)
    fieldNameMap["status_from_destination"] = "StatusFromDestination"
    fields["status_from_source"] = bindings.NewReferenceType(StatusBindingType)
    fieldNameMap["status_from_source"] = "StatusFromSource"
    fields["failed"] = bindings.NewBooleanType()
    fieldNameMap["failed"] = "Failed"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.check_info",fields, reflect.TypeOf(CheckInfo{}), fieldNameMap, validators)
}

func StatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewStringType()
    fieldNameMap["source"] = "Source"
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["latency"] = bindings.NewIntegerType()
    fieldNameMap["latency"] = "Latency"
    fields["failed"] = bindings.NewBooleanType()
    fieldNameMap["failed"] = "Failed"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.status",fields, reflect.TypeOf(Status{}), fieldNameMap, validators)
}



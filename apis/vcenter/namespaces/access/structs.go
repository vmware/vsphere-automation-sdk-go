/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Access.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package access

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Role`` enumeration class lists the default roles which can be associated with a subject on a domain on the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Role string

const (
    // This role allows modification of the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Role_EDIT Role = "EDIT"
    // This is a read-only role on the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Role_VIEW Role = "VIEW"
)

func (r Role) Role() bool {
    switch r {
        case Role_EDIT:
            return true
        case Role_VIEW:
            return true
        default:
            return false
    }
}




// The ``SubjectType`` enumeration class lists the types of subjects who can be associated with a ``Role`` on the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SubjectType string

const (
    // Single user. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SubjectType_USER SubjectType = "USER"
    // Group of users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     SubjectType_GROUP SubjectType = "GROUP"
)

func (s SubjectType) SubjectType() bool {
    switch s {
        case SubjectType_USER:
            return true
        case SubjectType_GROUP:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains the information about the access control of the subject on given domain on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    SubjectType SubjectType
    Role Role
}






// The ``CreateSpec`` class contains the specification required to create access control on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Type_ SubjectType
    Role Role
}






// The ``SetSpec`` class contains the specification required to set new access control on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SetSpec struct {
    Role Role
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["domain"] = bindings.NewStringType()
    fields["subject"] = bindings.NewStringType()
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["domain"] = "Domain"
    fieldNameMap["subject"] = "Subject"
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    paramsTypeMap["subject"] = bindings.NewStringType()
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["subject"] = bindings.NewStringType()
    pathParams["subject"] = "subject"
    pathParams["domain"] = "domain"
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["domain"] = bindings.NewStringType()
    fields["subject"] = bindings.NewStringType()
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["domain"] = "Domain"
    fieldNameMap["subject"] = "Subject"
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
    paramsTypeMap["subject"] = bindings.NewStringType()
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["subject"] = bindings.NewStringType()
    pathParams["subject"] = "subject"
    pathParams["domain"] = "domain"
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["domain"] = bindings.NewStringType()
    fields["subject"] = bindings.NewStringType()
    fields["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["domain"] = "Domain"
    fieldNameMap["subject"] = "Subject"
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    paramsTypeMap["subject"] = bindings.NewStringType()
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["subject"] = bindings.NewStringType()
    pathParams["subject"] = "subject"
    pathParams["domain"] = "domain"
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["domain"] = bindings.NewStringType()
    fields["subject"] = bindings.NewStringType()
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["domain"] = "Domain"
    fieldNameMap["subject"] = "Subject"
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
    paramsTypeMap["subject"] = bindings.NewStringType()
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["domain"] = bindings.NewStringType()
    paramsTypeMap["subject"] = bindings.NewStringType()
    pathParams["subject"] = "subject"
    pathParams["domain"] = "domain"
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subject_type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(SubjectType(SubjectType_USER)))
    fieldNameMap["subject_type"] = "SubjectType"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(Role(Role_EDIT)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.access.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(SubjectType(SubjectType_USER)))
    fieldNameMap["type"] = "Type_"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(Role(Role_EDIT)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.access.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(Role(Role_EDIT)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.access.set_spec",fields, reflect.TypeOf(SetSpec{}), fieldNameMap, validators)
}



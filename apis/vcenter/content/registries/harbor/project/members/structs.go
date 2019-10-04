/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Members.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package members

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``MemberType`` enumeration class describes the type of project member to be created for a Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type MemberType string

const (
    // Project member type for an individual user. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     MemberType_USER MemberType = "USER"
    // Project member type for a group of users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     MemberType_GROUP MemberType = "GROUP"
)

func (m MemberType) MemberType() bool {
    switch m {
        case MemberType_USER:
            return true
        case MemberType_GROUP:
            return true
        default:
            return false
    }
}




// The ``Role`` enumeration class describes roles available in a Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Role string

const (
    // This role allows image pull and push, repository and image deletion on an associated project,. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Role_MASTER Role = "MASTER"
    // This role allows image pull on an associated project. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Role_GUEST Role = "GUEST"
)

func (r Role) Role() bool {
    switch r {
        case Role_MASTER:
            return true
        case Role_GUEST:
            return true
        default:
            return false
    }
}




// The ``ConfigStatus`` enumeration class describes the status of reaching the desired configuration state for the Harbor project member. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigStatus string

const (
    // Harbor project member is being created or the configuration is being applied to the project member. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_PENDING ConfigStatus = "PENDING"
    // The configuration is being removed and Harbor project member is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_REMOVING ConfigStatus = "REMOVING"
    // Failed to create Harbor project member or apply the configuration to the project member, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_ERROR ConfigStatus = "ERROR"
    // Harbor project member is created or configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_READY ConfigStatus = "READY"
)

func (c ConfigStatus) ConfigStatus() bool {
    switch c {
        case ConfigStatus_PENDING:
            return true
        case ConfigStatus_REMOVING:
            return true
        case ConfigStatus_ERROR:
            return true
        case ConfigStatus_READY:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains basic Harbor project member information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Member string
    MemberName string
}






// The ``Info`` class contains detailed Harbor project member information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    MemberName string
    Type_ MemberType
    Role Role
    ConfigStatus ConfigStatus
    Message *std.LocalizableMessage
}






// The ``CreateSpec`` class contains the specification required to create project member for a specified Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    MemberName string
    Type_ MemberType
    Role Role
}






// The ``UpdateSpec`` class contains the specification required to set new project role for an existing project member. More fields could be added in future release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpdateSpec struct {
    Role *Role
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"AlreadyExists": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fields["member"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    fieldNameMap["member"] = "Member"
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
       map[string]int{"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fields["member"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    fieldNameMap["member"] = "Member"
    fieldNameMap["spec"] = "Spec"
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
       map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fields["member"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    fieldNameMap["member"] = "Member"
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
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["member"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
    fieldNameMap["member"] = "Member"
    fields["member_name"] = bindings.NewStringType()
    fieldNameMap["member_name"] = "MemberName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["member_name"] = bindings.NewStringType()
    fieldNameMap["member_name"] = "MemberName"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.member_type", reflect.TypeOf(MemberType(MemberType_USER)))
    fieldNameMap["type"] = "Type_"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(Role(Role_MASTER)))
    fieldNameMap["role"] = "Role"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_PENDING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_status",
        map[string][]bindings.FieldData {
            "ERROR": []bindings.FieldData {
                 bindings.NewFieldData("message", true),
            },
            "PENDING": []bindings.FieldData {},
            "REMOVING": []bindings.FieldData {},
            "READY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["member_name"] = bindings.NewStringType()
    fieldNameMap["member_name"] = "MemberName"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.member_type", reflect.TypeOf(MemberType(MemberType_USER)))
    fieldNameMap["type"] = "Type_"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(Role(Role_MASTER)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(Role(Role_MASTER))))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}



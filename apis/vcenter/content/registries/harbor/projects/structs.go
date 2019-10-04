/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Projects.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package projects

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/content/registries/harbor/project/members"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)



// The ``Scope`` enumeration class in a project defines access levels of the project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Scope string

const (
    // A Harbor project can be accessed by everyone. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Scope_PUBLIC Scope = "PUBLIC"
    // A Harbor project can only be accessed by assigned users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Scope_PRIVATE Scope = "PRIVATE"
)

func (s Scope) Scope() bool {
    switch s {
        case Scope_PUBLIC:
            return true
        case Scope_PRIVATE:
            return true
        default:
            return false
    }
}




// The ``ConfigStatus`` enumeration class describes the status of reaching the desired configuration state for the Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigStatus string

const (
    // Harbor project is being created or the configuration is being applied to the project. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_PENDING ConfigStatus = "PENDING"
    // The configuration is being removed and Harbor project is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_REMOVING ConfigStatus = "REMOVING"
    // Failed to create Harbor project or apply the configuration to the project, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_ERROR ConfigStatus = "ERROR"
    // Harbor project is created or configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
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





// The ``CreateSpec`` class defines the information required to create a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Name string
    Scope Scope
    Members []members.CreateSpec
}






// The ``Summary`` class contains basic information about a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Project string
    Name string
    Scope Scope
}






// The ``Info`` class contains detailed information about a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Name string
    ConfigStatus ConfigStatus
    Scope Scope
    CreationTime time.Time
    UpdateTime *time.Time
    AccessUrl *url.URL
    Message *std.LocalizableMessage
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
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
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
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
    fieldNameMap["registry"] = "Registry"
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
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func purgeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fieldNameMap["registry"] = "Registry"
    fieldNameMap["project"] = "Project"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func purgeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func purgeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(Scope(Scope_PUBLIC)))
    fieldNameMap["scope"] = "Scope"
    fields["members"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(members.CreateSpecBindingType), reflect.TypeOf([]members.CreateSpec{})))
    fieldNameMap["members"] = "Members"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["project"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
    fieldNameMap["project"] = "Project"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(Scope(Scope_PUBLIC)))
    fieldNameMap["scope"] = "Scope"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_PENDING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(Scope(Scope_PUBLIC)))
    fieldNameMap["scope"] = "Scope"
    fields["creation_time"] = bindings.NewDateTimeType()
    fieldNameMap["creation_time"] = "CreationTime"
    fields["update_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["update_time"] = "UpdateTime"
    fields["access_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["access_url"] = "AccessUrl"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_status",
        map[string][]bindings.FieldData {
            "READY": []bindings.FieldData {
                 bindings.NewFieldData("update_time", true),
                 bindings.NewFieldData("access_url", true),
            },
            "ERROR": []bindings.FieldData {
                 bindings.NewFieldData("message", true),
            },
            "PENDING": []bindings.FieldData {},
            "REMOVING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



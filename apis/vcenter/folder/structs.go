/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Folder.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package folder

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter folder
const Folder_RESOURCE_TYPE = "Folder"


// The ``Type`` enumeration class defines the type of a vCenter Server folder. The type of a folder determines what what kinds of children can be contained in the folder.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // A folder that can contain datacenters.
     Type_DATACENTER Type = "DATACENTER"
    // A folder that can contain datastores.
     Type_DATASTORE Type = "DATASTORE"
    // A folder that can contain compute resources (hosts and clusters).
     Type_HOST Type = "HOST"
    // A folder that can contain networkds.
     Type_NETWORK Type = "NETWORK"
    // A folder that can contain virtual machines.
     Type_VIRTUAL_MACHINE Type = "VIRTUAL_MACHINE"
)

func (t Type) Type() bool {
    switch t {
        case Type_DATACENTER:
            return true
        case Type_DATASTORE:
            return true
        case Type_HOST:
            return true
        case Type_NETWORK:
            return true
        case Type_VIRTUAL_MACHINE:
            return true
        default:
            return false
    }
}





// The ``FilterSpec`` class contains properties used to filter the results when listing folders (see Folder#list). If multiple properties are specified, only folders matching all of the properties match the filter.
 type FilterSpec struct {
    Folders map[string]bool
    Names map[string]bool
    Type_ *Type
    ParentFolders map[string]bool
    Datacenters map[string]bool
}






// The ``Summary`` class contains commonly used information about a folder.
 type Summary struct {
    Folder string
    Name string
    Type_ Type
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.folder.type", reflect.TypeOf(Type(Type_DATACENTER))))
    fieldNameMap["type"] = "Type_"
    fields["parent_folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["parent_folders"] = "ParentFolders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.folder.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewIdType([]string {"Folder"}, "")
    fieldNameMap["folder"] = "Folder"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.folder.type", reflect.TypeOf(Type(Type_DATACENTER)))
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.folder.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}



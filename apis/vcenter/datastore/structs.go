/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Datastore.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package datastore

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter datastore
const Datastore_RESOURCE_TYPE = "Datastore"


// The ``Type`` enumeration class defines the supported types of vCenter datastores.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // VMware File System (ESX Server only).
     Type_VMFS Type = "VMFS"
    // Network file system v3 (linux & esx servers only).
     Type_NFS Type = "NFS"
    // Network file system v4.1 (linux & esx servers only).
     Type_NFS41 Type = "NFS41"
    // Common Internet File System.
     Type_CIFS Type = "CIFS"
    // Virtual SAN (ESX Server only).
     Type_VSAN Type = "VSAN"
    // Flash Read Cache (ESX Server only).
     Type_VFFS Type = "VFFS"
    // vSphere Virtual Volume (ESX Server only).
     Type_VVOL Type = "VVOL"
)

func (t Type) Type() bool {
    switch t {
        case Type_VMFS:
            return true
        case Type_NFS:
            return true
        case Type_NFS41:
            return true
        case Type_CIFS:
            return true
        case Type_VSAN:
            return true
        case Type_VFFS:
            return true
        case Type_VVOL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a datastore.
 type Info struct {
    Name string
    Type_ Type
    Accessible bool
    FreeSpace *int64
    MultipleHostAccess bool
    ThinProvisioningSupported bool
}






// The ``FilterSpec`` class contains properties used to filter the results when listing datastores (see Datastore#list). If multiple properties are specified, only datastores matching all of the properties match the filter.
 type FilterSpec struct {
    Datastores map[string]bool
    Names map[string]bool
    Types map[Type]bool
    Folders map[string]bool
    Datacenters map[string]bool
}






// The ``Summary`` class contains commonly used information about a datastore.
 type Summary struct {
    Datastore string
    Name string
    Type_ Type
    FreeSpace *int64
    Capacity *int64
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
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
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Type(Type_VMFS)))
    fieldNameMap["type"] = "Type_"
    fields["accessible"] = bindings.NewBooleanType()
    fieldNameMap["accessible"] = "Accessible"
    fields["free_space"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["free_space"] = "FreeSpace"
    fields["multiple_host_access"] = bindings.NewBooleanType()
    fieldNameMap["multiple_host_access"] = "MultipleHostAccess"
    fields["thin_provisioning_supported"] = bindings.NewBooleanType()
    fieldNameMap["thin_provisioning_supported"] = "ThinProvisioningSupported"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastores"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datastore"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datastores"] = "Datastores"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Type(Type_VMFS))), reflect.TypeOf(map[Type]bool{})))
    fieldNameMap["types"] = "Types"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Type(Type_VMFS)))
    fieldNameMap["type"] = "Type_"
    fields["free_space"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["free_space"] = "FreeSpace"
    fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}



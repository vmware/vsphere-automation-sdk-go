/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Subscriptions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package subscriptions

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/content/library"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for Subscription resource.
const Subscriptions_RESOURCE_TYPE = "com.vmware.content.library.Subscriptions"


// The ``Location`` enumeration class defines the location of subscribed library relative to the published library.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Location string

const (
    // The subscribed library belongs to the same vCenter instance as the published library.
     Location_LOCAL Location = "LOCAL"
    // The subscribed library belongs to a different vCenter instance than the published library.
     Location_REMOTE Location = "REMOTE"
)

func (l Location) Location() bool {
    switch l {
        case Location_LOCAL:
            return true
        case Location_REMOTE:
            return true
        default:
            return false
    }
}





// The ``CreateSpecNewSubscribedLibrary`` class defines the information required to create a new subscribed library.
 type CreateSpecNewSubscribedLibrary struct {
    Name string
    Description *string
    StorageBackings []library.StorageBacking
    AutomaticSyncEnabled bool
    OnDemand bool
}






// The ``CreateSpecVcenter`` class defines information about the vCenter Server instance where the subscribed library associated with the subscription exists or will be created.
 type CreateSpecVcenter struct {
    Hostname string
    HttpsPort *int64
}






// The ``CreateSpecPlacement`` class defines the placement information for the subscribed library's virtual machine template library items. Storage location of the virtual machine template items is defined by the subscribed library's storage backing. This placement information needs to be compatible with the subscribed library's storage backing. The ``CreateSpecPlacement`` class is only applicable for the virtual machine template library items of the subscribed library.
 type CreateSpecPlacement struct {
    Folder *string
    Cluster *string
    ResourcePool *string
    Host *string
    Network *string
}






// The ``CreateSpecSubscribedLibrary`` class defines the subscribed library information used to create the subscription.
 type CreateSpecSubscribedLibrary struct {
    Target CreateSpecSubscribedLibrary_Target
    NewSubscribedLibrary *CreateSpecNewSubscribedLibrary
    SubscribedLibrary *string
    Location Location
    Vcenter *CreateSpecVcenter
    Placement *CreateSpecPlacement
}




    
    // The ``Target`` enumeration class defines the options related to the target subscribed library which will be associated with the subscription.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CreateSpecSubscribedLibrary_Target string

    const (
        // Create a new subscribed library.
         CreateSpecSubscribedLibrary_Target_CREATE_NEW CreateSpecSubscribedLibrary_Target = "CREATE_NEW"
        // Use the specified existing subscribed library.
         CreateSpecSubscribedLibrary_Target_USE_EXISTING CreateSpecSubscribedLibrary_Target = "USE_EXISTING"
    )

    func (t CreateSpecSubscribedLibrary_Target) CreateSpecSubscribedLibrary_Target() bool {
        switch t {
            case CreateSpecSubscribedLibrary_Target_CREATE_NEW:
                return true
            case CreateSpecSubscribedLibrary_Target_USE_EXISTING:
                return true
            default:
                return false
        }
    }



// The ``CreateSpec`` class defines the information required to create a new subscription of the published library.
 type CreateSpec struct {
    SubscribedLibrary CreateSpecSubscribedLibrary
}






// The ``Summary`` class contains commonly used information about the subscription.
 type Summary struct {
    Subscription string
    SubscribedLibrary string
    SubscribedLibraryName string
    SubscribedLibraryVcenterHostname *string
}






// The ``UpdateSpecVcenter`` class defines information about the vCenter Server instance where the subscribed library associated with the subscription exists. The ``UpdateSpecVcenter`` class is only applicable to subscribed library which exists on remote vCenter Server instance.
 type UpdateSpecVcenter struct {
    Hostname *string
    HttpsPort *int64
}






// The ``UpdateSpecPlacement`` class defines the placement information for the subscribed library's virtual machine template library items. Storage location of the virtual machine template items is defined by the subscribed library's storage backing. This placement information needs to be compatible with the subscribed library's storage backing. The ``UpdateSpecPlacement`` class is only applicable for the newly published virtual machine template library items of the subscribed library. Existing items will not be moved.
 type UpdateSpecPlacement struct {
    Folder *string
    Cluster *string
    ResourcePool *string
    Host *string
    Network *string
}






// The ``UpdateSpec`` class defines information required to update the subscription.
 type UpdateSpec struct {
    SubscribedLibraryVcenter *UpdateSpecVcenter
    SubscribedLibraryPlacement *UpdateSpecPlacement
}






// The ``VcenterInfo`` class contains information about the vCenter Server instance where the subscribed library associated with the subscription exists.
 type VcenterInfo struct {
    Hostname string
    HttpsPort *int64
    ServerGuid string
}






// The ``PlacementInfo`` class contains the placement information for the subscribed library's virtual machine template library items. The ``PlacementInfo`` class is only applicable for the virtual machine template library items of the subscribed library.
 type PlacementInfo struct {
    Folder *string
    Cluster *string
    ResourcePool *string
    Host *string
    Network *string
}






// The ``Info`` class contains information about the subscription.
 type Info struct {
    SubscribedLibrary string
    SubscribedLibraryName string
    SubscribedLibraryLocation Location
    SubscribedLibraryVcenter *VcenterInfo
    SubscribedLibraryPlacement PlacementInfo
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["library"] = "Library"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fields["subscription"] = bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
    fieldNameMap["library"] = "Library"
    fieldNameMap["subscription"] = "Subscription"
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
       map[string]int{"Error": 500,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library"] = "Library"
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
       map[string]int{"Error": 500,"InvalidElementType": 400,"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fields["subscription"] = bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["library"] = "Library"
    fieldNameMap["subscription"] = "Subscription"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fields["subscription"] = bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
    fieldNameMap["library"] = "Library"
    fieldNameMap["subscription"] = "Subscription"
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
       map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func CreateSpecNewSubscribedLibraryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["storage_backings"] = bindings.NewListType(bindings.NewReferenceType(library.StorageBackingBindingType), reflect.TypeOf([]library.StorageBacking{}))
    fieldNameMap["storage_backings"] = "StorageBackings"
    fields["automatic_sync_enabled"] = bindings.NewBooleanType()
    fieldNameMap["automatic_sync_enabled"] = "AutomaticSyncEnabled"
    fields["on_demand"] = bindings.NewBooleanType()
    fieldNameMap["on_demand"] = "OnDemand"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_new_subscribed_library",fields, reflect.TypeOf(CreateSpecNewSubscribedLibrary{}), fieldNameMap, validators)
}

func CreateSpecVcenterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_vcenter",fields, reflect.TypeOf(CreateSpecVcenter{}), fieldNameMap, validators)
}

func CreateSpecPlacementBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder:VCenter"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource:VCenter"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool:VCenter"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem:VCenter"}, ""))
    fieldNameMap["host"] = "Host"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["network"] = "Network"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_placement",fields, reflect.TypeOf(CreateSpecPlacement{}), fieldNameMap, validators)
}

func CreateSpecSubscribedLibraryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["target"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.create_spec_subscribed_library.target", reflect.TypeOf(CreateSpecSubscribedLibrary_Target(CreateSpecSubscribedLibrary_Target_CREATE_NEW)))
    fieldNameMap["target"] = "Target"
    fields["new_subscribed_library"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecNewSubscribedLibraryBindingType))
    fieldNameMap["new_subscribed_library"] = "NewSubscribedLibrary"
    fields["subscribed_library"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["subscribed_library"] = "SubscribedLibrary"
    fields["location"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.location", reflect.TypeOf(Location(Location_LOCAL)))
    fieldNameMap["location"] = "Location"
    fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecVcenterBindingType))
    fieldNameMap["vcenter"] = "Vcenter"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(CreateSpecPlacementBindingType))
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("target",
        map[string][]bindings.FieldData {
            "CREATE_NEW": []bindings.FieldData {
                 bindings.NewFieldData("new_subscribed_library", true),
            },
            "USE_EXISTING": []bindings.FieldData {
                 bindings.NewFieldData("subscribed_library", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("location",
        map[string][]bindings.FieldData {
            "REMOTE": []bindings.FieldData {
                 bindings.NewFieldData("vcenter", true),
            },
            "LOCAL": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_subscribed_library",fields, reflect.TypeOf(CreateSpecSubscribedLibrary{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscribed_library"] = bindings.NewReferenceType(CreateSpecSubscribedLibraryBindingType)
    fieldNameMap["subscribed_library"] = "SubscribedLibrary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscription"] = bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
    fieldNameMap["subscription"] = "Subscription"
    fields["subscribed_library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["subscribed_library"] = "SubscribedLibrary"
    fields["subscribed_library_name"] = bindings.NewStringType()
    fieldNameMap["subscribed_library_name"] = "SubscribedLibraryName"
    fields["subscribed_library_vcenter_hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["subscribed_library_vcenter_hostname"] = "SubscribedLibraryVcenterHostname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func UpdateSpecVcenterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["hostname"] = "Hostname"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec_vcenter",fields, reflect.TypeOf(UpdateSpecVcenter{}), fieldNameMap, validators)
}

func UpdateSpecPlacementBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder:VCenter"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource:VCenter"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool:VCenter"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem:VCenter"}, ""))
    fieldNameMap["host"] = "Host"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["network"] = "Network"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec_placement",fields, reflect.TypeOf(UpdateSpecPlacement{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscribed_library_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(UpdateSpecVcenterBindingType))
    fieldNameMap["subscribed_library_vcenter"] = "SubscribedLibraryVcenter"
    fields["subscribed_library_placement"] = bindings.NewOptionalType(bindings.NewReferenceType(UpdateSpecPlacementBindingType))
    fieldNameMap["subscribed_library_placement"] = "SubscribedLibraryPlacement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func VcenterInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["server_guid"] = bindings.NewStringType()
    fieldNameMap["server_guid"] = "ServerGuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.vcenter_info",fields, reflect.TypeOf(VcenterInfo{}), fieldNameMap, validators)
}

func PlacementInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder:VCenter"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource:VCenter"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool:VCenter"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem:VCenter"}, ""))
    fieldNameMap["host"] = "Host"
    fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["network"] = "Network"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscriptions.placement_info",fields, reflect.TypeOf(PlacementInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscribed_library"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["subscribed_library"] = "SubscribedLibrary"
    fields["subscribed_library_name"] = bindings.NewStringType()
    fieldNameMap["subscribed_library_name"] = "SubscribedLibraryName"
    fields["subscribed_library_location"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.location", reflect.TypeOf(Location(Location_LOCAL)))
    fieldNameMap["subscribed_library_location"] = "SubscribedLibraryLocation"
    fields["subscribed_library_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(VcenterInfoBindingType))
    fieldNameMap["subscribed_library_vcenter"] = "SubscribedLibraryVcenter"
    fields["subscribed_library_placement"] = bindings.NewReferenceType(PlacementInfoBindingType)
    fieldNameMap["subscribed_library_placement"] = "SubscribedLibraryPlacement"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("subscribed_library_location",
        map[string][]bindings.FieldData {
            "REMOTE": []bindings.FieldData {
                 bindings.NewFieldData("subscribed_library_vcenter", true),
            },
            "LOCAL": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.subscriptions.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



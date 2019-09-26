/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package content

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "time"
)



// The ``ConfigurationModel`` class defines the global settings of the Content Library Service.
type ConfigurationModel struct {
    AutomaticSyncEnabled *bool
    AutomaticSyncStartHour *int64
    AutomaticSyncStopHour *int64
    MaximumConcurrentItemSyncs *int64
}






// The LibraryModel class represents a Content Library resource model. 
//
//  A ``LibraryModel`` is a container for a set of items which represent a usable set of files. The Content Library Service allows for multiple libraries to be created with separate authorization and sharing policies. 
//
//  Each ``LibraryModel`` is a container for a set of library.ItemModel instances. Each item is a logical object in a library, which may have multiple files. 
//
//  A ``LibraryModel`` may be local or subscribed. A local library has its source of truth about items within this Content Library Service. Items may be added to or removed from the library. A local library may also be private or published. When published, the library is exposed by a network endpoint and can be used by another Content Library Service for synchronization. A private local library cannot be used for synchronization. 
//
//  A subscribed library is a library which gets its source of truth from another library that may be across a network in another Content Library Service. A subscribed library may have a different name and metadata from the library to which it subscribes, but the set of library items is always the same as those in the source library. Library items cannot be manually added to a subscribed library -- they can only be added by adding new items to the source library.
type LibraryModel struct {
    Id *string
    CreationTime *time.Time
    Description *string
    LastModifiedTime *time.Time
    LastSyncTime *time.Time
    Name *string
    StorageBackings []library.StorageBacking
    Type_ *LibraryModel_LibraryType
    OptimizationInfo *library.OptimizationInfo
    Version *string
    PublishInfo *library.PublishInfo
    SubscriptionInfo *library.SubscriptionInfo
    ServerGuid *string
}




    
    // The ``LibraryType`` enumeration class defines the type of a LibraryModel. 
    //
    //  The type of a library can be used to determine which additional services can be performed with a library.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type LibraryModel_LibraryType string

    const (
        // The library contents are defined and stored by the local Content Library Service installation. 
        //
        //  A local library can be retrieved and managed via the LocalLibrary.
         LibraryModel_LibraryType_LOCAL LibraryModel_LibraryType = "LOCAL"
        // The library synchronizes its items and content from another published library. 
        //
        //  A subscribed library can be retrieved and managed via the SubscribedLibrary.
         LibraryModel_LibraryType_SUBSCRIBED LibraryModel_LibraryType = "SUBSCRIBED"
    )

    func (l LibraryModel_LibraryType) LibraryModel_LibraryType() bool {
        switch l {
            case LibraryModel_LibraryType_LOCAL:
                return true
            case LibraryModel_LibraryType_SUBSCRIBED:
                return true
            default:
                return false
        }
    }







func ConfigurationModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["automatic_sync_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["automatic_sync_enabled"] = "AutomaticSyncEnabled"
    fields["automatic_sync_start_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["automatic_sync_start_hour"] = "AutomaticSyncStartHour"
    fields["automatic_sync_stop_hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["automatic_sync_stop_hour"] = "AutomaticSyncStopHour"
    fields["maximum_concurrent_item_syncs"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["maximum_concurrent_item_syncs"] = "MaximumConcurrentItemSyncs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.configuration_model",fields, reflect.TypeOf(ConfigurationModel{}), fieldNameMap, validators)
}

func LibraryModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["id"] = "Id"
    fields["creation_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["creation_time"] = "CreationTime"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["last_modified_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_modified_time"] = "LastModifiedTime"
    fields["last_sync_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_sync_time"] = "LastSyncTime"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["storage_backings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(library.StorageBackingBindingType), reflect.TypeOf([]library.StorageBacking{})))
    fieldNameMap["storage_backings"] = "StorageBackings"
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library_model.library_type", reflect.TypeOf(LibraryModel_LibraryType(LibraryModel_LibraryType_LOCAL))))
    fieldNameMap["type"] = "Type_"
    fields["optimization_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.OptimizationInfoBindingType))
    fieldNameMap["optimization_info"] = "OptimizationInfo"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    fields["publish_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.PublishInfoBindingType))
    fieldNameMap["publish_info"] = "PublishInfo"
    fields["subscription_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.SubscriptionInfoBindingType))
    fieldNameMap["subscription_info"] = "SubscriptionInfo"
    fields["server_guid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.VCenter"}, ""))
    fieldNameMap["server_guid"] = "ServerGuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library_model",fields, reflect.TypeOf(LibraryModel{}), fieldNameMap, validators)
}



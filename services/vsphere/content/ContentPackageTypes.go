/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.content.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package content

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"time"
)


// The ``ConfigurationModel`` class defines the global settings of the Content Library Service.
type ConfigurationModel struct {
    // Whether automatic synchronization is enabled. 
    //
    //  When automatic synchronization is enabled, the Content Library Service will automatically synchronize all subscribed libraries on a daily basis. Subscribed libraries with the library.SubscriptionInfo#automaticSyncEnabled flag turned on will be synchronized every hour between ConfigurationModel#automaticSyncStartHour and ConfigurationModel#automaticSyncStopHour.
	AutomaticSyncEnabled *bool
    // The hour at which the automatic synchronization will start. This value is between 0 (midnight) and 23 inclusive.
	AutomaticSyncStartHour *int64
    // The hour at which the automatic synchronization will stop. Any active synchronization operation will continue to run, however no new synchronization operations will be triggered after the stop hour. This value is between 0 (midnight) and 23 inclusive.
	AutomaticSyncStopHour *int64
    // The maximum allowed number of library items to synchronize concurrently from remote libraries. This must be a positive number. The service may not be able to guarantee the requested concurrency if there is no available capacity. 
    //
    //  This setting is global across all subscribed libraries.
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
    // An identifier which uniquely identifies this ``LibraryModel``.
	Id *string
    // The date and time when this library was created.
	CreationTime *time.Time
    // A human-readable description for this library.
	Description *string
    // The date and time when this library was last updated. 
    //
    //  This property is updated automatically when the library properties are changed. This property is not affected by adding, removing, or modifying a library item or its content within the library. Tagging the library or syncing the subscribed library does not alter this property.
	LastModifiedTime *time.Time
    // The date and time when this library was last synchronized. 
    //
    //  This property applies only to subscribed libraries. It is updated every time a synchronization is triggered on the library. The value is null for a local library.
	LastSyncTime *time.Time
    // The name of the library. 
    //
    //  A Library is identified by a human-readable name. Library names cannot be undefined or an empty string. Names do not have to be unique.
	Name *string
    // The list of default storage backings which are available for this library. 
    //
    //  A library.StorageBacking defines a default storage location which can be used to store files for library items in this library. Some library items, for instance, virtual machine template items, support files that may be distributed across various storage backings. One or more item files may or may not be located on the default storage backing. 
    //
    //  Multiple default storage locations are not currently supported but may become supported in future releases.
	StorageBackings []library.StorageBacking
    // The type (LOCAL, SUBSCRIBED) of this library. 
    //
    //  This value can be used to determine what additional services and information can be available for this library. This property is not used for the ``create`` and ``update`` methods. It will always be present in the result of a ``get`` method.
	Type_ *LibraryModelLibraryType
    // Defines various optimizations and optimization parameters applied to this library.
	OptimizationInfo *library.OptimizationInfo
    // A version number which is updated on metadata changes. This value allows clients to detect concurrent updates and prevent accidental clobbering of data. 
    //
    //  This value represents a number which is incremented every time library properties, such as name or description, are changed. It is not incremented by changes to a library item within the library, including adding or removing items. It is also not affected by tagging the library.
	Version *string
    // Defines how this library is published so that it can be subscribed to by a remote subscribed library. 
    //
    //  The library.PublishInfo defines where and how the metadata for this local library is accessible. A local library is only published publically if library.PublishInfo#published is ``true``.
	PublishInfo *library.PublishInfo
    // Defines the subscription behavior for this Library. 
    //
    //  The library.SubscriptionInfo defines how this subscribed library synchronizes to a remote source. Setting the value will determine the remote source to which the library synchronizes, and how. Changing the subscription will result in synchronizing to a new source. If the new source differs from the old one, the old library items and data will be lost. Setting library.SubscriptionInfo#automaticSyncEnabled to false will halt subscription but will not remove existing cached data.
	SubscriptionInfo *library.SubscriptionInfo
    // The unique identifier of the vCenter server where the library exists.
	ServerGuid *string
}

// The ``LibraryType`` enumeration class defines the type of a LibraryModel. 
//
//  The type of a library can be used to determine which additional services can be performed with a library.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryModelLibraryType string

const (
    // The library contents are defined and stored by the local Content Library Service installation. 
    //
    //  A local library can be retrieved and managed via the LocalLibrary.
	LibraryModelLibraryType_LOCAL LibraryModelLibraryType = "LOCAL"
    // The library synchronizes its items and content from another published library. 
    //
    //  A subscribed library can be retrieved and managed via the SubscribedLibrary.
	LibraryModelLibraryType_SUBSCRIBED LibraryModelLibraryType = "SUBSCRIBED"
)

func (l LibraryModelLibraryType) LibraryModelLibraryType() bool {
	switch l {
	case LibraryModelLibraryType_LOCAL:
		return true
	case LibraryModelLibraryType_SUBSCRIBED:
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
	return bindings.NewStructType("com.vmware.content.configuration_model", fields, reflect.TypeOf(ConfigurationModel{}), fieldNameMap, validators)
}

func LibraryModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
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
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library_model.library_type", reflect.TypeOf(LibraryModelLibraryType(LibraryModelLibraryType_LOCAL))))
	fieldNameMap["type"] = "Type_"
	fields["optimization_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.OptimizationInfoBindingType))
	fieldNameMap["optimization_info"] = "OptimizationInfo"
	fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["version"] = "Version"
	fields["publish_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.PublishInfoBindingType))
	fieldNameMap["publish_info"] = "PublishInfo"
	fields["subscription_info"] = bindings.NewOptionalType(bindings.NewReferenceType(library.SubscriptionInfoBindingType))
	fieldNameMap["subscription_info"] = "SubscriptionInfo"
	fields["server_guid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter"}, ""))
	fieldNameMap["server_guid"] = "ServerGuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library_model", fields, reflect.TypeOf(LibraryModel{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.library.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package library

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "net/url"
    "time"
)



// The ``ItemModel`` class represents a library item that has been stored in a library. 
//
//  A ``ItemModel`` represents a single logical unit to be managed within a content.LibraryModel. Items contain the actual content of a library, and their placement within a library determines policies that affect that content such as publishing. 
//
//  A library item can have a specified type, indicated with the ItemModel#type property. This property is associated with a Content Library Service plugin that supports specific types and provides additional services. The types available in a specific Content Library Service can be queried using the Type interface. Items of an unknown or unspecified type are treated generically. Because subscribed library catalogs are synchronized as is, subscribing to a remote Content Library Service effectively gives you a library with the functionality of the remote service's type adapter plugins, even if they are not installed locally. 
//
//  Items can be managed using the Item interface and, for items in subscribed libraries, the SubscribedItem interface.
type ItemModel struct {
    Id *string
    LibraryId *string
    ContentVersion *string
    CreationTime *time.Time
    Description *string
    LastModifiedTime *time.Time
    LastSyncTime *time.Time
    MetadataVersion *string
    Name *string
    Cached *bool
    Size *int64
    Type_ *string
    Version *string
    SourceId *string
}






// The ``OptimizationInfo`` class defines different optimizations and optimization parameters applied to particular library.
type OptimizationInfo struct {
    OptimizeRemotePublishing *bool
}






// The ``PublishInfo`` class defines how a local library is published publicly for synchronization to other libraries.
type PublishInfo struct {
    AuthenticationMethod *PublishInfo_AuthenticationMethod
    Published *bool
    PublishUrl *url.URL
    UserName *string
    Password *string
    CurrentPassword *string
    PersistJsonEnabled *bool
}




    
    // The ``AuthenticationMethod`` enumeration class indicates how a subscribed library should authenticate to the published library endpoint.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type PublishInfo_AuthenticationMethod string

    const (
        // Require HTTP Basic authentication matching a specified username and password.
         PublishInfo_AuthenticationMethod_BASIC PublishInfo_AuthenticationMethod = "BASIC"
        // Require no authentication.
         PublishInfo_AuthenticationMethod_NONE PublishInfo_AuthenticationMethod = "NONE"
    )

    func (a PublishInfo_AuthenticationMethod) PublishInfo_AuthenticationMethod() bool {
        switch a {
            case PublishInfo_AuthenticationMethod_BASIC:
                return true
            case PublishInfo_AuthenticationMethod_NONE:
                return true
            default:
                return false
        }
    }



// The ``SourceInfo`` class contains information about the source published library of a subscribed library.
type SourceInfo struct {
    SourceLibrary *string
    Subscription *string
}






// The ``StorageBacking`` class defines a storage location where content in a library will be stored. The storage location can either be a Datastore or Other type.
type StorageBacking struct {
    Type_ *StorageBacking_Type
    DatastoreId *string
    StorageUri *url.URL
}




    
    // The ``Type`` enumeration class specifies the type of the StorageBacking.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type StorageBacking_Type string

    const (
        // The content of the library will be stored on a datastore. 
        //
        //  These are vCenter Server managed datastores, and are logical containers that hide specifics of each storage device. Depending on the type of storage you use, datastores can be backed by the following file system formats: 
        //
        // * Virtual Machine File System (VMFS)
        // * Network File System (NFS)
        //
        //  
         StorageBacking_Type_DATASTORE StorageBacking_Type = "DATASTORE"
        // The content of the library will be stored on a remote file system. 
        //
        //  Supports the following remote file systems: 
        //
        // * NFS (on vCenter Server Appliance)
        // * SMB (on vCenter Server Appliance and vCenter Server for Windows)
        //
        //  
         StorageBacking_Type_OTHER StorageBacking_Type = "OTHER"
    )

    func (t StorageBacking_Type) StorageBacking_Type() bool {
        switch t {
            case StorageBacking_Type_DATASTORE:
                return true
            case StorageBacking_Type_OTHER:
                return true
            default:
                return false
        }
    }



// The ``SubscriptionInfo`` class defines the subscription behavior for a subscribed library.
type SubscriptionInfo struct {
    AuthenticationMethod *SubscriptionInfo_AuthenticationMethod
    AutomaticSyncEnabled *bool
    OnDemand *bool
    Password *string
    SslThumbprint *string
    SubscriptionUrl *url.URL
    UserName *string
    SourceInfo *SourceInfo
}




    
    // Indicate how the subscribed library should authenticate with the published library endpoint.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type SubscriptionInfo_AuthenticationMethod string

    const (
        // Require HTTP Basic authentication matching a specified username and password.
         SubscriptionInfo_AuthenticationMethod_BASIC SubscriptionInfo_AuthenticationMethod = "BASIC"
        // Require no authentication.
         SubscriptionInfo_AuthenticationMethod_NONE SubscriptionInfo_AuthenticationMethod = "NONE"
    )

    func (a SubscriptionInfo_AuthenticationMethod) SubscriptionInfo_AuthenticationMethod() bool {
        switch a {
            case SubscriptionInfo_AuthenticationMethod_BASIC:
                return true
            case SubscriptionInfo_AuthenticationMethod_NONE:
                return true
            default:
                return false
        }
    }







func ItemModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["id"] = "Id"
    fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["library_id"] = "LibraryId"
    fields["content_version"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, ""))
    fieldNameMap["content_version"] = "ContentVersion"
    fields["creation_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["creation_time"] = "CreationTime"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["last_modified_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_modified_time"] = "LastModifiedTime"
    fields["last_sync_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_sync_time"] = "LastSyncTime"
    fields["metadata_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["metadata_version"] = "MetadataVersion"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["cached"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cached"] = "Cached"
    fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["size"] = "Size"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["source_id"] = "SourceId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item_model",fields, reflect.TypeOf(ItemModel{}), fieldNameMap, validators)
}

func OptimizationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["optimize_remote_publishing"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["optimize_remote_publishing"] = "OptimizeRemotePublishing"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.optimization_info",fields, reflect.TypeOf(OptimizationInfo{}), fieldNameMap, validators)
}

func PublishInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.publish_info.authentication_method", reflect.TypeOf(PublishInfo_AuthenticationMethod(PublishInfo_AuthenticationMethod_BASIC))))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["published"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["published"] = "Published"
    fields["publish_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["publish_url"] = "PublishUrl"
    fields["user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["current_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["current_password"] = "CurrentPassword"
    fields["persist_json_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["persist_json_enabled"] = "PersistJsonEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.publish_info",fields, reflect.TypeOf(PublishInfo{}), fieldNameMap, validators)
}

func SourceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_library"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["source_library"] = "SourceLibrary"
    fields["subscription"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, ""))
    fieldNameMap["subscription"] = "Subscription"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.source_info",fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func StorageBackingBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.storage_backing.type", reflect.TypeOf(StorageBacking_Type(StorageBacking_Type_DATASTORE))))
    fieldNameMap["type"] = "Type_"
    fields["datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore_id"] = "DatastoreId"
    fields["storage_uri"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["storage_uri"] = "StorageUri"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "DATASTORE": []bindings.FieldData {
                 bindings.NewFieldData("datastore_id", true),
            },
            "OTHER": []bindings.FieldData {
                 bindings.NewFieldData("storage_uri", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.storage_backing",fields, reflect.TypeOf(StorageBacking{}), fieldNameMap, validators)
}

func SubscriptionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.subscription_info.authentication_method", reflect.TypeOf(SubscriptionInfo_AuthenticationMethod(SubscriptionInfo_AuthenticationMethod_BASIC))))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["automatic_sync_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["automatic_sync_enabled"] = "AutomaticSyncEnabled"
    fields["on_demand"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["on_demand"] = "OnDemand"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["subscription_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["subscription_url"] = "SubscriptionUrl"
    fields["user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user_name"] = "UserName"
    fields["source_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SourceInfoBindingType))
    fieldNameMap["source_info"] = "SourceInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.subscription_info",fields, reflect.TypeOf(SubscriptionInfo{}), fieldNameMap, validators)
}



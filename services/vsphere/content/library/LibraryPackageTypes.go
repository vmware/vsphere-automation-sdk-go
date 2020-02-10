/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
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
    // A unique identifier for this library item.
	Id *string
    // The identifier of the content.LibraryModel to which this item belongs.
	LibraryId *string
    // The latest version of the file content list of this library item.
	ContentVersion *string
    // The date and time when this library item was created.
	CreationTime *time.Time
    // A human-readable description for this library item.
	Description *string
    // The date and time when the metadata for this library item was last changed. 
    //
    //  This property is affected by changes to the properties or file content of this item. It is not modified by changes to the tags of the item, or by changes to the library which owns this item.
	LastModifiedTime *time.Time
    // The date and time when this library item was last synchronized. 
    //
    //  This property is updated every time a synchronization is triggered on the library item, including when a synchronization is triggered on the library to which this item belongs. The value is null for a library item that belongs to a local library.
	LastSyncTime *time.Time
    // A version number for the metadata of this library item. 
    //
    //  This value is incremented with each change to the metadata of this item. Changes to name, description, and so on will increment this value. The value is not incremented by changes to the content or tags of the item or the library which owns it.
	MetadataVersion *string
    // A human-readable name for this library item. 
    //
    //  The name may not be null or an empty string. The name does not have to be unique, even within the same library.
	Name *string
    // The status that indicates whether the library item is on disk or not. The library item is cached when all its files are on disk.
	Cached *bool
    // The library item size, in bytes. The size is the sum of the size used on the storage backing for all the files in the item. When the library item is not cached, the size is 0.
	Size *int64
    // An optional type identifier which indicates the type adapter plugin to use. 
    //
    //  This property may be set to a non-empty string value that corresponds to an identifier supported by a type adapter plugin present in the Content Library Service. A type adapter plugin, if present for the specified type, can provide additional information and services around the item content. A type adapter can guide the upload process by creating file entries that are in need of being uploaded to complete an item. 
    //
    //  The types and plugins supported by the Content Library Service can be queried using the Type interface.
	Type_ *string
    // A version number that is updated on metadata changes. This value is used to validate update requests to provide optimistic concurrency of changes. 
    //
    //  This value represents a number that is incremented every time library item properties, such as name or description, are changed. It is not incremented by changes to the file content of the library item, including adding or removing files. It is also not affected by tagging the library item.
	Version *string
    // The identifier of the ItemModel to which this item is synchronized to if the item belongs to a subscribed library. The value is null for a library item that belongs to a local library.
	SourceId *string
}

// The ``OptimizationInfo`` class defines different optimizations and optimization parameters applied to particular library.
type OptimizationInfo struct {
    // If set to ``true`` then library would be optimized for remote publishing. 
    //
    //  Turn it on if remote publishing is dominant use case for this library. Remote publishing means here that publisher and subscribers are not the part of the same ``Vcenter`` SSO domain. 
    //
    //  Any optimizations could be done as result of turning on this optimization during library creation. For example, library content could be stored in different format but optimizations are not limited to just storage format. 
    //
    //  Note, that value of this toggle could be set only during creation of the library and you would need to migrate your library in case you need to change this value (optimize the library for different use case).
	OptimizeRemotePublishing *bool
}

// The ``PublishInfo`` class defines how a local library is published publicly for synchronization to other libraries.
type PublishInfo struct {
    // Indicates how a subscribed library should authenticate (BASIC, NONE) to the published library endpoint.
	AuthenticationMethod *PublishInfoAuthenticationMethod
    // Whether the local library is published.
	Published *bool
    // The URL to which the library metadata is published by the Content Library Service. 
    //
    //  This value can be used to set the SubscriptionInfo#subscriptionUrl property when creating a subscribed library.
	PublishUrl *url.URL
    // The username to require for authentication.
	UserName *string
    // The new password to require for authentication.
	Password *string
    // The current password to verify. This property is available starting in vSphere 6.7.
	CurrentPassword *string
    // Whether library and library item metadata are persisted in the storage backing as JSON files. This flag only applies if the local library is published. 
    //
    //  Enabling JSON persistence allows you to synchronize a subscribed library manually instead of over HTTP. You copy the local library content and metadata to another storage backing manually and then create a subscribed library referencing the location of the library JSON file in the SubscriptionInfo#subscriptionUrl. When the subscribed library's storage backing matches the subscription URL, files do not need to be copied to the subscribed library. 
    //
    //  For a library backed by a datastore, the library JSON file will be stored at the path contentlib-{library_id}/lib.json on the datastore. 
    //
    //  For a library backed by a remote file system, the library JSON file will be stored at {library_id}/lib.json in the remote file system path.
	PersistJsonEnabled *bool
}

// The ``AuthenticationMethod`` enumeration class indicates how a subscribed library should authenticate to the published library endpoint.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PublishInfoAuthenticationMethod string

const (
    // Require HTTP Basic authentication matching a specified username and password.
	PublishInfoAuthenticationMethod_BASIC PublishInfoAuthenticationMethod = "BASIC"
    // Require no authentication.
	PublishInfoAuthenticationMethod_NONE PublishInfoAuthenticationMethod = "NONE"
)

func (a PublishInfoAuthenticationMethod) PublishInfoAuthenticationMethod() bool {
	switch a {
	case PublishInfoAuthenticationMethod_BASIC:
		return true
	case PublishInfoAuthenticationMethod_NONE:
		return true
	default:
		return false
	}
}


// The ``SourceInfo`` class contains information about the source published library of a subscribed library. This class was added in vSphere API 6.7.2.
type SourceInfo struct {
    // Identifier of the published library. This property was added in vSphere API 6.7.2.
	SourceLibrary *string
    // Identifier of the subscription associated with the subscribed library. This property was added in vSphere API 6.7.2.
	Subscription *string
}

// The ``StorageBacking`` class defines a storage location where content in a library will be stored. The storage location can either be a Datastore or Other type.
type StorageBacking struct {
    // Type (DATASTORE, OTHER) of StorageBacking.
	Type_ *StorageBackingType
    // Identifier of the datastore used to store the content in the library.
	DatastoreId *string
    // URI identifying the location used to store the content in the library. 
    //
    //  The following URI formats are supported: 
    //
    //  vSphere 6.5 
    //
    // * nfs://server/path?version=4 (for vCenter Server Appliance only) - Specifies an NFS Version 4 server.
    // * nfs://server/path (for vCenter Server Appliance only) - Specifies an NFS Version 3 server. The nfs://server:/path format is also supported.
    // * smb://server/path - Specifies an SMB server or Windows share.
    //
    //  
    //
    //  vSphere 6.0 Update 1 
    //
    // * nfs://server:/path (for vCenter Server Appliance only)
    // * file://unc-server/path (for vCenter Server for Windows only)
    // * file:///mount/point (for vCenter Server Appliance only) - Local file URIs are supported only when the path is a local mount point for an NFS file system. Use of file URIs is strongly discouraged. Instead, use an NFS URI to specify the remote file system.
    //
    //  
    //
    //  vSphere 6.0 
    //
    // * nfs://server:/path (for vCenter Server Appliance only)
    // * file://unc-server/path (for vCenter Server for Windows only)
    // * file:///path - Local file URIs are supported but strongly discouraged because it may interfere with the performance of vCenter Server.
	StorageUri *url.URL
}

// The ``Type`` enumeration class specifies the type of the StorageBacking.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type StorageBackingType string

const (
    // The content of the library will be stored on a datastore. 
    //
    //  These are vCenter Server managed datastores, and are logical containers that hide specifics of each storage device. Depending on the type of storage you use, datastores can be backed by the following file system formats: 
    //
    // * Virtual Machine File System (VMFS)
    // * Network File System (NFS)
    //
    //  
	StorageBackingType_DATASTORE StorageBackingType = "DATASTORE"
    // The content of the library will be stored on a remote file system. 
    //
    //  Supports the following remote file systems: 
    //
    // * NFS (on vCenter Server Appliance)
    // * SMB (on vCenter Server Appliance and vCenter Server for Windows)
    //
    //  
	StorageBackingType_OTHER StorageBackingType = "OTHER"
)

func (t StorageBackingType) StorageBackingType() bool {
	switch t {
	case StorageBackingType_DATASTORE:
		return true
	case StorageBackingType_OTHER:
		return true
	default:
		return false
	}
}


// The ``SubscriptionInfo`` class defines the subscription behavior for a subscribed library.
type SubscriptionInfo struct {
    // Indicate how the subscribed library should authenticate (BASIC, NONE) with the published library endpoint.
	AuthenticationMethod *SubscriptionInfoAuthenticationMethod
    // Whether the library should participate in automatic library synchronization. In order for automatic synchronization to happen, the global content.ConfigurationModel#automaticSyncEnabled option must also be true. The subscription is still active even when automatic synchronization is turned off, but synchronization is only activated with an explicit call to SubscribedLibrary#sync or SubscribedItem#sync. In other words, manual synchronization is still available even when automatic synchronization is disabled.
	AutomaticSyncEnabled *bool
    // Indicates whether a library item's content will be synchronized only on demand. 
    //
    //  If this is set to ``true``, then the library item's metadata will be synchronized but the item's content (its files) will not be synchronized. The Content Library Service will synchronize the content upon request only. This can cause the first use of the content to have a noticeable delay. 
    //
    //  Items without synchronized content can be forcefully synchronized in advance using the SubscribedItem#sync call with ``forceSyncContent`` set to true. Once content has been synchronized, the content can removed with the SubscribedItem#evict call. 
    //
    //  If this value is set to ``false``, all content will be synchronized in advance.
	OnDemand *bool
    // The password to use when authenticating. 
    //
    //  The password must be set when using a password-based authentication method; empty strings are not allowed.
	Password *string
    // An optional SHA-1 hash of the SSL certificate for the remote endpoint. 
    //
    //  If this value is defined the SSL certificate will be verified by comparing it to the SSL thumbprint. The SSL certificate must verify against the thumbprint. When specified, the standard certificate chain validation behavior is not used. The certificate chain is validated normally if this value is null.
	SslThumbprint *string
    // The URL of the endpoint where the metadata for the remotely published library is being served. 
    //
    //  This URL can be the PublishInfo#publishUrl of the published library (for example, https://server/path/lib.json). 
    //
    //  If the source content comes from a published library with PublishInfo#persistJsonEnabled, the subscription URL can be a URL pointing to the library JSON file on a datastore or remote file system. The supported formats are: 
    //
    //  vSphere 6.5 
    //
    // * ds:///vmfs/volumes/{uuid}/mylibrary/lib.json (for datastore)
    // * nfs://server/path/mylibrary/lib.json (for NFSv3 server on vCenter Server Appliance)
    // * nfs://server/path/mylibrary/lib.json?version=4 (for NFSv4 server on vCenter Server Appliance)
    // * smb://server/path/mylibrary/lib.json (for SMB server)
    //
    //  
    //
    //  vSphere 6.0 
    //
    // * file://server/mylibrary/lib.json (for UNC server on vCenter Server for Windows)
    // * file:///path/mylibrary/lib.json (for local file system)
    //
    //  
    //
    //  When you specify a DS subscription URL, the datastore must be on the same vCenter Server as the subscribed library. When you specify an NFS or SMB subscription URL, the StorageBacking#storageUri of the subscribed library must be on the same remote file server and should share a common parent path with the subscription URL.
	SubscriptionUrl *url.URL
    // The username to use when authenticating. 
    //
    //  The username must be set when using a password-based authentication method. Empty strings are allowed for usernames.
	UserName *string
    // Information about the source published library. This property will be set for a subscribed library which is associated with a subscription of the published library. This property was added in vSphere API 6.7.2.
	SourceInfo *SourceInfo
}

// Indicate how the subscribed library should authenticate with the published library endpoint.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SubscriptionInfoAuthenticationMethod string

const (
    // Require HTTP Basic authentication matching a specified username and password.
	SubscriptionInfoAuthenticationMethod_BASIC SubscriptionInfoAuthenticationMethod = "BASIC"
    // Require no authentication.
	SubscriptionInfoAuthenticationMethod_NONE SubscriptionInfoAuthenticationMethod = "NONE"
)

func (a SubscriptionInfoAuthenticationMethod) SubscriptionInfoAuthenticationMethod() bool {
	switch a {
	case SubscriptionInfoAuthenticationMethod_BASIC:
		return true
	case SubscriptionInfoAuthenticationMethod_NONE:
		return true
	default:
		return false
	}
}





func ItemModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["id"] = "Id"
	fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
	fieldNameMap["library_id"] = "LibraryId"
	fields["content_version"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, ""))
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
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item_model", fields, reflect.TypeOf(ItemModel{}), fieldNameMap, validators)
}

func OptimizationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["optimize_remote_publishing"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["optimize_remote_publishing"] = "OptimizeRemotePublishing"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.optimization_info", fields, reflect.TypeOf(OptimizationInfo{}), fieldNameMap, validators)
}

func PublishInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.publish_info.authentication_method", reflect.TypeOf(PublishInfoAuthenticationMethod(PublishInfoAuthenticationMethod_BASIC))))
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
	return bindings.NewStructType("com.vmware.content.library.publish_info", fields, reflect.TypeOf(PublishInfo{}), fieldNameMap, validators)
}

func SourceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_library"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
	fieldNameMap["source_library"] = "SourceLibrary"
	fields["subscription"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, ""))
	fieldNameMap["subscription"] = "Subscription"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.source_info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func StorageBackingBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.storage_backing.type", reflect.TypeOf(StorageBackingType(StorageBackingType_DATASTORE))))
	fieldNameMap["type"] = "Type_"
	fields["datastore_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore_id"] = "DatastoreId"
	fields["storage_uri"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["storage_uri"] = "StorageUri"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"DATASTORE": []bindings.FieldData{
				bindings.NewFieldData("datastore_id", true),
			},
			"OTHER": []bindings.FieldData{
				bindings.NewFieldData("storage_uri", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.storage_backing", fields, reflect.TypeOf(StorageBacking{}), fieldNameMap, validators)
}

func SubscriptionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.subscription_info.authentication_method", reflect.TypeOf(SubscriptionInfoAuthenticationMethod(SubscriptionInfoAuthenticationMethod_BASIC))))
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
	return bindings.NewStructType("com.vmware.content.library.subscription_info", fields, reflect.TypeOf(SubscriptionInfo{}), fieldNameMap, validators)
}



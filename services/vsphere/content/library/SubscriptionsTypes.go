/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Subscriptions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package library

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for Subscription resource. This constant field was added in vSphere API 6.7.2.
const Subscriptions_RESOURCE_TYPE = "com.vmware.content.library.Subscriptions"


// The ``Location`` enumeration class defines the location of subscribed library relative to the published library. This enumeration was added in vSphere API 6.7.2.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SubscriptionsLocation string

const (
    // The subscribed library belongs to the same vCenter instance as the published library. This constant field was added in vSphere API 6.7.2.
	SubscriptionsLocation_LOCAL SubscriptionsLocation = "LOCAL"
    // The subscribed library belongs to a different vCenter instance than the published library. This constant field was added in vSphere API 6.7.2.
	SubscriptionsLocation_REMOTE SubscriptionsLocation = "REMOTE"
)

func (l SubscriptionsLocation) SubscriptionsLocation() bool {
	switch l {
	case SubscriptionsLocation_LOCAL:
		return true
	case SubscriptionsLocation_REMOTE:
		return true
	default:
		return false
	}
}


// The ``CreateSpecNewSubscribedLibrary`` class defines the information required to create a new subscribed library. This class was added in vSphere API 6.7.2.
type SubscriptionsCreateSpecNewSubscribedLibrary struct {
    // Name of the subscribed library. This property was added in vSphere API 6.7.2.
	Name string
    // Description of the subscribed library. This property was added in vSphere API 6.7.2.
	Description *string
    // The list of default storage backings for this library. 
    //
    //  The list must contain exactly one storage backing. Multiple default storage locations are not currently supported but may become supported in future releases.. This property was added in vSphere API 6.7.2.
	StorageBackings []StorageBacking
    // Specifies whether the library should participate in automatic library synchronization. This property was added in vSphere API 6.7.2.
	AutomaticSyncEnabled bool
    // Specifies whether a library item's content will be synchronized only on demand. This property was added in vSphere API 6.7.2.
	OnDemand bool
}

// The ``CreateSpecVcenter`` class defines information about the vCenter Server instance where the subscribed library associated with the subscription exists or will be created. This class was added in vSphere API 6.7.2.
type SubscriptionsCreateSpecVcenter struct {
    // The hostname of the subscribed library's vCenter Server. This property was added in vSphere API 6.7.2.
	Hostname string
    // The HTTPS port of the vCenter Server instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	HttpsPort *int64
}

// The ``CreateSpecPlacement`` class defines the placement information for the subscribed library's virtual machine template library items. Storage location of the virtual machine template items is defined by the subscribed library's storage backing. This placement information needs to be compatible with the subscribed library's storage backing. The ``CreateSpecPlacement`` class is only applicable for the virtual machine template library items of the subscribed library. This class was added in vSphere API 6.7.2.
type SubscriptionsCreateSpecPlacement struct {
    // Virtual machine folder into which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	Folder *string
    // Cluster onto which the virtual machine template should be placed. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.7.2.
	Cluster *string
    // Resource pool into which the virtual machine template should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. This property was added in vSphere API 6.7.2.
	ResourcePool *string
    // Host onto which the virtual machine template should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.7.2.
	Host *string
    // Network that backs the virtual Ethernet adapters in the virtual machine template. This property was added in vSphere API 6.7.2.
	Network *string
}

// The ``CreateSpecSubscribedLibrary`` class defines the subscribed library information used to create the subscription. This class was added in vSphere API 6.7.2.
type SubscriptionsCreateSpecSubscribedLibrary struct {
    // Specifies whether the target subscribed library should be newly created or an existing subscribed library should be used. This property was added in vSphere API 6.7.2.
	Target SubscriptionsCreateSpecSubscribedLibraryTarget
    // Specification for creating a new subscribed library associated with the subscription. This property was added in vSphere API 6.7.2.
	NewSubscribedLibrary *SubscriptionsCreateSpecNewSubscribedLibrary
    // Identifier of the existing subscribed library to associate with the subscription. Only the subscribed libraries for which SubscriptionInfo#subscriptionUrl property is set to the PublishInfo#publishUrl of the published library can be associated with the subscription. This property was added in vSphere API 6.7.2.
	SubscribedLibrary *string
    // Location of the subscribed library relative to the published library. This property was added in vSphere API 6.7.2.
	Location SubscriptionsLocation
    // Specification for the subscribed library's vCenter Server instance. This property was added in vSphere API 6.7.2.
	Vcenter *SubscriptionsCreateSpecVcenter
    // Placement specification for the virtual machine template library items on the subscribed library. This property was added in vSphere API 6.7.2.
	Placement *SubscriptionsCreateSpecPlacement
}

// The ``Target`` enumeration class defines the options related to the target subscribed library which will be associated with the subscription. This enumeration was added in vSphere API 6.7.2.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SubscriptionsCreateSpecSubscribedLibraryTarget string

const (
    // Create a new subscribed library. This constant field was added in vSphere API 6.7.2.
	SubscriptionsCreateSpecSubscribedLibraryTarget_CREATE_NEW SubscriptionsCreateSpecSubscribedLibraryTarget = "CREATE_NEW"
    // Use the specified existing subscribed library. This constant field was added in vSphere API 6.7.2.
	SubscriptionsCreateSpecSubscribedLibraryTarget_USE_EXISTING SubscriptionsCreateSpecSubscribedLibraryTarget = "USE_EXISTING"
)

func (t SubscriptionsCreateSpecSubscribedLibraryTarget) SubscriptionsCreateSpecSubscribedLibraryTarget() bool {
	switch t {
	case SubscriptionsCreateSpecSubscribedLibraryTarget_CREATE_NEW:
		return true
	case SubscriptionsCreateSpecSubscribedLibraryTarget_USE_EXISTING:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class defines the information required to create a new subscription of the published library. This class was added in vSphere API 6.7.2.
type SubscriptionsCreateSpec struct {
    // Specification for the subscribed library to be associated with the subscription. This property was added in vSphere API 6.7.2.
	SubscribedLibrary SubscriptionsCreateSpecSubscribedLibrary
}

// The ``Summary`` class contains commonly used information about the subscription. This class was added in vSphere API 6.7.2.
type SubscriptionsSummary struct {
    // Identifier of the subscription. This property was added in vSphere API 6.7.2.
	Subscription string
    // Identifier of the subscribed library. This property was added in vSphere API 6.7.2.
	SubscribedLibrary string
    // Name of the subscribed library. This property was added in vSphere API 6.7.2.
	SubscribedLibraryName string
    // Hostname of the vCenter instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	SubscribedLibraryVcenterHostname *string
}

// The ``UpdateSpecVcenter`` class defines information about the vCenter Server instance where the subscribed library associated with the subscription exists. The ``UpdateSpecVcenter`` class is only applicable to subscribed library which exists on remote vCenter Server instance. This class was added in vSphere API 6.7.2.
type SubscriptionsUpdateSpecVcenter struct {
    // The hostname of the subscribed library's vCenter Server. This property was added in vSphere API 6.7.2.
	Hostname *string
    // The HTTPS port of the vCenter Server instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	HttpsPort *int64
}

// The ``UpdateSpecPlacement`` class defines the placement information for the subscribed library's virtual machine template library items. Storage location of the virtual machine template items is defined by the subscribed library's storage backing. This placement information needs to be compatible with the subscribed library's storage backing. The ``UpdateSpecPlacement`` class is only applicable for the newly published virtual machine template library items of the subscribed library. Existing items will not be moved. This class was added in vSphere API 6.7.2.
type SubscriptionsUpdateSpecPlacement struct {
    // Virtual machine folder into which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	Folder *string
    // Cluster onto which the virtual machine template should be placed. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``. If ``resourcePool`` or ``host`` is specified, it is recommended that this property be null. This property was added in vSphere API 6.7.2.
	Cluster *string
    // Resource pool into which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	ResourcePool *string
    // Host onto which the virtual machine template should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.7.2.
	Host *string
    // Network that backs the virtual Ethernet adapters in the virtual machine template. This property was added in vSphere API 6.7.2.
	Network *string
}

// The ``UpdateSpec`` class defines information required to update the subscription. This class was added in vSphere API 6.7.2.
type SubscriptionsUpdateSpec struct {
    // Specification for the subscribed library's vCenter Server instance. This property was added in vSphere API 6.7.2.
	SubscribedLibraryVcenter *SubscriptionsUpdateSpecVcenter
    // Placement specification for the virtual machine template items of the subscribed library. Updating this information will only affect new or updated items, existing items will not be moved. The entire placement configuration of the subscribed library will replaced by the new specification. This property was added in vSphere API 6.7.2.
	SubscribedLibraryPlacement *SubscriptionsUpdateSpecPlacement
}

// The ``VcenterInfo`` class contains information about the vCenter Server instance where the subscribed library associated with the subscription exists. This class was added in vSphere API 6.7.2.
type SubscriptionsVcenterInfo struct {
    // Hostname of the vCenter Server instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	Hostname string
    // The HTTPS port of the vCenter Server instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	HttpsPort *int64
    // The unique identifier of the vCenter Server where the subscribed library exists. This property was added in vSphere API 6.7.2.
	ServerGuid string
}

// The ``PlacementInfo`` class contains the placement information for the subscribed library's virtual machine template library items. The ``PlacementInfo`` class is only applicable for the virtual machine template library items of the subscribed library. This class was added in vSphere API 6.7.2.
type SubscriptionsPlacementInfo struct {
    // Virtual machine folder into which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	Folder *string
    // Cluster onto which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	Cluster *string
    // Resource pool into which the virtual machine template should be placed. This property was added in vSphere API 6.7.2.
	ResourcePool *string
    // Host onto which the virtual machine template should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.7.2.
	Host *string
    // Network that backs the virtual Ethernet adapters in the virtual machine template. This property was added in vSphere API 6.7.2.
	Network *string
}

// The ``Info`` class contains information about the subscription. This class was added in vSphere API 6.7.2.
type SubscriptionsInfo struct {
    // Identifier of the subscribed library associated with the subscription. This property was added in vSphere API 6.7.2.
	SubscribedLibrary string
    // Name of the subscribed library associated with the subscription. This property was added in vSphere API 6.7.2.
	SubscribedLibraryName string
    // Location of the subscribed library relative to the published library. This property was added in vSphere API 6.7.2.
	SubscribedLibraryLocation SubscriptionsLocation
    // Information about the vCenter Server instance where the subscribed library exists. This property was added in vSphere API 6.7.2.
	SubscribedLibraryVcenter *SubscriptionsVcenterInfo
    // Placement information about the subscribed library's virtual machine template items. This property was added in vSphere API 6.7.2.
	SubscribedLibraryPlacement SubscriptionsPlacementInfo
}



func subscriptionsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["spec"] = bindings.NewReferenceType(SubscriptionsCreateSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["library"] = "Library"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscriptionsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
}

func subscriptionsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["spec"] = bindings.NewReferenceType(SubscriptionsCreateSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["library"] = "Library"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func subscriptionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscriptionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscriptionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func subscriptionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library"] = "Library"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscriptionsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(SubscriptionsSummaryBindingType), reflect.TypeOf([]SubscriptionsSummary{}))
}

func subscriptionsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library"] = "Library"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidElementType": 400,"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func subscriptionsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fields["spec"] = bindings.NewReferenceType(SubscriptionsUpdateSpecBindingType)
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscriptionsUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func subscriptionsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fields["spec"] = bindings.NewReferenceType(SubscriptionsUpdateSpecBindingType)
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func subscriptionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func subscriptionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SubscriptionsInfoBindingType)
}

func subscriptionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["library"] = "Library"
	fieldNameMap["subscription"] = "Subscription"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func SubscriptionsCreateSpecNewSubscribedLibraryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["storage_backings"] = bindings.NewListType(bindings.NewReferenceType(StorageBackingBindingType), reflect.TypeOf([]StorageBacking{}))
	fieldNameMap["storage_backings"] = "StorageBackings"
	fields["automatic_sync_enabled"] = bindings.NewBooleanType()
	fieldNameMap["automatic_sync_enabled"] = "AutomaticSyncEnabled"
	fields["on_demand"] = bindings.NewBooleanType()
	fieldNameMap["on_demand"] = "OnDemand"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_new_subscribed_library", fields, reflect.TypeOf(SubscriptionsCreateSpecNewSubscribedLibrary{}), fieldNameMap, validators)
}

func SubscriptionsCreateSpecVcenterBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_vcenter", fields, reflect.TypeOf(SubscriptionsCreateSpecVcenter{}), fieldNameMap, validators)
}

func SubscriptionsCreateSpecPlacementBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder:VCenter"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource:VCenter"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool:VCenter"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem:VCenter"}, ""))
	fieldNameMap["host"] = "Host"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["network"] = "Network"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_placement", fields, reflect.TypeOf(SubscriptionsCreateSpecPlacement{}), fieldNameMap, validators)
}

func SubscriptionsCreateSpecSubscribedLibraryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.create_spec_subscribed_library.target", reflect.TypeOf(SubscriptionsCreateSpecSubscribedLibraryTarget(SubscriptionsCreateSpecSubscribedLibraryTarget_CREATE_NEW)))
	fieldNameMap["target"] = "Target"
	fields["new_subscribed_library"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsCreateSpecNewSubscribedLibraryBindingType))
	fieldNameMap["new_subscribed_library"] = "NewSubscribedLibrary"
	fields["subscribed_library"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
	fieldNameMap["subscribed_library"] = "SubscribedLibrary"
	fields["location"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.location", reflect.TypeOf(SubscriptionsLocation(SubscriptionsLocation_LOCAL)))
	fieldNameMap["location"] = "Location"
	fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsCreateSpecVcenterBindingType))
	fieldNameMap["vcenter"] = "Vcenter"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsCreateSpecPlacementBindingType))
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("target",
		map[string][]bindings.FieldData{
			"CREATE_NEW": []bindings.FieldData{
				bindings.NewFieldData("new_subscribed_library", true),
			},
			"USE_EXISTING": []bindings.FieldData{
				bindings.NewFieldData("subscribed_library", true),
			},
		},
	)
	validators = append(validators, uv1)
	uv2 := bindings.NewUnionValidator("location",
		map[string][]bindings.FieldData{
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("vcenter", true),
			},
			"LOCAL": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv2)
	return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec_subscribed_library", fields, reflect.TypeOf(SubscriptionsCreateSpecSubscribedLibrary{}), fieldNameMap, validators)
}

func SubscriptionsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscribed_library"] = bindings.NewReferenceType(SubscriptionsCreateSpecSubscribedLibraryBindingType)
	fieldNameMap["subscribed_library"] = "SubscribedLibrary"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.create_spec", fields, reflect.TypeOf(SubscriptionsCreateSpec{}), fieldNameMap, validators)
}

func SubscriptionsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["subscription"] = "Subscription"
	fields["subscribed_library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["subscribed_library"] = "SubscribedLibrary"
	fields["subscribed_library_name"] = bindings.NewStringType()
	fieldNameMap["subscribed_library_name"] = "SubscribedLibraryName"
	fields["subscribed_library_vcenter_hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subscribed_library_vcenter_hostname"] = "SubscribedLibraryVcenterHostname"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.summary", fields, reflect.TypeOf(SubscriptionsSummary{}), fieldNameMap, validators)
}

func SubscriptionsUpdateSpecVcenterBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["hostname"] = "Hostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec_vcenter", fields, reflect.TypeOf(SubscriptionsUpdateSpecVcenter{}), fieldNameMap, validators)
}

func SubscriptionsUpdateSpecPlacementBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder:VCenter"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource:VCenter"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool:VCenter"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem:VCenter"}, ""))
	fieldNameMap["host"] = "Host"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["network"] = "Network"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec_placement", fields, reflect.TypeOf(SubscriptionsUpdateSpecPlacement{}), fieldNameMap, validators)
}

func SubscriptionsUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscribed_library_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsUpdateSpecVcenterBindingType))
	fieldNameMap["subscribed_library_vcenter"] = "SubscribedLibraryVcenter"
	fields["subscribed_library_placement"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsUpdateSpecPlacementBindingType))
	fieldNameMap["subscribed_library_placement"] = "SubscribedLibraryPlacement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.update_spec", fields, reflect.TypeOf(SubscriptionsUpdateSpec{}), fieldNameMap, validators)
}

func SubscriptionsVcenterInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["server_guid"] = bindings.NewStringType()
	fieldNameMap["server_guid"] = "ServerGuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.vcenter_info", fields, reflect.TypeOf(SubscriptionsVcenterInfo{}), fieldNameMap, validators)
}

func SubscriptionsPlacementInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder:VCenter"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource:VCenter"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool:VCenter"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem:VCenter"}, ""))
	fieldNameMap["host"] = "Host"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["network"] = "Network"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.subscriptions.placement_info", fields, reflect.TypeOf(SubscriptionsPlacementInfo{}), fieldNameMap, validators)
}

func SubscriptionsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscribed_library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["subscribed_library"] = "SubscribedLibrary"
	fields["subscribed_library_name"] = bindings.NewStringType()
	fieldNameMap["subscribed_library_name"] = "SubscribedLibraryName"
	fields["subscribed_library_location"] = bindings.NewEnumType("com.vmware.content.library.subscriptions.location", reflect.TypeOf(SubscriptionsLocation(SubscriptionsLocation_LOCAL)))
	fieldNameMap["subscribed_library_location"] = "SubscribedLibraryLocation"
	fields["subscribed_library_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(SubscriptionsVcenterInfoBindingType))
	fieldNameMap["subscribed_library_vcenter"] = "SubscribedLibraryVcenter"
	fields["subscribed_library_placement"] = bindings.NewReferenceType(SubscriptionsPlacementInfoBindingType)
	fieldNameMap["subscribed_library_placement"] = "SubscribedLibraryPlacement"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("subscribed_library_location",
		map[string][]bindings.FieldData{
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("subscribed_library_vcenter", true),
			},
			"LOCAL": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.subscriptions.info", fields, reflect.TypeOf(SubscriptionsInfo{}), fieldNameMap, validators)
}


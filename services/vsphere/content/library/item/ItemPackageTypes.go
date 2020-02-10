/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.content.library.item.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/updatesession"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"net/url"
	"time"
)


// The ``TransferStatus`` enumeration class defines the transfer state of a file.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TransferStatus string

const (
    // Indicates that a file has been defined for a library item and its content needs to be uploaded.
	TransferStatus_WAITING_FOR_TRANSFER TransferStatus = "WAITING_FOR_TRANSFER"
    // Indicates that data is being transferred to the file.
	TransferStatus_TRANSFERRING TransferStatus = "TRANSFERRING"
    // Indicates that the file has been fully transferred and is ready to be used.
	TransferStatus_READY TransferStatus = "READY"
    // Indicates that the file is being validated (checksum, type adapters).
	TransferStatus_VALIDATING TransferStatus = "VALIDATING"
    // Indicates that there was an error transferring or validating the file.
	TransferStatus_ERROR TransferStatus = "ERROR"
)

func (t TransferStatus) TransferStatus() bool {
	switch t {
	case TransferStatus_WAITING_FOR_TRANSFER:
		return true
	case TransferStatus_TRANSFERRING:
		return true
	case TransferStatus_READY:
		return true
	case TransferStatus_VALIDATING:
		return true
	case TransferStatus_ERROR:
		return true
	default:
		return false
	}
}


// The ``DownloadSessionModel`` class provides information on an active DownloadSession resource.
type DownloadSessionModel struct {
    // The identifier of this download session.
	Id *string
    // The identifier of the library item whose content is being downloaded.
	LibraryItemId *string
    // The content version of the library item whose content is being downloaded. This value is the library.ItemModel#contentVersion at the time when the session is created for the library item.
	LibraryItemContentVersion *string
    // If the session is in the DownloadSessionModelState#State_ERROR status this property will have more details about the error.
	ErrorMessage *std.LocalizableMessage
    // The progress that has been made with the download. This property is to be updated by the client during the download process to indicate the progress of its work in completing the download. The initial progress is 0 until updated by the client. The maximum value is 100, which indicates that the download is complete.
	ClientProgress *int64
    // The current state (ACTIVE, CANCELED, ERROR) of the download session.
	State *DownloadSessionModelState
    // Indicates the time after which the session will expire. The session is guaranteed not to expire before this time.
	ExpirationTime *time.Time
}

// The state of the download session.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DownloadSessionModelState string

const (
    // The session is active. Individual files may be in the process of being transferred and may become ready for download at different times.
	DownloadSessionModelState_ACTIVE DownloadSessionModelState = "ACTIVE"
    // The session has been canceled. On-going downloads may fail. The session will stay in this state until it is either deleted by the user or automatically cleaned up by the Content Library Service.
	DownloadSessionModelState_CANCELED DownloadSessionModelState = "CANCELED"
    // Indicates there was an error during the session lifecycle.
	DownloadSessionModelState_ERROR DownloadSessionModelState = "ERROR"
)

func (s DownloadSessionModelState) DownloadSessionModelState() bool {
	switch s {
	case DownloadSessionModelState_ACTIVE:
		return true
	case DownloadSessionModelState_CANCELED:
		return true
	case DownloadSessionModelState_ERROR:
		return true
	default:
		return false
	}
}


// The ``TransferEndpoint`` class encapsulates a URI along with extra information about it.
type TransferEndpoint struct {
    // Transfer endpoint URI. The supported URI schemes are: ``http``, ``https``, ``file``, and ``ds``. 
    //
    //  An endpoint URI with the ``ds`` scheme specifies the location of the file on the datastore. The format of the datastore URI is: 
    //
    // * ds:///vmfs/volumes/uuid/path
    //
    //  
    //
    //  Some examples of valid file URI formats are: 
    //
    // * file:///path
    // * file:///C:/path
    // * file://unc-server/path
    //
    //  
    //
    //  When the transfer endpoint is a file or datastore location, the server can import the file directly from the storage backing without the overhead of streaming over HTTP.
	Uri url.URL
    // Thumbprint of the expected SSL certificate for this endpoint. Only used for HTTPS connections. The thumbprint is the SHA-1 hash of the DER encoding of the remote endpoint's SSL certificate. If set, the remote endpoint's SSL certificate is only accepted if it matches this thumbprint, and no other certificate validation is performed.
	SslCertificateThumbprint *string
}

// The ``UpdateSessionModel`` class provides information on an active UpdateSession resource.
type UpdateSessionModel struct {
    // The identifier of this update session.
	Id *string
    // The identifier of the library item to which content will be uploaded or removed.
	LibraryItemId *string
    // The content version of the library item whose content is being modified. This value is the library.ItemModel#contentVersion at the time when the session is created for the library item.
	LibraryItemContentVersion *string
    // If the session is in the UpdateSessionModelState#State_ERROR status this property will have more details about the error.
	ErrorMessage *std.LocalizableMessage
    // The progress that has been made with the upload. This property is to be updated by the client during the upload process to indicate the progress of its work in completing the upload. The initial progress is 0 until updated by the client. The maximum value is 100, which indicates that the update is complete.
	ClientProgress *int64
    // The current state (ACTIVE, DONE, ERROR, CANCELED) of the update session. This property was added in vSphere API 6.8.
	State *UpdateSessionModelState
    // Indicates the time after which the session will expire. The session is guaranteed not to expire earlier than this time.
	ExpirationTime *time.Time
    // A preview of the files currently being uploaded in the session. This property will be set only when the session is in the UpdateSessionModelState#State_ACTIVE. This property was added in vSphere API 6.8.
	PreviewInfo *updatesession.PreviewInfo
    // Indicates the update session behavior if warnings are raised in the session preview. Any warning which is raised by session preview but not ignored by the client will, by default, fail the update session. This property was added in vSphere API 6.8.
	WarningBehavior []updatesession.WarningBehavior
}

// The state of an update session.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type UpdateSessionModelState string

const (
    // The session is currently active. This is the initial state when the session is created. Files may be uploaded by the client or pulled by the Content Library Service at this stage.
	UpdateSessionModelState_ACTIVE UpdateSessionModelState = "ACTIVE"
    // The session is done and all its effects are now visible.
	UpdateSessionModelState_DONE UpdateSessionModelState = "DONE"
    // There was an error during the session.
	UpdateSessionModelState_ERROR UpdateSessionModelState = "ERROR"
    // The session has been canceled.
	UpdateSessionModelState_CANCELED UpdateSessionModelState = "CANCELED"
)

func (s UpdateSessionModelState) UpdateSessionModelState() bool {
	switch s {
	case UpdateSessionModelState_ACTIVE:
		return true
	case UpdateSessionModelState_DONE:
		return true
	case UpdateSessionModelState_ERROR:
		return true
	case UpdateSessionModelState_CANCELED:
		return true
	default:
		return false
	}
}





func DownloadSessionModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, ""))
	fieldNameMap["id"] = "Id"
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fields["library_item_content_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["library_item_content_version"] = "LibraryItemContentVersion"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["client_progress"] = "ClientProgress"
	fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.download_session_model.state", reflect.TypeOf(DownloadSessionModelState(DownloadSessionModelState_ACTIVE))))
	fieldNameMap["state"] = "State"
	fields["expiration_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["expiration_time"] = "ExpirationTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.download_session_model", fields, reflect.TypeOf(DownloadSessionModel{}), fieldNameMap, validators)
}

func TransferEndpointBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["uri"] = bindings.NewUriType()
	fieldNameMap["uri"] = "Uri"
	fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.transfer_endpoint", fields, reflect.TypeOf(TransferEndpoint{}), fieldNameMap, validators)
}

func UpdateSessionModelBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, ""))
	fieldNameMap["id"] = "Id"
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fields["library_item_content_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["library_item_content_version"] = "LibraryItemContentVersion"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["client_progress"] = "ClientProgress"
	fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.update_session_model.state", reflect.TypeOf(UpdateSessionModelState(UpdateSessionModelState_ACTIVE))))
	fieldNameMap["state"] = "State"
	fields["expiration_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["expiration_time"] = "ExpirationTime"
	fields["preview_info"] = bindings.NewOptionalType(bindings.NewReferenceType(updatesession.PreviewInfoBindingType))
	fieldNameMap["preview_info"] = "PreviewInfo"
	fields["warning_behavior"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(updatesession.WarningBehaviorBindingType), reflect.TypeOf([]updatesession.WarningBehavior{})))
	fieldNameMap["warning_behavior"] = "WarningBehavior"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("state",
		map[string][]bindings.FieldData{
			"ACTIVE": []bindings.FieldData{
				bindings.NewFieldData("preview_info", true),
			},
			"DONE": []bindings.FieldData{},
			"ERROR": []bindings.FieldData{},
			"CANCELED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.item.update_session_model", fields, reflect.TypeOf(UpdateSessionModel{}), fieldNameMap, validators)
}



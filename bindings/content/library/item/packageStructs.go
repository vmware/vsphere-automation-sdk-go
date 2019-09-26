/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.library.item.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library/item/updatesession"
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
    Id *string
    LibraryItemId *string
    LibraryItemContentVersion *string
    ErrorMessage *std.LocalizableMessage
    ClientProgress *int64
    State *DownloadSessionModel_State
    ExpirationTime *time.Time
}




    
    // The state of the download session.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DownloadSessionModel_State string

    const (
        // The session is active. Individual files may be in the process of being transferred and may become ready for download at different times.
         DownloadSessionModel_State_ACTIVE DownloadSessionModel_State = "ACTIVE"
        // The session has been canceled. On-going downloads may fail. The session will stay in this state until it is either deleted by the user or automatically cleaned up by the Content Library Service.
         DownloadSessionModel_State_CANCELED DownloadSessionModel_State = "CANCELED"
        // Indicates there was an error during the session lifecycle.
         DownloadSessionModel_State_ERROR DownloadSessionModel_State = "ERROR"
    )

    func (s DownloadSessionModel_State) DownloadSessionModel_State() bool {
        switch s {
            case DownloadSessionModel_State_ACTIVE:
                return true
            case DownloadSessionModel_State_CANCELED:
                return true
            case DownloadSessionModel_State_ERROR:
                return true
            default:
                return false
        }
    }



// The ``TransferEndpoint`` class encapsulates a URI along with extra information about it.
type TransferEndpoint struct {
    Uri url.URL
    SslCertificateThumbprint *string
}






// The ``UpdateSessionModel`` class provides information on an active UpdateSession resource.
type UpdateSessionModel struct {
    Id *string
    LibraryItemId *string
    LibraryItemContentVersion *string
    ErrorMessage *std.LocalizableMessage
    ClientProgress *int64
    State *UpdateSessionModel_State
    ExpirationTime *time.Time
    PreviewInfo *updatesession.PreviewInfo
    WarningBehavior []updatesession.WarningBehavior
}




    
    // The state of an update session.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type UpdateSessionModel_State string

    const (
        // The session is currently active. This is the initial state when the session is created. Files may be uploaded by the client or pulled by the Content Library Service at this stage.
         UpdateSessionModel_State_ACTIVE UpdateSessionModel_State = "ACTIVE"
        // The session is done and all its effects are now visible.
         UpdateSessionModel_State_DONE UpdateSessionModel_State = "DONE"
        // There was an error during the session.
         UpdateSessionModel_State_ERROR UpdateSessionModel_State = "ERROR"
        // The session has been canceled.
         UpdateSessionModel_State_CANCELED UpdateSessionModel_State = "CANCELED"
    )

    func (s UpdateSessionModel_State) UpdateSessionModel_State() bool {
        switch s {
            case UpdateSessionModel_State_ACTIVE:
                return true
            case UpdateSessionModel_State_DONE:
                return true
            case UpdateSessionModel_State_ERROR:
                return true
            case UpdateSessionModel_State_CANCELED:
                return true
            default:
                return false
        }
    }







func DownloadSessionModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.item.DownloadSession"}, ""))
    fieldNameMap["id"] = "Id"
    fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fields["library_item_content_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["library_item_content_version"] = "LibraryItemContentVersion"
    fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["error_message"] = "ErrorMessage"
    fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["client_progress"] = "ClientProgress"
    fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.download_session_model.state", reflect.TypeOf(DownloadSessionModel_State(DownloadSessionModel_State_ACTIVE))))
    fieldNameMap["state"] = "State"
    fields["expiration_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["expiration_time"] = "ExpirationTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.download_session_model",fields, reflect.TypeOf(DownloadSessionModel{}), fieldNameMap, validators)
}

func TransferEndpointBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uri"] = bindings.NewUriType()
    fieldNameMap["uri"] = "Uri"
    fields["ssl_certificate_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_certificate_thumbprint"] = "SslCertificateThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.transfer_endpoint",fields, reflect.TypeOf(TransferEndpoint{}), fieldNameMap, validators)
}

func UpdateSessionModelBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.item.UpdateSession"}, ""))
    fieldNameMap["id"] = "Id"
    fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fields["library_item_content_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["library_item_content_version"] = "LibraryItemContentVersion"
    fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["error_message"] = "ErrorMessage"
    fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["client_progress"] = "ClientProgress"
    fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.update_session_model.state", reflect.TypeOf(UpdateSessionModel_State(UpdateSessionModel_State_ACTIVE))))
    fieldNameMap["state"] = "State"
    fields["expiration_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["expiration_time"] = "ExpirationTime"
    fields["preview_info"] = bindings.NewOptionalType(bindings.NewReferenceType(updatesession.PreviewInfoBindingType))
    fieldNameMap["preview_info"] = "PreviewInfo"
    fields["warning_behavior"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(updatesession.WarningBehaviorBindingType), reflect.TypeOf([]updatesession.WarningBehavior{})))
    fieldNameMap["warning_behavior"] = "WarningBehavior"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "ACTIVE": []bindings.FieldData {
                 bindings.NewFieldData("preview_info", true),
            },
            "DONE": []bindings.FieldData {},
            "ERROR": []bindings.FieldData {},
            "CANCELED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.item.update_session_model",fields, reflect.TypeOf(UpdateSessionModel{}), fieldNameMap, validators)
}



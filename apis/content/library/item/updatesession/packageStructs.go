/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.content.library.item.updatesession.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package updatesession

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``WarningType`` enumeration class defines the warnings which can be raised during the update session.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type WarningType string

const (
    // The certificate used for signing the content is self-signed.
     WarningType_SELF_SIGNED_CERTIFICATE WarningType = "SELF_SIGNED_CERTIFICATE"
    // The certificate used for signing the content is expired.
     WarningType_EXPIRED_CERTIFICATE WarningType = "EXPIRED_CERTIFICATE"
    // The certificate used for signing the content is not yet valid.
     WarningType_NOT_YET_VALID_CERTIFICATE WarningType = "NOT_YET_VALID_CERTIFICATE"
    // The certificate used for signing the content is not trusted.
     WarningType_UNTRUSTED_CERTIFICATE WarningType = "UNTRUSTED_CERTIFICATE"
)

func (w WarningType) WarningType() bool {
    switch w {
        case WarningType_SELF_SIGNED_CERTIFICATE:
            return true
        case WarningType_EXPIRED_CERTIFICATE:
            return true
        case WarningType_NOT_YET_VALID_CERTIFICATE:
            return true
        case WarningType_UNTRUSTED_CERTIFICATE:
            return true
        default:
            return false
    }
}





// The ``CertificateInfo`` class contains information about the public key certificate used to sign the content.
type CertificateInfo struct {
    Issuer string
    Subject string
    SelfSigned bool
    X509 string
}






// The ``PreviewInfo`` class contains information about the files being uploaded in the update session.
type PreviewInfo struct {
    State PreviewInfo_State
    CertificateInfo *CertificateInfo
    Warnings []PreviewWarningInfo
}




    
    // The ``State`` enumeration class defines the state of the update session's preview.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type PreviewInfo_State string

    const (
        // There are no files in the update session OR a preview is not possible for the files currently in the update session. However, preview may be possible after metadata files such as OVF descriptor are added to the session. In this case the state will transition to ``PREPARING``.
         PreviewInfo_State_UNAVAILABLE PreviewInfo_State = "UNAVAILABLE"
        // Preview is not possible for this update session. This state is reached when there are no metadata files in the update session and user invokes a session complete operation.
         PreviewInfo_State_NOT_APPLICABLE PreviewInfo_State = "NOT_APPLICABLE"
        // A preview is being prepared for the files currently in the update session. This state is reached when the applicable metadata files are added to the update session but their content is not fully uploaded yet. For OVF item type, this state indicates that the OVF descriptor file is currently being uploaded.
         PreviewInfo_State_PREPARING PreviewInfo_State = "PREPARING"
        // Preview is available for this update session. It is possible to review certificate details and warnings, if any. This state is reached when the applicable metadata files in the session have been fully uploaded.
         PreviewInfo_State_AVAILABLE PreviewInfo_State = "AVAILABLE"
    )

    func (s PreviewInfo_State) PreviewInfo_State() bool {
        switch s {
            case PreviewInfo_State_UNAVAILABLE:
                return true
            case PreviewInfo_State_NOT_APPLICABLE:
                return true
            case PreviewInfo_State_PREPARING:
                return true
            case PreviewInfo_State_AVAILABLE:
                return true
            default:
                return false
        }
    }



// The ``PreviewWarningInfo`` class provides information about the warnings which are raised during the update session preview.
type PreviewWarningInfo struct {
    Type_ WarningType
    Message std.LocalizableMessage
    Ignored bool
}






// The ``WarningBehavior`` class defines the session behavior if the warning is raised during the update session.
type WarningBehavior struct {
    Type_ WarningType
    Ignored bool
}










func CertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["subject"] = bindings.NewStringType()
    fieldNameMap["subject"] = "Subject"
    fields["self_signed"] = bindings.NewBooleanType()
    fieldNameMap["self_signed"] = "SelfSigned"
    fields["x509"] = bindings.NewStringType()
    fieldNameMap["x509"] = "X509"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.certificate_info",fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func PreviewInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.preview_info.state", reflect.TypeOf(PreviewInfo_State(PreviewInfo_State_UNAVAILABLE)))
    fieldNameMap["state"] = "State"
    fields["certificate_info"] = bindings.NewOptionalType(bindings.NewReferenceType(CertificateInfoBindingType))
    fieldNameMap["certificate_info"] = "CertificateInfo"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PreviewWarningInfoBindingType), reflect.TypeOf([]PreviewWarningInfo{})))
    fieldNameMap["warnings"] = "Warnings"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "AVAILABLE": []bindings.FieldData {
                 bindings.NewFieldData("certificate_info", false),
                 bindings.NewFieldData("warnings", true),
            },
            "UNAVAILABLE": []bindings.FieldData {},
            "NOT_APPLICABLE": []bindings.FieldData {},
            "PREPARING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.preview_info",fields, reflect.TypeOf(PreviewInfo{}), fieldNameMap, validators)
}

func PreviewWarningInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.warning_type", reflect.TypeOf(WarningType(WarningType_SELF_SIGNED_CERTIFICATE)))
    fieldNameMap["type"] = "Type_"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["ignored"] = bindings.NewBooleanType()
    fieldNameMap["ignored"] = "Ignored"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.preview_warning_info",fields, reflect.TypeOf(PreviewWarningInfo{}), fieldNameMap, validators)
}

func WarningBehaviorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.warning_type", reflect.TypeOf(WarningType(WarningType_SELF_SIGNED_CERTIFICATE)))
    fieldNameMap["type"] = "Type_"
    fields["ignored"] = bindings.NewBooleanType()
    fieldNameMap["ignored"] = "Ignored"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.updatesession.warning_behavior",fields, reflect.TypeOf(WarningBehavior{}), fieldNameMap, validators)
}



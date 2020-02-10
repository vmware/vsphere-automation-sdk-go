/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
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


// The ``WarningType`` enumeration class defines the warnings which can be raised during the update session. This enumeration was added in vSphere API 6.8.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type WarningType string

const (
    // The certificate used for signing the content is self-signed. This constant field was added in vSphere API 6.8.
	WarningType_SELF_SIGNED_CERTIFICATE WarningType = "SELF_SIGNED_CERTIFICATE"
    // The certificate used for signing the content is expired. This constant field was added in vSphere API 6.8.
	WarningType_EXPIRED_CERTIFICATE WarningType = "EXPIRED_CERTIFICATE"
    // The certificate used for signing the content is not yet valid. This constant field was added in vSphere API 6.8.
	WarningType_NOT_YET_VALID_CERTIFICATE WarningType = "NOT_YET_VALID_CERTIFICATE"
    // The certificate used for signing the content is not trusted. This constant field was added in vSphere API 6.8.
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


// The ``CertificateInfo`` class contains information about the public key certificate used to sign the content. This class was added in vSphere API 6.8.
type CertificateInfo struct {
    // Certificate issuer. For example: /C=US/ST=California/L=Palo Alto/O=VMware, Inc. This property was added in vSphere API 6.8.
	Issuer string
    // Certificate subject. For example: C=US/ST=Massachusetts/L=Hopkinton/O=EMC Corporation/OU=EMC Avamar/CN=EMC Corporation. This property was added in vSphere API 6.8.
	Subject string
    // Whether the certificate is self-signed. This property was added in vSphere API 6.8.
	SelfSigned bool
    // The X509 representation of the certificate. This property was added in vSphere API 6.8.
	X509 string
}

// The ``PreviewInfo`` class contains information about the files being uploaded in the update session. This class was added in vSphere API 6.8.
type PreviewInfo struct {
    // Indicates the state of the preview of the update session. This property was added in vSphere API 6.8.
	State PreviewInfoState
    // The certificate information of the signed update session content. This property was added in vSphere API 6.8.
	CertificateInfo *CertificateInfo
    // The list of warnings raised for this update session. Any warning which is not ignored by the client will, by default, fail the update session during session complete operation. This property was added in vSphere API 6.8.
	Warnings []PreviewWarningInfo
}

// The ``State`` enumeration class defines the state of the update session's preview. This enumeration was added in vSphere API 6.8.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PreviewInfoState string

const (
    // There are no files in the update session OR a preview is not possible for the files currently in the update session. However, preview may be possible after metadata files such as OVF descriptor are added to the session. In this case the state will transition to ``PREPARING``. This constant field was added in vSphere API 6.8.
	PreviewInfoState_UNAVAILABLE PreviewInfoState = "UNAVAILABLE"
    // Preview is not possible for this update session. This state is reached when there are no metadata files in the update session and user invokes a session complete operation. This constant field was added in vSphere API 6.8.
	PreviewInfoState_NOT_APPLICABLE PreviewInfoState = "NOT_APPLICABLE"
    // A preview is being prepared for the files currently in the update session. This state is reached when the applicable metadata files are added to the update session but their content is not fully uploaded yet. For OVF item type, this state indicates that the OVF descriptor file is currently being uploaded. This constant field was added in vSphere API 6.8.
	PreviewInfoState_PREPARING PreviewInfoState = "PREPARING"
    // Preview is available for this update session. It is possible to review certificate details and warnings, if any. This state is reached when the applicable metadata files in the session have been fully uploaded. This constant field was added in vSphere API 6.8.
	PreviewInfoState_AVAILABLE PreviewInfoState = "AVAILABLE"
)

func (s PreviewInfoState) PreviewInfoState() bool {
	switch s {
	case PreviewInfoState_UNAVAILABLE:
		return true
	case PreviewInfoState_NOT_APPLICABLE:
		return true
	case PreviewInfoState_PREPARING:
		return true
	case PreviewInfoState_AVAILABLE:
		return true
	default:
		return false
	}
}


// The ``PreviewWarningInfo`` class provides information about the warnings which are raised during the update session preview. This class was added in vSphere API 6.8.
type PreviewWarningInfo struct {
    // The warning type raised during preview of the update session. This property was added in vSphere API 6.8.
	Type_ WarningType
    // The message specifying more details about the warning. This property was added in vSphere API 6.8.
	Message std.LocalizableMessage
    // Indicates if this warning will be ignored during session complete operation. This property was added in vSphere API 6.8.
	Ignored bool
}

// The ``WarningBehavior`` class defines the session behavior if the warning is raised during the update session. This class was added in vSphere API 6.8.
type WarningBehavior struct {
    // The warning type which may be raised during the update session. This property was added in vSphere API 6.8.
	Type_ WarningType
    // Indicates if this warning will be ignored during session complete operation. This property was added in vSphere API 6.8.
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
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.certificate_info", fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func PreviewInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["state"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.preview_info.state", reflect.TypeOf(PreviewInfoState(PreviewInfoState_UNAVAILABLE)))
	fieldNameMap["state"] = "State"
	fields["certificate_info"] = bindings.NewOptionalType(bindings.NewReferenceType(CertificateInfoBindingType))
	fieldNameMap["certificate_info"] = "CertificateInfo"
	fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PreviewWarningInfoBindingType), reflect.TypeOf([]PreviewWarningInfo{})))
	fieldNameMap["warnings"] = "Warnings"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("state",
		map[string][]bindings.FieldData{
			"AVAILABLE": []bindings.FieldData{
				bindings.NewFieldData("certificate_info", false),
				bindings.NewFieldData("warnings", true),
			},
			"UNAVAILABLE": []bindings.FieldData{},
			"NOT_APPLICABLE": []bindings.FieldData{},
			"PREPARING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.preview_info", fields, reflect.TypeOf(PreviewInfo{}), fieldNameMap, validators)
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
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.preview_warning_info", fields, reflect.TypeOf(PreviewWarningInfo{}), fieldNameMap, validators)
}

func WarningBehaviorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.warning_type", reflect.TypeOf(WarningType(WarningType_SELF_SIGNED_CERTIFICATE)))
	fieldNameMap["type"] = "Type_"
	fields["ignored"] = bindings.NewBooleanType()
	fieldNameMap["ignored"] = "Ignored"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.warning_behavior", fields, reflect.TypeOf(WarningBehavior{}), fieldNameMap, validators)
}



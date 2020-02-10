/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.appliance.update.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"time"
)


// The ``ServiceInfo`` class describes a service to be stopped and started during the update installation. This class was added in vSphere API 6.7.
type ServiceInfo struct {
    // Service ID. This property was added in vSphere API 6.7.
	Service string
    // Service description. This property was added in vSphere API 6.7.
	Description std.LocalizableMessage
}

// The ``CommonInfo`` class defines common update information. This class was added in vSphere API 6.7.
type CommonInfo struct {
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5". This property was added in vSphere API 6.7.
	Description std.LocalizableMessage
    // Update priority. This property was added in vSphere API 6.7.
	Priority CommonInfoPriority
    // Update severity. This property was added in vSphere API 6.7.
	Severity CommonInfoSeverity
    // Update category. This property was added in vSphere API 6.7.
	UpdateType CommonInfoCategory
    // Update release date. This property was added in vSphere API 6.7.
	ReleaseDate time.Time
    // Flag indicating whether reboot is required after update. This property was added in vSphere API 6.7.
	RebootRequired bool
    // Download Size of update in Megabytes. This property was added in vSphere API 6.7.
	Size int64
}

// The ``Priority`` enumeration class defines the update installation priority recommendations. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CommonInfoPriority string

const (
    // Install ASAP. This constant field was added in vSphere API 6.7.
	CommonInfoPriority_HIGH CommonInfoPriority = "HIGH"
    // Install at the earliest convenience. This constant field was added in vSphere API 6.7.
	CommonInfoPriority_MEDIUM CommonInfoPriority = "MEDIUM"
    // Install at your discretion. This constant field was added in vSphere API 6.7.
	CommonInfoPriority_LOW CommonInfoPriority = "LOW"
)

func (p CommonInfoPriority) CommonInfoPriority() bool {
	switch p {
	case CommonInfoPriority_HIGH:
		return true
	case CommonInfoPriority_MEDIUM:
		return true
	case CommonInfoPriority_LOW:
		return true
	default:
		return false
	}
}


// The ``Severity`` enumeration class defines the severity of the issues fixed in the update. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CommonInfoSeverity string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation. The exploitation results in the complete compromise of confidentiality, integrity, and availability of user data and/or processing resources without user interaction. Exploitation could be leveraged to propagate an Internet worm or execute arbitrary code between Virtual Machines and/or the Host Operating System. This constant field was added in vSphere API 6.7.
	CommonInfoSeverity_CRITICAL CommonInfoSeverity = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers. This rating also applies to those vulnerabilities which could lead to the complete compromise of availability when exploitation is by a remote unauthenticated attacker from the Internet or through a breach of virtual machine isolation. This constant field was added in vSphere API 6.7.
	CommonInfoSeverity_IMPORTANT CommonInfoSeverity = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources. This constant field was added in vSphere API 6.7.
	CommonInfoSeverity_MODERATE CommonInfoSeverity = "MODERATE"
    // All other issues that have a security impact. Vulnerabilities where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact. This constant field was added in vSphere API 6.7.
	CommonInfoSeverity_LOW CommonInfoSeverity = "LOW"
)

func (s CommonInfoSeverity) CommonInfoSeverity() bool {
	switch s {
	case CommonInfoSeverity_CRITICAL:
		return true
	case CommonInfoSeverity_IMPORTANT:
		return true
	case CommonInfoSeverity_MODERATE:
		return true
	case CommonInfoSeverity_LOW:
		return true
	default:
		return false
	}
}


// The ``Category`` enumeration class defines update type. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CommonInfoCategory string

const (
    // Fixes vulnerabilities, doesn't change functionality. This constant field was added in vSphere API 6.7.
	CommonInfoCategory_SECURITY CommonInfoCategory = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality. This constant field was added in vSphere API 6.7.
	CommonInfoCategory_FIX CommonInfoCategory = "FIX"
    // Changes product functionality. This constant field was added in vSphere API 6.7.
	CommonInfoCategory_UPDATE CommonInfoCategory = "UPDATE"
    // Introduces new features, significantly changes product functionality. This constant field was added in vSphere API 6.7.
	CommonInfoCategory_UPGRADE CommonInfoCategory = "UPGRADE"
)

func (c CommonInfoCategory) CommonInfoCategory() bool {
	switch c {
	case CommonInfoCategory_SECURITY:
		return true
	case CommonInfoCategory_FIX:
		return true
	case CommonInfoCategory_UPDATE:
		return true
	case CommonInfoCategory_UPGRADE:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains the essential information about the update. This class was added in vSphere API 6.7.
type Summary struct {
    // Version in form of X.Y.Z.P. e.g. 6.5.1.5400. This property was added in vSphere API 6.7.
	Version string
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5". This property was added in vSphere API 6.7.
	Description std.LocalizableMessage
    // Update priority. This property was added in vSphere API 6.7.
	Priority CommonInfoPriority
    // Update severity. This property was added in vSphere API 6.7.
	Severity CommonInfoSeverity
    // Update category. This property was added in vSphere API 6.7.
	UpdateType CommonInfoCategory
    // Update release date. This property was added in vSphere API 6.7.
	ReleaseDate time.Time
    // Flag indicating whether reboot is required after update. This property was added in vSphere API 6.7.
	RebootRequired bool
    // Download Size of update in Megabytes. This property was added in vSphere API 6.7.
	Size int64
}




func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.service"}, "")
	fieldNameMap["service"] = "Service"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.service_info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

func CommonInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfoPriority(CommonInfoPriority_HIGH)))
	fieldNameMap["priority"] = "Priority"
	fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfoSeverity(CommonInfoSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfoCategory(CommonInfoCategory_SECURITY)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.common_info", fields, reflect.TypeOf(CommonInfo{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfoPriority(CommonInfoPriority_HIGH)))
	fieldNameMap["priority"] = "Priority"
	fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfoSeverity(CommonInfoSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfoCategory(CommonInfoCategory_SECURITY)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.summary", fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}



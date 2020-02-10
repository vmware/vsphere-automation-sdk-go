/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.deployment.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package deployment

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"time"
)


// The ``ApplianceType`` enumeration class defines the vCenter appliance types. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ApplianceType string

const (
    // vCenter Server Appliance with an embedded Platform Services Controller. This constant field was added in vSphere API 6.7.
	ApplianceType_VCSA_EMBEDDED ApplianceType = "VCSA_EMBEDDED"
    // vCenter Server Appliance with an external Platform Services Controller. This constant field was added in vSphere API 6.7.
	ApplianceType_VCSA_EXTERNAL ApplianceType = "VCSA_EXTERNAL"
    // An external Platform Services Controller. This constant field was added in vSphere API 6.7.
	ApplianceType_PSC_EXTERNAL ApplianceType = "PSC_EXTERNAL"
)

func (a ApplianceType) ApplianceType() bool {
	switch a {
	case ApplianceType_VCSA_EMBEDDED:
		return true
	case ApplianceType_VCSA_EXTERNAL:
		return true
	case ApplianceType_PSC_EXTERNAL:
		return true
	default:
		return false
	}
}


// The ``ApplianceState`` enumeration class defines the various states the vCenter Appliance can be in. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ApplianceState string

const (
    // The appliance is in the process of being initialized and not ready for configuration. This constant field was added in vSphere API 6.7.
	ApplianceState_NOT_INITIALIZED ApplianceState = "NOT_INITIALIZED"
    // The appliance is initialized and ready to be configured. This constant field was added in vSphere API 6.7.
	ApplianceState_INITIALIZED ApplianceState = "INITIALIZED"
    // The appliance is in the process of being configured. This constant field was added in vSphere API 6.7.
	ApplianceState_CONFIG_IN_PROGRESS ApplianceState = "CONFIG_IN_PROGRESS"
    // The deployment script has raised a question and is waiting for an answer to continue with the appliance configuration. This constant field was added in vSphere API 6.7.
	ApplianceState_QUESTION_RAISED ApplianceState = "QUESTION_RAISED"
    // The appliance configuration has failed. This constant field was added in vSphere API 6.7.
	ApplianceState_FAILED ApplianceState = "FAILED"
    // The appliance has been configured. This constant field was added in vSphere API 6.7.
	ApplianceState_CONFIGURED ApplianceState = "CONFIGURED"
)

func (a ApplianceState) ApplianceState() bool {
	switch a {
	case ApplianceState_NOT_INITIALIZED:
		return true
	case ApplianceState_INITIALIZED:
		return true
	case ApplianceState_CONFIG_IN_PROGRESS:
		return true
	case ApplianceState_QUESTION_RAISED:
		return true
	case ApplianceState_FAILED:
		return true
	case ApplianceState_CONFIGURED:
		return true
	default:
		return false
	}
}


// The ``Operation`` enumeration class defines the supported vCenter appliance deployment operations. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Operation string

const (
    // The appliance installation is in progress. This constant field was added in vSphere API 6.7.
	Operation_INSTALL Operation = "INSTALL"
    // The appliance upgrade is in progress. This constant field was added in vSphere API 6.7.
	Operation_UPGRADE Operation = "UPGRADE"
    // The appliance migration is in progress. This constant field was added in vSphere API 6.7.
	Operation_MIGRATE Operation = "MIGRATE"
    // The appliance restoration is in progress. This constant field was added in vSphere API 6.7.
	Operation_RESTORE Operation = "RESTORE"
    // The appliance is being rolled back to an unconfigured state. This constant field was added in vSphere API 6.7.
	Operation_ROLLBACK Operation = "ROLLBACK"
)

func (o Operation) Operation() bool {
	switch o {
	case Operation_INSTALL:
		return true
	case Operation_UPGRADE:
		return true
	case Operation_MIGRATE:
		return true
	case Operation_RESTORE:
		return true
	case Operation_ROLLBACK:
		return true
	default:
		return false
	}
}


// The ``VerificationMode`` enumeration class defines the verification modes for SSL certificates or SSH connections. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type VerificationMode string

const (
    // No verification will be performed. This constant field was added in vSphere API 6.7.
	VerificationMode_NONE VerificationMode = "NONE"
    // Passed thumbprint will be used for verification. This constant field was added in vSphere API 6.7.
	VerificationMode_THUMBPRINT VerificationMode = "THUMBPRINT"
)

func (v VerificationMode) VerificationMode() bool {
	switch v {
	case VerificationMode_NONE:
		return true
	case VerificationMode_THUMBPRINT:
		return true
	default:
		return false
	}
}


// The ``CheckStatus`` enumeration class defines the status of the checks. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CheckStatus string

const (
    // All checks have completed successfully. This constant field was added in vSphere API 6.7.
	CheckStatus_SUCCESS CheckStatus = "SUCCESS"
    // A fatal error was encountered when running the sanity checks. This constant field was added in vSphere API 6.7.
	CheckStatus_FAILED CheckStatus = "FAILED"
)

func (c CheckStatus) CheckStatus() bool {
	switch c {
	case CheckStatus_SUCCESS:
		return true
	case CheckStatus_FAILED:
		return true
	default:
		return false
	}
}


// The ``HistoryMigrationOption`` enumeration class defines the vCenter history migration option choices. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HistoryMigrationOption string

const (
    // Only event data and task data will be migrated along with the core data. This constant field was added in vSphere API 6.7.
	HistoryMigrationOption_EVENTS_TASKS HistoryMigrationOption = "EVENTS_TASKS"
    // All history data will be migrated along with the core data. This constant field was added in vSphere API 6.7.
	HistoryMigrationOption_ALL HistoryMigrationOption = "ALL"
)

func (h HistoryMigrationOption) HistoryMigrationOption() bool {
	switch h {
	case HistoryMigrationOption_EVENTS_TASKS:
		return true
	case HistoryMigrationOption_ALL:
		return true
	default:
		return false
	}
}


// The ``Notification`` class contains properties to describe any info/warning/error messages that Tasks can raise. This class was added in vSphere API 6.7.
type Notification struct {
    // The notification id. This property was added in vSphere API 6.7.
	Id string
    // The time the notification was raised/found. This property was added in vSphere API 6.7.
	Time *time.Time
    // The notification message. This property was added in vSphere API 6.7.
	Message std.LocalizableMessage
    // The resolution message, if any. This property was added in vSphere API 6.7.
	Resolution *std.LocalizableMessage
}

// The ``Notifications`` class contains info/warning/error messages that can be reported be the task. This class was added in vSphere API 6.7.
type Notifications struct {
    // Info notification messages reported. This property was added in vSphere API 6.7.
	Info []Notification
    // Warning notification messages reported. This property was added in vSphere API 6.7.
	Warnings []Notification
    // Error notification messages reported. This property was added in vSphere API 6.7.
	Errors []Notification
}

// The ``StandaloneSpec`` class contains information used to configure a standalone embedded vCenter Server appliance. This class was added in vSphere API 6.7.
type StandaloneSpec struct {
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // The SSO domain name to be used to configure this appliance. This property was added in vSphere API 6.7.
	SsoDomainName *string
}

// The ``StandalonePscSpec`` class contains information used to configure a standalone PSC appliance. This class was added in vSphere API 6.7.
type StandalonePscSpec struct {
    // The SSO site name used for this PSC. This property was added in vSphere API 6.7.
	SsoSiteName *string
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // The SSO domain name to be used to configure this appliance. This property was added in vSphere API 6.7.
	SsoDomainName *string
}

// The ``ReplicatedSpec`` class contains information used to check if the configuring vCenter Server can be replicated to the remote PSC. This class was added in vSphere API 6.7.
type ReplicatedSpec struct {
    // The IP address or DNS resolvable name of the partner PSC appliance. This property was added in vSphere API 6.7.
	PartnerHostname string
    // The HTTPS port of the external PSC appliance. This property was added in vSphere API 6.7.
	HttpsPort *int64
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // SHA1 thumbprint of the server SSL certificate will be used for verification. This property was added in vSphere API 6.7.
	SslThumbprint *string
    // SSL verification should be enabled or disabled. This property was added in vSphere API 6.7.
	SslVerify *bool
}

// The ``ReplicatedPscSpec`` class contains information used to check if the configuring PSC can be replicated to the remote PSC. This class was added in vSphere API 6.7.
type ReplicatedPscSpec struct {
    // The SSO sitename that will be used in PSC replication. This property was added in vSphere API 6.7.
	SsoSiteName *string
    // The IP address or DNS resolvable name of the partner PSC appliance. This property was added in vSphere API 6.7.
	PartnerHostname string
    // The HTTPS port of the external PSC appliance. This property was added in vSphere API 6.7.
	HttpsPort *int64
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // SHA1 thumbprint of the server SSL certificate will be used for verification. This property was added in vSphere API 6.7.
	SslThumbprint *string
    // SSL verification should be enabled or disabled. This property was added in vSphere API 6.7.
	SslVerify *bool
}

// The ``RemotePscSpec`` class contains information used to configure an external vCenter Server that registers with a remote PSC. This class was added in vSphere API 6.7.
type RemotePscSpec struct {
    // The IP address or DNS resolvable name of the remote PSC to which this configuring vCenter Server will be registered to. This property was added in vSphere API 6.7.
	PscHostname string
    // The HTTPS port of the external PSC appliance. This property was added in vSphere API 6.7.
	HttpsPort *int64
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // SHA1 thumbprint of the server SSL certificate will be used for verification when ssl_verify field is set to true. This property was added in vSphere API 6.7.
	SslThumbprint *string
    // SSL verification should be enabled or disabled. If ``sslVerify`` is true and and ``sslThumbprint`` is null, the CA certificate will be used for verification. If ``sslVerify`` is true and ``sslThumbprint`` is set then the thumbprint will be used for verification. No verification will be performed if ``sslVerify`` value is set to false. This property was added in vSphere API 6.7.
	SslVerify *bool
}

// The ``CheckInfo`` class describes the result of the appliance deployment check. This class was added in vSphere API 6.7.
type CheckInfo struct {
    // Status of the check. This property was added in vSphere API 6.7.
	Status CheckStatus
    // Result of the check. This property was added in vSphere API 6.7.
	Result *Notifications
}

// The ``HistoryMigrationSpec`` enumeration class defines how vCenter history data will be migrated. vCenter History data includes 
//
// * Statistics
// * Events
// * Tasks
//
// . This class was added in vSphere API 6.7.
type HistoryMigrationSpec struct {
    // Defines what part of vCenter historical data will be migrated along with core data. This property was added in vSphere API 6.7.
	DataSet HistoryMigrationOption
    // Defines how vCenter history will be migrated. If set to true, vCenter history will be migrated separately after successful upgrade or migration, otherwise it will be migrated along with core data during the upgrade or migration process. This property was added in vSphere API 6.7.
	DeferImport *bool
}

// The ``LocationSpec`` class is used to pass the container ESXi or vCenter server of the VM to patch the size of this appliance. This class was added in vSphere API 6.7.
type LocationSpec struct {
    // The IP address or DNS resolvable name of the container. This property was added in vSphere API 6.7.
	Hostname string
    // The HTTPS port of the container. This property was added in vSphere API 6.7.
	HttpsPort *int64
    // SHA1 thumbprint of the server SSL certificate will be used for verification. This property was added in vSphere API 6.7.
	SslThumbprint *string
    // SSL verification should be enabled or disabled. If ``sslVerify`` is true and and ``sslThumbprint`` is null, the CA certificate will be used for verification. If ``sslVerify`` is true and ``sslThumbprint`` is set then the thumbprint will be used for verification. No verification will be performed if ``sslVerify`` value is set to false. This property was added in vSphere API 6.7.
	SslVerify *bool
    // The administrator account on the host. This property was added in vSphere API 6.7.
	Username string
    // The administrator account password. This property was added in vSphere API 6.7.
	Password string
}




func NotificationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["time"] = "Time"
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["resolution"] = "Resolution"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.notification", fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
}

func NotificationsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
	fieldNameMap["info"] = "Info"
	fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
	fieldNameMap["warnings"] = "Warnings"
	fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
	fieldNameMap["errors"] = "Errors"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.notifications", fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}

func StandaloneSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sso_admin_password"] = bindings.NewStringType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["sso_domain_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain_name"] = "SsoDomainName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.standalone_spec", fields, reflect.TypeOf(StandaloneSpec{}), fieldNameMap, validators)
}

func StandalonePscSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_site_name"] = "SsoSiteName"
	fields["sso_admin_password"] = bindings.NewStringType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["sso_domain_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_domain_name"] = "SsoDomainName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.standalone_psc_spec", fields, reflect.TypeOf(StandalonePscSpec{}), fieldNameMap, validators)
}

func ReplicatedSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["partner_hostname"] = bindings.NewStringType()
	fieldNameMap["partner_hostname"] = "PartnerHostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["sso_admin_password"] = bindings.NewStringType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssl_verify"] = "SslVerify"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.replicated_spec", fields, reflect.TypeOf(ReplicatedSpec{}), fieldNameMap, validators)
}

func ReplicatedPscSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["sso_site_name"] = "SsoSiteName"
	fields["partner_hostname"] = bindings.NewStringType()
	fieldNameMap["partner_hostname"] = "PartnerHostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["sso_admin_password"] = bindings.NewStringType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssl_verify"] = "SslVerify"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.replicated_psc_spec", fields, reflect.TypeOf(ReplicatedPscSpec{}), fieldNameMap, validators)
}

func RemotePscSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["psc_hostname"] = bindings.NewStringType()
	fieldNameMap["psc_hostname"] = "PscHostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["sso_admin_password"] = bindings.NewStringType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssl_verify"] = "SslVerify"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.remote_psc_spec", fields, reflect.TypeOf(RemotePscSpec{}), fieldNameMap, validators)
}

func CheckInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.deployment.check_status", reflect.TypeOf(CheckStatus(CheckStatus_SUCCESS)))
	fieldNameMap["status"] = "Status"
	fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(NotificationsBindingType))
	fieldNameMap["result"] = "Result"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.check_info", fields, reflect.TypeOf(CheckInfo{}), fieldNameMap, validators)
}

func HistoryMigrationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data_set"] = bindings.NewEnumType("com.vmware.vcenter.deployment.history_migration_option", reflect.TypeOf(HistoryMigrationOption(HistoryMigrationOption_EVENTS_TASKS)))
	fieldNameMap["data_set"] = "DataSet"
	fields["defer_import"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["defer_import"] = "DeferImport"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.history_migration_spec", fields, reflect.TypeOf(HistoryMigrationSpec{}), fieldNameMap, validators)
}

func LocationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssl_verify"] = "SslVerify"
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.location_spec", fields, reflect.TypeOf(LocationSpec{}), fieldNameMap, validators)
}



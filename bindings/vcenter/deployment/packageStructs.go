/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
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


// The ``ApplianceType`` enumeration class defines the vCenter appliance types.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceType string

const (
    // vCenter Server Appliance with an embedded Platform Services Controller.
     ApplianceType_VCSA_EMBEDDED ApplianceType = "VCSA_EMBEDDED"
    // vCenter Server Appliance with an external Platform Services Controller.
     ApplianceType_VCSA_EXTERNAL ApplianceType = "VCSA_EXTERNAL"
    // An external Platform Services Controller.
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




// The ``ApplianceState`` enumeration class defines the various states the vCenter Appliance can be in.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceState string

const (
    // The appliance is in the process of being initialized and not ready for configuration.
     ApplianceState_NOT_INITIALIZED ApplianceState = "NOT_INITIALIZED"
    // The appliance is initialized and ready to be configured.
     ApplianceState_INITIALIZED ApplianceState = "INITIALIZED"
    // The appliance is in the process of being configured.
     ApplianceState_CONFIG_IN_PROGRESS ApplianceState = "CONFIG_IN_PROGRESS"
    // The deployment script has raised a question and is waiting for an answer to continue with the appliance configuration.
     ApplianceState_QUESTION_RAISED ApplianceState = "QUESTION_RAISED"
    // The appliance configuration has failed.
     ApplianceState_FAILED ApplianceState = "FAILED"
    // The appliance has been configured.
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




// The ``ApplianceSize`` enumeration class defines the vCenter Server Appliance sizes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceSize string

const (
    // Appliance size of 'tiny'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_TINY ApplianceSize = "TINY"
    // Appliance size of 'small'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_SMALL ApplianceSize = "SMALL"
    // Appliance size of 'medium'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_MEDIUM ApplianceSize = "MEDIUM"
    // Appliance size of 'large'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_LARGE ApplianceSize = "LARGE"
    // Appliance size of 'extra large'. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_XLARGE ApplianceSize = "XLARGE"
)

func (a ApplianceSize) ApplianceSize() bool {
    switch a {
        case ApplianceSize_TINY:
            return true
        case ApplianceSize_SMALL:
            return true
        case ApplianceSize_MEDIUM:
            return true
        case ApplianceSize_LARGE:
            return true
        case ApplianceSize_XLARGE:
            return true
        default:
            return false
    }
}




// The ``Operation`` enumeration class defines the supported vCenter appliance deployment operations.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type Operation string

const (
    // The appliance installation is in progress.
     Operation_INSTALL Operation = "INSTALL"
    // The appliance upgrade is in progress.
     Operation_UPGRADE Operation = "UPGRADE"
    // The appliance migration is in progress.
     Operation_MIGRATE Operation = "MIGRATE"
    // The appliance restoration is in progress.
     Operation_RESTORE Operation = "RESTORE"
    // The appliance is being rolled back to an unconfigured state.
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




// The ``VerificationMode`` enumeration class defines the verification modes for SSL certificates or SSH connections.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type VerificationMode string

const (
    // No verification will be performed.
     VerificationMode_NONE VerificationMode = "NONE"
    // Passed thumbprint will be used for verification.
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




// The ``CheckStatus`` enumeration class defines the status of the checks.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type CheckStatus string

const (
    // All checks have completed successfully.
     CheckStatus_SUCCESS CheckStatus = "SUCCESS"
    // A fatal error was encountered when running the sanity checks.
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




// The ``HistoryMigrationOption`` enumeration class defines the vCenter history migration option choices.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type HistoryMigrationOption string

const (
    // Only event data and task data will be migrated along with the core data.
     HistoryMigrationOption_EVENTS_TASKS HistoryMigrationOption = "EVENTS_TASKS"
    // All history data will be migrated along with the core data.
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





// The ``Notification`` class contains properties to describe any info/warning/error messages that Tasks can raise.
type Notification struct {
    Id string
    Time *time.Time
    Message std.LocalizableMessage
    Resolution *std.LocalizableMessage
}






// The ``Notifications`` class contains info/warning/error messages that can be reported be the task.
type Notifications struct {
    Info []Notification
    Warnings []Notification
    Errors []Notification
}






// The ``StandaloneSpec`` class contains information used to configure a standalone embedded vCenter Server appliance.
type StandaloneSpec struct {
    SsoAdminPassword string
    SsoDomainName *string
}






// The ``StandalonePscSpec`` class contains information used to configure a standalone PSC appliance.
type StandalonePscSpec struct {
    SsoSiteName *string
    SsoAdminPassword string
    SsoDomainName *string
}






// The ``ReplicatedSpec`` class contains information used to check if the configuring vCenter Server can be replicated to the remote PSC.
type ReplicatedSpec struct {
    PartnerHostname string
    HttpsPort *int64
    SsoAdminPassword string
    SslThumbprint *string
    SslVerify *bool
}






// The ``ReplicatedPscSpec`` class contains information used to check if the configuring PSC can be replicated to the remote PSC.
type ReplicatedPscSpec struct {
    SsoSiteName *string
    PartnerHostname string
    HttpsPort *int64
    SsoAdminPassword string
    SslThumbprint *string
    SslVerify *bool
}






// The ``RemotePscSpec`` class contains information used to configure an external vCenter Server that registers with a remote PSC.
type RemotePscSpec struct {
    PscHostname string
    HttpsPort *int64
    SsoAdminPassword string
    SslThumbprint *string
    SslVerify *bool
}






// The ``DataMigrationEstimate`` {class contains estimated time and disk space required for the vCenter Server database migration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DataMigrationEstimate struct {
    EstimatedExportTime int64
    EstimatedImportTime int64
    RequiredFreeDiskSpaceOnSource float64
}






// The ``DataMigrationInfo`` {class contains the disk space requirements and time estimates for the different choices available to migrate the vCenter Server data. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DataMigrationInfo struct {
    Core DataMigrationEstimate
    CoreEventsTasks DataMigrationEstimate
    All DataMigrationEstimate
    CoreEventsTasksWithDeferred *DataMigrationEstimate
    AllWithDeferred *DataMigrationEstimate
}






// The ``SourceInfo`` {class contains the information about the source vCenter Server and the database migration options. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SourceInfo struct {
    Hostname string
    Version string
    DeploymentType ApplianceType
    DeploymentSize ApplianceSize
    SsoDomainName string
    ActiveDirectoryDomain string
    DnsServers []string
    DataMigrationInfo *DataMigrationInfo
}






// The ``CheckInfo`` class describes the result of the appliance deployment check.
type CheckInfo struct {
    Status CheckStatus
    Result *Notifications
    SourceInfo *SourceInfo
}






// The ``HistoryMigrationSpec`` enumeration class defines how vCenter history data will be migrated. vCenter History data includes 
//
// * Statistics
// * Events
// * Tasks
type HistoryMigrationSpec struct {
    DataSet HistoryMigrationOption
    DeferImport *bool
}






// The ``LocationSpec`` class is used to pass the container ESXi or vCenter server of the VM to patch the size of this appliance.
type LocationSpec struct {
    Hostname string
    HttpsPort *int64
    SslThumbprint *string
    SslVerify *bool
    Username string
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.notification",fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.notifications",fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}

func StandaloneSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewStringType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.standalone_spec",fields, reflect.TypeOf(StandaloneSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.standalone_psc_spec",fields, reflect.TypeOf(StandalonePscSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.replicated_spec",fields, reflect.TypeOf(ReplicatedSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.replicated_psc_spec",fields, reflect.TypeOf(ReplicatedPscSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.remote_psc_spec",fields, reflect.TypeOf(RemotePscSpec{}), fieldNameMap, validators)
}

func DataMigrationEstimateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["estimated_export_time"] = bindings.NewIntegerType()
    fieldNameMap["estimated_export_time"] = "EstimatedExportTime"
    fields["estimated_import_time"] = bindings.NewIntegerType()
    fieldNameMap["estimated_import_time"] = "EstimatedImportTime"
    fields["required_free_disk_space_on_source"] = bindings.NewDoubleType()
    fieldNameMap["required_free_disk_space_on_source"] = "RequiredFreeDiskSpaceOnSource"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.data_migration_estimate",fields, reflect.TypeOf(DataMigrationEstimate{}), fieldNameMap, validators)
}

func DataMigrationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["core"] = bindings.NewReferenceType(DataMigrationEstimateBindingType)
    fieldNameMap["core"] = "Core"
    fields["core_events_tasks"] = bindings.NewReferenceType(DataMigrationEstimateBindingType)
    fieldNameMap["core_events_tasks"] = "CoreEventsTasks"
    fields["all"] = bindings.NewReferenceType(DataMigrationEstimateBindingType)
    fieldNameMap["all"] = "All"
    fields["core_events_tasks_with_deferred"] = bindings.NewOptionalType(bindings.NewReferenceType(DataMigrationEstimateBindingType))
    fieldNameMap["core_events_tasks_with_deferred"] = "CoreEventsTasksWithDeferred"
    fields["all_with_deferred"] = bindings.NewOptionalType(bindings.NewReferenceType(DataMigrationEstimateBindingType))
    fieldNameMap["all_with_deferred"] = "AllWithDeferred"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.data_migration_info",fields, reflect.TypeOf(DataMigrationInfo{}), fieldNameMap, validators)
}

func SourceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["deployment_type"] = bindings.NewEnumType("com.vmware.vcenter.deployment.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EMBEDDED)))
    fieldNameMap["deployment_type"] = "DeploymentType"
    fields["deployment_size"] = bindings.NewEnumType("com.vmware.vcenter.deployment.appliance_size", reflect.TypeOf(ApplianceSize(ApplianceSize_TINY)))
    fieldNameMap["deployment_size"] = "DeploymentSize"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["active_directory_domain"] = bindings.NewStringType()
    fieldNameMap["active_directory_domain"] = "ActiveDirectoryDomain"
    fields["dns_servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["data_migration_info"] = bindings.NewOptionalType(bindings.NewReferenceType(DataMigrationInfoBindingType))
    fieldNameMap["data_migration_info"] = "DataMigrationInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.source_info",fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

func CheckInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.deployment.check_status", reflect.TypeOf(CheckStatus(CheckStatus_SUCCESS)))
    fieldNameMap["status"] = "Status"
    fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(NotificationsBindingType))
    fieldNameMap["result"] = "Result"
    fields["source_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SourceInfoBindingType))
    fieldNameMap["source_info"] = "SourceInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.check_info",fields, reflect.TypeOf(CheckInfo{}), fieldNameMap, validators)
}

func HistoryMigrationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["data_set"] = bindings.NewEnumType("com.vmware.vcenter.deployment.history_migration_option", reflect.TypeOf(HistoryMigrationOption(HistoryMigrationOption_EVENTS_TASKS)))
    fieldNameMap["data_set"] = "DataSet"
    fields["defer_import"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["defer_import"] = "DeferImport"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.history_migration_spec",fields, reflect.TypeOf(HistoryMigrationSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.location_spec",fields, reflect.TypeOf(LocationSpec{}), fieldNameMap, validators)
}



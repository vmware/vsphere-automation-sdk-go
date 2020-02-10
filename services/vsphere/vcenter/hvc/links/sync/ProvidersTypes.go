/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package sync

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)

// Resource type for Sync Providers. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const Providers_RESOURCE_TYPE = "com.vmware.vcenter.hvc.links.sync.Providers"


// The ``Status`` enumeration class defines valid sync status. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersStatus string

const (
    // If Sync was successful. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersStatus_SUCCEEDED ProvidersStatus = "SUCCEEDED"
    // If Sync failed. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersStatus_FAILED ProvidersStatus = "FAILED"
    // If Sync has not been triggered. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersStatus_NO_SYNC_FOUND ProvidersStatus = "NO_SYNC_FOUND"
)

func (s ProvidersStatus) ProvidersStatus() bool {
	switch s {
	case ProvidersStatus_SUCCEEDED:
		return true
	case ProvidersStatus_FAILED:
		return true
	case ProvidersStatus_NO_SYNC_FOUND:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about sync for a provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersInfo struct {
    // Last sync time for the provider. This indicates the last time that either a background sync or a force sync was started for the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	LastSyncTime *time.Time
    // Last Sync status for the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Status ProvidersStatus
    // Sync Polling interval between local and remote replicas for the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	PollingIntervalInSeconds int64
    // Returns information on the forced sync for the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CurrentSessionInfo *ProvidersSessionInfo
    // Localizable messages associated with sync status. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	StatusMessage *std.LocalizableMessage
}

// The ``SessionInfo`` class contains sync session information. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersSessionInfo struct {
    // Sync stage for the session. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Stage ProvidersSessionInfoStage
    // Completed work for the session. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CompletedWork int64
    // Total work for the session. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TotalWork int64
    // Time at which forced sync session was completed. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CompletionTime *time.Time
    // Time at which force sync was initiated. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	StartTime time.Time
    // Exception message if there is a sync failure on forced sync. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Exception *std.LocalizableMessage
}

// The ``Stage`` class defines the different stages of Sync. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersSessionInfoStage string

const (
    // Changes are being detected on the source replica. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_CHANGE_DETECTION ProvidersSessionInfoStage = "CHANGE_DETECTION"
    // Changes from the source replica are being enumerated. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_CHANGE_ENUMERATION ProvidersSessionInfoStage = "CHANGE_ENUMERATION"
    // Changes are being applied to the destination replica. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_CHANGE_APPLICATION ProvidersSessionInfoStage = "CHANGE_APPLICATION"
    // Sync has completed. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_COMPLETED ProvidersSessionInfoStage = "COMPLETED"
    // Sync failed. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_FAILED ProvidersSessionInfoStage = "FAILED"
    // Session is waiting for progress to be set. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ProvidersSessionInfoStage_WAITING ProvidersSessionInfoStage = "WAITING"
)

func (s ProvidersSessionInfoStage) ProvidersSessionInfoStage() bool {
	switch s {
	case ProvidersSessionInfoStage_CHANGE_DETECTION:
		return true
	case ProvidersSessionInfoStage_CHANGE_ENUMERATION:
		return true
	case ProvidersSessionInfoStage_CHANGE_APPLICATION:
		return true
	case ProvidersSessionInfoStage_COMPLETED:
		return true
	case ProvidersSessionInfoStage_FAILED:
		return true
	case ProvidersSessionInfoStage_WAITING:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains information about a provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersSummary struct {
    // Sync provider id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Provider string
}

// The ``Credentials`` interface specifies user credentials to make a successful connection to remote endpoint. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersCredentials struct {
    // Name of the user to authenticate. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	UserName string
    // Password for the user. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Password string
}



func providersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fieldNameMap["link"] = "Link"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ProvidersSummaryBindingType), reflect.TypeOf([]ProvidersSummary{}))
}

func providersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fieldNameMap["link"] = "Link"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500,"Unauthorized": 403})
}

func providersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.sync.Providers"}, "")
	fieldNameMap["link"] = "Link"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProvidersInfoBindingType)
}

func providersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.sync.Providers"}, "")
	fieldNameMap["link"] = "Link"
	fieldNameMap["provider"] = "Provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403})
}

func providersStartInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.sync.Providers"}, "")
	fieldNameMap["link"] = "Link"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersStartOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func providersStartRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.sync.Providers"}, "")
	fieldNameMap["link"] = "Link"
	fieldNameMap["provider"] = "Provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403,"ResourceBusy": 500})
}


func ProvidersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_sync_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["last_sync_time"] = "LastSyncTime"
	fields["status"] = bindings.NewEnumType("com.vmware.vcenter.hvc.links.sync.providers.status", reflect.TypeOf(ProvidersStatus(ProvidersStatus_SUCCEEDED)))
	fieldNameMap["status"] = "Status"
	fields["polling_interval_in_seconds"] = bindings.NewIntegerType()
	fieldNameMap["polling_interval_in_seconds"] = "PollingIntervalInSeconds"
	fields["current_session_info"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersSessionInfoBindingType))
	fieldNameMap["current_session_info"] = "CurrentSessionInfo"
	fields["status_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["status_message"] = "StatusMessage"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
		map[string][]bindings.FieldData{
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("status_message", true),
			},
			"SUCCEEDED": []bindings.FieldData{},
			"NO_SYNC_FOUND": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.info", fields, reflect.TypeOf(ProvidersInfo{}), fieldNameMap, validators)
}

func ProvidersSessionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["stage"] = bindings.NewEnumType("com.vmware.vcenter.hvc.links.sync.providers.session_info.stage", reflect.TypeOf(ProvidersSessionInfoStage(ProvidersSessionInfoStage_CHANGE_DETECTION)))
	fieldNameMap["stage"] = "Stage"
	fields["completed_work"] = bindings.NewIntegerType()
	fieldNameMap["completed_work"] = "CompletedWork"
	fields["total_work"] = bindings.NewIntegerType()
	fieldNameMap["total_work"] = "TotalWork"
	fields["completion_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["completion_time"] = "CompletionTime"
	fields["start_time"] = bindings.NewDateTimeType()
	fieldNameMap["start_time"] = "StartTime"
	fields["exception"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["exception"] = "Exception"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("stage",
		map[string][]bindings.FieldData{
			"FAILED": []bindings.FieldData{
				bindings.NewFieldData("completion_time", true),
				bindings.NewFieldData("exception", true),
			},
			"COMPLETED": []bindings.FieldData{
				bindings.NewFieldData("completion_time", true),
			},
			"CHANGE_DETECTION": []bindings.FieldData{},
			"CHANGE_ENUMERATION": []bindings.FieldData{},
			"CHANGE_APPLICATION": []bindings.FieldData{},
			"WAITING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.session_info", fields, reflect.TypeOf(ProvidersSessionInfo{}), fieldNameMap, validators)
}

func ProvidersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.links.sync.Providers"}, "")
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.summary", fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

func ProvidersCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.credentials", fields, reflect.TypeOf(ProvidersCredentials{}), fieldNameMap, validators)
}


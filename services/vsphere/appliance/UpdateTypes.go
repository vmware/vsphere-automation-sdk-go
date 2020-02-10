/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Update.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``State`` enumeration class defines the various states the appliance update can be in. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type UpdateState string

const (
    // The appliance is up to date. This constant field was added in vSphere API 6.7.
	UpdateState_UP_TO_DATE UpdateState = "UP_TO_DATE"
    // A new update is available. This constant field was added in vSphere API 6.7.
	UpdateState_UPDATES_PENDING UpdateState = "UPDATES_PENDING"
    // The appliance update is in progress of downloading an update. This constant field was added in vSphere API 6.7.
	UpdateState_STAGE_IN_PROGRESS UpdateState = "STAGE_IN_PROGRESS"
    // The appliance update is in progress of installing an update. This constant field was added in vSphere API 6.7.
	UpdateState_INSTALL_IN_PROGRESS UpdateState = "INSTALL_IN_PROGRESS"
    // The appliance update failed and cannot recover. This constant field was added in vSphere API 6.7.
	UpdateState_INSTALL_FAILED UpdateState = "INSTALL_FAILED"
    // The appliance update failed and recovery is in progress. This constant field was added in vSphere API 6.7.
	UpdateState_ROLLBACK_IN_PROGRESS UpdateState = "ROLLBACK_IN_PROGRESS"
)

func (s UpdateState) UpdateState() bool {
	switch s {
	case UpdateState_UP_TO_DATE:
		return true
	case UpdateState_UPDATES_PENDING:
		return true
	case UpdateState_STAGE_IN_PROGRESS:
		return true
	case UpdateState_INSTALL_IN_PROGRESS:
		return true
	case UpdateState_INSTALL_FAILED:
		return true
	case UpdateState_ROLLBACK_IN_PROGRESS:
		return true
	default:
		return false
	}
}


// The ``Info`` class describes the state of the appliance update. This class was added in vSphere API 6.7.
type UpdateInfo struct {
    // State of the appliance update. This property was added in vSphere API 6.7.
	State UpdateState
    // The running or completed update task. This property was added in vSphere API 6.7.
	Task *TaskInfo
    // Version of base appliance if state is UP_TO_DATE Version of update being staged or installed if state is INSTALL_IN_PROGRESS or STAGE_IN_PROGRESS Version of update staged if state is UPDATES_PENDING Version of update failed if state is INSTALL_FAILED or ROLLBACK_IN_PROGRESS. This property was added in vSphere API 6.7.
	Version string
    // Timestamp of latest query to update repository. This property was added in vSphere API 6.7.
	LatestQueryTime *time.Time
}



func updateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(UpdateInfoBindingType)
}

func updateGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func updateCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateCancelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func UpdateInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["state"] = bindings.NewEnumType("com.vmware.appliance.update.state", reflect.TypeOf(UpdateState(UpdateState_UP_TO_DATE)))
	fieldNameMap["state"] = "State"
	fields["task"] = bindings.NewOptionalType(bindings.NewReferenceType(TaskInfoBindingType))
	fieldNameMap["task"] = "Task"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["latest_query_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["latest_query_time"] = "LatestQueryTime"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("state",
		map[string][]bindings.FieldData{
			"UP_TO_DATE": []bindings.FieldData{},
			"UPDATES_PENDING": []bindings.FieldData{},
			"STAGE_IN_PROGRESS": []bindings.FieldData{},
			"INSTALL_IN_PROGRESS": []bindings.FieldData{},
			"INSTALL_FAILED": []bindings.FieldData{},
			"ROLLBACK_IN_PROGRESS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.update.info", fields, reflect.TypeOf(UpdateInfo{}), fieldNameMap, validators)
}


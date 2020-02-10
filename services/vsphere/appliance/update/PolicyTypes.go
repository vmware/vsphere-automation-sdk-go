/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``DayOfWeek`` enumeration class defines the set of days. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PolicyDayOfWeek string

const (
    // Monday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_MONDAY PolicyDayOfWeek = "MONDAY"
    // Tuesday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_TUESDAY PolicyDayOfWeek = "TUESDAY"
    // Wednesday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_WEDNESDAY PolicyDayOfWeek = "WEDNESDAY"
    // Thursday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_THURSDAY PolicyDayOfWeek = "THURSDAY"
    // Friday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_FRIDAY PolicyDayOfWeek = "FRIDAY"
    // Saturday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_SATURDAY PolicyDayOfWeek = "SATURDAY"
    // Sunday. This constant field was added in vSphere API 6.7.
	PolicyDayOfWeek_SUNDAY PolicyDayOfWeek = "SUNDAY"
)

func (d PolicyDayOfWeek) PolicyDayOfWeek() bool {
	switch d {
	case PolicyDayOfWeek_MONDAY:
		return true
	case PolicyDayOfWeek_TUESDAY:
		return true
	case PolicyDayOfWeek_WEDNESDAY:
		return true
	case PolicyDayOfWeek_THURSDAY:
		return true
	case PolicyDayOfWeek_FRIDAY:
		return true
	case PolicyDayOfWeek_SATURDAY:
		return true
	case PolicyDayOfWeek_SUNDAY:
		return true
	default:
		return false
	}
}


// The ``Time`` class defines weekday and time the automatic check for new updates will be run. This class was added in vSphere API 6.7.
type PolicyTime struct {
    // weekday to check for updates. This property was added in vSphere API 6.7.
	Day PolicyDayOfWeek
    // Hour: 0-24. This property was added in vSphere API 6.7.
	Hour int64
    // Minute: 0-59. This property was added in vSphere API 6.7.
	Minute int64
}

// The ``Info`` class defines automatic update checking and staging policy. This class was added in vSphere API 6.7.
type PolicyInfo struct {
    // Current appliance update custom repository URL. This property was added in vSphere API 6.7.
	CustomURL *string
    // Current appliance update default repository URL. This property was added in vSphere API 6.7.
	DefaultURL string
    // Username for the update repository. This property was added in vSphere API 6.7.
	Username *string
    // Schedule when the automatic check will be run. This property was added in vSphere API 6.7.
	CheckSchedule []PolicyTime
    // Automatically stage the latest update if available. This property was added in vSphere API 6.7.
	AutoStage bool
    // Is the appliance updated automatically. If map with bool value the appliance may ignore the check schedule or auto-stage settings. This property was added in vSphere API 6.7.
	AutoUpdate bool
    // Whether API client should allow the user to start update manually. This property was added in vSphere API 6.7.
	ManualControl bool
}

// The ``Config`` class defines automatic update checking and staging policy. This class was added in vSphere API 6.7.
type PolicyConfig struct {
    // Current appliance update repository URL. This property was added in vSphere API 6.7.
	CustomURL *string
    // Username for the update repository. This property was added in vSphere API 6.7.
	Username *string
    // Password for the update repository. This property was added in vSphere API 6.7.
	Password *string
    // Schedule when the automatic check will be run. This property was added in vSphere API 6.7.
	CheckSchedule []PolicyTime
    // Automatically stage the latest update if available. This property was added in vSphere API 6.7.
	AutoStage bool
}



func policyGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PolicyInfoBindingType)
}

func policyGetRestMetadata() protocol.OperationRestMetadata {
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

func policySetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewReferenceType(PolicyConfigBindingType)
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policySetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func policySetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["policy"] = bindings.NewReferenceType(PolicyConfigBindingType)
	fieldNameMap["policy"] = "Policy"
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


func PolicyTimeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["day"] = bindings.NewEnumType("com.vmware.appliance.update.policy.day_of_week", reflect.TypeOf(PolicyDayOfWeek(PolicyDayOfWeek_MONDAY)))
	fieldNameMap["day"] = "Day"
	fields["hour"] = bindings.NewIntegerType()
	fieldNameMap["hour"] = "Hour"
	fields["minute"] = bindings.NewIntegerType()
	fieldNameMap["minute"] = "Minute"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.policy.time", fields, reflect.TypeOf(PolicyTime{}), fieldNameMap, validators)
}

func PolicyInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["custom_URL"] = "CustomURL"
	fields["default_URL"] = bindings.NewStringType()
	fieldNameMap["default_URL"] = "DefaultURL"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(PolicyTimeBindingType), reflect.TypeOf([]PolicyTime{}))
	fieldNameMap["check_schedule"] = "CheckSchedule"
	fields["auto_stage"] = bindings.NewBooleanType()
	fieldNameMap["auto_stage"] = "AutoStage"
	fields["auto_update"] = bindings.NewBooleanType()
	fieldNameMap["auto_update"] = "AutoUpdate"
	fields["manual_control"] = bindings.NewBooleanType()
	fieldNameMap["manual_control"] = "ManualControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.policy.info", fields, reflect.TypeOf(PolicyInfo{}), fieldNameMap, validators)
}

func PolicyConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["custom_URL"] = "CustomURL"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["password"] = "Password"
	fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(PolicyTimeBindingType), reflect.TypeOf([]PolicyTime{}))
	fieldNameMap["check_schedule"] = "CheckSchedule"
	fields["auto_stage"] = bindings.NewBooleanType()
	fieldNameMap["auto_stage"] = "AutoStage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.policy.config", fields, reflect.TypeOf(PolicyConfig{}), fieldNameMap, validators)
}


/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: LocalAccounts.
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



// The ``Info`` class defines the local account properties. This class was added in vSphere API 6.7.
type LocalAccountsInfo struct {
    // Full name of the user. This property was added in vSphere API 6.7.
	Fullname *string
    // Email address of the local account. This property was added in vSphere API 6.7.
	Email *string
    // User roles. This property was added in vSphere API 6.7.
	Roles []string
    // Flag indicating if the account is enabled. This property was added in vSphere API 6.7.
	Enabled bool
    // Is the user password set. This property was added in vSphere API 6.7.
	HasPassword bool
    // Date and time password was changed. This property was added in vSphere API 6.7.
	LastPasswordChange *time.Time
    // Date when the account's password will expire. This property was added in vSphere API 6.7.
	PasswordExpiresAt *time.Time
    // Date and time account will be locked after password expiration. This property was added in vSphere API 6.7.
	InactiveAt *time.Time
    // Minimum number of days between password change. This property was added in vSphere API 6.7.
	MinDaysBetweenPasswordChange *int64
    // Maximum number of days between password change. This property was added in vSphere API 6.7.
	MaxDaysBetweenPasswordChange *int64
    // Number of days of warning before password expires. This property was added in vSphere API 6.7.
	WarnDaysBeforePasswordExpiration *int64
}

// The ``Config`` class defines the information required for the account. This class was added in vSphere API 6.7.
type LocalAccountsConfig struct {
    // Password. This property was added in vSphere API 6.7.
	Password string
    // Old password of the user (required in case of the password change, not required if superAdmin user changes the password of the other user). This property was added in vSphere API 6.7.
	OldPassword *string
    // Full name of the user. This property was added in vSphere API 6.7.
	FullName *string
    // Email address of the local account. This property was added in vSphere API 6.7.
	Email *string
    // User roles. This property was added in vSphere API 6.7.
	Roles []string
    // Flag indicating if the account is enabled. This property was added in vSphere API 6.7.
	Enabled *bool
    // Flag indicating if the account password expires. This property was added in vSphere API 6.7.
	PasswordExpires *bool
    // Date when the account's password will expire. This property was added in vSphere API 6.7.
	PasswordExpiresAt *time.Time
    // Flag indicating if the account will be locked after password expiration. This property was added in vSphere API 6.7.
	InactiveAfterPasswordExpiration *bool
    // Number of days after password expiration before the account will be locked. This property was added in vSphere API 6.7.
	DaysAfterPasswordExpiration *int64
    // Minimum number of days between password change. This property was added in vSphere API 6.7.
	MinDaysBetweenPasswordChange *int64
    // Maximum number of days between password change. This property was added in vSphere API 6.7.
	MaxDaysBetweenPasswordChange *int64
    // Number of days of warning before password expires. This property was added in vSphere API 6.7.
	WarnDaysBeforePasswordExpiration *int64
}

// The ``UpdateConfig`` class defines the fields that might be updated. This class was added in vSphere API 6.7.
type LocalAccountsUpdateConfig struct {
    // Password. This property was added in vSphere API 6.7.
	Password *string
    // Old password of the user (required in case of the password change, not required if superAdmin user changes the password of the other user). This property was added in vSphere API 6.7.
	OldPassword *string
    // Full name of the user. This property was added in vSphere API 6.7.
	FullName *string
    // Email address of the local account. This property was added in vSphere API 6.7.
	Email *string
    // User roles. This property was added in vSphere API 6.7.
	Roles []string
    // Flag indicating if the account is enabled. This property was added in vSphere API 6.7.
	Enabled *bool
    // Flag indicating if the account password expires. This property was added in vSphere API 6.7.
	PasswordExpires *bool
    // Date when the account's password will expire. This property was added in vSphere API 6.7.
	PasswordExpiresAt *time.Time
    // Flag indicating if the account will be locked after password expiration. This property was added in vSphere API 6.7.
	InactiveAfterPasswordExpiration *bool
    // Number of days after password expiration before the account will be locked. This property was added in vSphere API 6.7.
	DaysAfterPasswordExpiration *int64
    // Minimum number of days between password change. This property was added in vSphere API 6.7.
	MinDaysBetweenPasswordChange *int64
    // Maximum number of days between password change. This property was added in vSphere API 6.7.
	MaxDaysBetweenPasswordChange *int64
    // Number of days of warning before password expires. This property was added in vSphere API 6.7.
	WarnDaysBeforePasswordExpiration *int64
}



func localAccountsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LocalAccountsInfoBindingType)
}

func localAccountsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
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
		map[string]int{"NotFound": 404,"Error": 500})
}

func localAccountsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, ""), reflect.TypeOf([]string{}))
}

func localAccountsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500})
}

func localAccountsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localAccountsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500})
}

func localAccountsSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localAccountsSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
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
		map[string]int{"NotFound": 404,"Error": 500})
}

func localAccountsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsUpdateConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localAccountsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fields["config"] = bindings.NewReferenceType(LocalAccountsUpdateConfigBindingType)
	fieldNameMap["username"] = "Username"
	fieldNameMap["config"] = "Config"
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
		map[string]int{"NotFound": 404,"Error": 500})
}

func localAccountsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fieldNameMap["username"] = "Username"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localAccountsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func localAccountsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["username"] = bindings.NewIdType([]string{"com.vmware.appliance.local_accounts"}, "")
	fieldNameMap["username"] = "Username"
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
		map[string]int{"NotFound": 404,"Error": 500})
}


func LocalAccountsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fullname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["fullname"] = "Fullname"
	fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["email"] = "Email"
	fields["roles"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["roles"] = "Roles"
	fields["enabled"] = bindings.NewBooleanType()
	fieldNameMap["enabled"] = "Enabled"
	fields["has_password"] = bindings.NewBooleanType()
	fieldNameMap["has_password"] = "HasPassword"
	fields["last_password_change"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["last_password_change"] = "LastPasswordChange"
	fields["password_expires_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["password_expires_at"] = "PasswordExpiresAt"
	fields["inactive_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["inactive_at"] = "InactiveAt"
	fields["min_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_days_between_password_change"] = "MinDaysBetweenPasswordChange"
	fields["max_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_days_between_password_change"] = "MaxDaysBetweenPasswordChange"
	fields["warn_days_before_password_expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["warn_days_before_password_expiration"] = "WarnDaysBeforePasswordExpiration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.local_accounts.info", fields, reflect.TypeOf(LocalAccountsInfo{}), fieldNameMap, validators)
}

func LocalAccountsConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	fields["old_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["old_password"] = "OldPassword"
	fields["full_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["full_name"] = "FullName"
	fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["email"] = "Email"
	fields["roles"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["roles"] = "Roles"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["password_expires"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["password_expires"] = "PasswordExpires"
	fields["password_expires_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["password_expires_at"] = "PasswordExpiresAt"
	fields["inactive_after_password_expiration"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["inactive_after_password_expiration"] = "InactiveAfterPasswordExpiration"
	fields["days_after_password_expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["days_after_password_expiration"] = "DaysAfterPasswordExpiration"
	fields["min_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_days_between_password_change"] = "MinDaysBetweenPasswordChange"
	fields["max_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_days_between_password_change"] = "MaxDaysBetweenPasswordChange"
	fields["warn_days_before_password_expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["warn_days_before_password_expiration"] = "WarnDaysBeforePasswordExpiration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.local_accounts.config", fields, reflect.TypeOf(LocalAccountsConfig{}), fieldNameMap, validators)
}

func LocalAccountsUpdateConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["password"] = "Password"
	fields["old_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["old_password"] = "OldPassword"
	fields["full_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["full_name"] = "FullName"
	fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["email"] = "Email"
	fields["roles"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["roles"] = "Roles"
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["enabled"] = "Enabled"
	fields["password_expires"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["password_expires"] = "PasswordExpires"
	fields["password_expires_at"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["password_expires_at"] = "PasswordExpiresAt"
	fields["inactive_after_password_expiration"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["inactive_after_password_expiration"] = "InactiveAfterPasswordExpiration"
	fields["days_after_password_expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["days_after_password_expiration"] = "DaysAfterPasswordExpiration"
	fields["min_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_days_between_password_change"] = "MinDaysBetweenPasswordChange"
	fields["max_days_between_password_change"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_days_between_password_change"] = "MaxDaysBetweenPasswordChange"
	fields["warn_days_before_password_expiration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["warn_days_before_password_expiration"] = "WarnDaysBeforePasswordExpiration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.local_accounts.update_config", fields, reflect.TypeOf(LocalAccountsUpdateConfig{}), fieldNameMap, validators)
}


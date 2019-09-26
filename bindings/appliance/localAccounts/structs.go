/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LocalAccounts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package localAccounts

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)




// The ``Info`` class defines the local account properties.
 type Info struct {
    Fullname *string
    Email *string
    Roles []string
    Enabled bool
    HasPassword bool
    LastPasswordChange *time.Time
    PasswordExpiresAt *time.Time
    InactiveAt *time.Time
    MinDaysBetweenPasswordChange *int64
    MaxDaysBetweenPasswordChange *int64
    WarnDaysBeforePasswordExpiration *int64
}






// The ``Config`` class defines the information required for the account.
 type Config struct {
    Password string
    OldPassword *string
    FullName *string
    Email *string
    Roles []string
    Enabled *bool
    PasswordExpires *bool
    PasswordExpiresAt *time.Time
    InactiveAfterPasswordExpiration *bool
    DaysAfterPasswordExpiration *int64
    MinDaysBetweenPasswordChange *int64
    MaxDaysBetweenPasswordChange *int64
    WarnDaysBeforePasswordExpiration *int64
}






// The ``UpdateConfig`` class defines the fields that might be updated.
 type UpdateConfig struct {
    Password *string
    OldPassword *string
    FullName *string
    Email *string
    Roles []string
    Enabled *bool
    PasswordExpires *bool
    PasswordExpiresAt *time.Time
    InactiveAfterPasswordExpiration *bool
    DaysAfterPasswordExpiration *int64
    MinDaysBetweenPasswordChange *int64
    MaxDaysBetweenPasswordChange *int64
    WarnDaysBeforePasswordExpiration *int64
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.local_accounts"}, ""), reflect.TypeOf([]string{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewIdType([]string {"com.vmware.appliance.local_accounts"}, "")
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["username"] = "Username"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewIdType([]string {"com.vmware.appliance.local_accounts"}, "")
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["username"] = "Username"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewIdType([]string {"com.vmware.appliance.local_accounts"}, "")
    fields["config"] = bindings.NewReferenceType(UpdateConfigBindingType)
    fieldNameMap["username"] = "Username"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewIdType([]string {"com.vmware.appliance.local_accounts"}, "")
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["fullname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["fullname"] = "Fullname"
    fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["email"] = "Email"
    fields["roles"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{}))
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
    return bindings.NewStructType("com.vmware.appliance.local_accounts.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ConfigBindingType() bindings.BindingType {
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
    fields["roles"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{}))
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
    return bindings.NewStructType("com.vmware.appliance.local_accounts.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func UpdateConfigBindingType() bindings.BindingType {
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
    fields["roles"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.roles"}, ""), reflect.TypeOf([]string{})))
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
    return bindings.NewStructType("com.vmware.appliance.local_accounts.update_config",fields, reflect.TypeOf(UpdateConfig{}), fieldNameMap, validators)
}



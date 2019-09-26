/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: User.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package user

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``UserAccountStatus`` enumeration class Defines status of user accounts
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UserAccountStatus string

const (
    // The user account is disabled.
     UserAccountStatus_disabled UserAccountStatus = "disabled"
    // The user account is enabled.
     UserAccountStatus_enabled UserAccountStatus = "enabled"
)

func (u UserAccountStatus) UserAccountStatus() bool {
    switch u {
        case UserAccountStatus_disabled:
            return true
        case UserAccountStatus_enabled:
            return true
        default:
            return false
    }
}




// ``UserPasswordStatus`` enumeration class Defines state of user password
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UserPasswordStatus string

const (
    // No password has been set
     UserPasswordStatus_notset UserPasswordStatus = "notset"
    // The password has expired.
     UserPasswordStatus_expired UserPasswordStatus = "expired"
    // The password is still valid.
     UserPasswordStatus_valid UserPasswordStatus = "valid"
)

func (u UserPasswordStatus) UserPasswordStatus() bool {
    switch u {
        case UserPasswordStatus_notset:
            return true
        case UserPasswordStatus_expired:
            return true
        case UserPasswordStatus_valid:
            return true
        default:
            return false
    }
}




// ``UserRole`` enumeration class Defines user roles for appliance
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UserRole string

const (
    // Able to configure the appliance.
     UserRole_admin UserRole = "admin"
    // Able to read the appliance configuration.
     UserRole_operator UserRole = "operator"
    // Able to configure the appliance, manage local accounts and use the BASH shell
     UserRole_superAdmin UserRole = "superAdmin"
)

func (u UserRole) UserRole() bool {
    switch u {
        case UserRole_admin:
            return true
        case UserRole_operator:
            return true
        case UserRole_superAdmin:
            return true
        default:
            return false
    }
}





// ``UserConfigGet`` class Structure defines a user configuration for user.get API.
 type UserConfigGet struct {
    Username string
    Role UserRole
    Fullname string
    Status UserAccountStatus
    Passwordstatus UserPasswordStatus
    Email string
}






// ``UserConfig`` class Structure that defines a new user configuration for CLI.
 type UserConfig struct {
    Username string
    Role UserRole
    Fullname string
    Status UserAccountStatus
    Email string
}






// ``NewUserConfig`` class Structure that defines a new user configuration.
 type NewUserConfig struct {
    Username string
    Role *UserRole
    Password string
    Fullname *string
    Email *string
}









func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
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
       map[string]int{"Error": 500})
}


func addInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(NewUserConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func addOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func addRestMetadata() protocol.OperationRestMetadata {
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


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(UserConfigBindingType)
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
       map[string]int{"Error": 500})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(UserConfigGetBindingType), reflect.TypeOf([]UserConfigGet{}))
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UserConfigGetBindingType)
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
       map[string]int{"Error": 500})
}



func UserConfigGetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(UserRole(UserRole_admin)))
    fieldNameMap["role"] = "Role"
    fields["fullname"] = bindings.NewStringType()
    fieldNameMap["fullname"] = "Fullname"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_account_status", reflect.TypeOf(UserAccountStatus(UserAccountStatus_disabled)))
    fieldNameMap["status"] = "Status"
    fields["passwordstatus"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_password_status", reflect.TypeOf(UserPasswordStatus(UserPasswordStatus_notset)))
    fieldNameMap["passwordstatus"] = "Passwordstatus"
    fields["email"] = bindings.NewStringType()
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.user_config_get",fields, reflect.TypeOf(UserConfigGet{}), fieldNameMap, validators)
}

func UserConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(UserRole(UserRole_admin)))
    fieldNameMap["role"] = "Role"
    fields["fullname"] = bindings.NewStringType()
    fieldNameMap["fullname"] = "Fullname"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_account_status", reflect.TypeOf(UserAccountStatus(UserAccountStatus_disabled)))
    fieldNameMap["status"] = "Status"
    fields["email"] = bindings.NewStringType()
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.user_config",fields, reflect.TypeOf(UserConfig{}), fieldNameMap, validators)
}

func NewUserConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.appliance.techpreview.localaccounts.user.user_role", reflect.TypeOf(UserRole(UserRole_admin))))
    fieldNameMap["role"] = "Role"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["fullname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["fullname"] = "Fullname"
    fields["email"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["email"] = "Email"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.localaccounts.user.new_user_config",fields, reflect.TypeOf(NewUserConfig{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.std.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package std

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "time"
)



// The AuthenticationScheme class defines constants for authentication scheme identifiers for authentication mechanisms present in the vAPI infrastructure shipped by VMware. 
//
//  A third party extension can define and implements it's own authentication mechanism and define a constant in a different IDL file.
type AuthenticationScheme struct {
}

// Indicates that the request doesn't need any authentication.
const AuthenticationScheme_NO_AUTHENTICATION = "com.vmware.vapi.std.security.no_authentication"
// Indicates that the security context in a request is using a SAML bearer token based authentication scheme. 
//
//  In this scheme, the following pieces of information has to be passed in the SecurityContext structure in the execution context of the request: 
//
// * The scheme identifier: com.vmware.vapi.std.security.saml_bearer_token
// * The token itself
//
//  
//
//  Sample security context in JSON format that matches the specification: ``{
// 'schemeId': 'com.vmware.vapi.std.security.saml_bearer_token',
// 'token': 'the token itself'
// }`` vAPI runtime provide convenient factory methods that take SAML bearer token and to create the security context that conforms to the above mentioned format.
const AuthenticationScheme_SAML_BEARER_TOKEN = "com.vmware.vapi.std.security.saml_bearer_token"
// Indicates that the security context in a request is using a SAML holder-of-key token based authentication scheme. 
//
//  In this scheme, the following pieces of information has to be passed in the SecurityContext structure in the execution context of the request: 
//
// * The scheme identifier: com.vmware.vapi.std.security.saml_hok_token
// * Signature of the request: This includes - algorithm used for signing the request, SAML holder of key token and signature digest
// * Request timestamp: This includes the ``created`` and ``expires`` timestamp of the request. The timestamp should match the following format - YYYY-MM-DDThh:mm:ss.sssZ (e.g. 1878-03-03T19:20:30.451Z).
//
//  
//
//  Sample security context in JSON format that matches the specification: ``{
// 'schemeId': 'com.vmware.vapi.std.security.saml_hok_token',
// 'signature': {
// 'alg': 'RS256',
// 'samlToken': ...,
// 'value': ...,``, 'timestamp': { 'created': '2012-10-26T12:24:18.941Z', 'expires': '2012-10-26T12:44:18.941Z', } } } vAPI runtime provide convenient factory methods that take SAML holder of key token and private key to create the security context that conforms to the above mentioned format.
const AuthenticationScheme_SAML_HOK_TOKEN = "com.vmware.vapi.std.security.saml_hok_token"
// Indicates that the security context in a request is using a session identifier based authentication scheme. 
//
//  In this scheme, the following pieces of information has to be passed in the SecurityContext structure in the execution context of the request: 
//
// * The scheme identifier - com.vmware.vapi.std.security.session_id
// * Valid session identifier - This is usually returned by a login method of a session manager interface for a particular vAPI service of this authentication scheme
//
//  Sample security context in JSON format that matches the specification: ``{
// 'schemeId': 'com.vmware.vapi.std.security.session_id',
// 'sessionId': ....,
// }`` vAPI runtime provides convenient factory methods that take session identifier as input parameter and create a security context that conforms to the above format.
const AuthenticationScheme_SESSION_ID = "com.vmware.vapi.std.security.session_id"
// Indicates that the security context in a request is using username/password based authentication scheme. 
//
//  In this scheme, the following pieces of information has to be passed in the SecurityContext structure in the execution context of the request: 
//
// * The scheme identifier - com.vmware.vapi.std.security.user_pass
// * Username
// * Password
//
//  
//
//  Sample security context in JSON format that matches the specification: ``{
// 'schemeId': 'com.vmware.vapi.std.security.user_pass',
// 'userName': ....,
// 'password': ...
// }`` 
//  vAPI runtime provides convenient factory methods that take username and password as input parameters and create a security context that conforms to the above format.
const AuthenticationScheme_USER_PASSWORD = "com.vmware.vapi.std.security.user_pass"
// Indicates that the security context in a request is using OAuth2 based authentication scheme. 
//
//  In this scheme, the following pieces of information has to be passed in the SecurityContext structure in the execution context of the request: 
//
// * The scheme identifier - com.vmware.vapi.std.security.oauth
// * Valid OAuth2 access token - This is usually acquired by OAuth2 Authorization Server after successful authentication of the end user.
//
//  
//
//  Sample security context in JSON format that matches the specification: ``{
// 'schemeId': 'com.vmware.vapi.std.security.oauth',
// 'accesstoken': ....
// }`` 
//  vAPI runtime provides convenient factory methods that takes OAuth2 access token as input parameter and creates a security context that conforms to the above format.
const AuthenticationScheme_OAUTH_ACCESS_TOKEN = "com.vmware.vapi.std.security.oauth"





// The ``DynamicID`` class represents an identifier for a resource of an arbitrary type.
type DynamicID struct {
    Type_ string
    Id string
}






// The ``LocalizableMessage`` class represents localizable string and message template. Interfaces include one or more localizable message templates in the exceptions they report so that clients can display diagnostic messages in the native language of the user. Interfaces can include localizable strings in the data returned from methods to allow clients to display localized status information in the native language of the user.
type LocalizableMessage struct {
    Id string
    DefaultMessage string
    Args []string
    Params map[string]LocalizationParam
    Localized *string
}






// This class holds a single message parameter and formatting settings for it. The class has fields for string, int64, float64, date time and nested messages. Only one will be used depending on the type of data sent. For date, float64 and int64 it is possible to set additional formatting details. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type LocalizationParam struct {
    S *string
    Dt *time.Time
    I *int64
    D *float64
    L *NestedLocalizableMessage
    Format *LocalizationParam_DateTimeFormat
    Precision *int64
}




    
    // The ``DateTimeFormat`` enumeration class lists possible date and time formatting options. It combines the Unicode CLDR format types - full, long, medium and short with 3 different presentations - date only, time only and combined date and time presentation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type LocalizationParam_DateTimeFormat string

    const (
        // The date and time value will be formatted as short date, for example *2019-01-28*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_SHORT_DATE LocalizationParam_DateTimeFormat = "SHORT_DATE"
        // The date and time value will be formatted as medium date, for example *2019 Jan 28*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_MED_DATE LocalizationParam_DateTimeFormat = "MED_DATE"
        // The date and time value will be formatted as long date, for example *2019 Jan 28*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_LONG_DATE LocalizationParam_DateTimeFormat = "LONG_DATE"
        // The date and time value will be formatted as full date, for example *2019 Jan 28, Mon*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_FULL_DATE LocalizationParam_DateTimeFormat = "FULL_DATE"
        // The date and time value will be formatted as short time, for example *12:59*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_SHORT_TIME LocalizationParam_DateTimeFormat = "SHORT_TIME"
        // The date and time value will be formatted as medium time, for example *12:59:01*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_MED_TIME LocalizationParam_DateTimeFormat = "MED_TIME"
        // The date and time value will be formatted as long time, for example *12:59:01 Z*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_LONG_TIME LocalizationParam_DateTimeFormat = "LONG_TIME"
        // The date and time value will be formatted as full time, for example *12:59:01 Z*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_FULL_TIME LocalizationParam_DateTimeFormat = "FULL_TIME"
        // The date and time value will be formatted as short date and time, for example *2019-01-28 12:59*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_SHORT_DATE_TIME LocalizationParam_DateTimeFormat = "SHORT_DATE_TIME"
        // The date and time value will be formatted as medium date and time, for example *2019 Jan 28 12:59:01*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_MED_DATE_TIME LocalizationParam_DateTimeFormat = "MED_DATE_TIME"
        // The date and time value will be formatted as long date and time, for example *2019 Jan 28 12:59:01 Z*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_LONG_DATE_TIME LocalizationParam_DateTimeFormat = "LONG_DATE_TIME"
        // The date and time value will be formatted as full date and time, for example *2019 Jan 28, Mon 12:59:01 Z*. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         LocalizationParam_DateTimeFormat_FULL_DATE_TIME LocalizationParam_DateTimeFormat = "FULL_DATE_TIME"
    )

    func (d LocalizationParam_DateTimeFormat) LocalizationParam_DateTimeFormat() bool {
        switch d {
            case LocalizationParam_DateTimeFormat_SHORT_DATE:
                return true
            case LocalizationParam_DateTimeFormat_MED_DATE:
                return true
            case LocalizationParam_DateTimeFormat_LONG_DATE:
                return true
            case LocalizationParam_DateTimeFormat_FULL_DATE:
                return true
            case LocalizationParam_DateTimeFormat_SHORT_TIME:
                return true
            case LocalizationParam_DateTimeFormat_MED_TIME:
                return true
            case LocalizationParam_DateTimeFormat_LONG_TIME:
                return true
            case LocalizationParam_DateTimeFormat_FULL_TIME:
                return true
            case LocalizationParam_DateTimeFormat_SHORT_DATE_TIME:
                return true
            case LocalizationParam_DateTimeFormat_MED_DATE_TIME:
                return true
            case LocalizationParam_DateTimeFormat_LONG_DATE_TIME:
                return true
            case LocalizationParam_DateTimeFormat_FULL_DATE_TIME:
                return true
            default:
                return false
        }
    }



// The ``NestedLocalizableMessage`` class represents a nested within a parameter localizable string or message template. This class is useful for modeling composite messages. Such messages are necessary to do correct pluralization of phrases, represent lists of several items etc. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NestedLocalizableMessage struct {
    Id string
    Params map[string]LocalizationParam
}










func AuthenticationSchemeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.authentication_scheme",fields, reflect.TypeOf(AuthenticationScheme{}), fieldNameMap, validators)
}

func DynamicIDBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["id"] = bindings.NewIdType(nil, "type")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.dynamic_ID",fields, reflect.TypeOf(DynamicID{}), fieldNameMap, validators)
}

func LocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["default_message"] = bindings.NewStringType()
    fieldNameMap["default_message"] = "DefaultMessage"
    fields["args"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["args"] = "Args"
    fields["params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(LocalizationParamBindingType),reflect.TypeOf(map[string]LocalizationParam{})))
    fieldNameMap["params"] = "Params"
    fields["localized"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["localized"] = "Localized"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.localizable_message",fields, reflect.TypeOf(LocalizableMessage{}), fieldNameMap, validators)
}

func LocalizationParamBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["s"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["s"] = "S"
    fields["dt"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["dt"] = "Dt"
    fields["i"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["i"] = "I"
    fields["d"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["d"] = "D"
    fields["l"] = bindings.NewOptionalType(bindings.NewReferenceType(NestedLocalizableMessageBindingType))
    fieldNameMap["l"] = "L"
    fields["format"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.localization_param.date_time_format", reflect.TypeOf(LocalizationParam_DateTimeFormat(LocalizationParam_DateTimeFormat_SHORT_DATE))))
    fieldNameMap["format"] = "Format"
    fields["precision"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["precision"] = "Precision"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.localization_param",fields, reflect.TypeOf(LocalizationParam{}), fieldNameMap, validators)
}

func NestedLocalizableMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(LocalizationParamBindingType),reflect.TypeOf(map[string]LocalizationParam{})))
    fieldNameMap["params"] = "Params"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.nested_localizable_message",fields, reflect.TypeOf(NestedLocalizableMessage{}), fieldNameMap, validators)
}



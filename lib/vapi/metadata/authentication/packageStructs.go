/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.metadata.authentication.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package authentication

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``AuthenticationInfo`` class describes the authentication information. Authentication information could be specified for a package element, service elenent or an operation element. 
//
//  Using the authentication scheme information, a client invoking an API call from any interface can figure out what kind of credentials are needed for that API call.
type AuthenticationInfo struct {
    SchemeType AuthenticationInfo_SchemeType
    SessionManager *string
    Scheme string
}




    
    // The ``SchemeType`` enumeration class provides enumeration constants for the set of valid authentication scheme types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type AuthenticationInfo_SchemeType string

    const (
        // Indicates that the scheme is a session less authentication scheme, the user is authenticated on every method. There is no explicit session establishment.
         AuthenticationInfo_SchemeType_SESSIONLESS AuthenticationInfo_SchemeType = "SESSIONLESS"
        // Indicates that the scheme is a session aware authentication scheme. It requires an explicit login before executing a method and logout when a session terminates. A interface might choose to have a session aware scheme if it wants to associate some state corresponding to the user until the user logs out or if it wants to mitigate the cost of authenticating the user on every method.
         AuthenticationInfo_SchemeType_SESSION_AWARE AuthenticationInfo_SchemeType = "SESSION_AWARE"
    )

    func (s AuthenticationInfo_SchemeType) AuthenticationInfo_SchemeType() bool {
        switch s {
            case AuthenticationInfo_SchemeType_SESSIONLESS:
                return true
            case AuthenticationInfo_SchemeType_SESSION_AWARE:
                return true
            default:
                return false
        }
    }



// The ``ComponentData`` class contains the authentication information of the component along with its fingerprint.
type ComponentData struct {
    Info ComponentInfo
    Fingerprint string
}






// The ``ComponentInfo`` class contains authentication information of a component element. 
//
//  For an explanation of authentication information contained within component elements, see Component.
type ComponentInfo struct {
    Packages map[string]PackageInfo
}






// The ``OperationInfo`` class contains authentication information of an operation element.
type OperationInfo struct {
    Schemes []AuthenticationInfo
}






// The ``PackageInfo`` class contains authentication information of a package element. 
//
//  For an explanation of authentication information contained within package elements, see Package.
type PackageInfo struct {
    Schemes []AuthenticationInfo
    Services map[string]ServiceInfo
}






// The ``ServiceInfo`` class contains authentication information of a service element. 
//
//  For an explanation of authentication information contained within service elements, see Service.
type ServiceInfo struct {
    Schemes []AuthenticationInfo
    Operations map[string]OperationInfo
}










func AuthenticationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["scheme_type"] = bindings.NewEnumType("com.vmware.vapi.metadata.authentication.authentication_info.scheme_type", reflect.TypeOf(AuthenticationInfo_SchemeType(AuthenticationInfo_SchemeType_SESSIONLESS)))
    fieldNameMap["scheme_type"] = "SchemeType"
    fields["session_manager"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["session_manager"] = "SessionManager"
    fields["scheme"] = bindings.NewStringType()
    fieldNameMap["scheme"] = "Scheme"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("scheme_type",
        map[string][]bindings.FieldData {
            "SESSION_AWARE": []bindings.FieldData {
                 bindings.NewFieldData("session_manager", true),
            },
            "SESSIONLESS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.authentication_info",fields, reflect.TypeOf(AuthenticationInfo{}), fieldNameMap, validators)
}

func ComponentDataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewReferenceType(ComponentInfoBindingType)
    fieldNameMap["info"] = "Info"
    fields["fingerprint"] = bindings.NewStringType()
    fieldNameMap["fingerprint"] = "Fingerprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.component_data",fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType),reflect.TypeOf(map[string]PackageInfo{}))
    fieldNameMap["packages"] = "Packages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.component_info",fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
    fieldNameMap["schemes"] = "Schemes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.operation_info",fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
    fieldNameMap["schemes"] = "Schemes"
    fields["services"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType),reflect.TypeOf(map[string]ServiceInfo{}))
    fieldNameMap["services"] = "Services"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.package_info",fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
    fieldNameMap["schemes"] = "Schemes"
    fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType),reflect.TypeOf(map[string]OperationInfo{}))
    fieldNameMap["operations"] = "Operations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.authentication.service_info",fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}



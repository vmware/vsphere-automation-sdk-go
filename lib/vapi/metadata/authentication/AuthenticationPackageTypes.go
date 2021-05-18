// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.authentication.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package authentication

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``AuthenticationInfo`` class describes the authentication information. Authentication information could be specified for a package element, service elenent or an operation element.
//
//  Using the authentication scheme information, a client invoking an API call from any interface can figure out what kind of credentials are needed for that API call.
type AuthenticationInfo struct {
	// The type of the authentication scheme.
	SchemeType AuthenticationInfoSchemeTypeEnum
	// In a session aware authentication scheme, a session manager is required that supports ``create``, ``delete`` and ``keepAlive`` methods. The fully qualified interface name of the session manager is provided in AuthenticationInfo#sessionManager property. This interface is responsible for handling sessions.
	SessionManager *string
	// String identifier of the authentication scheme.
	//
	//  Following are the supported authentication schemes by the infrastructure:
	//
	// * The identifier ``com.vmware.vapi.std.security.saml_hok_token`` for SAML holder of key token based authentication mechanism.
	// * The identifier ``com.vmware.vapi.std.security.bearer_token`` for SAML bearer token based authentication mechanism.
	// * The identifier ``com.vmware.vapi.std.security.session_id`` for session based authentication mechanism.
	// * The identifier ``com.vmware.vapi.std.security.user_pass`` for username and password based authentication mechanism.
	Scheme string
}

func (s *AuthenticationInfo) GetType__() bindings.BindingType {
	return AuthenticationInfoBindingType()
}

func (s *AuthenticationInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for AuthenticationInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``SchemeType`` enumeration class provides enumeration constants for the set of valid authentication scheme types.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type AuthenticationInfoSchemeTypeEnum string

const (
	// Indicates that the scheme is a session less authentication scheme, the user is authenticated on every method. There is no explicit session establishment.
	AuthenticationInfoSchemeType_SESSIONLESS AuthenticationInfoSchemeTypeEnum = "SESSIONLESS"
	// Indicates that the scheme is a session aware authentication scheme. It requires an explicit login before executing a method and logout when a session terminates. A interface might choose to have a session aware scheme if it wants to associate some state corresponding to the user until the user logs out or if it wants to mitigate the cost of authenticating the user on every method.
	AuthenticationInfoSchemeType_SESSION_AWARE AuthenticationInfoSchemeTypeEnum = "SESSION_AWARE"
)

func (s AuthenticationInfoSchemeTypeEnum) AuthenticationInfoSchemeTypeEnum() bool {
	switch s {
	case AuthenticationInfoSchemeType_SESSIONLESS:
		return true
	case AuthenticationInfoSchemeType_SESSION_AWARE:
		return true
	default:
		return false
	}
}

// The ``ComponentData`` class contains the authentication information of the component along with its fingerprint.
type ComponentData struct {
	// Authentication information of the component. This includes information about all the packages in the component.
	Info ComponentInfo
	// Fingerprint of the metadata of the component.
	//
	//  Authentication information could change when there is an infrastructure update. Since the data present in ComponentData#info could be quite large, ``fingerprint`` provides a convenient way to check if the data for a particular component is updated.
	//
	//  You should store the fingerprint associated with a component. After an update, by invoking the Component#fingerprint method, you can retrieve the new fingerprint for the component. If the new fingerprint and the previously stored fingerprint do not match, clients can then use the Component#get to retrieve the new authentication information for the component.
	Fingerprint string
}

func (s *ComponentData) GetType__() bindings.BindingType {
	return ComponentDataBindingType()
}

func (s *ComponentData) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentData._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ComponentInfo`` class contains authentication information of a component element.
//
//  For an explanation of authentication information contained within component elements, see Component.
type ComponentInfo struct {
	// Authentication information of all the package elements. The key in the map is the identifier of the package element and the value in the map is the authentication information for the package element.
	//
	//  For an explanation of authentication information containment within package elements, see Package.
	Packages map[string]PackageInfo
}

func (s *ComponentInfo) GetType__() bindings.BindingType {
	return ComponentInfoBindingType()
}

func (s *ComponentInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ComponentInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``OperationInfo`` class contains authentication information of an operation element.
type OperationInfo struct {
	// List of authentication schemes used by an operation element. The authentication scheme specified on the service element corresponding to this operation element is ignored.
	Schemes []AuthenticationInfo
}

func (s *OperationInfo) GetType__() bindings.BindingType {
	return OperationInfoBindingType()
}

func (s *OperationInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for OperationInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``PackageInfo`` class contains authentication information of a package element.
//
//  For an explanation of authentication information contained within package elements, see Package.
type PackageInfo struct {
	// List of authentication schemes to be used for all the operation elements contained in this package element. If a particular service or operation element has no explicit authentications defined in the authentication defintion file, these authentication schemes are used for authenticating the user.
	Schemes []AuthenticationInfo
	// Information about all service elements contained in this package element that contain authentication information. The key in the map is the identifier of the service element and the value in the map is the authentication information for the service element.
	//
	//  For an explanation of authentication information containment within service elements, see Service.
	Services map[string]ServiceInfo
}

func (s *PackageInfo) GetType__() bindings.BindingType {
	return PackageInfoBindingType()
}

func (s *PackageInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for PackageInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ServiceInfo`` class contains authentication information of a service element.
//
//  For an explanation of authentication information contained within service elements, see Service.
type ServiceInfo struct {
	// List of authentication schemes to be used for all the operation elements contained in this service element. The authentication scheme specified on the package element corresponding to this service element is ignored.
	Schemes []AuthenticationInfo
	// Information about all operation elements contained in this service element that contain authentication information. The key in the map is the identifier of the operation element and the value in the map is the authentication information for the operation element.
	//
	//  For an explanation of containment of authentication information within operation elements, see Operation.
	Operations map[string]OperationInfo
}

func (s *ServiceInfo) GetType__() bindings.BindingType {
	return ServiceInfoBindingType()
}

func (s *ServiceInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for ServiceInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func AuthenticationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["scheme_type"] = bindings.NewEnumType("com.vmware.vapi.metadata.authentication.authentication_info.scheme_type", reflect.TypeOf(AuthenticationInfoSchemeTypeEnum(AuthenticationInfoSchemeType_SESSIONLESS)))
	fieldNameMap["scheme_type"] = "SchemeType"
	fields["session_manager"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["session_manager"] = "SessionManager"
	fields["scheme"] = bindings.NewStringType()
	fieldNameMap["scheme"] = "Scheme"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("scheme_type",
		map[string][]bindings.FieldData{
			"SESSION_AWARE": []bindings.FieldData{
				bindings.NewFieldData("session_manager", true),
			},
			"SESSIONLESS": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.authentication_info", fields, reflect.TypeOf(AuthenticationInfo{}), fieldNameMap, validators)
}

func ComponentDataBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["info"] = bindings.NewReferenceType(ComponentInfoBindingType)
	fieldNameMap["info"] = "Info"
	fields["fingerprint"] = bindings.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.component_data", fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packages"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.package"}, ""), bindings.NewReferenceType(PackageInfoBindingType), reflect.TypeOf(map[string]PackageInfo{}))
	fieldNameMap["packages"] = "Packages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
	fieldNameMap["schemes"] = "Schemes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.operation_info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
	fieldNameMap["schemes"] = "Schemes"
	fields["services"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf(map[string]ServiceInfo{}))
	fieldNameMap["services"] = "Services"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.package_info", fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["schemes"] = bindings.NewListType(bindings.NewReferenceType(AuthenticationInfoBindingType), reflect.TypeOf([]AuthenticationInfo{}))
	fieldNameMap["schemes"] = "Schemes"
	fields["operations"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), bindings.NewReferenceType(OperationInfoBindingType), reflect.TypeOf(map[string]OperationInfo{}))
	fieldNameMap["operations"] = "Operations"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.authentication.service_info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

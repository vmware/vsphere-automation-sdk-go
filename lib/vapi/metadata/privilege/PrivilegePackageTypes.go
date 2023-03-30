// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.privilege.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package privilege

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentData`` class contains the privilege information of the component along with its fingerprint.
type ComponentData struct {
	// Privilege information of the component. This includes information about all the packages in the component.
	Info ComponentInfo
	// Fingerprint of the metadata of the component.
	//
	//  Privilege information could change when there is an infrastructure update. Since the data present in ComponentData#info could be quite large, ``fingerprint`` provides a convenient way to check if the data for a particular component is updated.
	//
	//  You should store the fingerprint associated with a component. After an update, by invoking the Component#fingerprint method, you can retrieve the new fingerprint for the component. If the new fingerprint and the previously stored fingerprint do not match, clients can then use the Component#get to retrieve the new privilege information for the component.
	Fingerprint string
}

func (s *ComponentData) GetType__() vapiBindings_.BindingType {
	return ComponentDataBindingType()
}

func (s *ComponentData) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ComponentData._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ComponentInfo`` class contains the privilege information of a component element.
//
//  For an explanation of privilege information contained within component elements, see Component.
type ComponentInfo struct {
	// Privilege information of all the package elements. The key in the map is the identifier of the package element and the value in the map is the privilege information for the package element.
	//
	//  For an explanation of privilege information containment within package elements, see Package.
	Packages map[string]PackageInfo
}

func (s *ComponentInfo) GetType__() vapiBindings_.BindingType {
	return ComponentInfoBindingType()
}

func (s *ComponentInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ComponentInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``OperationInfo`` class contains privilege information of an operation element.
//
//  For an explanation of containment within operation elements, see Operation.
type OperationInfo struct {
	// List of all privileges assigned to the operation element.
	Privileges []string
	// Privilege information of all the parameter elements of the operation element.
	//
	//  For an explanation of containment of privilege information within parameter elements, see PrivilegeInfo.
	PrivilegeInfo []PrivilegeInfo
}

func (s *OperationInfo) GetType__() vapiBindings_.BindingType {
	return OperationInfoBindingType()
}

func (s *OperationInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OperationInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``PackageInfo`` class contains the privilege information of a package element.
//
//  For an explanation of privilege information contained within package elements, see Package.
type PackageInfo struct {
	// List of default privileges to be used for all the operations present in this package. If a particular operation element has no explicit privileges defined in the privilege definition file, these privileges are used for enforcing authorization.
	Privileges []string
	// Information about all service elements contained in this package element that contain privilege information. The key in the map is the identifier of the service element and the value in the map is the privilege information for the service element.
	//
	//  For an explanation of privilege information containment within service elements, see Service.
	Services map[string]ServiceInfo
}

func (s *PackageInfo) GetType__() vapiBindings_.BindingType {
	return PackageInfoBindingType()
}

func (s *PackageInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PackageInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``PrivilegeInfo`` class contains the privilege information for a parameter element in an operation element.
type PrivilegeInfo struct {
	// The ``propertyPath`` points to an entity that is used in the operation element. An entity can either be present in one of the parameter elements or if a parameter is a structure element, it could also be present in one of the field elements.
	//
	//  If the privilege is assigned to an entity used in the parameter, ``propertyPath`` will just contain the name of the parameter field. If the privilege is assigned to an entity in one of the field elements of a parameter element that is a structure element, then ``propertyPath`` will contain a path to the field element starting from the parameter name.
	PropertyPath string
	// List of privileges assigned to the entity that is being referred by PrivilegeInfo#propertyPath.
	Privileges []string
}

func (s *PrivilegeInfo) GetType__() vapiBindings_.BindingType {
	return PrivilegeInfoBindingType()
}

func (s *PrivilegeInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for PrivilegeInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The ``ServiceInfo`` class contains privilege information of a service element.
//
//  For an explanation of privilege information contained within service elements, see Service.
type ServiceInfo struct {
	// Information about all operation elements contained in this service element that contain privilege information. The key in the map is the identifier of the operation element and the value in the map is the privilege information for the operation element.
	//
	//  For an explanation of containment of privilege information within operation elements, see Operation.
	Operations map[string]OperationInfo
}

func (s *ServiceInfo) GetType__() vapiBindings_.BindingType {
	return ServiceInfoBindingType()
}

func (s *ServiceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for ServiceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ComponentDataBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["info"] = vapiBindings_.NewReferenceType(ComponentInfoBindingType)
	fieldNameMap["info"] = "Info"
	fields["fingerprint"] = vapiBindings_.NewStringType()
	fieldNameMap["fingerprint"] = "Fingerprint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.component_data", fields, reflect.TypeOf(ComponentData{}), fieldNameMap, validators)
}

func ComponentInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["packages"] = vapiBindings_.NewMapType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.package"}, ""), vapiBindings_.NewReferenceType(PackageInfoBindingType), reflect.TypeOf(map[string]PackageInfo{}))
	fieldNameMap["packages"] = "Packages"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func OperationInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["privileges"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["privileges"] = "Privileges"
	fields["privilege_info"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(PrivilegeInfoBindingType), reflect.TypeOf([]PrivilegeInfo{}))
	fieldNameMap["privilege_info"] = "PrivilegeInfo"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.operation_info", fields, reflect.TypeOf(OperationInfo{}), fieldNameMap, validators)
}

func PackageInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["privileges"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["privileges"] = "Privileges"
	fields["services"] = vapiBindings_.NewMapType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, ""), vapiBindings_.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf(map[string]ServiceInfo{}))
	fieldNameMap["services"] = "Services"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.package_info", fields, reflect.TypeOf(PackageInfo{}), fieldNameMap, validators)
}

func PrivilegeInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["property_path"] = vapiBindings_.NewStringType()
	fieldNameMap["property_path"] = "PropertyPath"
	fields["privileges"] = vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["privileges"] = "Privileges"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.privilege_info", fields, reflect.TypeOf(PrivilegeInfo{}), fieldNameMap, validators)
}

func ServiceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["operations"] = vapiBindings_.NewMapType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, ""), vapiBindings_.NewReferenceType(OperationInfoBindingType), reflect.TypeOf(map[string]OperationInfo{}))
	fieldNameMap["operations"] = "Operations"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.privilege.service_info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

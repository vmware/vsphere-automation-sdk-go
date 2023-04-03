// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package vapi

import (
	vapiMetadataAuthentication_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/authentication"
	vapiMetadataCli_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/cli"
	vapiMetadataMetamodel_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
	vapiMetadataPrivilege_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
	vapiMetadataRouting_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentInfo`` class holds component metadata of the different metadata types for an API component. The class allows any combination of metadata types to be aggregated into one instance.
type ComponentInfo struct {
	// The metamodel component data
	Metamodel vapiMetadataMetamodel_.ComponentInfo
	// The CLI component data
	Cli *vapiMetadataCli_.ComponentInfo
	// The authentication component data
	Authentication *vapiMetadataAuthentication_.ComponentInfo
	// The routing component data
	Routing *vapiMetadataRouting_.ComponentInfo
	// The privilege component data
	Privilege *vapiMetadataPrivilege_.ComponentInfo
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

// The ``MetadataInfo`` is a class which holds a map of the available metadata aggregated in a ComponentInfo class.
type MetadataInfo struct {
	// Version of the current ``MetadataInfo`` class. Property value changes when the content of the ``MetadataInfo`` or referenced classes changes. This enables class processing adjustments.
	Version string
	// Component information of all available components. The key in the map is the identifier of the component and the value is the aggregated ComponentInfo.
	Metadata map[string]ComponentInfo
}

func (s *MetadataInfo) GetType__() vapiBindings_.BindingType {
	return MetadataInfoBindingType()
}

func (s *MetadataInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for MetadataInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ComponentInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["metamodel"] = vapiBindings_.NewReferenceType(vapiMetadataMetamodel_.ComponentInfoBindingType)
	fieldNameMap["metamodel"] = "Metamodel"
	fields["cli"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiMetadataCli_.ComponentInfoBindingType))
	fieldNameMap["cli"] = "Cli"
	fields["authentication"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiMetadataAuthentication_.ComponentInfoBindingType))
	fieldNameMap["authentication"] = "Authentication"
	fields["routing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiMetadataRouting_.ComponentInfoBindingType))
	fieldNameMap["routing"] = "Routing"
	fields["privilege"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vapiMetadataPrivilege_.ComponentInfoBindingType))
	fieldNameMap["privilege"] = "Privilege"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func MetadataInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = vapiBindings_.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["metadata"] = vapiBindings_.NewMapType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.component"}, ""), vapiBindings_.NewReferenceType(ComponentInfoBindingType), reflect.TypeOf(map[string]ComponentInfo{}))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata_info", fields, reflect.TypeOf(MetadataInfo{}), fieldNameMap, validators)
}

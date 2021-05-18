// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package vapi

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/authentication"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/cli"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/metamodel"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentInfo`` class holds component metadata of the different metadata types for an API component. The class allows any combination of metadata types to be aggregated into one instance.
type ComponentInfo struct {
	// The metamodel component data
	Metamodel metamodel.ComponentInfo
	// The CLI component data
	Cli *cli.ComponentInfo
	// The authentication component data
	Authentication *authentication.ComponentInfo
	// The routing component data
	Routing *routing.ComponentInfo
	// The privilege component data
	Privilege *privilege.ComponentInfo
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

// The ``MetadataInfo`` is a class which holds a map of the available metadata aggregated in a ComponentInfo class.
type MetadataInfo struct {
	// Version of the current ``MetadataInfo`` class. Property value changes when the content of the ``MetadataInfo`` or referenced classes changes. This enables class processing adjustments.
	Version string
	// Component information of all available components. The key in the map is the identifier of the component and the value is the aggregated ComponentInfo.
	Metadata map[string]ComponentInfo
}

func (s *MetadataInfo) GetType__() bindings.BindingType {
	return MetadataInfoBindingType()
}

func (s *MetadataInfo) GetDataValue__() (data.DataValue, []error) {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		log.Errorf("Error in ConvertToVapi for MetadataInfo._GetDataValue method - %s",
			bindings.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["metamodel"] = bindings.NewReferenceType(metamodel.ComponentInfoBindingType)
	fieldNameMap["metamodel"] = "Metamodel"
	fields["cli"] = bindings.NewOptionalType(bindings.NewReferenceType(cli.ComponentInfoBindingType))
	fieldNameMap["cli"] = "Cli"
	fields["authentication"] = bindings.NewOptionalType(bindings.NewReferenceType(authentication.ComponentInfoBindingType))
	fieldNameMap["authentication"] = "Authentication"
	fields["routing"] = bindings.NewOptionalType(bindings.NewReferenceType(routing.ComponentInfoBindingType))
	fieldNameMap["routing"] = "Routing"
	fields["privilege"] = bindings.NewOptionalType(bindings.NewReferenceType(privilege.ComponentInfoBindingType))
	fieldNameMap["privilege"] = "Privilege"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

func MetadataInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["metadata"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vapi.component"}, ""), bindings.NewReferenceType(ComponentInfoBindingType), reflect.TypeOf(map[string]ComponentInfo{}))
	fieldNameMap["metadata"] = "Metadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata_info", fields, reflect.TypeOf(MetadataInfo{}), fieldNameMap, validators)
}

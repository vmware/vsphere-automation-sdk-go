// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.cli.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentInfo`` is an aggregated class for CLI commands and namespaces information.
type ComponentInfo struct {
	// Information for all CLI namespaces of a component
	Namespaces []NamespaceInfo
	// Information for all CLI commands of a component
	Commands []CommandInfo
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

func ComponentInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespaces"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(NamespaceInfoBindingType), reflect.TypeOf([]NamespaceInfo{}))
	fieldNameMap["namespaces"] = "Namespaces"
	fields["commands"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(CommandInfoBindingType), reflect.TypeOf([]CommandInfo{}))
	fieldNameMap["commands"] = "Commands"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.cli.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

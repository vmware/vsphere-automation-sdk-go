// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.cli.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cli

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// The ``ComponentInfo`` is an aggregated class for CLI commands and namespaces information.
type ComponentInfo struct {
	// Information for all CLI namespaces of a component
	Namespaces []NamespaceInfo
	// Information for all CLI commands of a component
	Commands []CommandInfo
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

func ComponentInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespaces"] = bindings.NewListType(bindings.NewReferenceType(NamespaceInfoBindingType), reflect.TypeOf([]NamespaceInfo{}))
	fieldNameMap["namespaces"] = "Namespaces"
	fields["commands"] = bindings.NewListType(bindings.NewReferenceType(CommandInfoBindingType), reflect.TypeOf([]CommandInfo{}))
	fieldNameMap["commands"] = "Commands"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vapi.metadata.cli.component_info", fields, reflect.TypeOf(ComponentInfo{}), fieldNameMap, validators)
}

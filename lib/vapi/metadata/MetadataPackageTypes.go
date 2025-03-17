// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.metadata.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package metadata

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"net/url"
	"reflect"
)

// The “SourceType“ enumeration class defines the types of sources for API metadata. You specify the type of source when adding a metadata source to a metadata service.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SourceTypeEnum string

const (
	// Indicates the metadata source is a JSON file.
	SourceType_FILE SourceTypeEnum = "FILE"
	// Indicates the metadata source is a remote server.
	SourceType_REMOTE SourceTypeEnum = "REMOTE"
)

func (s SourceTypeEnum) SourceTypeEnum() bool {
	switch s {
	case SourceType_FILE:
		return true
	case SourceType_REMOTE:
		return true
	default:
		return false
	}
}

// The “LyfecycleInfo“ class contains information about the lifecycle of an API element.
// **Warning:** this class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type LifecycleInfo struct {
	// Indicates whether the API element is deprecated.
	// **Warning:** this property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Deprecated bool
}

func (s *LifecycleInfo) GetType__() vapiBindings_.BindingType {
	return LifecycleInfoBindingType()
}

func (s *LifecycleInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for LifecycleInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// The “SourceCreateSpec“ class contains the registration information for a metadata source.
type SourceCreateSpec struct {
	// English language human readable description of the source.
	Description string
	// Type of the metadata source.
	Type_ SourceTypeEnum
	// Absolute file path of the metamodel metadata file that has the metamodel information about one component element.
	Filepath *string
	// Connection information of the remote server. This should be of the format http(s)://IP:port/namespace.
	//
	//  The remote server should contain the interfaces in com.vmware.vapi.metadata.metamodel package. It could expose metamodel information of one or more components.
	Address *url.URL
}

func (s *SourceCreateSpec) GetType__() vapiBindings_.BindingType {
	return SourceCreateSpecBindingType()
}

func (s *SourceCreateSpec) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SourceCreateSpec._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Metadata source info
type SourceInfo struct {
	// Type of the metadata source
	Type_ SourceTypeEnum
	// Name of the metadata source file
	FileName *string
	// Address of the remote metadata source
	RemoteAddr *string
	// Message protocol to be used
	MsgProtocol *string
}

func (s *SourceInfo) GetType__() vapiBindings_.BindingType {
	return SourceInfoBindingType()
}

func (s *SourceInfo) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SourceInfo._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func LifecycleInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deprecated"] = vapiBindings_.NewBooleanType()
	fieldNameMap["deprecated"] = "Deprecated"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.lifecycle_info", fields, reflect.TypeOf(LifecycleInfo{}), fieldNameMap, validators)
}

func SourceCreateSpecBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = vapiBindings_.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(SourceTypeEnum(SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("type",
		map[string][]vapiBindings_.FieldData{
			"FILE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("filepath", true),
			},
			"REMOTE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.source_create_spec", fields, reflect.TypeOf(SourceCreateSpec{}), fieldNameMap, validators)
}

func SourceInfoBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = vapiBindings_.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(SourceTypeEnum(SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["file_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	fields["remote_addr"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["remote_addr"] = "RemoteAddr"
	fields["msg_protocol"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["msg_protocol"] = "MsgProtocol"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("type",
		map[string][]vapiBindings_.FieldData{
			"FILE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("file_name", true),
			},
			"REMOTE": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("remote_addr", true),
				vapiBindings_.NewFieldData("msg_protocol", true),
			},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.metadata.source_info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}

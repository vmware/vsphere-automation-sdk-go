/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vapi.metadata.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package metadata

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"net/url"
)


// The ``SourceType`` enumeration class defines the types of sources for API metadata. You specify the type of source when adding a metadata source to a metadata service.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SourceType string

const (
    // Indicates the metadata source is a JSON file.
	SourceType_FILE SourceType = "FILE"
    // Indicates the metadata source is a remote server.
	SourceType_REMOTE SourceType = "REMOTE"
)

func (s SourceType) SourceType() bool {
	switch s {
	case SourceType_FILE:
		return true
	case SourceType_REMOTE:
		return true
	default:
		return false
	}
}


// The ``SourceCreateSpec`` class contains the registration information for a metadata source.
type SourceCreateSpec struct {
    // English language human readable description of the source.
	Description string
    // Type of the metadata source.
	Type_ SourceType
    // Absolute file path of the metamodel metadata file that has the metamodel information about one component element.
	Filepath *string
    // Connection information of the remote server. This should be of the format http(s)://IP:port/namespace. 
    //
    //  The remote server should contain the interfaces in com.vmware.vapi.metadata.metamodel package. It could expose metamodel information of one or more components.
	Address *url.URL
}

// Metadata source info
type SourceInfo struct {
    // Type of the metadata source
	Type_ SourceType
    // Name of the metadata source file
	FileName *string
    // Address of the remote metadata source
	RemoteAddr *string
    // Message protocol to be used
	MsgProtocol *string
}




func SourceCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(SourceType(SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["filepath"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["filepath"] = "Filepath"
	fields["address"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["address"] = "Address"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("filepath", true),
			},
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("address", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.source_create_spec", fields, reflect.TypeOf(SourceCreateSpec{}), fieldNameMap, validators)
}

func SourceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vapi.metadata.source_type", reflect.TypeOf(SourceType(SourceType_FILE)))
	fieldNameMap["type"] = "Type_"
	fields["file_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_name"] = "FileName"
	fields["remote_addr"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["remote_addr"] = "RemoteAddr"
	fields["msg_protocol"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["msg_protocol"] = "MsgProtocol"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"FILE": []bindings.FieldData{
				bindings.NewFieldData("file_name", true),
			},
			"REMOTE": []bindings.FieldData{
				bindings.NewFieldData("remote_addr", true),
				bindings.NewFieldData("msg_protocol", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vapi.metadata.source_info", fields, reflect.TypeOf(SourceInfo{}), fieldNameMap, validators)
}



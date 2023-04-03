// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vapi.std.interposition.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package interposition

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
)

// Information about an interposed request for operation invocation. All interposers would receive an instance of this structure as an input parameter.
type InvocationRequest struct {
	// Fully qualified name of the service which contains the interposed operation. In canonical format. For example org.example.hello.
	ServiceId string
	// Name of the interposed operation. In canonical format. For example say_hello.
	OperationId string
	// Input of the interposed operation.
	OperationInput vapiData_.DataValue
	// User which started the interposed operation.
	User *SecurityPrincipal
	// Groups of the user who started the interposed operation. Would be empty if there is no authentication information.
	Groups []SecurityPrincipal
}

func (s *InvocationRequest) GetType__() vapiBindings_.BindingType {
	return InvocationRequestBindingType()
}

func (s *InvocationRequest) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InvocationRequest._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Information about the result from an interposed operation invocation. All POST interposers will receive an instance of this structure.
type InvocationResult struct {
	// Type of the invocation result.
	ResultType InvocationResultResultTypeEnum
	// Normal result value.
	Output vapiData_.DataValue
	// Error result value.
	Error_ *vapiData_.StructValue
}

func (s *InvocationResult) GetType__() vapiBindings_.BindingType {
	return InvocationResultBindingType()
}

func (s *InvocationResult) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for InvocationResult._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

// Type of the invocation result.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type InvocationResultResultTypeEnum string

const (
	// Normal invocation result.
	InvocationResultResultType_NORMAL_RESULT InvocationResultResultTypeEnum = "NORMAL_RESULT"
	// Error invocation result.
	InvocationResultResultType_ERROR_RESULT InvocationResultResultTypeEnum = "ERROR_RESULT"
)

func (r InvocationResultResultTypeEnum) InvocationResultResultTypeEnum() bool {
	switch r {
	case InvocationResultResultType_NORMAL_RESULT:
		return true
	case InvocationResultResultType_ERROR_RESULT:
		return true
	default:
		return false
	}
}

// VMODL equivalent of com.vmware.vapi.security.PrincipalId.
type SecurityPrincipal struct {
	// Principal name.
	Name string
	// Principal domain.
	Domain *string
}

func (s *SecurityPrincipal) GetType__() vapiBindings_.BindingType {
	return SecurityPrincipalBindingType()
}

func (s *SecurityPrincipal) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for SecurityPrincipal._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func InvocationRequestBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fields["operation_id"] = vapiBindings_.NewStringType()
	fieldNameMap["operation_id"] = "OperationId"
	fields["operation_input"] = vapiBindings_.NewOpaqueType()
	fieldNameMap["operation_input"] = "OperationInput"
	fields["user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(SecurityPrincipalBindingType))
	fieldNameMap["user"] = "User"
	fields["groups"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(SecurityPrincipalBindingType), reflect.TypeOf([]SecurityPrincipal{}))
	fieldNameMap["groups"] = "Groups"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.std.interposition.invocation_request", fields, reflect.TypeOf(InvocationRequest{}), fieldNameMap, validators)
}

func InvocationResultBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["result_type"] = vapiBindings_.NewEnumType("com.vmware.vapi.std.interposition.invocation_result.result_type", reflect.TypeOf(InvocationResultResultTypeEnum(InvocationResultResultType_NORMAL_RESULT)))
	fieldNameMap["result_type"] = "ResultType"
	fields["output"] = vapiBindings_.NewOptionalType(vapiBindings_.NewOpaqueType())
	fieldNameMap["output"] = "Output"
	fields["error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDynamicStructType(nil))
	fieldNameMap["error"] = "Error_"
	var validators = []vapiBindings_.Validator{}
	uv1 := vapiBindings_.NewUnionValidator("result_type",
		map[string][]vapiBindings_.FieldData{
			"NORMAL_RESULT": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("output", true),
			},
			"ERROR_RESULT": []vapiBindings_.FieldData{
				vapiBindings_.NewFieldData("error", true),
			},
		},
	)
	validators = append(validators, uv1)
	return vapiBindings_.NewStructType("com.vmware.vapi.std.interposition.invocation_result", fields, reflect.TypeOf(InvocationResult{}), fieldNameMap, validators)
}

func SecurityPrincipalBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["domain"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain"] = "Domain"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vapi.std.interposition.security_principal", fields, reflect.TypeOf(SecurityPrincipal{}), fieldNameMap, validators)
}

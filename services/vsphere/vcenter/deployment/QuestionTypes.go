/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Question.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package deployment

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``QuestionType`` enumeration class defines the type of the question raised. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type QuestionQuestionType string

const (
    // Question with answer values Yes/No. This constant field was added in vSphere API 6.7.
	QuestionQuestionType_YES_NO QuestionQuestionType = "YES_NO"
    // Question with answer values Ok/Cancel. This constant field was added in vSphere API 6.7.
	QuestionQuestionType_OK_CANCEL QuestionQuestionType = "OK_CANCEL"
    // Question with answer values Abort/Retry/Ignore. This constant field was added in vSphere API 6.7.
	QuestionQuestionType_ABORT_RETRY_IGNORE QuestionQuestionType = "ABORT_RETRY_IGNORE"
)

func (q QuestionQuestionType) QuestionQuestionType() bool {
	switch q {
	case QuestionQuestionType_YES_NO:
		return true
	case QuestionQuestionType_OK_CANCEL:
		return true
	case QuestionQuestionType_ABORT_RETRY_IGNORE:
		return true
	default:
		return false
	}
}


// The ``Question`` class contains properties to describe a deployment question. This class was added in vSphere API 6.7.
type QuestionQuestion struct {
    // Id of the question raised. This property was added in vSphere API 6.7.
	Id string
    // Message describing the question. This property was added in vSphere API 6.7.
	Question std.LocalizableMessage
    // Type of the question raised. This property was added in vSphere API 6.7.
	Type_ QuestionQuestionType
    // Default answer value. This property was added in vSphere API 6.7.
	DefaultAnswer string
    // Possible answers values. This property was added in vSphere API 6.7.
	PossibleAnswers []string
}

// The ``Info`` class contains properties to describe questions raised during the deployment process. This class was added in vSphere API 6.7.
type QuestionInfo struct {
    // One or more questions raised during the deployment. This property was added in vSphere API 6.7.
	Questions []QuestionQuestion
}

// The ``AnswerSpec`` class contains properties to describe the answer to a raised question. This class was added in vSphere API 6.7.
type QuestionAnswerSpec struct {
    // Id of the question being answered. This property was added in vSphere API 6.7.
	QuestionId string
    // The answer value. This property was added in vSphere API 6.7.
	AnswerVal string
}



func questionGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func questionGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(QuestionInfoBindingType)
}

func questionGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400,"InternalServerError": 500})
}

func questionAnswerInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(QuestionAnswerSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func questionAnswerOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func questionAnswerRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(QuestionAnswerSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"InternalServerError": 500})
}


func QuestionQuestionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["question"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["question"] = "Question"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.deployment.question.question_type", reflect.TypeOf(QuestionQuestionType(QuestionQuestionType_YES_NO)))
	fieldNameMap["type"] = "Type_"
	fields["default_answer"] = bindings.NewStringType()
	fieldNameMap["default_answer"] = "DefaultAnswer"
	fields["possible_answers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["possible_answers"] = "PossibleAnswers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.question.question", fields, reflect.TypeOf(QuestionQuestion{}), fieldNameMap, validators)
}

func QuestionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["questions"] = bindings.NewListType(bindings.NewReferenceType(QuestionQuestionBindingType), reflect.TypeOf([]QuestionQuestion{}))
	fieldNameMap["questions"] = "Questions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.question.info", fields, reflect.TypeOf(QuestionInfo{}), fieldNameMap, validators)
}

func QuestionAnswerSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["question_id"] = bindings.NewStringType()
	fieldNameMap["question_id"] = "QuestionId"
	fields["answer_val"] = bindings.NewStringType()
	fieldNameMap["answer_val"] = "AnswerVal"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.question.answer_spec", fields, reflect.TypeOf(QuestionAnswerSpec{}), fieldNameMap, validators)
}


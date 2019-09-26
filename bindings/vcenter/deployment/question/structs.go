/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Question.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package question

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``QuestionType`` enumeration class defines the type of the question raised.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type QuestionType string

const (
    // Question with answer values Yes/No.
     QuestionType_YES_NO QuestionType = "YES_NO"
    // Question with answer values Ok/Cancel.
     QuestionType_OK_CANCEL QuestionType = "OK_CANCEL"
    // Question with answer values Abort/Retry/Ignore.
     QuestionType_ABORT_RETRY_IGNORE QuestionType = "ABORT_RETRY_IGNORE"
)

func (q QuestionType) QuestionType() bool {
    switch q {
        case QuestionType_YES_NO:
            return true
        case QuestionType_OK_CANCEL:
            return true
        case QuestionType_ABORT_RETRY_IGNORE:
            return true
        default:
            return false
    }
}





// The ``Question`` class contains properties to describe a deployment question.
 type Question struct {
    Id string
    Question std.LocalizableMessage
    Type_ QuestionType
    DefaultAnswer string
    PossibleAnswers []string
}






// The ``Info`` class contains properties to describe questions raised during the deployment process.
 type Info struct {
    Questions []Question
}






// The ``AnswerSpec`` class contains properties to describe the answer to a raised question.
 type AnswerSpec struct {
    QuestionId string
    AnswerVal string
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400,"InternalServerError": 500})
}


func answerInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(AnswerSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func answerOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func answerRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"InternalServerError": 500})
}



func QuestionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["question"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["question"] = "Question"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.deployment.question.question_type", reflect.TypeOf(QuestionType(QuestionType_YES_NO)))
    fieldNameMap["type"] = "Type_"
    fields["default_answer"] = bindings.NewStringType()
    fieldNameMap["default_answer"] = "DefaultAnswer"
    fields["possible_answers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["possible_answers"] = "PossibleAnswers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.question.question",fields, reflect.TypeOf(Question{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["questions"] = bindings.NewListType(bindings.NewReferenceType(QuestionBindingType), reflect.TypeOf([]Question{}))
    fieldNameMap["questions"] = "Questions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.question.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func AnswerSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["question_id"] = bindings.NewStringType()
    fieldNameMap["question_id"] = "QuestionId"
    fields["answer_val"] = bindings.NewStringType()
    fieldNameMap["answer_val"] = "AnswerVal"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.question.answer_spec",fields, reflect.TypeOf(AnswerSpec{}), fieldNameMap, validators)
}



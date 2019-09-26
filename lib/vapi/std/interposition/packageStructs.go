/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.std.interposition.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package interposition

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)



// Information about an interposed request for operation invocation. All interposers would receive an instance of this structure as an input parameter.
type InvocationRequest struct {
    ServiceId string
    OperationId string
    OperationInput data.DataValue
    User *SecurityPrincipal
    Groups []SecurityPrincipal
}






// Information about the result from an interposed operation invocation. All POST interposers will receive an instance of this structure.
type InvocationResult struct {
    ResultType InvocationResult_ResultType
    Output data.DataValue
    Error *data.StructValue
}




    
    // Type of the invocation result.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type InvocationResult_ResultType string

    const (
        // Normal invocation result.
         InvocationResult_ResultType_NORMAL_RESULT InvocationResult_ResultType = "NORMAL_RESULT"
        // Error invocation result.
         InvocationResult_ResultType_ERROR_RESULT InvocationResult_ResultType = "ERROR_RESULT"
    )

    func (r InvocationResult_ResultType) InvocationResult_ResultType() bool {
        switch r {
            case InvocationResult_ResultType_NORMAL_RESULT:
                return true
            case InvocationResult_ResultType_ERROR_RESULT:
                return true
            default:
                return false
        }
    }



// VMODL equivalent of com.vmware.vapi.security.PrincipalId.
type SecurityPrincipal struct {
    Name string
    Domain *string
}










func InvocationRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewStringType()
    fieldNameMap["service_id"] = "ServiceId"
    fields["operation_id"] = bindings.NewStringType()
    fieldNameMap["operation_id"] = "OperationId"
    fields["operation_input"] = bindings.NewOpaqueType()
    fieldNameMap["operation_input"] = "OperationInput"
    fields["user"] = bindings.NewOptionalType(bindings.NewReferenceType(SecurityPrincipalBindingType))
    fieldNameMap["user"] = "User"
    fields["groups"] = bindings.NewListType(bindings.NewReferenceType(SecurityPrincipalBindingType), reflect.TypeOf([]SecurityPrincipal{}))
    fieldNameMap["groups"] = "Groups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.interposition.invocation_request",fields, reflect.TypeOf(InvocationRequest{}), fieldNameMap, validators)
}

func InvocationResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["result_type"] = bindings.NewEnumType("com.vmware.vapi.std.interposition.invocation_result.result_type", reflect.TypeOf(InvocationResult_ResultType(InvocationResult_ResultType_NORMAL_RESULT)))
    fieldNameMap["result_type"] = "ResultType"
    fields["output"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["output"] = "Output"
    fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("result_type",
        map[string][]bindings.FieldData {
            "NORMAL_RESULT": []bindings.FieldData {
                 bindings.NewFieldData("output", true),
            },
            "ERROR_RESULT": []bindings.FieldData {
                 bindings.NewFieldData("error", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vapi.std.interposition.invocation_result",fields, reflect.TypeOf(InvocationResult{}), fieldNameMap, validators)
}

func SecurityPrincipalBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.interposition.security_principal",fields, reflect.TypeOf(SecurityPrincipal{}), fieldNameMap, validators)
}



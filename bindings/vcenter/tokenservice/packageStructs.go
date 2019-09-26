/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.tokenservice.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tokenservice

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)



// The ``Error`` exception describes theproperties common to all standard exceptions. 
//
//  This exception serves two purposes: 
//
// #. It is the exception that clients in many programming languages can catch to handle all standard exceptions. Typically those clients will display one or more of the localizable messages from Error#messages to a human.
// #. It is the exception that methods can report when they need to report some exception, but the exception doesn't fit into any other standard exception, and in fact the only reasonable way for a client to react to the exception is to display the message(s) to a human.
type Error struct {
    Messages []std.LocalizableMessage
    Data *data.StructValue
}




func NewError() *Error {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = "null"
    return &Error{Messages:messages, ErrorType:&discriminatingValue}
}

func (Error Error) Error() string {
    return "com.vmware.vcenter.tokenservice.error"
}



// The ``InvalidGrant`` exception indicates that provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client. **Warning:** This class is available as technical preview. It may be changed in a future release.
type InvalidGrant struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue
}




func NewInvalidGrant() *InvalidGrant {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = "null"
    return &InvalidGrant{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidGrant InvalidGrant) Error() string {
    return "com.vmware.vcenter.tokenservice.invalid_grant"
}



// The ``InvalidRequest`` exception indicates that tokenExchange.ExchangeSpec is missing a required parameter, includes an unsupported parameter value (other than tokenExchange.ExchangeSpec#grant_type). **Warning:** This class is available as technical preview. It may be changed in a future release.
type InvalidRequest struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue
}




func NewInvalidRequest() *InvalidRequest {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = "null"
    return &InvalidRequest{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidRequest InvalidRequest) Error() string {
    return "com.vmware.vcenter.tokenservice.invalid_request"
}



// The ``InvalidScope`` exception indicates requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner. **Warning:** This class is available as technical preview. It may be changed in a future release.
type InvalidScope struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue
}




func NewInvalidScope() *InvalidScope {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = "null"
    return &InvalidScope{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidScope InvalidScope) Error() string {
    return "com.vmware.vcenter.tokenservice.invalid_scope"
}







func ErrorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    return bindings.NewErrorType("com.vmware.vcenter.tokenservice.error",fields,reflect.TypeOf(Error{}), fieldNameMap)
}

func InvalidGrantBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_grant",fields,reflect.TypeOf(InvalidGrant{}), fieldNameMap)
}

func InvalidRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_request",fields,reflect.TypeOf(InvalidRequest{}), fieldNameMap)
}

func InvalidScopeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_scope",fields,reflect.TypeOf(InvalidScope{}), fieldNameMap)
}



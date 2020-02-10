/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
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
    // Stack of one or more localizable messages for human exception consumers. 
    //
    //  The message at the top of the stack (first in the list) describes the exception from the perspective of the method the client invoked. Each subsequent message in the stack describes the "cause" of the prior message.
	Messages []std.LocalizableMessage
    // Data to facilitate clients responding to the method reporting a standard exception to indicating that it was unable to complete successfully. 
    //
    //  Methods may provide data that clients can use when responding to exceptions. Since the data that clients need may be specific to the context of the method reporting the exception, different methods that report the same exception may provide different data in the exception. The documentation for each each method will describe what, if any, data it provides for each exception it reports. The null, null, and null classes are intended as possible values for this property. std.DynamicID may also be useful as a value for this property (although that is not its primary purpose). Some interfaces may provide their own specific classes for use as the value of this property when reporting exceptions from their methods.
	Data *data.StructValue
}

func (Error Error) Error() string {
	return "com.vmware.vcenter.tokenservice.error"
}

// The ``InvalidGrant`` exception indicates that provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type InvalidGrant struct {
    // Stack of one or more localizable messages for human exception consumers. 
    //
    //  The message at the top of the stack (first in the list) describes the exception from the perspective of the method the client invoked. Each subsequent message in the stack describes the "cause" of the prior message.
	Messages []std.LocalizableMessage
    // Data to facilitate clients responding to the method reporting a standard exception to indicating that it was unable to complete successfully. 
    //
    //  Methods may provide data that clients can use when responding to exceptions. Since the data that clients need may be specific to the context of the method reporting the exception, different methods that report the same exception may provide different data in the exception. The documentation for each each method will describe what, if any, data it provides for each exception it reports. The null, null, and null classes are intended as possible values for this property. std.DynamicID may also be useful as a value for this property (although that is not its primary purpose). Some interfaces may provide their own specific classes for use as the value of this property when reporting exceptions from their methods.
	Data *data.StructValue
}

func (InvalidGrant InvalidGrant) Error() string {
	return "com.vmware.vcenter.tokenservice.invalid_grant"
}

// The ``InvalidRequest`` exception indicates that TokenExchangeExchangeSpec is missing a required parameter, includes an unsupported parameter value (other than TokenExchangeExchangeSpec#grant_type). **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type InvalidRequest struct {
    // Stack of one or more localizable messages for human exception consumers. 
    //
    //  The message at the top of the stack (first in the list) describes the exception from the perspective of the method the client invoked. Each subsequent message in the stack describes the "cause" of the prior message.
	Messages []std.LocalizableMessage
    // Data to facilitate clients responding to the method reporting a standard exception to indicating that it was unable to complete successfully. 
    //
    //  Methods may provide data that clients can use when responding to exceptions. Since the data that clients need may be specific to the context of the method reporting the exception, different methods that report the same exception may provide different data in the exception. The documentation for each each method will describe what, if any, data it provides for each exception it reports. The null, null, and null classes are intended as possible values for this property. std.DynamicID may also be useful as a value for this property (although that is not its primary purpose). Some interfaces may provide their own specific classes for use as the value of this property when reporting exceptions from their methods.
	Data *data.StructValue
}

func (InvalidRequest InvalidRequest) Error() string {
	return "com.vmware.vcenter.tokenservice.invalid_request"
}

// The ``InvalidScope`` exception indicates requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type InvalidScope struct {
    // Stack of one or more localizable messages for human exception consumers. 
    //
    //  The message at the top of the stack (first in the list) describes the exception from the perspective of the method the client invoked. Each subsequent message in the stack describes the "cause" of the prior message.
	Messages []std.LocalizableMessage
    // Data to facilitate clients responding to the method reporting a standard exception to indicating that it was unable to complete successfully. 
    //
    //  Methods may provide data that clients can use when responding to exceptions. Since the data that clients need may be specific to the context of the method reporting the exception, different methods that report the same exception may provide different data in the exception. The documentation for each each method will describe what, if any, data it provides for each exception it reports. The null, null, and null classes are intended as possible values for this property. std.DynamicID may also be useful as a value for this property (although that is not its primary purpose). Some interfaces may provide their own specific classes for use as the value of this property when reporting exceptions from their methods.
	Data *data.StructValue
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
	return bindings.NewErrorType("com.vmware.vcenter.tokenservice.error", fields,reflect.TypeOf(Error{}), fieldNameMap)
}

func InvalidGrantBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["data"] = "Data"
	return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_grant", fields,reflect.TypeOf(InvalidGrant{}), fieldNameMap)
}

func InvalidRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["data"] = "Data"
	return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_request", fields,reflect.TypeOf(InvalidRequest{}), fieldNameMap)
}

func InvalidScopeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
	fieldNameMap["data"] = "Data"
	return bindings.NewErrorType("com.vmware.vcenter.tokenservice.invalid_scope", fields,reflect.TypeOf(InvalidScope{}), fieldNameMap)
}



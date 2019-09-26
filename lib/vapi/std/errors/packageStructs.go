/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vapi.std.errors.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package errors

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)



// The ``AlreadyExists`` exception indicates that an attempt was made to create an entity but the entity already exists. Typically the entity has a name or identifier that is required to be unique in some context, but there is already an entity with that name or identifier in that context. 
//
//  Examples: 
//
// * Trying to create a new tag category when a tag category with the specified name already exists.
// * Trying to create a new tag in tag category when a tag with the specified name already exists the tag category.
// * Trying to create a LUN with a specific UUID on a node (for replication purposes) when a LUN with that UUID already exists on the node.
// * Trying to create a file in a directory or move or copy a file to a directory when a file with that name already exists in the directory.
type AlreadyExists struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewAlreadyExists() *AlreadyExists {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_ALREADY_EXISTS
    return &AlreadyExists{Messages:messages, ErrorType:&discriminatingValue}
}

func (AlreadyExists AlreadyExists) Error() string {
    return "com.vmware.vapi.std.errors.already_exists"
}



// The ``AlreadyInDesiredState`` exception indicates that an attempt to change the state of a resource or service had no effect because the resource or service is already in the desired state. 
//
//  Examples: 
//
// * Trying to power on a virtual machine that is already powered on.
type AlreadyInDesiredState struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewAlreadyInDesiredState() *AlreadyInDesiredState {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_ALREADY_IN_DESIRED_STATE
    return &AlreadyInDesiredState{Messages:messages, ErrorType:&discriminatingValue}
}

func (AlreadyInDesiredState AlreadyInDesiredState) Error() string {
    return "com.vmware.vapi.std.errors.already_in_desired_state"
}



// The ``ArgumentLocations`` class describes which part(s) of the input to the method caused the exception. 
//
//  Some types of exceptions are caused by the value of one of the inputs to the method, possibly due to an interaction with other inputs to the method. This class is intended to be used as the payload to identify those inputs when the method reports exceptions like InvalidArgument or NotFound. See Error#data.
type ArgumentLocations struct {
    Primary string
    Secondary []string
}






// The ``Canceled`` exception indicates that the method canceled itself in response to an explicit request to do so. Methods being "canceled" for other reasons (for example the client connection was closed, a time out occured, or due to excessive resource consumption) should not report this exception. 
//
//  Examples: 
//
// * A user is monitoring the progress of the method in a GUI and sees that it is likely to take longer than he is willing to wait and clicks the cancel button.
// * A user invokes the method using a command-line tool and decides that she didn't really want to invoke that method, and presses CTRL-c.
//
//  
//
//  Counterexamples: 
//
// * The client's connection to the server was closed. Reporting an exception is pointless since the client will not receive the error response because the connection has been closed.
// * The request is taking longer than some amount of time. The TimedOut exception would be reported if the time was specified as part of the input or is documented in the API contract.
type Canceled struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewCanceled() *Canceled {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_CANCELED
    return &Canceled{Messages:messages, ErrorType:&discriminatingValue}
}

func (Canceled Canceled) Error() string {
    return "com.vmware.vapi.std.errors.canceled"
}



// The ``ConcurrentChange`` exception indicates that a data structure, entity, or resource has been modified since some earlier point in time. Typically this happens when the client is doing the *write* portion of a read-modify-write sequence and indicates that it wants the server to notify it if the data in the server has changed after it did the *read*, so that it can avoid overwriting that change inadvertantly.
type ConcurrentChange struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewConcurrentChange() *ConcurrentChange {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_CONCURRENT_CHANGE
    return &ConcurrentChange{Messages:messages, ErrorType:&discriminatingValue}
}

func (ConcurrentChange ConcurrentChange) Error() string {
    return "com.vmware.vapi.std.errors.concurrent_change"
}



// The ``Error`` exception describes theproperties common to all standard exceptions. 
//
//  This exception serves two purposes: 
//
// #. It is the exception that clients in many programming languages can catch to handle all standard exceptions. Typically those clients will display one or more of the localizable messages from Error#messages to a human.
// #. It is the exception that methods can report when they need to report some exception, but the exception doesn't fit into any other standard exception, and in fact the only reasonable way for a client to react to the exception is to display the message(s) to a human.
type Error struct {
    Messages []std.LocalizableMessage
    Data *data.StructValue
    ErrorType *Error_Type
}




func NewError() *Error {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_ERROR
    return &Error{Messages:messages, ErrorType:&discriminatingValue}
}

func (Error Error) Error() string {
    return "com.vmware.vapi.std.errors.error"
}

    
    // Enumeration of all standard errors. Used as discriminator in protocols that have no standard means for transporting the error type, e.g. REST.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Error_Type string

    const (
        // Discriminator for the Error type.
         Error_Type_ERROR Error_Type = "ERROR"
        // Discriminator for the AlreadyExists type.
         Error_Type_ALREADY_EXISTS Error_Type = "ALREADY_EXISTS"
        // Discriminator for the AlreadyInDesiredState type.
         Error_Type_ALREADY_IN_DESIRED_STATE Error_Type = "ALREADY_IN_DESIRED_STATE"
        // Discriminator for the Canceled type.
         Error_Type_CANCELED Error_Type = "CANCELED"
        // Discriminator for the ConcurrentChange type.
         Error_Type_CONCURRENT_CHANGE Error_Type = "CONCURRENT_CHANGE"
        // Discriminator for the FeatureInUse type.
         Error_Type_FEATURE_IN_USE Error_Type = "FEATURE_IN_USE"
        // Discriminator for the InternalServerError type.
         Error_Type_INTERNAL_SERVER_ERROR Error_Type = "INTERNAL_SERVER_ERROR"
        // Discriminator for the InvalidArgument type.
         Error_Type_INVALID_ARGUMENT Error_Type = "INVALID_ARGUMENT"
        // Discriminator for the InvalidElementConfiguration type.
         Error_Type_INVALID_ELEMENT_CONFIGURATION Error_Type = "INVALID_ELEMENT_CONFIGURATION"
        // Discriminator for the InvalidElementType type.
         Error_Type_INVALID_ELEMENT_TYPE Error_Type = "INVALID_ELEMENT_TYPE"
        // Discriminator for the InvalidRequest type.
         Error_Type_INVALID_REQUEST Error_Type = "INVALID_REQUEST"
        // Discriminator for the NotAllowedInCurrentState type.
         Error_Type_NOT_ALLOWED_IN_CURRENT_STATE Error_Type = "NOT_ALLOWED_IN_CURRENT_STATE"
        // Discriminator for the NotFound type.
         Error_Type_NOT_FOUND Error_Type = "NOT_FOUND"
        // Discriminator for the OperationNotFound type.
         Error_Type_OPERATION_NOT_FOUND Error_Type = "OPERATION_NOT_FOUND"
        // Discriminator for the ResourceBusy type.
         Error_Type_RESOURCE_BUSY Error_Type = "RESOURCE_BUSY"
        // Discriminator for the ResourceInUse type.
         Error_Type_RESOURCE_IN_USE Error_Type = "RESOURCE_IN_USE"
        // Discriminator for the ResourceInaccessible type.
         Error_Type_RESOURCE_INACCESSIBLE Error_Type = "RESOURCE_INACCESSIBLE"
        // Discriminator for the ServiceUnavailable type.
         Error_Type_SERVICE_UNAVAILABLE Error_Type = "SERVICE_UNAVAILABLE"
        // Discriminator for the TimedOut type.
         Error_Type_TIMED_OUT Error_Type = "TIMED_OUT"
        // Discriminator for the UnableToAllocateResource type.
         Error_Type_UNABLE_TO_ALLOCATE_RESOURCE Error_Type = "UNABLE_TO_ALLOCATE_RESOURCE"
        // Discriminator for the Unauthenticated type.
         Error_Type_UNAUTHENTICATED Error_Type = "UNAUTHENTICATED"
        // Discriminator for the Unauthorized type.
         Error_Type_UNAUTHORIZED Error_Type = "UNAUTHORIZED"
        // Discriminator for the UnexpectedInput type.
         Error_Type_UNEXPECTED_INPUT Error_Type = "UNEXPECTED_INPUT"
        // Discriminator for the Unsupported type.
         Error_Type_UNSUPPORTED Error_Type = "UNSUPPORTED"
        // Discriminator for the UnverifiedPeer type.
         Error_Type_UNVERIFIED_PEER Error_Type = "UNVERIFIED_PEER"
    )

    func (t Error_Type) Error_Type() bool {
        switch t {
            case Error_Type_ERROR:
                return true
            case Error_Type_ALREADY_EXISTS:
                return true
            case Error_Type_ALREADY_IN_DESIRED_STATE:
                return true
            case Error_Type_CANCELED:
                return true
            case Error_Type_CONCURRENT_CHANGE:
                return true
            case Error_Type_FEATURE_IN_USE:
                return true
            case Error_Type_INTERNAL_SERVER_ERROR:
                return true
            case Error_Type_INVALID_ARGUMENT:
                return true
            case Error_Type_INVALID_ELEMENT_CONFIGURATION:
                return true
            case Error_Type_INVALID_ELEMENT_TYPE:
                return true
            case Error_Type_INVALID_REQUEST:
                return true
            case Error_Type_NOT_ALLOWED_IN_CURRENT_STATE:
                return true
            case Error_Type_NOT_FOUND:
                return true
            case Error_Type_OPERATION_NOT_FOUND:
                return true
            case Error_Type_RESOURCE_BUSY:
                return true
            case Error_Type_RESOURCE_IN_USE:
                return true
            case Error_Type_RESOURCE_INACCESSIBLE:
                return true
            case Error_Type_SERVICE_UNAVAILABLE:
                return true
            case Error_Type_TIMED_OUT:
                return true
            case Error_Type_UNABLE_TO_ALLOCATE_RESOURCE:
                return true
            case Error_Type_UNAUTHENTICATED:
                return true
            case Error_Type_UNAUTHORIZED:
                return true
            case Error_Type_UNEXPECTED_INPUT:
                return true
            case Error_Type_UNSUPPORTED:
                return true
            case Error_Type_UNVERIFIED_PEER:
                return true
            default:
                return false
        }
    }



// The ``FeatureInUse`` exception indicates that an action cannot be completed because a feature is in use. 
//
//  Examples: 
//
// * Trying to disable snapshots on a virtual machine which has a snapshot.
// * Trying to downgrade a license that has licensed features that are in use.
type FeatureInUse struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewFeatureInUse() *FeatureInUse {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_FEATURE_IN_USE
    return &FeatureInUse{Messages:messages, ErrorType:&discriminatingValue}
}

func (FeatureInUse FeatureInUse) Error() string {
    return "com.vmware.vapi.std.errors.feature_in_use"
}



// The ``FileLocations`` class identifies the file(s) that caused the method to report the exception. 
//
//  Some types of exceptions are caused by a problem with one or more files. This class is intended to be used as the payload to identify those files when the method reports exceptions like NotFound. See Error#data.
type FileLocations struct {
    Primary string
    Secondary []string
}






// The ``InternalServerError`` exception indicates that the server encounters an unexpected condition that prevented it from fulfilling the request. 
//
//  This exception is reported by the API infrastructure, so it could occur in response to the invocation of any method. 
//
//  Examples: 
//
// * The method returns a value whose type doesn't match the type type the method says it should return.
// * The method reports an exception that is not included in the list of exceptions the method says that it can report.
type InternalServerError struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewInternalServerError() *InternalServerError {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_INTERNAL_SERVER_ERROR
    return &InternalServerError{Messages:messages, ErrorType:&discriminatingValue}
}

func (InternalServerError InternalServerError) Error() string {
    return "com.vmware.vapi.std.errors.internal_server_error"
}



// The ``InvalidArgument`` exception indicates that the values received for one or more parameters are not acceptable. 
//
//  This exception is reported by the API infrastructure, so it could occur in response to the invocation of any method. It may also be reported as the result of method-specific validation. 
//
//  Examples: 
//
// * A parameter has a value that is not of the expected type.
// * A parameter has a value that is not in the required range.
// * A parameter has a value that is not one of the specifically allowed strings.
// * One property of a class is the tag for a tagged union, and has a specific value but another property of the class that is required to be specified when the tag has that value is not specified, or another property of the class that is required to be unspecified when the tag has that value is specified.
//
//  
//
//  Counterexamples: 
//
// * Trying to create a new tag in tag category when a tag with the specified name already exists the tag category. The AlreadyExists exception would be used instead.
// * Invoke the method to retrieve information about a virtual machine, passing an id that does not identify an existing virtual machine. The NotFound exception would be used instead.
// * Attempt to put a virtual machine into a folder that can only contain hosts. The InvalidElementType exception would be used instead.
// * Attempt to attach a SCSI virtual disk to an IDE port. The InvalidElementType exception would be used instead.
type InvalidArgument struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewInvalidArgument() *InvalidArgument {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_INVALID_ARGUMENT
    return &InvalidArgument{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidArgument InvalidArgument) Error() string {
    return "com.vmware.vapi.std.errors.invalid_argument"
}



// The ``InvalidElementConfiguration`` exception indicates that an attempt to modify the configuration of an element or a group containing the element failed due to the configuraton of the element. A typical case is when the method is am attempt to change the group membership of the element fails, in which case a configuration change on the element may allow the group membership change to succeed. 
//
//  Examples: 
//
// * Attempt to move a host with a fault tolerant virtual machine out of a cluster (i.e. make the host a standalone host).
// * Attempt to remove a host from a DRS cluster without putting the host into maintenance mode.
type InvalidElementConfiguration struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewInvalidElementConfiguration() *InvalidElementConfiguration {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_INVALID_ELEMENT_CONFIGURATION
    return &InvalidElementConfiguration{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidElementConfiguration InvalidElementConfiguration) Error() string {
    return "com.vmware.vapi.std.errors.invalid_element_configuration"
}



// The ``InvalidElementType`` exception indicates that the server was unable to fulfil the request because an element of a specific type cannot be a member of particular group. 
//
//  This exception could be reported, for example, if an attempt is made to put an element into the wrong type of container. 
//
//  Examples: 
//
// * Attempt to put a virtual machine into a folder that can only contain hosts.
// * Attempt to attach a SCSI virtual disk to an IDE port.
//
//  Counterexamples: 
//
// * A parameter has a value that is not of the expected type. The InvalidArgument exception would be used instead.
type InvalidElementType struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewInvalidElementType() *InvalidElementType {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_INVALID_ELEMENT_TYPE
    return &InvalidElementType{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidElementType InvalidElementType) Error() string {
    return "com.vmware.vapi.std.errors.invalid_element_type"
}



// The ``InvalidRequest`` exception indicates that the request is malformed in such a way that the server is unable to process it. 
//
//  Examples: 
//
// * The XML in a SOAP request is not well-formed so the server cannot parse the request.
// * The XML in a SOAP request is well-formed but does not match the structure required by the SOAP specification.
// * A JSON-RPC request is not valid JSON.
// * The JSON sent in a JSON-RPC request is not a valid JSON-RPC Request object.
// * The Request object from a JSON-RPC request does not match the structure required by the API infrastructure.
//
//  
//
//  Counterexamples: 
//
// * The parameter has a value that is not with the required range. The InvalidArgument exception would be used instead.
// * The name of the method specified in the request doesn't not match any known method. The NotFound exception would be used instead.
//
//  
//
//  Some transport protocols (for example JSON-RPC) include their own mechanism for reporting these kinds of errors, and the API infrastructure for a programming language may expose the errors using a language specific mechanism, so this exception might not be used.
type InvalidRequest struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewInvalidRequest() *InvalidRequest {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_INVALID_REQUEST
    return &InvalidRequest{Messages:messages, ErrorType:&discriminatingValue}
}

func (InvalidRequest InvalidRequest) Error() string {
    return "com.vmware.vapi.std.errors.invalid_request"
}



// The ``NotAllowedInCurrentState`` exception indicates that the requested method is not allowed with a resource or service in its current state. This could be because the method is performing a configuration change that is not allowed in the current state or because method itself is not allowed in the current state. 
//
//  Examples: 
//
// * Trying to add a virtual device that cannot be hot plugged to a running virtual machine.
// * Trying to upgrade the virtual hardware version for a suspended virtual machine.
// * Trying to power off, reset, or suspend a virtual machine that is not powered on.
//
//  
//
//  Counterexamples: 
//
// * Trying to power off a virtual machine that is in the process of being powered on. The ResourceBusy exception would be used instead.
type NotAllowedInCurrentState struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewNotAllowedInCurrentState() *NotAllowedInCurrentState {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_NOT_ALLOWED_IN_CURRENT_STATE
    return &NotAllowedInCurrentState{Messages:messages, ErrorType:&discriminatingValue}
}

func (NotAllowedInCurrentState NotAllowedInCurrentState) Error() string {
    return "com.vmware.vapi.std.errors.not_allowed_in_current_state"
}



// The ``NotFound`` exception indicates that a specified element could not be found. 
//
//  Examples: 
//
// * Invoke the method to retrieve information about a virtual machine, passing an id that does not identify an existing virtual machine.
// * Invoke the method to modify the configuration of a virtual nic, passing an id that does not identify an existing virtual nic in the specified virtual machine.
// * Invoke the method to remove a vswitch, passing an id that does not identify an existing vswitch.
type NotFound struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewNotFound() *NotFound {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_NOT_FOUND
    return &NotFound{Messages:messages, ErrorType:&discriminatingValue}
}

func (NotFound NotFound) Error() string {
    return "com.vmware.vapi.std.errors.not_found"
}



// The ``OperationNotFound`` exception indicates that the method specified in the request could not be found. 
//
//  Every API request specifies a service identifier and an operation identifier along with the parameters. If the API infrastructure is unable to find the requested interface or method it reports this exception. 
//
//  This exception can be reported by the API infrastructure for any method, but it is specific to the API infrastructure, and should never be reported by the implementation of any method. 
//
//  Examples: 
//
// * A client provides an invalid service or operation identifier when invoking the method using a dynamic interface (for example REST).
// * A client invokes the method from a interface, but that interface has not been installed.
//
//  
//
//  Counterexamples: 
//
// * A client invokes a task scheduling method, but provides an invalid service identifier or operation identifier. The NotFound exception would be used instead.
type OperationNotFound struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewOperationNotFound() *OperationNotFound {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_OPERATION_NOT_FOUND
    return &OperationNotFound{Messages:messages, ErrorType:&discriminatingValue}
}

func (OperationNotFound OperationNotFound) Error() string {
    return "com.vmware.vapi.std.errors.operation_not_found"
}



// The ``ResourceBusy`` exception indicates that the method could not be completed because a resource it needs is busy. 
//
//  Examples: 
//
// * Trying to power off a virtual machine that is in the process of being powered on.
//
//  
//
//  Counterexamples: 
//
// * Trying to remove a VMFS datastore when there is a virtual machine registered on any host attached to the datastore. The ResourceInUse exception would be used instead.
type ResourceBusy struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewResourceBusy() *ResourceBusy {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_RESOURCE_BUSY
    return &ResourceBusy{Messages:messages, ErrorType:&discriminatingValue}
}

func (ResourceBusy ResourceBusy) Error() string {
    return "com.vmware.vapi.std.errors.resource_busy"
}



// The ``ResourceInUse`` exception indicates that the method could not be completed because a resource is in use. 
//
//  Examples: 
//
// * Trying to remove a VMFS datastore when the is a virtual machine registered on any host attached to the datastore.
// * Trying to add a virtual switch if the physical network adapter being bridged is already in use.
//
//  
//
//  Counterexamples: 
//
// * Trying to power off a virtual machine that is in the process of being powered on. The ResourceBusy exception would be used instead.
type ResourceInUse struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewResourceInUse() *ResourceInUse {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_RESOURCE_IN_USE
    return &ResourceInUse{Messages:messages, ErrorType:&discriminatingValue}
}

func (ResourceInUse ResourceInUse) Error() string {
    return "com.vmware.vapi.std.errors.resource_in_use"
}



// The ``ResourceInaccessible`` exception indicates that the method could not be completed because an entity is not accessible. 
//
//  Examples: 
//
// * Attempt to invoke some method on a virtual machine when the virtual machine's configuration file is not accessible (for example due to a storage APD condition).
//
//  
//
//  Counterexamples: 
//
// * Attempt to invoke some method when the server is too busy. The ServiceUnavailable exception would be used instead.
// * Attempt to invoke some method when the server is undergoing maintenance. The ServiceUnavailable exception would be used instead.
// * Some method fails to contact VMware Tools running inside the virtual machine. The ServiceUnavailable exception would be used instead.
type ResourceInaccessible struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewResourceInaccessible() *ResourceInaccessible {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_RESOURCE_INACCESSIBLE
    return &ResourceInaccessible{Messages:messages, ErrorType:&discriminatingValue}
}

func (ResourceInaccessible ResourceInaccessible) Error() string {
    return "com.vmware.vapi.std.errors.resource_inaccessible"
}



// The ``ServiceUnavailable`` exception indicates that the interface is unavailable. 
//
//  Examples: 
//
// * Attempt to invoke a method when the server is too busy.
// * Attempt to invoke a method when the server is undergoing maintenance.
// * An method fails to contact VMware Tools running inside the virtual machine.
//
//  
//
//  Counterexamples: 
//
// * A client provides an invalid service or operation identifier when invoking the method using a dynamic interface (for example REST). The OperationNotFound exception would be used instead.
// * A client invokes the method from the interface, but that interface has not been installed. The OperationNotFound exception would be used instead.
type ServiceUnavailable struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewServiceUnavailable() *ServiceUnavailable {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_SERVICE_UNAVAILABLE
    return &ServiceUnavailable{Messages:messages, ErrorType:&discriminatingValue}
}

func (ServiceUnavailable ServiceUnavailable) Error() string {
    return "com.vmware.vapi.std.errors.service_unavailable"
}



// The ``TimedOut`` exception indicates that the method did not complete within the allowed amount of time. The allowed amount of time might be: 
//
// * provided by the client as an input parameter.
// * a fixed limit of the interface implementation that is a documented part of the contract of the interface.
// * a configurable limit used by the implementation of the interface.
// * a dynamic limit computed by the implementation of the interface.
//
//  The method may or may not complete after the ``TimedOut`` exception was reported. 
//
//  Examples: 
//
// * The method was unable to complete within the timeout duration specified by a parameter of the method.
//
//  
//
//  Counterexamples: 
//
// * A server implementation that puts requests into a queue before dispatching them might delete a request from the queue if it doesn't get dispatched within *n* minutes. The ServiceUnavailable exception would be used instead.
type TimedOut struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewTimedOut() *TimedOut {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_TIMED_OUT
    return &TimedOut{Messages:messages, ErrorType:&discriminatingValue}
}

func (TimedOut TimedOut) Error() string {
    return "com.vmware.vapi.std.errors.timed_out"
}



// The ``TransientIndication`` class indicates whether or not the exception is transient. 
//
//  Some types of exceptions are transient in certain situtations and not transient in other situtations. This exception payload can be used to indicate to clients whether a particular exception is transient. See Error#data.
type TransientIndication struct {
    IsTransient bool
}






// The ``UnableToAllocateResource`` exception indicates that the method failed because it was unable to allocate or acquire a required resource. 
//
//  Examples: 
//
// * Trying to power on a virtual machine when there are not enough licenses to do so.
// * Trying to power on a virtual machine that would violate a resource usage policy.
//
//  
//
//  Counterexamples: 
//
// * Trying to power off a virtual machine that is in the process of being powered on. A ResourceBusy exception would be used instead.
// * Trying to remove a VMFS datastore when the is a virtual machine registered on any host attached to the datastore. The ResourceInUse exception would be used instead.
// * Trying to add a virtual switch if the physical network adapter being bridged is already in use. The ResourceInUse exception would be used instead.
// * Attempt to invoke some method on a virtual machine when the virtual machine's configuration file is not accessible (for example due to a storage APD condition). The ResourceInaccessible exception would be used instead.
type UnableToAllocateResource struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewUnableToAllocateResource() *UnableToAllocateResource {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNABLE_TO_ALLOCATE_RESOURCE
    return &UnableToAllocateResource{Messages:messages, ErrorType:&discriminatingValue}
}

func (UnableToAllocateResource UnableToAllocateResource) Error() string {
    return "com.vmware.vapi.std.errors.unable_to_allocate_resource"
}



// The ``Unauthenticated`` exception indicates that the method requires authentication and the user is not authenticated. 
//
//  API requests may include a security context containing user credentials. For example, the user credentials could be a SAML token, a user name and password, or the session identifier for a previously established session. 
//
//  Examples: 
//
// * The SAML token in the request's security context has expired.
// * The user name and password in the request's security context are invalid.
// * The session identifier in the request's security context identifies a session that has expired.
//
//  Counterexamples: 
//
// * The user is authenticated but isn't authorized to perform the requested method. The Unauthorized exception would be used instead.
//
//  
//
//  For security reasons, the Error#data property in this exception is null, and the Error#messages property in this exception does not disclose which part of the security context is correct or incorrect. For example the messages would not disclose whether a username or a password is valid or invalid, but only that a combination of username and password is invalid.
type Unauthenticated struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
    Challenge *string
}




func NewUnauthenticated() *Unauthenticated {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNAUTHENTICATED
    return &Unauthenticated{Messages:messages, ErrorType:&discriminatingValue}
}

func (Unauthenticated Unauthenticated) Error() string {
    return "com.vmware.vapi.std.errors.unauthenticated"
}



// The ``Unauthorized`` exception indicates that the user is not authorized to perform the method. 
//
//  API requests may include a security context containing user credentials. For example, the user credentials could be a SAML token, a user name and password, or the session identifier for a previously established session. Invoking the method may require that the user identified by those credentials has particular privileges on the method or on one or more resource identifiers passed to the method. 
//
//  Examples: 
//
// * The method requires that the user have one or more privileges on the method, but the user identified by the credentials in the security context does not have the required privileges.
// * The method requires that the user have one or more privileges on a resource identifier passed to the method, but the user identified by the credentials in the security context does not have the required privileges.
//
//  
//
//  
//
//  Counterexamples: 
//
// * The SAML token in the request's security context has expired. A Unauthenticated exception would be used instead.
// * The user name and password in the request's security context are invalid. The Unauthenticated exception would be used instead.
// * The session identifier in the request's security context identifies a session that has expired. The Unauthenticated exception would be used instead.
//
//  
//
//  For security reasons, the Error#data property in this exception is null, and the Error#messages property in this exception does not disclose why the user is not authorized to perform the method. For example the messages would not disclose which privilege the user did not have or which resource identifier the user did not have the required privilege to access. The API documentation should indicate what privileges are required.
type Unauthorized struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewUnauthorized() *Unauthorized {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNAUTHORIZED
    return &Unauthorized{Messages:messages, ErrorType:&discriminatingValue}
}

func (Unauthorized Unauthorized) Error() string {
    return "com.vmware.vapi.std.errors.unauthorized"
}



// The ``UnexpectedInput`` exception indicates that the request contained a parameter or property whose name is not known by the server. 
//
//  Every method expects parameters with known names. Some of those parameters may be (or contain) classes, and the method expects those classes to contain properties with known names. If the method receives parameters or properties with names that is does not expect, this exception may be reported. 
//
//  This exception can be reported by the API infrastructure for any method, but it is specific to the API infrastructure, and should never be reported by the implementation of any method. 
//
//  Examples: 
//
// * A client using stubs generated from the interface specification for version2 of a interface invokes the method passing one or more parameters that were added in version2, but they are communicating with a server that only supports version1 of the interface.
// * A client provides an unexpected parameter or property name when invoking the method using a dynamic interface (for example REST).
type UnexpectedInput struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewUnexpectedInput() *UnexpectedInput {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNEXPECTED_INPUT
    return &UnexpectedInput{Messages:messages, ErrorType:&discriminatingValue}
}

func (UnexpectedInput UnexpectedInput) Error() string {
    return "com.vmware.vapi.std.errors.unexpected_input"
}



// The ``Unsupported`` exception indicates that the method is not supported by the interface. 
//
//  Examples: 
//
// * Trying to hot-plug a CPU when the current configuration of the VM does not support hot-plugging of CPUs.
// * Trying to change the memory size to a value that is not within the acceptable guest memory bounds supported by the virtual machine's host.
type Unsupported struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewUnsupported() *Unsupported {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNSUPPORTED
    return &Unsupported{Messages:messages, ErrorType:&discriminatingValue}
}

func (Unsupported Unsupported) Error() string {
    return "com.vmware.vapi.std.errors.unsupported"
}



// The ``UnverifiedPeer`` exception indicates that an attempt to connect to an unknown or not-yet-trusted endpoint failed because the system was unable to verify the identity of the endpoint. 
//
//  Typically the {Error#data property of this error will contain information that can be presented to a human to allow them to decide whether to trust the endpoint. If they decide to trust the endpoint, the request can be resubmitted with an indication that the endpoint should be trusted. 
//
//  Examples: 
//
// * The client provides an IP address or URL of an endpoint the system should communicate with using an SSL connection, but the endpoint's SSL certificate is self-signed, expired, or otherwise not trustworthy.
// * The client provides an IP address of a host the system should communicate with using ssh, but ssh doesn't recognize the public key of the host
//
//  
type UnverifiedPeer struct {

     Messages []std.LocalizableMessage

     Data *data.StructValue

     ErrorType *Error_Type
}




func NewUnverifiedPeer() *UnverifiedPeer {
    var messages = []std.LocalizableMessage{}
    var discriminatingValue = Error_Type_UNVERIFIED_PEER
    return &UnverifiedPeer{Messages:messages, ErrorType:&discriminatingValue}
}

func (UnverifiedPeer UnverifiedPeer) Error() string {
    return "com.vmware.vapi.std.errors.unverified_peer"
}







func AlreadyExistsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.already_exists",fields,reflect.TypeOf(AlreadyExists{}), fieldNameMap)
}

func AlreadyInDesiredStateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.already_in_desired_state",fields,reflect.TypeOf(AlreadyInDesiredState{}), fieldNameMap)
}

func ArgumentLocationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["primary"] = bindings.NewStringType()
    fieldNameMap["primary"] = "Primary"
    fields["secondary"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["secondary"] = "Secondary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.errors.argument_locations",fields, reflect.TypeOf(ArgumentLocations{}), fieldNameMap, validators)
}

func CanceledBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.canceled",fields,reflect.TypeOf(Canceled{}), fieldNameMap)
}

func ConcurrentChangeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.concurrent_change",fields,reflect.TypeOf(ConcurrentChange{}), fieldNameMap)
}

func ErrorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.error",fields,reflect.TypeOf(Error{}), fieldNameMap)
}

func FeatureInUseBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.feature_in_use",fields,reflect.TypeOf(FeatureInUse{}), fieldNameMap)
}

func FileLocationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["primary"] = bindings.NewStringType()
    fieldNameMap["primary"] = "Primary"
    fields["secondary"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["secondary"] = "Secondary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.errors.file_locations",fields, reflect.TypeOf(FileLocations{}), fieldNameMap, validators)
}

func InternalServerErrorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.internal_server_error",fields,reflect.TypeOf(InternalServerError{}), fieldNameMap)
}

func InvalidArgumentBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.invalid_argument",fields,reflect.TypeOf(InvalidArgument{}), fieldNameMap)
}

func InvalidElementConfigurationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.invalid_element_configuration",fields,reflect.TypeOf(InvalidElementConfiguration{}), fieldNameMap)
}

func InvalidElementTypeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.invalid_element_type",fields,reflect.TypeOf(InvalidElementType{}), fieldNameMap)
}

func InvalidRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.invalid_request",fields,reflect.TypeOf(InvalidRequest{}), fieldNameMap)
}

func NotAllowedInCurrentStateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.not_allowed_in_current_state",fields,reflect.TypeOf(NotAllowedInCurrentState{}), fieldNameMap)
}

func NotFoundBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.not_found",fields,reflect.TypeOf(NotFound{}), fieldNameMap)
}

func OperationNotFoundBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.operation_not_found",fields,reflect.TypeOf(OperationNotFound{}), fieldNameMap)
}

func ResourceBusyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.resource_busy",fields,reflect.TypeOf(ResourceBusy{}), fieldNameMap)
}

func ResourceInUseBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.resource_in_use",fields,reflect.TypeOf(ResourceInUse{}), fieldNameMap)
}

func ResourceInaccessibleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.resource_inaccessible",fields,reflect.TypeOf(ResourceInaccessible{}), fieldNameMap)
}

func ServiceUnavailableBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.service_unavailable",fields,reflect.TypeOf(ServiceUnavailable{}), fieldNameMap)
}

func TimedOutBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.timed_out",fields,reflect.TypeOf(TimedOut{}), fieldNameMap)
}

func TransientIndicationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["is_transient"] = bindings.NewBooleanType()
    fieldNameMap["is_transient"] = "IsTransient"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.errors.transient_indication",fields, reflect.TypeOf(TransientIndication{}), fieldNameMap, validators)
}

func UnableToAllocateResourceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unable_to_allocate_resource",fields,reflect.TypeOf(UnableToAllocateResource{}), fieldNameMap)
}

func UnauthenticatedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    fields["challenge"]= bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["challenge"] = "Challenge"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unauthenticated",fields,reflect.TypeOf(Unauthenticated{}), fieldNameMap)
}

func UnauthorizedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unauthorized",fields,reflect.TypeOf(Unauthorized{}), fieldNameMap)
}

func UnexpectedInputBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unexpected_input",fields,reflect.TypeOf(UnexpectedInput{}), fieldNameMap)
}

func UnsupportedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unsupported",fields,reflect.TypeOf(Unsupported{}), fieldNameMap)
}

func UnverifiedPeerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"]= bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["data"]= bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["data"] = "Data"
    fields["error_type"]= bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.std.errors.error.type", reflect.TypeOf(Error_Type(Error_Type_ERROR))))
    fieldNameMap["error_type"] = "ErrorType"
    return bindings.NewErrorType("com.vmware.vapi.std.errors.unverified_peer",fields,reflect.TypeOf(UnverifiedPeer{}), fieldNameMap)
}



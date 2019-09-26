/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tickets.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tickets

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)



// The ``Type`` enumeration class defines the types of console tickets. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // Virtual machine remote console ticket. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Type_VMRC Type = "VMRC"
    // Web socket console ticket. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Type_WEBMKS Type = "WEBMKS"
)

func (t Type) Type() bool {
    switch t {
        case Type_VMRC:
            return true
        case Type_WEBMKS:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class defines the information used to create the virtual machine console ticket. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type CreateSpec struct {
    Type_ Type
}






// The ``Summary`` class contains commonly used information about the virtual machine console ticket. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Summary struct {
    Ticket url.URL
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SummaryBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.console.tickets.type", reflect.TypeOf(Type(Type_VMRC)))
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.console.tickets.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ticket"] = bindings.NewUriType()
    fieldNameMap["ticket"] = "Ticket"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.console.tickets.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}



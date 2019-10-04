/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Status.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package status

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class contains information that describes the status of the service.
 type Info struct {
    Health Info_Health
    Details []std.LocalizableMessage
}




    
    // The ``Health`` enumeration class defines the possible service health states.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Info_Health string

    const (
        // No status available.
         Info_Health_NONE Info_Health = "NONE"
        // Attestation is functioning normally.
         Info_Health_OK Info_Health = "OK"
        // Attestation is functioning, however there is an issue that requires attention.
         Info_Health_WARNING Info_Health = "WARNING"
        // Attestation is not functioning.
         Info_Health_ERROR Info_Health = "ERROR"
    )

    func (h Info_Health) Info_Health() bool {
        switch h {
            case Info_Health_NONE:
                return true
            case Info_Health_OK:
                return true
            case Info_Health_WARNING:
                return true
            case Info_Health_ERROR:
                return true
            default:
                return false
        }
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
      "GET",
      "/esx/attestation/status",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health"] = bindings.NewEnumType("com.vmware.esx.attestation.status.info.health", reflect.TypeOf(Info_Health(Info_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.status.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



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



// The ``Health`` enumeration class defines the possible health states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Health string

const (
    // NONE. 
    //
    //  No status available.
     Health_NONE Health = "NONE"
    // OK. 
    //
    //  Health is normal.
     Health_OK Health = "OK"
    // WARNING 
    //
    //  Health is normal, however there is an issue that requires attention.
     Health_WARNING Health = "WARNING"
    // ERROR 
    //
    //  Not healthy.
     Health_ERROR Health = "ERROR"
)

func (h Health) Health() bool {
    switch h {
        case Health_NONE:
            return true
        case Health_OK:
            return true
        case Health_WARNING:
            return true
        case Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``ServerInfo`` class contains properties that describe the status of a key server.
 type ServerInfo struct {
    Health Health
    Details []std.LocalizableMessage
    ClientTrustServer bool
    ServerTrustClient bool
    Name string
}






// The ``Info`` class contains properties that describe the status of the Key Provider.
 type Info struct {
    Health Health
    Details []std.LocalizableMessage
    Servers []ServerInfo
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}/status",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func ServerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health"] = bindings.NewEnumType("com.vmware.esx.kms.providers.status.health", reflect.TypeOf(Health(Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    fields["client_trust_server"] = bindings.NewBooleanType()
    fieldNameMap["client_trust_server"] = "ClientTrustServer"
    fields["server_trust_client"] = bindings.NewBooleanType()
    fieldNameMap["server_trust_client"] = "ServerTrustClient"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.status.server_info",fields, reflect.TypeOf(ServerInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health"] = bindings.NewEnumType("com.vmware.esx.kms.providers.status.health", reflect.TypeOf(Health(Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ServerInfoBindingType), reflect.TypeOf([]ServerInfo{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.status.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Database.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package database

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class contains information about the health of the the database. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Status Info_Status
    Messages []Message
}




    
    // The ``Status`` enumeration class describes the health of the database. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Info_Status string

    const (
        // The database is corrupted and vCenter server functionality will be impacted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_Status_UNHEALTHY Info_Status = "UNHEALTHY"
        // The database has issues but the impact on vCenter Server is low. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_Status_DEGRADED Info_Status = "DEGRADED"
        // The database is healthy. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_Status_HEALTHY Info_Status = "HEALTHY"
    )

    func (s Info_Status) Info_Status() bool {
        switch s {
            case Info_Status_UNHEALTHY:
                return true
            case Info_Status_DEGRADED:
                return true
            case Info_Status_HEALTHY:
                return true
            default:
                return false
        }
    }



// The ``Message`` class contains a database health message along with its severity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Message struct {
    Severity Message_Severity
    Message std.LocalizableMessage
}




    
    // The ``MessageSeverity`` enumeration class defines the levels of severity for a message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Message_Severity string

    const (
        // Error message. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_ERROR Message_Severity = "ERROR"
        // Warning message. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_WARNING Message_Severity = "WARNING"
    )

    func (s Message_Severity) Message_Severity() bool {
        switch s {
            case Message_Severity_ERROR:
                return true
            case Message_Severity_WARNING:
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
       map[string]int{"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.health.database.info.status", reflect.TypeOf(Info_Status(Info_Status_UNHEALTHY)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health.database.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.health.database.message.severity", reflect.TypeOf(Message_Severity(Message_Severity_ERROR)))
    fieldNameMap["severity"] = "Severity"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health.database.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}



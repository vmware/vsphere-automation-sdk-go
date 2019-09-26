/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ProbeImportSession.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package probeImportSession

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/ovf"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/ovf/importSession"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``State`` enumeration class defines the states of ProbeImportSession.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // State of an import transfer that does not have any files available. The transfer needs the OVF descriptor to continue. If this is a PUSH transfer, the client must upload the OVF descriptor, and the transfer file list has one file info entry with a URL to which the client must upload the OVF descriptor using HTTP PUT. For pull transfers (including content library), the server is in the process of retrieving the OVF descriptor. 
    //
    //  Transition to the next state is done when the server has retrieved the complete OVF content and parsed it.
     State_PROBE_IMPORT_OVF_TRANSFER State = "PROBE_IMPORT_OVF_TRANSFER"
    // The file list contains a number of message bundles that need to be transferred to the server. If this is a PUSH transfer, the client must PUT the requested files to the server. 
    //
    //  In case the OVF descriptor does not specify any bundles this state is skipped. 
    //
    //  Transition to next state is done when the complete content of all message bundles has been retrieved by the server.
     State_PROBE_IMPORT_MSG_BUNDLES_TRANSFER State = "PROBE_IMPORT_MSG_BUNDLES_TRANSFER"
    // The server can be queried for OVF parameters, and the client can specify instantiation parameters. 
    //
    //  Specifying an OVF instantiation parameter might affect other OVF instantiation parameters.
     State_PROBE_IMPORT_SELECTING_OVF_PARAMS State = "PROBE_IMPORT_SELECTING_OVF_PARAMS"
    // The transfer failed.
     State_PROBE_IMPORT_ERROR State = "PROBE_IMPORT_ERROR"
)

func (s State) State() bool {
    switch s {
        case State_PROBE_IMPORT_OVF_TRANSFER:
            return true
        case State_PROBE_IMPORT_MSG_BUNDLES_TRANSFER:
            return true
        case State_PROBE_IMPORT_SELECTING_OVF_PARAMS:
            return true
        case State_PROBE_IMPORT_ERROR:
            return true
        default:
            return false
    }
}





// The ``Info`` class represents an import session.
 type Info struct {
    State State
    Files []ovf.OvfFileInfo
    Errors []ovf.OvfError
    Warnings []ovf.OvfWarning
    Information []ovf.OvfInfo
}









func createProbeImportSessionInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(importSession.CreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createProbeImportSessionOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
}

func createProbeImportSessionRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fieldNameMap["id"] = "Id"
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
       map[string]int{"NotFound": 404})
}


func tryInstantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ovf.OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tryInstantiateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(importSession.OvfValidationResultBindingType)
}

func tryInstantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.probe_import_session.state", reflect.TypeOf(State(State_PROBE_IMPORT_OVF_TRANSFER)))
    fieldNameMap["state"] = "State"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfFileInfoBindingType), reflect.TypeOf([]ovf.OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfErrorBindingType), reflect.TypeOf([]ovf.OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfWarningBindingType), reflect.TypeOf([]ovf.OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(ovf.OvfInfoBindingType), reflect.TypeOf([]ovf.OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.probe_import_session.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



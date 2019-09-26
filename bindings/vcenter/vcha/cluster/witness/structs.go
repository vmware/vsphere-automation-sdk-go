/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Witness.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package witness

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vcha"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``CheckSpec`` class contains placement information for validation.
 type CheckSpec struct {
    VcSpec *vcha.CredentialsSpec
    Placement vcha.PlacementSpec
}






// The ``CheckResult`` class contains the warnings and errors that will occur during the clone operation.
 type CheckResult struct {
    Warnings []std.LocalizableMessage
    Errors []std.LocalizableMessage
}






// The ``RedeploySpec`` class contains the redeploy specification.
 type RedeploySpec struct {
    VcSpec *vcha.CredentialsSpec
    Placement vcha.PlacementSpec
    HaIp *vcha.IpSpec
}









func checkInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CheckSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CheckResultBindingType)
}

func checkRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"UnverifiedPeer": 400,"NotFound": 404,"InvalidElementConfiguration": 400,"NotAllowedInCurrentState": 400,"Unauthorized": 403,"Error": 500})
}


func redeployInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(RedeploySpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func redeployOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func redeployRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}



func CheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.check_spec",fields, reflect.TypeOf(CheckSpec{}), fieldNameMap, validators)
}

func CheckResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.check_result",fields, reflect.TypeOf(CheckResult{}), fieldNameMap, validators)
}

func RedeploySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vc_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.CredentialsSpecBindingType))
    fieldNameMap["vc_spec"] = "VcSpec"
    fields["placement"] = bindings.NewReferenceType(vcha.PlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    fields["ha_ip"] = bindings.NewOptionalType(bindings.NewReferenceType(vcha.IpSpecBindingType))
    fieldNameMap["ha_ip"] = "HaIp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.witness.redeploy_spec",fields, reflect.TypeOf(RedeploySpec{}), fieldNameMap, validators)
}



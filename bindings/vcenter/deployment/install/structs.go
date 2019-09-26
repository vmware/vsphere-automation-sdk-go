/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Install.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package install

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``VcsaEmbeddedSpec`` class contains information used to configure an embedded standalone or replicated vCenter Server.
 type VcsaEmbeddedSpec struct {
    Standalone *deployment.StandaloneSpec
    Replicated *deployment.ReplicatedSpec
    CeipEnabled bool
}






// The ``InstallSpec`` class contains information used to configure the appliance installation.
 type InstallSpec struct {
    VcsaEmbedded VcsaEmbeddedSpec
    AutoAnswer *bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InstallSpecBindingType)
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}


func checkInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(InstallSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutputType() bindings.BindingType {
    return bindings.NewReferenceType(deployment.CheckInfoBindingType)
}

func checkRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}


func startInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(InstallSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func startOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func startRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}


func cancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cancelOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func cancelRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}



func VcsaEmbeddedSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.StandaloneSpecBindingType))
    fieldNameMap["standalone"] = "Standalone"
    fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.ReplicatedSpecBindingType))
    fieldNameMap["replicated"] = "Replicated"
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.install.vcsa_embedded_spec",fields, reflect.TypeOf(VcsaEmbeddedSpec{}), fieldNameMap, validators)
}

func InstallSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vcsa_embedded"] = bindings.NewReferenceType(VcsaEmbeddedSpecBindingType)
    fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
    fields["auto_answer"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_answer"] = "AutoAnswer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.install.install_spec",fields, reflect.TypeOf(InstallSpec{}), fieldNameMap, validators)
}



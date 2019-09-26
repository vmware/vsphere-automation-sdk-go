/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Migrate.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package migrate

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/deployment"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``VcsaEmbeddedSpec`` class contains information used to migrate an embedded vCenter Server for Windows to embedded vCenter Server appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcsaEmbeddedSpec struct {
    CeipEnabled bool
}






// The ``PscSpec`` class contains information used to migrate a windows Platform Service Controller to Platform Service Controller appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PscSpec struct {
    CeipEnabled bool
}






// The ``SourceVcWindows`` class contains information about the windows vCenter Server that is going to be migrated. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SourceVcWindows struct {
    Hostname string
    Username string
    Password string
}






// The ``MigrationAssistantSpec`` class contains the information needed to connect to the Migration Assistant that is running on the source windows vCenter Server machine. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type MigrationAssistantSpec struct {
    HttpsPort *int64
    SslThumbprint string
}






// The ``MigrateSpec`` class contains the fields to migrate an existing vCenter Server for Windows to an appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type MigrateSpec struct {
    SourceVcWindows SourceVcWindows
    ExistingMigrationAssistant MigrationAssistantSpec
    History *deployment.HistoryMigrationSpec
    VcsaEmbedded *VcsaEmbeddedSpec
    Psc *PscSpec
    ActiveDirectory *ActiveDirectorySpec
    AutoAnswer *bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MigrateSpecBindingType)
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
    fields["spec"] = bindings.NewReferenceType(MigrateSpecBindingType)
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
    fields["spec"] = bindings.NewReferenceType(MigrateSpecBindingType)
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
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.vcsa_embedded_spec",fields, reflect.TypeOf(VcsaEmbeddedSpec{}), fieldNameMap, validators)
}

func PscSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.psc_spec",fields, reflect.TypeOf(PscSpec{}), fieldNameMap, validators)
}

func SourceVcWindowsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.source_vc_windows",fields, reflect.TypeOf(SourceVcWindows{}), fieldNameMap, validators)
}

func MigrationAssistantSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["ssl_thumbprint"] = bindings.NewStringType()
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.migration_assistant_spec",fields, reflect.TypeOf(MigrationAssistantSpec{}), fieldNameMap, validators)
}

func MigrateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_vc_windows"] = bindings.NewReferenceType(SourceVcWindowsBindingType)
    fieldNameMap["source_vc_windows"] = "SourceVcWindows"
    fields["existing_migration_assistant"] = bindings.NewReferenceType(MigrationAssistantSpecBindingType)
    fieldNameMap["existing_migration_assistant"] = "ExistingMigrationAssistant"
    fields["history"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.HistoryMigrationSpecBindingType))
    fieldNameMap["history"] = "History"
    fields["vcsa_embedded"] = bindings.NewOptionalType(bindings.NewReferenceType(VcsaEmbeddedSpecBindingType))
    fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
    fields["psc"] = bindings.NewOptionalType(bindings.NewReferenceType(PscSpecBindingType))
    fieldNameMap["psc"] = "Psc"
    fields["active_directory"] = bindings.NewOptionalType(bindings.NewReferenceType(ActiveDirectorySpecBindingType))
    fieldNameMap["active_directory"] = "ActiveDirectory"
    fields["auto_answer"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_answer"] = "AutoAnswer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.migrate_spec",fields, reflect.TypeOf(MigrateSpec{}), fieldNameMap, validators)
}



/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Upgrade.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package upgrade

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/deployment"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``VcsaEmbeddedSpec`` class contains information used to upgrade a Embedded vCenter Server appliance.
 type VcsaEmbeddedSpec struct {
    CeipEnabled bool
}






// The ``PscSpec`` class contains information used to upgrade a Platform Service Controller appliance.
 type PscSpec struct {
    CeipEnabled bool
}






// The ``SourceApplianceSpec`` class contains information used to connect to the appliance used as the source for an upgrade.
 type SourceApplianceSpec struct {
    Hostname string
    HttpsPort *int64
    SslThumbprint *string
    SslVerify *bool
    SsoAdminUsername string
    SsoAdminPassword string
    RootPassword string
    SshVerify *bool
    SshThumbprint *string
}






// The ``UpgradeSpec`` class contains information used to configure the appliance upgrade.
 type UpgradeSpec struct {
    SourceAppliance SourceApplianceSpec
    SourceLocation deployment.LocationSpec
    History *deployment.HistoryMigrationSpec
    VcsaEmbedded *VcsaEmbeddedSpec
    Psc *PscSpec
    AutoAnswer *bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UpgradeSpecBindingType)
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
    fields["spec"] = bindings.NewReferenceType(UpgradeSpecBindingType)
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
    fields["spec"] = bindings.NewReferenceType(UpgradeSpecBindingType)
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
    return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.vcsa_embedded_spec",fields, reflect.TypeOf(VcsaEmbeddedSpec{}), fieldNameMap, validators)
}

func PscSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.psc_spec",fields, reflect.TypeOf(PscSpec{}), fieldNameMap, validators)
}

func SourceApplianceSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["sso_admin_username"] = bindings.NewStringType()
    fieldNameMap["sso_admin_username"] = "SsoAdminUsername"
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["root_password"] = bindings.NewSecretType()
    fieldNameMap["root_password"] = "RootPassword"
    fields["ssh_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssh_verify"] = "SshVerify"
    fields["ssh_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssh_thumbprint"] = "SshThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.source_appliance_spec",fields, reflect.TypeOf(SourceApplianceSpec{}), fieldNameMap, validators)
}

func UpgradeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_appliance"] = bindings.NewReferenceType(SourceApplianceSpecBindingType)
    fieldNameMap["source_appliance"] = "SourceAppliance"
    fields["source_location"] = bindings.NewReferenceType(deployment.LocationSpecBindingType)
    fieldNameMap["source_location"] = "SourceLocation"
    fields["history"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.HistoryMigrationSpecBindingType))
    fieldNameMap["history"] = "History"
    fields["vcsa_embedded"] = bindings.NewOptionalType(bindings.NewReferenceType(VcsaEmbeddedSpecBindingType))
    fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
    fields["psc"] = bindings.NewOptionalType(bindings.NewReferenceType(PscSpecBindingType))
    fieldNameMap["psc"] = "Psc"
    fields["auto_answer"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_answer"] = "AutoAnswer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.upgrade_spec",fields, reflect.TypeOf(UpgradeSpec{}), fieldNameMap, validators)
}



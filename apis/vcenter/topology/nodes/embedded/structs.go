/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Embedded.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package embedded

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``DecommissionSpec`` class contains information about the vCenter Server node to be decommissioned. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DecommissionSpec struct {
    SsoAdminUsername string
    SsoAdminPassword string
    SslThumbprint *string
    SslVerify *bool
}









func decommissionInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter.name"}, "")
    fields["spec"] = bindings.NewReferenceType(DecommissionSpecBindingType)
    fields["only_precheck"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fields["repair_replication"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["hostname"] = "Hostname"
    fieldNameMap["spec"] = "Spec"
    fieldNameMap["only_precheck"] = "OnlyPrecheck"
    fieldNameMap["repair_replication"] = "RepairReplication"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func decommissionOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func decommissionRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unsupported": 400,"InvalidArgument": 400,"UnverifiedPeer": 400,"Unauthenticated": 401})
}



func DecommissionSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_username"] = bindings.NewStringType()
    fieldNameMap["sso_admin_username"] = "SsoAdminUsername"
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.topology.nodes.embedded.decommission_spec",fields, reflect.TypeOf(DecommissionSpec{}), fieldNameMap, validators)
}



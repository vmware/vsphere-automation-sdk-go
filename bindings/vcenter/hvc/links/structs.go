/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Links.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package links

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Summary`` class contains information about the hybrid link. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Summary struct {
    Link string
    DisplayName string
}






// The ``CreateSpec`` class is the specification used for the hybrid link creation. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type CreateSpec struct {
    PscHostname string
    Port *string
    DomainName string
    Username string
    Password string
    SslThumbprint *string
    AdminGroups map[string]bool
}






// The ``CertificateInfo`` class contains information about the SSL certificate for a destination PSC endpoint. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type CertificateInfo struct {
    SslThumbprint string
}






// The ``Info`` class contains information about link. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    ConnectionHealthStatus *Info_HealthStatus
    HealthStatusMessage *std.LocalizableMessage
}




    
    // The ``HealthStatus`` enumeration class defines the possible states for health of a link. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Info_HealthStatus string

    const (
        // Connection is healthy. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_HealthStatus_HEALTHY Info_HealthStatus = "HEALTHY"
        // Connection issues will need to be remediated. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_HealthStatus_UNHEALTHY Info_HealthStatus = "UNHEALTHY"
    )

    func (h Info_HealthStatus) Info_HealthStatus() bool {
        switch h {
            case Info_HealthStatus_HEALTHY:
                return true
            case Info_HealthStatus_UNHEALTHY:
                return true
            default:
                return false
        }
    }



// The ``Credentials`` interface specifies user credentials to make a successful connection to remote endpoint. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Credentials struct {
    UserName string
    Password string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unsupported": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fieldNameMap["link"] = "Link"
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Error": 500})
}


func deleteWithCredentialsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fields["credentials"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsBindingType))
    fieldNameMap["link"] = "Link"
    fieldNameMap["credentials"] = "Credentials"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteWithCredentialsOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteWithCredentialsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Error": 500})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fieldNameMap["link"] = "Link"
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403,"ResourceBusy": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fieldNameMap["link"] = "Link"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["psc_hostname"] = bindings.NewStringType()
    fieldNameMap["psc_hostname"] = "PscHostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["port"] = "Port"
    fields["domain_name"] = bindings.NewStringType()
    fieldNameMap["domain_name"] = "DomainName"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["admin_groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["admin_groups"] = "AdminGroups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func CertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ssl_thumbprint"] = bindings.NewStringType()
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.certificate_info",fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection_health_status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.hvc.links.info.health_status", reflect.TypeOf(Info_HealthStatus(Info_HealthStatus_HEALTHY))))
    fieldNameMap["connection_health_status"] = "ConnectionHealthStatus"
    fields["health_status_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["health_status_message"] = "HealthStatusMessage"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("connection_health_status",
        map[string][]bindings.FieldData {
            "UNHEALTHY": []bindings.FieldData {
                 bindings.NewFieldData("health_status_message", true),
            },
            "HEALTHY": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CredentialsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.credentials",fields, reflect.TypeOf(Credentials{}), fieldNameMap, validators)
}



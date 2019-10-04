/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VcIdentity.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcIdentity

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/certificate_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``UpdateSpec`` class contains data to update the local vCenter identity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpdateSpec struct {
    Label *string
    Type_ *string
    SigningCertChain *certificate_management.X509CertChain
    Oauth2ResponseEndpoint *string
}






// The ``Info`` class contains data that represents a local vCenter identity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Domain string
    Label string
    Type_ string
    UpnSuffixes []string
    SigningCertChains []certificate_management.X509CertChain
    Oauth2ResponseEndpoint *string
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
       map[string]int{"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"InvalidArgument": 400})
}



func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["signing_cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType))
    fieldNameMap["signing_cert_chain"] = "SigningCertChain"
    fields["oauth2_response_endpoint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["oauth2_response_endpoint"] = "Oauth2ResponseEndpoint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.vc_identity.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["upn_suffixes"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["upn_suffixes"] = "UpnSuffixes"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType), reflect.TypeOf([]certificate_management.X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    fields["oauth2_response_endpoint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["oauth2_response_endpoint"] = "Oauth2ResponseEndpoint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.vc_identity.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



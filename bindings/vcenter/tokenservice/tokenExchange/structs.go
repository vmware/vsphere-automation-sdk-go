/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TokenExchange.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tokenExchange

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Constant field indicates that token exchange grant type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_TOKEN_EXCHANGE_GRANT = "urn:ietf:params:oauth:grant-type:token-exchange"
// Constant field indicates OAuth 2.0 access token type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_ACCESS_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:access_token"
// Constant field indicates OAuth 2.0 refresh token type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_REFRESH_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:refresh_token"
// Constant field indicates OIDC ID token type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_ID_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:id_token"
// Constant field indicates base64-encoded SAML 1.1 token type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_SAML1_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:saml1"
// Constant field indicates base64-encoded SAML 2.0 token type. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_SAML2_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:saml2"
// Constant field indicates that the security token is a bearer token. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_BEARER_TOKEN_METHOD_TYPE = "Bearer"
// Constant field indicates Info#token_type identifier is not applicable in that context. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const TokenExchange_N_A_TOKEN_METHOD_TYPE = "N_A"



// The ``ExchangeSpec`` class contains arguments required for token exchange. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type ExchangeSpec struct {
    GrantType string
    Resource *string
    Audience *string
    Scope *string
    RequestedTokenType *string
    SubjectToken string
    SubjectTokenType string
    ActorToken *string
    ActorTokenType *string
}






// The ``Info`` class contains data that represents successful token exchange response. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Info struct {
    AccessToken string
    IssuedTokenType string
    TokenType string
    ExpiresIn *int64
    Scope *string
    RefreshToken *string
}









func exchangeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ExchangeSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exchangeOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func exchangeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403})
}



func ExchangeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["grant_type"] = bindings.NewStringType()
    fieldNameMap["grant_type"] = "GrantType"
    fields["resource"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource"] = "Resource"
    fields["audience"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["audience"] = "Audience"
    fields["scope"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["scope"] = "Scope"
    fields["requested_token_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["requested_token_type"] = "RequestedTokenType"
    fields["subject_token"] = bindings.NewStringType()
    fieldNameMap["subject_token"] = "SubjectToken"
    fields["subject_token_type"] = bindings.NewStringType()
    fieldNameMap["subject_token_type"] = "SubjectTokenType"
    fields["actor_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["actor_token"] = "ActorToken"
    fields["actor_token_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["actor_token_type"] = "ActorTokenType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tokenservice.token_exchange.exchange_spec",fields, reflect.TypeOf(ExchangeSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["access_token"] = bindings.NewStringType()
    fieldNameMap["access_token"] = "AccessToken"
    fields["issued_token_type"] = bindings.NewStringType()
    fieldNameMap["issued_token_type"] = "IssuedTokenType"
    fields["token_type"] = bindings.NewStringType()
    fieldNameMap["token_type"] = "TokenType"
    fields["expires_in"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["expires_in"] = "ExpiresIn"
    fields["scope"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["scope"] = "Scope"
    fields["refresh_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["refresh_token"] = "RefreshToken"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tokenservice.token_exchange.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}



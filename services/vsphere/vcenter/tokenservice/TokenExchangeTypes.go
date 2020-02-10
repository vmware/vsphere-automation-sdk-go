/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: TokenExchange.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tokenservice

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Constant field indicates that token exchange grant type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_TOKEN_EXCHANGE_GRANT = "urn:ietf:params:oauth:grant-type:token-exchange"
// Constant field indicates OAuth 2.0 access token type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_ACCESS_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:access_token"
// Constant field indicates OAuth 2.0 refresh token type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_REFRESH_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:refresh_token"
// Constant field indicates OIDC ID token type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_ID_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:id_token"
// Constant field indicates base64-encoded SAML 1.1 token type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_SAML1_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:saml1"
// Constant field indicates base64-encoded SAML 2.0 token type. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_SAML2_TOKEN_TYPE = "urn:ietf:params:oauth:token-type:saml2"
// Constant field indicates that the security token is a bearer token. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_BEARER_TOKEN_METHOD_TYPE = "Bearer"
// Constant field indicates TokenExchangeInfo#token_type identifier is not applicable in that context. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const TokenExchange_N_A_TOKEN_METHOD_TYPE = "N_A"


// The ``ExchangeSpec`` class contains arguments required for token exchange. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TokenExchangeExchangeSpec struct {
    // The value of TokenExchange#TokenExchange_TOKEN_EXCHANGE_GRANT indicates that a token exchange is being performed. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	GrantType string
    // Indicates the location of the target service or resource where the client intends to use the requested security token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Resource *string
    // The logical name of the target service where the client intends to use the requested security token. This serves a purpose similar to the TokenExchangeExchangeSpec#resource parameter, but with the client providing a logical name rather than a location. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Audience *string
    // A list of space-delimited, case-sensitive strings, that allow the client to specify the desired scope of the requested security token in the context of the service or resource where the token will be used. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Scope *string
    // An identifier for the type of the requested security token. If the requested type is unspecified, the issued token type is at the discretion of the server and may be dictated by knowledge of the requirements of the service or resource indicated by the TokenExchangeExchangeSpec#resource or TokenExchangeExchangeSpec#audience parameter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	RequestedTokenType *string
    // A security token that represents the identity of the party on behalf of whom exchange is being made. Typically, the subject of this token will be the subject of the security token issued. Token is base64-encoded. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	SubjectToken string
    // An identifier, that indicates the type of the security token in the TokenExchangeExchangeSpec#subject_token parameter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	SubjectTokenType string
    // A security token that represents the identity of the acting party. Typically, this will be the party that is authorized to use the requested security token and act on behalf of the subject. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ActorToken *string
    // An identifier, that indicates the type of the security token in the TokenExchangeExchangeSpec#actor_token parameter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ActorTokenType *string
}

// The ``Info`` class contains data that represents successful token exchange response. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TokenExchangeInfo struct {
    // The security token issued by the server in response to the token exchange request. Token is base64-encoded. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AccessToken string
    // An identifier, that indicates the type of the security token in the TokenExchangeInfo#access_token parameter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	IssuedTokenType string
    // A case-insensitive value specifying the method of using the access token issued. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TokenType string
    // The validity lifetime, in seconds, of the token issued by the server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ExpiresIn *int64
    // Scope of the issued security token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Scope *string
    // A refresh token can be issued in cases where the client of the token exchange needs the ability to access a resource even when the original credential is no longer valid. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	RefreshToken *string
}



func tokenExchangeExchangeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(TokenExchangeExchangeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tokenExchangeExchangeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TokenExchangeInfoBindingType)
}

func tokenExchangeExchangeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(TokenExchangeExchangeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403})
}


func TokenExchangeExchangeSpecBindingType() bindings.BindingType {
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
	return bindings.NewStructType("com.vmware.vcenter.tokenservice.token_exchange.exchange_spec", fields, reflect.TypeOf(TokenExchangeExchangeSpec{}), fieldNameMap, validators)
}

func TokenExchangeInfoBindingType() bindings.BindingType {
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
	return bindings.NewStructType("com.vmware.vcenter.tokenservice.token_exchange.info", fields, reflect.TypeOf(TokenExchangeInfo{}), fieldNameMap, validators)
}


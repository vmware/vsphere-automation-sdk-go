// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"reflect"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

// JSONSsoSigner is used for signing Json request messages.
type JSONSsoSigner struct {
}

func NewJSONSsoSigner() *JSONSsoSigner {
	return &JSONSsoSigner{}
}

// Process signs the input JSON request message.
// The message is signed using user's private key. The digest and saml
// token is then added to the security context block of the execution
// context. A timestamp is also added to guard against replay attacks
// Sample input security context:
//
//	{
//	   'schemeId': 'SAML_TOKEN',
//	   'privateKey': <PRIVATE_KEY>,
//	   'samlToken': <SAML_TOKEN>,
//	   'signatureAlgorithm': <ALGORITHM>,
//	}
//
// Security context block before signing:
//
//	{
//	   'schemeId': 'SAML_TOKEN',
//	   'signatureAlgorithm': <ALGORITHM>,
//	   'timestamp': {
//	       'created': '2012-10-26T12:24:18.941Z',
//	       'expires': '2012-10-26T12:44:18.941Z',
//	}
//
// Security context block after signing:
//
//	{
//	   'schemeId': 'SAML_TOKEN',
//	   'signatureAlgorithm': <ALGORITHM>,
//	   'signature': {
//	       'samlToken': <SAML_TOKEN>,
//	       'value': <DIGEST>
//	   }
//	   'timestamp': {
//	       'created': '2012-10-26T12:24:18.941Z',
//	       'expires': '2012-10-26T12:44:18.941Z',
//	   }
//	}
func (j *JSONSsoSigner) Process(jsonRequestBody map[string]interface{}, serviceId string, operationId string, inputValue data.DataValue, ctx *core.ExecutionContext) error {
	if ctx == nil || ctx.SecurityContext() == nil {
		// anonymous requests have no security context
		return nil
	}
	if SAML_HOK_SCHEME_ID != ctx.SecurityContext().Property(AUTHENTICATION_SCHEME_ID) {
		return nil
	}

	securityContext := ctx.SecurityContext().GetAllProperties()

	samlToken, err := getSamlToken(securityContext)
	if err != nil {
		return nil
	}

	jwsAlgorithm, hashAlgorithm, err := getSignatureAlgorithm(securityContext)
	if err != nil {
		return err
	}

	privateKey, err := getPrivateKey(securityContext)
	if err != nil {
		return err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Errorf("Security Context for SAML HoK contains private key of unsupported type '%v'", reflect.TypeOf(privateKey))
		return l10n.NewRuntimeErrorNoParam("vapi.security.sso.pvtkey.invalid")
	}

	// The JSON body which will be canonicalized and signed needs to have
	// a security context with fields 'schemeId', 'signatureAlgorithm' and 'timestamp'.
	// The SAML token is not included in the content to be signed.
	newSecurityContext := map[string]interface{}{}
	newSecurityContext[AUTHENTICATION_SCHEME_ID] = SAML_HOK_SCHEME_ID
	newSecurityContext[TIMESTAMP] = GenerateRequestTimeStamp()
	newSecurityContext[SIGNATURE_ALGORITHM] = jwsAlgorithm
	err = SetSecurityContext(jsonRequestBody, newSecurityContext)
	if err != nil {
		return err
	}

	jce := JSONCanonicalEncoder{}
	toSign, canonErr := jce.Marshal(jsonRequestBody)
	if canonErr != nil {
		return canonErr
	}

	signedBytes, err := Sign(toSign, hashAlgorithm, rsaPrivateKey)
	if err != nil {
		return err
	}
	sig := base64.StdEncoding.EncodeToString(signedBytes)
	// Include the SAML token and the signature before sending the request
	newSecurityContext[SIGNATURE] = map[string]interface{}{SAML_TOKEN: samlToken, DIGEST: sig}
	return nil
}

func getSamlToken(securityContext map[string]interface{}) (string, error) {
	tokenProperty, ok := securityContext[SAML_TOKEN]
	if !ok {
		log.Error("Security Context for SAML HoK does not contain SAML token")
		return "", l10n.NewRuntimeErrorNoParam("vapi.security.sso.samltoken.invalid")
	}
	samlToken, ok := tokenProperty.(string)
	if !ok {
		log.Errorf("Security Context for SAML HoK contains SAML token of unsupported type '%v'", reflect.TypeOf(tokenProperty))
		return "", l10n.NewRuntimeErrorNoParam("vapi.security.sso.samltoken.invalid")
	}
	return samlToken, nil
}

func getSignatureAlgorithm(securityContext map[string]interface{}) (string, crypto.Hash, error) {
	algoProperty, ok := securityContext[SIGNATURE_ALGORITHM]
	if !ok {
		log.Error("Security Context for SAML HoK does not contain signature algorithm")
		return "", 0, l10n.NewRuntimeErrorNoParam("vapi.security.sso.signature.invalid")
	}
	jwsAlgorithm, ok := algoProperty.(string)
	if !ok {
		log.Errorf("Security Context for SAML HoK contains signature algorithm of unsupported type '%v'", reflect.TypeOf(algoProperty))
		return "", 0, l10n.NewRuntimeErrorNoParam("vapi.security.sso.signature.invalid")
	}
	algorithm, ok := algorithmMap[jwsAlgorithm]
	if !ok {
		log.Errorf("Security Context for SAML HoK contains unsupported signature algorithm: '%v'", jwsAlgorithm)
		return "", 0, l10n.NewRuntimeErrorNoParam("vapi.security.sso.hash.invalid")
	}
	return jwsAlgorithm, algorithm, nil
}

func getPrivateKey(securityContext map[string]interface{}) (crypto.PrivateKey, error) {
	keyProperty, ok := securityContext[PRIVATE_KEY]
	if !ok {
		log.Error("Security Context for SAML HoK does not contain private key")
		return nil, l10n.NewRuntimeErrorNoParam("vapi.security.sso.pvtkey.invalid")
	}
	pemKey, isPemKey := keyProperty.(string)
	if isPemKey {
		return ParsePrivateKey(preparePrivateKey(pemKey))
	}
	return keyProperty, nil
}

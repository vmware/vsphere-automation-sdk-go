// Copyright (c) 2023-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package contextbuilder

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
)

const samlAuthScheme = "SIGN"
const samlTokenParam = "token"
const samlSignatureParam = "signature"

var samlAuthHeaderPattern = regexp.MustCompile("^SIGN\\s+")

var _ SecurityContextBuilder = (*SamlSecurityContextBuilderImpl)(nil)

type SamlSecurityContextBuilderImpl struct {
}

func NewSamlSecurityContextBuilderImpl() *SamlSecurityContextBuilderImpl {
	return &SamlSecurityContextBuilderImpl{}
}

func (b *SamlSecurityContextBuilderImpl) BuildSecurityContext(r *http.Request) (core.SecurityContext, error) {
	// see https://wiki.eng.vmware.com/SSO/REST
	// see https://wiki.eng.vmware.com/VAPI/APIEndpoint/REST_API_Authentication
	headers, ok := r.Header[lib.AUTH_HEADER]
	if !ok {
		return nil, errors.New("missing HTTP Authorization header for SAML")
	}
	params, err := parseSamlAuthParams(headers)
	if err != nil {
		return nil, err
	}
	token, ok := params[samlTokenParam]
	if !ok || token == "" {
		return nil, errors.New("HTTP Authorization header does not contain a SAML token")
	}
	_, ok = params[samlSignatureParam]
	if ok {
		return nil, errors.New("SAML holder-of-key tokens are not supported")
	}
	gzipbytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	gzipReader, err := gzip.NewReader(bytes.NewBuffer(gzipbytes))
	if err != nil {
		return nil, err
	}
	strbytes, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}
	err = gzipReader.Close()
	if err != nil {
		return nil, err
	}
	samlToken := string(strbytes)
	secCtx := core.NewSecurityContextImpl()
	secCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SAML_BEARER_SCHEME_ID)
	secCtx.SetProperty(security.SAML_TOKEN, samlToken)
	return secCtx, nil
}

func (b *SamlSecurityContextBuilderImpl) CanHandle(r *http.Request) bool {
	headers, ok := r.Header[lib.AUTH_HEADER]
	if !ok {
		return false
	}
	if !samlAuthHeaderPattern.MatchString(headers[0]) {
		return false
	}
	for _, h := range headers {
		if strings.Contains(h, samlSignatureParam) {
			// no support for SAML holder-of-key yet
			return false
		}

	}
	return true
}

func parseSamlAuthParams(authHeaders []string) (map[string]string, error) {
	paramBuilders := make(map[string]*strings.Builder)
	for _, h := range authHeaders {
		paramStr := strings.TrimPrefix(h, samlAuthScheme)
		for _, part := range strings.Split(paramStr, ",") {
			keyvalue := strings.TrimSpace(part)
			if keyvalue == "" {
				continue
			}
			indexOfEqual := strings.IndexByte(keyvalue, '=')
			if indexOfEqual <= 0 {
				return nil, errors.New("malformed HTTP Authorization header for SAML")
			}
			key := strings.TrimSpace(keyvalue[0:indexOfEqual])
			val := strings.TrimSpace(keyvalue[indexOfEqual+1:])
			builder, ok := paramBuilders[key]
			if !ok {
				builder = &strings.Builder{}
				paramBuilders[key] = builder
			}
			unquotedVal := strings.TrimSuffix(strings.TrimPrefix(val, `"`), `"`)
			builder.WriteString(unquotedVal)
		}
	}
	params := make(map[string]string, len(paramBuilders))
	for k, v := range paramBuilders {
		params[k] = v.String()
	}
	return params, nil
}

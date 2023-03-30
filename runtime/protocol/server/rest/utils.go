/* Copyright © 2019-2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	httpStatusCodes "github.com/vmware/vsphere-automation-sdk-go/runtime/lib/rest"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

var runtimePropertiesToVapiErrorMap = map[string]string{
	"vapi.protocol.server.rest.invalid_json_content":   "com.vmware.vapi.std.errors.invalid_argument",
	"vapi.protocol.server.rest.unsupported_media_type": "com.vmware.vapi.std.errors.invalid_request",
	"vapi.protocol.server.rest.param.invalid_argument": "com.vmware.vapi.std.errors.invalid_argument",
	"vapi.protocol.server.rest.response.not_structure": "com.vmware.vapi.std.errors.internal_server_error",
}

func returnError(err []error, rw http.ResponseWriter) {
	if l10nErr, ok := err[0].(*l10n.Error); ok {
		if val, ok := runtimePropertiesToVapiErrorMap[l10nErr.ID()]; ok {
			returnBadRequest(val, rw, err)
			return
		}
	}
	returnBadRequest("com.vmware.vapi.std.errors.invalid_argument", rw, err)
}

func returnBadRequest(errorStr string, rw http.ResponseWriter, dataErr []error) {
	var errorValue *data.ErrorValue
	var status int

	errorValue = bindings.CreateErrorValueFromMessages(bindings.ERROR_MAP[errorStr], dataErr)
	status = httpStatusCodes.VAPI_TO_HTTP_ERROR_MAP[errorStr]
	responseBody, err := getResponseBody(errorValue)
	if err != nil {
		log.Errorf("Error while setting error response body: %s", err)
	}
	rw.Header().Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)
	rw.WriteHeader(status)
	rw.Write([]byte(responseBody))
}

func GetQuery(request *http.Request) map[string][]*string {
	v, _ := ParseVAPIQuery(request.URL.RawQuery)
	return v
}

func ParseVAPIQuery(query string) (m map[string][]*string, err error) {
	m = make(map[string][]*string)
	query = PrepareVAPIQuery(query)
	for query != "" {
		key := query
		if i := strings.Index(key, "&"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		var value *string
		value = nil
		if i := strings.Index(key, "="); i >= 0 {
			valueStr := key[i+1:]
			key = key[:i]
			value = &valueStr
		}
		key, err1 := url.QueryUnescape(key)
		if err != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		if value != nil {
			var strValue string
			strValue, err1 = url.QueryUnescape(*value)
			if err1 != nil {
				if err == nil {
					err = err1
				}
				continue
			}
			value = &strValue
		}
		m[key] = append(m[key], value)
	}
	return m, err
}

func PrepareVAPIQuery(query string) string {
	// Since + char in Go query parser is QueryUnescape-d to space char ' '
	// we URL encode it here to suppress that behavior
	// once parsed query keys and values are QueryDecoded so they get converted back
	query = strings.ReplaceAll(query, "+", "%2B")
	return query
}

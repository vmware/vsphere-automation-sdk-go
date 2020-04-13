/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	httpStatusCodes "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib/rest"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"net/http"
)

var runtimePropertiesToVapiErrorMap = map[string]string{
	"com.vmware.vapi.rest.unsupported_media_type":             "com.vmware.vapi.std.errors.invalid_request",
	"vapi.protocol.server.rest.param.internal_server_error": "com.vmware.vapi.std.errors.internal_server_error",
}

func returnError(err []error, rw http.ResponseWriter) {
	if l10n_err, ok := err[0].(*l10n.Error); ok {
		if val, ok := runtimePropertiesToVapiErrorMap[l10n_err.ID()]; ok {
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
	responseBody, err := setResponseBody(errorValue)
	if err != nil {
		log.Errorf("Error while setting error response body: %s", err)
	}
	rw.WriteHeader(status)
	rw.Write([]byte(responseBody))
}

/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest/contextbuilder"
)

type RequestHandler struct {
	nextApiProvider core.APIProvider
	opMetadata      protocol.OperationMetadata
	appCtxBuilders  []contextbuilder.ApplicationContextBuilder
	secCtxBuilders  []contextbuilder.SecurityContextBuilder
}

func NewRequestHandler(opMetadata protocol.OperationMetadata, nextApiProvider core.APIProvider) *RequestHandler {
	rh := RequestHandler{nextApiProvider: nextApiProvider, opMetadata: opMetadata}
	return &rh
}

func (rh *RequestHandler) SetApplicationContextBuilders(appCtxBuilders []contextbuilder.ApplicationContextBuilder) {
	rh.appCtxBuilders = appCtxBuilders
}

func (rh *RequestHandler) SetSecurityContextBuilders(secCtxBuilders []contextbuilder.SecurityContextBuilder) {
	rh.secCtxBuilders = secCtxBuilders
}

func (rh *RequestHandler) BuildExecutionContext(r *http.Request) (*core.ExecutionContext, []error) {
	var secCtx core.SecurityContext
	var appCtx *core.ApplicationContext
	var err error

	for _, appCtxBuilder := range rh.appCtxBuilders {
		if appCtxBuilder.CanHandle(r) {
			appCtx, err = appCtxBuilder.BuildApplicationContext(r)
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.security.rest.context.error",
					map[string]string{"err": err.Error()})}
			}
			break
		}
	}

	for _, secCtxBuilder := range rh.secCtxBuilders {
		if secCtxBuilder.CanHandle(r) {
			secCtx, err = secCtxBuilder.BuildSecurityContext(r)
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.security.rest.context.error",
					map[string]string{"err": err.Error()})}
			}
			break
		}
	}

	if secCtx == nil {
		log.Debugf("No security context builder can handle the request")
		secCtx = core.NewSecurityContextImpl()
	}

	return core.NewExecutionContext(appCtx, secCtx), nil
}

func (rh *RequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	msg := "Calling " + rh.opMetadata.MethodDefinition().Identifier().FullyQualifiedName() + mux.Vars(r)["id"]
	log.Debugf(msg)

	var errs []error
	opRestMetaData := rh.opMetadata.RestMetadata()
	inputDataVal := data.NewStructValue("operation-input", nil)

	// Step 1: parse uri parameters, query parameters, headers and request Body to build inputDataVal.
	errs = ParsePathParams(r, opRestMetaData, inputDataVal)
	if errs != nil {
		log.Errorf("Error while processing path params: %s", errs)
		returnError(errs, rw)
		return
	}

	errs = ParseHeaderParams(r, opRestMetaData, inputDataVal)
	if errs != nil {
		log.Errorf("Error while processing header params: %s", errs)
		returnError(errs, rw)
		return
	}

	errs = ParseQueryParams(r, opRestMetaData, inputDataVal)
	if errs != nil {
		log.Errorf("Error while processing query params: %s", errs)
		returnError(errs, rw)
		return
	}

	// @BodyField and @Body are mutually exclusive
	bodyParamName := opRestMetaData.BodyParamActualName()
	if bodyParamName == "" {
		errs := ParseBodyFieldParams(r, opRestMetaData, inputDataVal)
		if errs != nil {
			log.Errorf("Error: %s", errs)
			returnError(errs, rw)
			return
		}
	} else {
		errs := ParseBodyParams(r, bodyParamName, opRestMetaData, inputDataVal)
		if errs != nil {
			log.Errorf("Error while processing body params: %s", errs)
			returnError(errs, rw)
			return
		}
	}

	// Step 2: Build application and security context
	var ctx *core.ExecutionContext
	ctx, errs = rh.BuildExecutionContext(r)
	if errs != nil {
		log.Errorf("Error while creation of execution context", errs)
		returnError(errs, rw)
		return
	}

	// Step 3: invoke apiprovider (local provider)
	methodId := rh.opMetadata.MethodDefinition().Identifier()
	methodResult := rh.nextApiProvider.Invoke(
		methodId.InterfaceIdentifier().Name(),
		methodId.Name(),
		inputDataVal,
		ctx)

	// Step 4: process output
	var statusCode int
	var status string
	var resp string

	errs = SetResponseHeader(methodResult, opRestMetaData.ResultHeadersNameMap(), opRestMetaData.ErrorHeadersNameMap(), rw.Header())
	if errs != nil {
		log.Debugf(" header errs : ", errs)
		returnError(errs, rw)
		return
	}

	statusCode, status, errs = CreateResponseStatus(methodResult, opRestMetaData)
	if errs != nil {
		log.Debugf(" Failure :: statusCode, status: ", statusCode, status)
		returnError(errs, rw)
		return
	}

	resp, errs = CreateResponseBody(methodResult, opRestMetaData)
	if errs != nil {
		log.Debugf(" response body errs : ", errs)
		returnError(errs, rw)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Write([]byte(resp))
}

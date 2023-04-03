/* Copyright © 2019, 2021-2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rest/contextbuilder"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/tracing"
)

type RequestHandler struct {
	nextApiProvider core.APIProvider
	opMetadata      protocol.OperationMetadata
	appCtxBuilders  []contextbuilder.ApplicationContextBuilder
	secCtxBuilders  []contextbuilder.SecurityContextBuilder
	tracer          tracing.StartSpan
}

type RequestHandlerOption func(*RequestHandler)

// WithTracerRest option allows NewRequestHandler to set API server tracer that will be used to extract http headers and
// create server spans for incoming calls.
func WithTracerRest(tracer tracing.StartSpan) RequestHandlerOption {
	return func(handler *RequestHandler) {
		handler.tracer = tracer
	}
}

func WithApplicationContextBuilders(appCtxBuilders []contextbuilder.ApplicationContextBuilder) RequestHandlerOption {
	return func(handler *RequestHandler) {
		handler.appCtxBuilders = appCtxBuilders
	}
}

func WithSecurityContextBuilders(secCtxBuilders []contextbuilder.SecurityContextBuilder) RequestHandlerOption {
	return func(handler *RequestHandler) {
		handler.secCtxBuilders = secCtxBuilders
	}
}

func NewRequestHandler(opMetadata protocol.OperationMetadata, nextApiProvider core.APIProvider, opts ...RequestHandlerOption) *RequestHandler {
	rh := &RequestHandler{
		nextApiProvider: nextApiProvider,
		opMetadata:      opMetadata,
		tracer:          tracing.NoopTracer,
	}
	for _, opt := range opts {
		opt(rh)
	}
	return rh
}

// Deprecated: use WithApplicationContextBuilders builder option instead
func (rh *RequestHandler) SetApplicationContextBuilders(appCtxBuilders []contextbuilder.ApplicationContextBuilder) {
	rh.appCtxBuilders = appCtxBuilders
}

// Deprecated: use WithSecurityContextBuilders builder option instead
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

	methodId := rh.opMetadata.MethodDefinition().Identifier()
	serviceId := methodId.InterfaceIdentifier().Name()
	operationId := methodId.Name()

	if vmwTask := r.URL.Query().Get(lib.TaskRESTQueryKey); vmwTask == "true" {
		operationId += lib.TaskInvocationString
	}

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

	tracerContext, serverSpan := rh.tracer(serviceId, operationId, tracing.Rest2018, r)
	defer serverSpan.Finish()
	goContext := context.WithValue(tracerContext, core.ResponseTypeKey, core.OnlyMonoResponse)
	ctx.WithContext(goContext)

	// Step 3: invoke apiprovider (local provider)
	methodResult := rh.nextApiProvider.Invoke(
		serviceId,
		operationId,
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
	_, e := rw.Write([]byte(resp))
	if e != nil {
		log.Debugf("Error writing rest response: %v", e)
	}
}

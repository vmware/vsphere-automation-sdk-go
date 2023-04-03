/* Copyright © 2021-2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package opentracing

import (
	"context"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/tracing"
)

const (
	GolangErrorType = "golang error type"
	ComponentTag    = "VAPI"
)

// OperationDelimiter is used to delimit the name of the service from operation name. The code uses dot by default as
// that works with Wavefront. Wavefront does not like columns, forward slashes etc. dot, dash, alphanumeric seem ok.
var OperationDelimiter string = "."

// NewDefault initializes a default opentracing tracer. Use the opentracing global tracer.
// Deprecated: Use package 'otel'
func NewDefault() tracing.StartSpan {
	return NewWithTracer(opentracing.GlobalTracer())
}

// Deprecated: Use package 'otel'
func NewWithTracer(tracer opentracing.Tracer) tracing.StartSpan {
	return func(serviceId string, operationId string, protocol tracing.WireProtocol, r *http.Request) (context.Context, tracing.ServerSpan) {
		spanName := serviceId + OperationDelimiter + operationId
		wireContext, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Debug("Failed to extract opentracing headers")
		}

		// Create the span referring to the RPC client if available.
		// If wireContext == nil, a root span will be created.
		// ext.RPCServerOption() sets tag span.kind=server
		serverSpan := tracer.StartSpan(
			spanName,
			ext.RPCServerOption(wireContext))

		//Set tag for the component
		serverSpan.SetTag("component", ComponentTag)

		//Set headers, if header is not present the tag will be set to ""
		serverSpan.SetTag("http.user_agent", r.Header.Get("User-Agent"))
		serverSpan.SetTag("peer.address", peerAddress(r))
		serverSpan.SetTag("wire.protocol", protocol)
		serverSpan.SetTag("is_internal", isInternal(r))

		span := NewServerSpan(serverSpan)

		ctx := opentracing.ContextWithSpan(r.Context(), serverSpan)

		return ctx, span
	}
}

func peerAddress(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}
	ip := trimPort(r.RemoteAddr)
	return ip
}

// see bora/vim/lib/vmacore/http/request.cpp#GetUserBindingEx()
func isInternal(r *http.Request) bool {
	ip := trimPort(r.RemoteAddr)
	if !isLocalAddress(ip) {
		return false
	}
	xff := r.Header.Get("X-Forwarded-For")
	if xff == "" {
		return true
	}
	for _, tok := range strings.Split(xff, ",") {
		ip := strings.TrimSpace(tok)
		if !isLocalAddress(ip) {
			return false
		}
	}
	return true
}

func trimPort(hostAndPort string) string {
	lastColon := strings.LastIndexByte(hostAndPort, ':')
	if lastColon < 0 {
		return hostAndPort
	}
	return hostAndPort[:lastColon]
}

func isLocalAddress(ip string) bool {
	return strings.HasPrefix(ip, "127.") || ip == "::1"
}

// Copyright (c) 2022-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package otel

import (
	"context"
	"net/http"
	"strings"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	GolangErrorType = "golang error type"
	ComponentVapi   = "VAPI"

	// OperationDelimiter is used to delimit the name of the service from operation name. The code uses dot by default as
	// that works with Wavefront. Wavefront does not like colons, forward slashes etc. dot, dash, alphanumeric seem ok.
	OperationDelimiter string = "."

	TaskSpanNameSuffix string = "lro"

	ComponentKey    = attribute.Key("component")
	UserAgentKey    = attribute.Key("http.user_agent")
	PeerAddressKey  = attribute.Key("peer.address")
	WireProtocolKey = attribute.Key("wire.protocol")
	IsInternalKey   = attribute.Key("is_internal")
	ErrorTypeKey    = attribute.Key("error.type")
	LroKey          = attribute.Key("lro")

	UserAgentHeader    = "User-Agent"
	ForwardedForHeader = "X-Forwarded-For"

	operationInfoCtxKey = "com.vmware.vapi.server.otel.operation.info"
)

var componentVapiAttr = ComponentKey.String(ComponentVapi)

// Used to transfer knowledge (as a context value) about an operation from
// StartSpan to StartTaskSpan.
type operationInfo struct {
	userAgent    string
	wireProtocol tracing.WireProtocol
	isInternal   bool
}

// NewWithOpenTelemetry prepares distributed tracing based on OpenTelemetry.
// A service on VCSA should use Jaeger exporter and Jaeger propagator:
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/jaeger
// https://pkg.go.dev/go.opentelemetry.io/contrib/propagators/jaeger
func NewWithOpenTelemetry(tracer trace.Tracer, propagator propagation.TextMapPropagator) tracing.StartSpan {
	return func(serviceId string, operationId string, protocol tracing.WireProtocol, r *http.Request) (context.Context, tracing.ServerSpan) {
		spanName := serviceId + OperationDelimiter + operationId
		wireCtx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		userAgent := r.Header.Get(UserAgentHeader)
		peerAddress := peerAddress(r)
		isInternal := isInternal(r)
		traceCtx, serverSpan := tracer.Start(
			wireCtx,
			spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				componentVapiAttr,
				UserAgentKey.String(userAgent),
				PeerAddressKey.String(peerAddress),
				WireProtocolKey.String(string(protocol)),
				IsInternalKey.Bool(isInternal)))
		span := NewServerSpan(serverSpan)
		traceCtxEx := context.WithValue(
			traceCtx,
			operationInfoCtxKey,
			&operationInfo{userAgent, protocol, isInternal})
		return traceCtxEx, span
	}
}

// NewTaskTracer prepares distributed tracing based on OpenTelemetry for tasks
// (i.e. long-running operations or "lro" for short).
// A service on VCSA should use Jaeger exporter and Jaeger propagator:
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/jaeger
// https://pkg.go.dev/go.opentelemetry.io/contrib/propagators/jaeger
func NewTaskTracer(tracer trace.Tracer, propagator propagation.TextMapPropagator) tracing.StartTaskSpan {
	return func(ctx context.Context, serviceId string, operationId string, taskId string) (context.Context, tracing.ServerSpan) {
		spanName := serviceId + OperationDelimiter + operationId + OperationDelimiter + TaskSpanNameSuffix
		var attributes []attribute.KeyValue
		attributes = append(attributes, LroKey.String(taskId))
		if info, ok := ctx.Value(operationInfoCtxKey).(*operationInfo); ok {
			attributes = append(attributes, UserAgentKey.String(info.userAgent))
			attributes = append(attributes, WireProtocolKey.String(string(info.wireProtocol)))
			attributes = append(attributes, IsInternalKey.Bool(info.isInternal))
		}
		traceCtx, serverSpan := tracer.Start(
			ctx,
			spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attributes...))
		span := NewServerSpan(serverSpan)
		return traceCtx, span
	}
}

func peerAddress(r *http.Request) string {
	xff := r.Header.Get(ForwardedForHeader)
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
	xff := r.Header.Get(ForwardedForHeader)
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

/* Copyright © 2022-2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */
package otel

import (
	"reflect"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

type serverSpan struct {
	span trace.Span
}

func NewServerSpan(span trace.Span) *serverSpan {
	return &serverSpan{span}
}

func (s *serverSpan) Finish() {
	s.span.End()
}

func (s *serverSpan) LogError(err error) {
	setErrorTags(err, s.span)
}

func (s *serverSpan) LogVapiError(value *data.ErrorValue) {
	errorType := getErrorType(value)
	errorMessage := getErrorMessage(value)
	s.span.SetStatus(codes.Error, errorType)
	// Do not set attribute "error"=true, because the Jaeger exporter does it automatically,
	// and we would get a duplicate tag warning.
	// See https://github.com/open-telemetry/opentelemetry-go/blob/5568a3072367bb8ac4d2d9e759e7deb5a975b9f5/exporters/jaeger/jaeger.go#L135
	s.span.SetAttributes(ErrorTypeKey.String(errorType))
	s.span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(
		semconv.ExceptionTypeKey.String(errorType),
		semconv.ExceptionMessageKey.String(errorMessage)))
}

func setErrorTags(e error, span trace.Span) {
	var errorType string
	var errorMessage string
	if vapiError, ok := e.(bindings.Structure); ok {
		errorType = e.Error()
		errorMessage = extractDefaultMessages(vapiError)
	} else {
		errorType = GolangErrorType
		errorMessage = e.Error()
	}
	span.SetStatus(codes.Error, errorType)
	span.SetAttributes(ErrorTypeKey.String(errorType))
	span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(
		semconv.ExceptionTypeKey.String(errorType),
		semconv.ExceptionMessageKey.String(errorMessage)))
}

func extractDefaultMessages(vapiError bindings.Structure) string {
	val := reflect.ValueOf(vapiError).Elem()
	messages := val.FieldByName("Messages")
	if !messages.IsValid() {
		return ""
	}
	if localizableMessages, ok := messages.Interface().([]bindings.LocalizableMessage); ok {
		var result string
		for _, msg := range localizableMessages {
			result += msg.DefaultMessage + "\n"
		}
		return result
	}
	return ""
}

func getErrorType(errorValue *data.ErrorValue) string {
	return errorValue.Name()
}

func getErrorMessage(errorValue *data.ErrorValue) string {
	messageField, err := errorValue.Field(bindings.MESSAGES_FIELD_NAME)
	if err != nil {
		return ""
	}

	messageList, ok := messageField.(*data.ListValue)
	if !ok || len(messageList.List()) == 0 {
		return ""
	}

	var result string
	for _, msg := range messageList.List() {
		structValue, ok := msg.(*data.StructValue)
		if !ok {
			continue
		}
		defaultMessage, err := structValue.Field(bindings.DEFAULT_MESSAGE_FIELD_NAME)
		if err != nil && defaultMessage != nil {
			continue
		}
		stringDefaultMessage, ok := defaultMessage.(*data.StringValue)
		if !ok {
			continue
		}
		result += stringDefaultMessage.Value()
	}

	return result
}

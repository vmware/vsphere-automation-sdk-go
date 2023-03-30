/* Copyright © 2021-2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package opentracing

import (
	"reflect"

	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
)

type serverSpan struct {
	span opentracing.Span
}

func NewServerSpan(span opentracing.Span) *serverSpan {
	return &serverSpan{span}
}

func (s *serverSpan) Finish() {
	s.span.Finish()
}

func (s *serverSpan) LogError(err error) {
	setErrorTags(err, s.span)
}

func (s *serverSpan) LogVapiError(value *data.ErrorValue) {
	errorType := getErrorType(value)
	errorMessage := getErrorMessage(value)
	s.span.SetTag("error", true)
	s.span.SetTag("error.type", errorType)
	s.span.LogFields(opentracinglog.String("error.what", errorMessage))
}

func setErrorTags(e error, span opentracing.Span) {
	var errorType string
	var errorMessage string
	if vapiError, ok := e.(bindings.Structure); ok {
		errorType = e.Error()
		errorMessage = extractDefaultMessages(vapiError)
	} else {
		errorType = GolangErrorType
		errorMessage = e.Error()
	}
	span.SetTag("error", true)
	span.SetTag("error.type", errorType)
	span.LogFields(opentracinglog.String("error.what", errorMessage))
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

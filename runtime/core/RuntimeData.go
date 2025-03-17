// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

import "net/http"

// RequestProcessor defines contract function for accessing and modifying
// http request object
type RequestProcessor func(*http.Request) error

// ResponseAcceptor defines contract function for accessing and getting
// information from http response object
type ResponseAcceptor func(*http.Response)

// RuntimeData holds custom runtime information
type RuntimeData struct {
	requestProcessors []RequestProcessor
	responseAcceptors []ResponseAcceptor
}

// NewRuntimeData creates instance of RuntimeData object
func NewRuntimeData(requestProcessors []RequestProcessor,
	responseAcceptors []ResponseAcceptor) *RuntimeData {
	return &RuntimeData{
		requestProcessors: requestProcessors,
		responseAcceptors: responseAcceptors}
}

// GetRequestProcessors returns slice of request processing functions
// executed right before making request to the server
func (r *RuntimeData) GetRequestProcessors() []RequestProcessor {
	if r == nil {
		return nil
	}
	return r.requestProcessors
}

// GetResponseAcceptors returns slice of response accepting functions
// executed right after a response from the server
func (r *RuntimeData) GetResponseAcceptors() []ResponseAcceptor {
	if r == nil {
		return nil
	}
	return r.responseAcceptors
}

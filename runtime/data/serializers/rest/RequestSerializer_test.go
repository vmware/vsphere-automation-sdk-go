package rest_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/common"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/rest"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

func TestEmptyInputNoMetadata(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	_, err := rest.SerializeRequests(inputVal, nil, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Rest metadata cannot be nil", err.Error())
}

func TestStringInput(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val": data.NewStringValue("string1")})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUrlPathWithoutPathVariable(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val": data.NewStringValue("string1")})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/bar"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUrlPathWithIntegerVarilable(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"int_val_url": data.NewIntegerValue(100)})
	urlTemplate := "/foo/{intValUrl}"
	pathVariables := map[string]string{"int_val_url": "intValUrl"}
	metadata := protocol.NewOperationRestMetadata(
		nil, pathVariables, nil, nil, "", "GET", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/100"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUrlPathWithPathVarilable(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":     data.NewStringValue("string1"),
		"string_val_url": data.NewStringValue("string2")})
	urlTemplate := "/foo/{stringValUrl}"
	PathVariables := map[string]string{"string_val_url": "stringValUrl"}
	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariables, nil, nil, "", "GET", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/string2"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUrlPathWithQueryParm(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val": data.NewStringValue("string1")})
	PathVariable := map[string]string{"string_val_url": "stringValUrl"}
	urlTemplate := "/foo/bar?action=start"
	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariable, map[string]string{}, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/bar?action=start"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestQueryParms(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"action1_val": data.NewStringValue("start"),
		"action2_val": data.NewStringValue("stop")})
	queryParams := map[string]string{
		"action1_val": "action1",
		"action2_val": "action2",
	}
	urlTemplate := "/foo/bar"
	metadata := protocol.NewOperationRestMetadata(
		nil, nil, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPaths := []string{"/foo/bar?action1=start&action2=stop", "/foo/bar?action2=stop&action1=start"}
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Contains(t, expectedURLPaths, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestQueryParmsBoollean(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"action1_val": data.NewBooleanValue(false)})
	queryParams := map[string]string{
		"action1_val": "action1",
	}
	urlTemplate := "/foo/bar"
	metadata := protocol.NewOperationRestMetadata(
		nil, nil, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/bar?action1=false"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestQueryParmsOptional(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"action1_val": data.NewOptionalValue(data.NewStringValue("start")),
		"action2_val": data.NewOptionalValue(nil)})
	queryParams := map[string]string{
		"action1_val": "action1",
		"action2_val": "action2",
	}
	urlTemplate := "/foo/bar"
	metadata := protocol.NewOperationRestMetadata(
		nil, nil, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/bar?action1=start"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestInvalidQueryParms(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"action1_val": data.NewOptionalValue(data.NewStringValue("start")),
		"action2_val": data.NewListValue()})
	queryParams := map[string]string{
		"action1_val": "action1",
		"action2_val": "action2",
	}
	urlTemplate := "/foo/bar"
	metadata := protocol.NewOperationRestMetadata(
		nil, nil, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	_, err := rest.SerializeRequests(inputVal, nil, &metadata)
	assert.NotNil(t, err)
	assert.Equal(t, "Data value of type 'LIST' is not supported to be serialized as part of REST request URL", err.Error())
}

func TestUrlPathWithQueryParamAndOtherQueryParams(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"string_val":  data.NewStringValue("string1"),
		"action2_val": data.NewStringValue("stop")})
	queryParams := map[string]string{
		"action2_val": "action2",
	}
	urlTemplate := "/foo/bar?action1=start"
	metadata := protocol.NewOperationRestMetadata(
		nil, nil, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPaths := []string{
		"/foo/bar?action1=start&action2=stop",
		"/foo/bar?action2=stop&action1=start"}
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Contains(t, expectedURLPaths, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestRequestBodyAnnotation(t *testing.T) {
	specVal := data.NewStructValue("spec", map[string]data.DataValue{
		"string_val": data.NewStringValue("string1")})
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"spec_val": specVal})
	PathVariable := map[string]string{}
	queryParams := map[string]string{
		"action1_val": "action1",
		"action2_val": "action2",
	}
	urlTemplate := "/foo"
	requestBodyParam := "spec_val"
	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariable, queryParams, nil, requestBodyParam, "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestAllAnnotation(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"spec_val":       data.NewStructValue("spec", map[string]data.DataValue{"string_val": data.NewStringValue("string1")}),
		"string_val_url": data.NewStringValue("string2"),
		"action_val":     data.NewStringValue("stop")})
	PathVariable := map[string]string{
		"string_val_url": "stringValUrl",
	}
	queryParams := map[string]string{
		"action_val": "action",
	}
	urlTemplate := "/foo/{stringValUrl}"
	requestBodyParam := "spec_val"

	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariable, queryParams, nil, requestBodyParam, "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo/string2?action=stop"
	expectedBody := "{\"string_val\":\"string1\"}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUsernameSecurityContext(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	securityCtx := security.NewUserPasswordSecurityContext("user", "pwd")
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	result, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	b64Credentials := base64.StdEncoding.EncodeToString([]byte("user:pwd"))
	authorizationVal := b64Credentials
	expectedBody := "{}"
	expectedHeaders := map[string]string{"Authorization": "Basic " + authorizationVal}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())

	// Test password is nil
	securityCtxNoPwd := core.NewSecurityContextImpl()
	securityCtxNoPwd.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	securityCtxNoPwd.SetProperty(security.USER_KEY, "user")
	securityCtxNoPwd.SetProperty(security.PASSWORD_KEY, nil)
	ctx = core.NewExecutionContext(appCtx, securityCtxNoPwd)
	result, err = rest.SerializeRequests(inputVal, ctx, &metadata)
	b64Credentials = base64.StdEncoding.EncodeToString([]byte("user:"))
	authorizationVal = b64Credentials
	expectedBody = "{}"
	expectedHeaders = map[string]string{"Authorization": "Basic " + authorizationVal}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestUsernameSecurityContextNegative(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	appCtx := common.NewDefaultApplicationContext()

	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.USER_PASSWORD_SCHEME_ID)
	ctx := core.NewExecutionContext(appCtx, securityCtx)

	securityCtx.SetProperty(security.USER_KEY, 1)
	_, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	assert.Equal(t, "Invalid type for userName, expected type string, actual type int", err.Error())
}

func TestSessionSecurityContext(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	sessionID := "some-session-id"
	securityCtx := security.NewSessionSecurityContext(sessionID)
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	result, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	expectedBody := "{}"
	expectedHeaders := map[string]string{security.SESSION_ID_KEY: sessionID}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())

	// Test session ID is nil
	securityCtxNoSessionID := core.NewSecurityContextImpl()
	securityCtxNoSessionID.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	securityCtxNoSessionID.SetProperty(security.SESSION_ID, nil)
	appCtx = common.NewDefaultApplicationContext()
	ctx = core.NewExecutionContext(appCtx, securityCtxNoSessionID)
	result, err = rest.SerializeRequests(inputVal, ctx, &metadata)
	expectedBody = "{}"
	expectedHeaders = map[string]string{security.SESSION_ID_KEY: ""}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestSessionSecurityContextNegative(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.SESSION_SCHEME_ID)
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)

	securityCtx.SetProperty(security.SESSION_ID, 123)
	_, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	assert.Equal(t, "Invalid type for sessionId, expected type string, actual type int",
		err.Error())
}

func TestOauthSecurityContext(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	bearerToken := "some-token"
	securityCtx := security.NewOauthSecurityContext(bearerToken)
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)
	result, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	expectedBody := "{}"
	expectedHeaders := map[string]string{security.CSP_AUTH_TOKEN_KEY: bearerToken}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())

	// Test token is nil
	securityCtxNoToken := core.NewSecurityContextImpl()
	securityCtxNoToken.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.OAUTH_SCHEME_ID)
	securityCtxNoToken.SetProperty(security.ACCESS_TOKEN, nil)
	appCtx = common.NewDefaultApplicationContext()
	ctx = core.NewExecutionContext(appCtx, securityCtxNoToken)
	result, err = rest.SerializeRequests(inputVal, ctx, &metadata)
	expectedBody = "{}"
	expectedHeaders = map[string]string{security.CSP_AUTH_TOKEN_KEY: ""}
	assert.Nil(t, err)
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestOauthSecurityContextNegative(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	securityCtx := core.NewSecurityContextImpl()
	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, security.OAUTH_SCHEME_ID)
	appCtx := common.NewDefaultApplicationContext()
	ctx := core.NewExecutionContext(appCtx, securityCtx)

	securityCtx.SetProperty(security.ACCESS_TOKEN, 123)
	_, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	assert.Equal(t, "Invalid type for accessToken, expected type string, actual type int", err.Error())
}

func TestSecurityContextInvalidScheme(t *testing.T) {
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{})
	metadata := protocol.NewOperationRestMetadata(
		nil, map[string]string{}, nil, nil, "", "GET", "/foo/bar", nil, 200, nil, nil)
	appCtx := common.NewDefaultApplicationContext()
	securityCtx := core.NewSecurityContextImpl()
	ctx := core.NewExecutionContext(appCtx, securityCtx)

	securityCtx.SetProperty(security.AUTHENTICATION_SCHEME_ID, "InvalidScheme")
	_, err := rest.SerializeRequests(inputVal, ctx, &metadata)
	assert.Equal(t, "Invalid authentication scheme ID, "+
		"expected one of [com.vmware.vapi.std.security.user_pass, "+
		"com.vmware.vapi.std.security.session_id, "+
		"com.vmware.vapi.std.security.oauth], "+
		"actual value InvalidScheme", err.Error())

	securityCtx2 := core.NewSecurityContextImpl()
	ctx2 := core.NewExecutionContext(appCtx, securityCtx2)
	_, err = rest.SerializeRequests(inputVal, ctx2, &metadata)
	assert.Equal(t, "Invalid authentication scheme ID, "+
		"expected one of [com.vmware.vapi.std.security.user_pass, "+
		"com.vmware.vapi.std.security.session_id, "+
		"com.vmware.vapi.std.security.oauth], "+
		"actual value %!s(<nil>)", err.Error())
}

func TestBody(t *testing.T) {
	specVal := data.NewStructValue("spec", map[string]data.DataValue{
		"string_val": data.NewStringValue("string1")})
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"spec_val": specVal})
	PathVariable := map[string]string{}
	queryParams := map[string]string{}
	urlTemplate := "/foo"
	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariable, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedURLPath := "/foo"
	expectedBody := "{\"spec_val\":{\"string_val\":\"string1\"}}"
	expectedHeaders := make(map[string]string)
	assert.Nil(t, err)
	assert.Equal(t, expectedURLPath, result.URLPath())
	assert.Equal(t, expectedHeaders, result.InputHeaders())
	assert.Equal(t, expectedBody, result.RequestBody())
}

func TestBodyOptionalNil(t *testing.T) {
	optionalVal := data.NewOptionalValue(data.NewStringValue("abc"))
	nilOptionalVal := data.NewOptionalValue(nil)
	specVal := data.NewStructValue("spec", map[string]data.DataValue{
		"name":        optionalVal,
		"description": nilOptionalVal})
	inputVal := data.NewStructValue("operation-input", map[string]data.DataValue{
		"spec_val": specVal})
	PathVariable := map[string]string{}
	queryParams := map[string]string{}
	urlTemplate := "/foo"
	metadata := protocol.NewOperationRestMetadata(
		nil, PathVariable, queryParams, nil, "", "POST", urlTemplate, nil, 200, nil, nil)
	result, err := rest.SerializeRequests(inputVal, nil, &metadata)
	expectedBody := "{\"spec_val\":{\"name\":\"abc\"}}"
	assert.Nil(t, err)
	assert.Equal(t, expectedBody, result.RequestBody())
}

package rest

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type Request struct {
	urlPath      string
	inputHeaders map[string]string
	requestBody  string
}

func NewRequest(urlPath string, inputHeaders map[string]string, requestBody string) *Request {
	return &Request{urlPath: urlPath, inputHeaders: inputHeaders, requestBody: requestBody}
}

func (result *Request) URLPath() string {
	return result.urlPath
}

func (result *Request) SetURLPath(path string) {
	result.urlPath = path
}

func (result *Request) InputHeaders() map[string]string {
	return result.inputHeaders
}

func (result *Request) SetInputHeaders(headers map[string]string) {
	result.inputHeaders = headers
}

func (result *Request) RequestBody() string {
	return result.requestBody
}

func (result *Request) SetRequestBody(body string) {
	result.requestBody = body
}

// SerializeRequests serializes a request as a REST request
// Return Request with urlPath, inputHeaders and requestBody
func SerializeRequests(inputValue *data.StructValue, ctx *core.ExecutionContext,
	metadata *protocol.OperationRestMetadata) (*Request, error) {
	result, err := SerializeInput(inputValue, metadata)
	if err != nil {
		log.Error(err)
		return result, err
	}
	var secCtx core.SecurityContext
	if ctx != nil {
		secCtx = ctx.SecurityContext()
	}

	authzHeaders, err := getAuthorizationHeaders(secCtx)
	if err != nil {
		return nil, err
	}
	result.SetInputHeaders(authzHeaders)

	return result, nil
}

// SerializeInput serializes the input value
// Return Request with urlPath, inputHeaders and requestBody
func SerializeInput(inputValue *data.StructValue, metadata *protocol.OperationRestMetadata) (*Request, error) {
	if metadata == nil {
		return nil, l10n.NewRuntimeErrorNoParam("vapi.data.serializers.rest.metadata.value.nil")
	}

	var requestBodyInputValue data.DataValue
	var err error
	requestBodyParam := metadata.BodyParamActualName()
	if requestBodyParam != "" {
		requestBodyInputValue, err = inputValue.Field(requestBodyParam)
		if err != nil {
			return nil, err
		}
	} else {
		requestBodyInputValue = data.NewStructValue(inputValue.Name(), map[string]data.DataValue{})
	}

	pathVariableFields := map[string]string{}
	queryParamFields := map[string]string{}
	headerFields := map[string]string{}

	pathVariableFieldNames := metadata.PathVariableFieldNames()
	queryParamFieldNames := metadata.QueryParamFieldNames()
	headerFieldNames := metadata.HeaderFieldNames()
	for fieldName, fieldValue := range inputValue.Fields() {
		if contains(pathVariableFieldNames, fieldName) {
			pathVariableFields, err = replaceFieldValue(pathVariableFields, fieldName, fieldValue)
		} else if contains(queryParamFieldNames, fieldName) {
			queryParamFields, err = replaceFieldValue(queryParamFields, fieldName, fieldValue)
		} else if contains(headerFieldNames, fieldName) {
			headerFields, err = replaceFieldValue(headerFields, fieldName, fieldValue)
		} else if requestBodyParam == "" {
			requestBodyInputValue.(*data.StructValue).SetField(fieldName, fieldValue)
		}
		if err != nil {
			return nil, err
		}
	}

	dataValueToJSONEncoder := cleanjson.NewDataValueToJsonEncoder()
	requestBodyStr, err := dataValueToJSONEncoder.Encode(requestBodyInputValue)
	if err != nil {
		return nil, err
	}

	urlPath := metadata.GetUrlPath(pathVariableFields, queryParamFields)
	return NewRequest(urlPath, headerFields, requestBodyStr), nil
}

func replaceFieldValue(fields map[string]string, fieldName string, fieldValue data.DataValue) (map[string]string, error) {
	fieldStr, err := convertDataValueToString(fieldValue)
	if err != nil {
		return nil, err
	}
	if fieldStr != "" {
		fields[fieldName] = fieldStr
	}
	return fields, nil
}

// Serialize data value to string for use in Request URL
func convertDataValueToString(dataValue data.DataValue) (string, error) {
	dataValueToJSONEncoder := cleanjson.NewDataValueToJsonEncoder()
	switch reflect.TypeOf(dataValue) {
	case data.OptionalValuePtr:
		optionalVal := dataValue.(*data.OptionalValue)
		if optionalVal.IsSet() {
			return convertDataValueToString(optionalVal.Value())
		}
		return "", nil
	case data.BoolValuePtr:
		return dataValueToJSONEncoder.Encode(dataValue)
	case data.StringValuePtr:
		encodedStr, err := dataValueToJSONEncoder.Encode(dataValue)
		if err != nil {
			log.Error(err)
			return encodedStr, err
		}
		encodedStr, err = strconv.Unquote(encodedStr)
		if err != nil {
			log.Error(err)
			return encodedStr, err
		}
		return encodedStr, nil
	case data.IntegerValuePtr:
		return dataValueToJSONEncoder.Encode(dataValue)
	default:
		return "", l10n.NewRuntimeError(
			"vapi.data.serializers.rest.unsupported_data_value",
			map[string]string{"type": dataValue.Type().String()})
	}
}

// Get the authorization headers for the corresponding security context
func getAuthorizationHeaders(securityContext core.SecurityContext) (map[string]string, error) {
	if securityContext == nil {
		log.Info("securityContext is nil")
		return map[string]string{}, nil
	}

	switch securityContext.Property(security.AUTHENTICATION_SCHEME_ID) {
	case security.USER_PASSWORD_SCHEME_ID:
		return getUsernameCtxHeaders(securityContext)
	case security.SESSION_SCHEME_ID:
		return getSessionCtxHeaders(securityContext)
	case security.OAUTH_SCHEME_ID:
		return getOauthCtxHeaders(securityContext)
	default:
		return nil, fmt.Errorf("Invalid authentication scheme ID, expected one of [%s, %s, %s], actual value %s",
			security.USER_PASSWORD_SCHEME_ID, security.SESSION_SCHEME_ID, security.OAUTH_SCHEME_ID,
			securityContext.Property(security.AUTHENTICATION_SCHEME_ID))
	}
}

// Get the authorization headers for username security context
func getUsernameCtxHeaders(securityContext core.SecurityContext) (map[string]string, error) {
	username, err := getSecurityCtxStrValue(securityContext, security.USER_KEY)
	if err != nil {
		return nil, err
	}

	password, err := getSecurityCtxStrValue(securityContext, security.PASSWORD_KEY)
	if err != nil {
		return nil, err
	}

	credentialString := fmt.Sprintf("%s:%s", username, password)
	base64EncodedVal := base64.StdEncoding.EncodeToString([]byte(credentialString))
	return map[string]string{"Authorization": "Basic " + base64EncodedVal}, nil
}

// Get the authorization headers for session security context
func getSessionCtxHeaders(securityContext core.SecurityContext) (map[string]string, error) {
	sessionID, err := getSecurityCtxStrValue(securityContext, security.SESSION_ID)
	if err != nil {
		return nil, err
	}
	return map[string]string{security.SESSION_ID_KEY: sessionID}, nil
}

//  Get the authorization headers for oauth security context.
func getOauthCtxHeaders(securityContext core.SecurityContext) (map[string]string, error) {
	oauthToken, err := getSecurityCtxStrValue(securityContext, security.ACCESS_TOKEN)
	if err != nil {
		return nil, err
	}
	return map[string]string{security.CSP_AUTH_TOKEN_KEY: oauthToken}, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getSecurityCtxStrValue(securityContext core.SecurityContext, propKey string) (string, error) {
	securityContextMap := securityContext.GetAllProperties()
	var propVal interface{}
	if propValue, ok := securityContextMap[propKey]; !ok {
		err := fmt.Errorf("%s is not present in the security context", propKey)
		log.Error(err)
		return "", err
	} else {
		propVal = propValue
	}

	if propVal == nil {
		log.Debugf("%s value is nil, use \"\" instead", propKey)
		return "", nil
	}

	if propValueStr, ok := propVal.(string); ok {
		return propValueStr, nil
	}

	err := fmt.Errorf("Invalid type for %s, expected type string, actual type %s",
		propKey, reflect.TypeOf(propVal).String())
	log.Error(err)
	return "", err
}

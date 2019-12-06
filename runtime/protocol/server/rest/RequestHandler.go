/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"github.com/gorilla/mux"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib/rest"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"net/http"
	"strings"
	"unicode"
)

type RequestHandler struct {
	nextApiProvider core.APIProvider
	opMetadata      protocol.OperationMetadata
}

func NewRequestHandler(opMetadata protocol.OperationMetadata, nextApiProvider core.APIProvider) *RequestHandler {
	rh := RequestHandler{nextApiProvider: nextApiProvider, opMetadata: opMetadata}
	return &rh
}

func (rh *RequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	msg := "Calling " + rh.opMetadata.MethodDefinition().Identifier().FullyQualifiedName() + mux.Vars(r)["id"]
	log.Debugf(msg)

	var err []error
	opRestMetaData := rh.opMetadata.RestMetadata()
	inputDataVal := data.NewStructValue("operation-input", nil)

	// Step 1: parse uri parameters, query parameters, headers and request Body to build inputDataVal.
	err = ParsePathParams(r, opRestMetaData, inputDataVal)
	if err != nil {
		log.Errorf("Error while processing path params: %s", err)
		returnBadRequest(rw, err)
		return
	}

	err = ParseHeaderParams(r, opRestMetaData, inputDataVal)
	if err != nil {
		log.Errorf("Error while processing header params: %s", err)
		returnBadRequest(rw, err)
		return
	}

	err = ParseQueryParams(r, opRestMetaData, inputDataVal)
	if err != nil {
		log.Errorf("Error while processing query params: %s", err)
		returnBadRequest(rw, err)
		return
	}

	// @BodyField and @Body are mutually exclusive
	bodyParamName := opRestMetaData.BodyParamActualName()
	if bodyParamName == "" {
		err := ParseBodyFieldParams(r, opRestMetaData, inputDataVal)
		if err != nil {
			log.Errorf("Error: %s", err)
			returnBadRequest(rw, err)
			return
		}
	} else {
		err := ParseBodyParams(r, bodyParamName, opRestMetaData, inputDataVal)
		if err != nil {
			log.Errorf("Error while processing body params: %s", err)
			returnBadRequest(rw, err)
			return
		}
	}

	// Step 2: Build application and security context
	var ctx *core.ExecutionContext = core.NewExecutionContext(nil, nil)

	// Step 3: invoke apiprovider (local provider)
	methodId := rh.opMetadata.MethodDefinition().Identifier()
	methodResult := rh.nextApiProvider.Invoke(
		methodId.InterfaceIdentifier().Name(),
		methodId.Name(),
		inputDataVal,
		ctx)

	// Step 4: process output
	if methodResult.IsSuccess() {
		response := WriteToHeader(rw, opRestMetaData, methodResult.Output())
		responseBody, err := setResponseBody(response)
		if err != nil {
			log.Errorf("Error while setting output response body: %s", err)
		}
		rw.Write([]byte(responseBody))
	} else {
		errorNameToStatusMap := opRestMetaData.ErrorCodeMap()
		if len(errorNameToStatusMap) > 0 {
			// TODO: getErrorStructureName should not be used, its a temporary measure.
			errorName := getErrorStructureName(methodResult.Error().Name())
			if httpErrorCode, ok := errorNameToStatusMap[errorName]; ok {
				rw.WriteHeader(httpErrorCode)
			} else {
				rw.WriteHeader(rest.HTTP_500_INTERNAL_SERVER_ERROR)
			}
		} else {
			errorName := methodResult.Error().Name()
			if errCode, ok := rest.VAPI_TO_HTTP_ERROR_MAP[errorName]; ok {
				rw.WriteHeader(errCode)
			} else {
				rw.WriteHeader(rest.HTTP_500_INTERNAL_SERVER_ERROR)
			}
		}
		responseBody, err := setResponseBody(methodResult.Error())
		if err != nil {
			log.Errorf("Error while setting error response body: %s", err)
		}
		rw.Write([]byte(responseBody))
	}
}

func returnBadRequest(rw http.ResponseWriter, dataErr []error) {
	errorValue := bindings.CreateErrorValueFromMessages(bindings.INVALID_ARGUMENT_ERROR_DEF, dataErr)
	responseBody, err := setResponseBody(errorValue)
	if err != nil {
		log.Errorf("Error while setting error response body: %s", err)
	}
	rw.WriteHeader(rest.HTTP_400_BAD_REQUEST)
	rw.Write([]byte(responseBody))
}

func CreateResponseAndWriteToHeader(rw http.ResponseWriter, opRestMetaData protocol.OperationRestMetadata, methodOutput data.DataValue) data.DataValue {
	responseHeaderNameMap := opRestMetaData.ResultHeadersNameMap()
	resStruct, ok := methodOutput.(*data.StructValue)
	if !ok || (len(responseHeaderNameMap) == 0) {
		return methodOutput
	}
	structName := resStruct.Name()
	structFields := resStruct.Fields()

	output := data.NewStructValue(structName, nil)

	for field, value := range structFields {
		if headerfield, ok := responseHeaderNameMap[field]; ok {
			rw.Header().Set(headerfield, dataValueToString(value))
		} else {
			output.SetField(field, value)
		}
	}
	return output
}

func WriteToHeader(rw http.ResponseWriter, opRestMetaData protocol.OperationRestMetadata, methodOutput data.DataValue) data.DataValue {

	output := CreateResponseAndWriteToHeader(rw, opRestMetaData, methodOutput)

	// Make sure all the header content are written before WriteHeader(int) is called
	var statusCode int = opRestMetaData.SuccessCode()
	for _, code := range rest.SUCCESSFUL_STATUS_CODES {
		if statusCode == code {
			rw.WriteHeader(statusCode)
			return output
		}
	}

	rw.WriteHeader(rest.HTTP_200_OK)
	return output
}

func dataValueToString(value data.DataValue) string {
	if headerValueString, ok := value.(*data.StringValue); ok {
		return headerValueString.Value()
	} else if headerValueBLOB, ok := value.(*data.BlobValue); ok {
		return string(headerValueBLOB.Value())
	} else if headerValueSecret, ok := value.(*data.SecretValue); ok {
		return headerValueSecret.Value()
	} else if headerValueInteger, ok := value.(*data.IntegerValue); ok {
		return headerValueInteger.String()
	} else if headerValueOptional, ok := value.(*data.OptionalValue); ok {
		if headerValueOptional.IsSet() {
			return dataValueToString(headerValueOptional.Value())
		}
		return ""
	} else {
		log.Errorf("Header Value of type %s DataValue is not supported", value.Type())
	}
	return ""
}

func getErrorStructureName(name string) string {
	temp := strings.Split(name, ".")
	res := temp[len(temp)-1]
	return ConvertToCamelCase(res)
}

func ConvertToCamelCase(str string) string {
	res := ""
	upper := true
	for _, r := range str {
		if unicode.IsLetter(r) {
			if upper {
				res += string(unicode.ToUpper(r))
			} else {
				res += string(r)
			}
			upper = false
		} else {
			upper = true
		}
	}
	return res
}

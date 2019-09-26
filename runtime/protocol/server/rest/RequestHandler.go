package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"log"
	"net/http"
)

type RequestHandler struct {
	nextApiProvider   core.APIProvider
	opMetadata        protocol.OperationMetadata
}

func NewRequestHandler(opMetadata protocol.OperationMetadata, nextApiProvider core.APIProvider) *RequestHandler {
	rh := RequestHandler{nextApiProvider: nextApiProvider, opMetadata:opMetadata}
	return &rh
}

func (rh *RequestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	msg := "Calling " + rh.opMetadata.MethodDefinition().Identifier().FullyQualifiedName() + mux.Vars(r)["id"]
	log.Printf(msg)
	json.NewEncoder(rw).Encode(msg)
	return

	// Step 1: parse uri parameters, query parameters, headers, request body  to build input Dataval.
	var inputDataVal data.DataValue = nil

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
	//rest.ResponseWriter(methodResult, rh.opMetadata, rw)
	//ResponseWriter(result core.MethodResult, metadata protocol.OperationMetadata, rw http.ResponseWriter)
	if methodResult.IsSuccess() {
		//rw.WriteHeader(o.successCode)
		//Add other response headers if there are any..
		//cleanJsonData := dataValConverter.ConvertDataValueToCleanJson(methodResult.Output())
		//rw.Write([]byte(cleanJsonData))
	} else {
		//errorName := methodResult.Error().Name()
		//httpErrorCode := o.errorNameToStatusMap[errorName]
		//rw.WriteHeader(httpErrorCode)
		//cleanJsonData := dataValConverter.ConvertDataValueToCleanJson(methodResult.Error())
		//rw.Write([]byte(cleanJsonData))
	}
}
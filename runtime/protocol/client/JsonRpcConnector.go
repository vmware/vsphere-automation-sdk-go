package client

import (
	"bytes"
	"encoding/json"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/common"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type JsonRpcConnector struct {
	url                string
	httpClient         http.Client
	securityContext    core.SecurityContext
	appContext         *core.ApplicationContext
	typeConverter      *bindings.TypeConverter
	connectionMetadata map[string]interface{}
}

func NewJsonRpcConnector(url string, client http.Client) *JsonRpcConnector {
	typeConverter := bindings.NewTypeConverter()
	typeConverter.SetMode(bindings.JSONRPC)
	return &JsonRpcConnector{url: url, httpClient: client, typeConverter: typeConverter}
}

func (j *JsonRpcConnector) ApplicationContext() *core.ApplicationContext {
	return j.appContext
}

func (j *JsonRpcConnector) SetApplicationContext(ctx *core.ApplicationContext) {
	j.appContext = ctx
}

func (j *JsonRpcConnector) SecurityContext() core.SecurityContext {
	return j.securityContext
}

func (j *JsonRpcConnector) SetSecurityContext(ctx core.SecurityContext) {
	j.securityContext = ctx
}

func (j *JsonRpcConnector) NewExecutionContext() *core.ExecutionContext {
	if j.appContext == nil {
		j.appContext = common.NewDefaultApplicationContext()
	} else {
		common.InsertOperationId(j.appContext)
	}
	return core.NewExecutionContext(j.appContext, j.securityContext)
}

func (j *JsonRpcConnector) GetApiProvider() core.APIProvider {
	return j
}

func (j *JsonRpcConnector) TypeConverter() *bindings.TypeConverter {
	return j.typeConverter
}

func (j *JsonRpcConnector) SetConnectionMetadata(connectionMetadata map[string]interface{}) {
	j.connectionMetadata = connectionMetadata
}

func (j *JsonRpcConnector) ConnectionMetadata() map[string]interface{} {
	return j.connectionMetadata
}

func (j *JsonRpcConnector) Invoke(serviceId string, operationId string, inputValue data.DataValue, ctx *core.ExecutionContext) core.MethodResult {
	if ctx == nil {
		ctx = j.NewExecutionContext()
	}
	if !ctx.ApplicationContext().HasProperty(lib.OPID) {
		common.InsertOperationId(ctx.ApplicationContext())
	}
	opId := ctx.ApplicationContext().GetProperty(lib.OPID)

	var params = make(map[string]interface{})
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	var encodedInput, encodingError = jsonRpcEncoder.Encode(inputValue)
	if encodingError != nil {
		log.Error("Error encoding input")
		log.Error(encodingError)
	}
	encodedInputMap := make(map[string]interface{})
	var errx = json.Unmarshal(encodedInput, &encodedInputMap)
	if errx != nil {
		log.Error("Error unmarshalling inputValue")
		log.Error(errx)
	}
	log.Debugf("Invoking with input %+v", string(encodedInput))
	params[lib.REQUEST_INPUT] = encodedInputMap
	params[lib.EXECUTION_CONTEXT] = msg.NewExecutionContextSerializer(ctx)
	params[lib.REQUEST_OPERATION_ID] = operationId
	params[lib.REQUEST_SERVICE_ID] = serviceId
	var jsonRpcRequest = msg.NewJsonRpc20Request(lib.JSONRPC_VERSION, lib.JSONRPC_INVOKE, params, opId, false)
	var jsonRpcRequestSerializer = msg.NewJsonRpc20RequestSerializer(jsonRpcRequest)
	var requestBytes, _ = jsonRpcRequestSerializer.MarshalJSON()

	req, err := http.NewRequest("POST", j.url, bytes.NewBuffer(requestBytes))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)
	req.Header.Set(lib.VAPI_SERVICE_HEADER, serviceId)
	req.Header.Set(lib.VAPI_OPERATION_HEADER, operationId)
	j.copyContextsToHeaders(ctx, req)
	response, err := j.httpClient.Do(req)
	if err != nil {
		if strings.HasSuffix(err.Error(), "connection refused") {
			err := l10n.NewRuntimeErrorNoParam("vapi.server.unavailable")
			err_val := bindings.CreateErrorValueFromMessages(bindings.SERVICE_UNAVAILABLE_ERROR_DEF, []error{err})
			return core.NewMethodResult(nil, err_val)
		} else if strings.HasSuffix(err.Error(), "i/o timeout") {
			err := l10n.NewRuntimeErrorNoParam("vapi.server.timedout")
			err_val := bindings.CreateErrorValueFromMessages(bindings.TIMEDOUT_ERROR_DEF, []error{err})
			return core.NewMethodResult(nil, err_val)
		}
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	jsonRpcDecoder := msg.NewJsonRpcDecoder()
	jsonRpcResponse, error := jsonRpcDecoder.DeSerializeResponse(resp)
	if error != nil {
		log.Error(error)
	}
	if jsonRpcResponse.Result() != nil {
		methodResult, err := jsonRpcDecoder.DeSerializeMethodResult(jsonRpcResponse.Result().(map[string]interface{}))
		if err != nil {
			log.Error(err)
		}
		return methodResult
	}
	return core.MethodResult{}
}

func (j *JsonRpcConnector) copyContextsToHeaders(ctx *core.ExecutionContext, req *http.Request) {
	appCtx := ctx.ApplicationContext()
	secCtx := ctx.SecurityContext()

	if appCtx != nil {
		for key, value := range appCtx.GetAllProperties() {
			keyLowerCase := strings.ToLower(key)
			switch keyLowerCase {
			case lib.HTTP_USER_AGENT_HEADER:
				req.Header.Set(lib.HTTP_USER_AGENT_HEADER, *value)
			case lib.HTTP_ACCEPT_LANGUAGE:
				req.Header.Set(lib.HTTP_ACCEPT_LANGUAGE, *value)
			default:
				req.Header.Set(lib.VAPI_HEADER_PREFIX+keyLowerCase, *value)
			}
		}
	}

	if secCtx != nil {
		if secCtx.Property(security.AUTHENTICATION_SCHEME_ID) == security.SESSION_SCHEME_ID {
			if sessionId, ok := secCtx.Property(security.SESSION_ID).(string); ok {
				req.Header.Set(lib.VAPI_SESSION_HEADER, sessionId)
			} else {
				log.Errorf("Invalid session ID in security context. Skipping setting request header. Expected string but was %s",
					reflect.TypeOf(secCtx.Property(security.SESSION_ID)))
			}
		}
	}
}

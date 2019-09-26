package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/test"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/local"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func getInitializedLocalProvider() *local.LocalProvider {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)

	localProvider.AddInterface(nil)
	return localProvider
}

var noOperationId = []byte(`{  
   "id":"2",
   "jsonrpc":"2.0",
   "method":"invoke",
   "params":{  
      "ctx":{
      	
      },
      "input":{  
         "STRUCTURE":{}  
      },
      "serviceId":"sample.first.toy_VM"
   }
}`)

var noServiceId = []byte(`{  
   "id":"2",
   "jsonrpc":"2.0",
   "method":"invoke",
   "params":{  
      "ctx":{
      	
      },
      "input":{  
         "STRUCTURE":{}  
      },
      "operationId":"create"
   }
}`)

var noExecutionContext = []byte(`{  
   "id":"2",
   "jsonrpc":"2.0",
   "method":"invoke",
   "params":{
      "input":{  
         "STRUCTURE":{}  
      },
      "operationId":"create",
      "serviceId":"sample.first.toy_VM"
   }
}`)

var noInput = []byte(`{  
   "id":"2",
   "jsonrpc":"2.0",
   "method":"invoke",
   "params":{
      "ctx":{
          "securityCtx":{
 			 "sessionId": "xyz"
		   },
          "appCtx":{
          }
      },
      "operationId":"create",
      "serviceId":"sample.first.toy_VM"
   }
}`)

func TestNewJsonRpcErrorNoServiceId(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, jsonRpcHandler)
	server.Start()

	req, err := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noServiceId))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var jsonError = data[lib.METHOD_RESULT_ERROR]
	assert.Equal(t, "params does not have key: serviceId", jsonError.(map[string]interface{})["data"])
}

func TestNewJsonRpcErrorNoOperationId(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, jsonRpcHandler)
	server.Start()

	req, err := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noOperationId))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var jsonError = data[lib.METHOD_RESULT_ERROR]
	assert.Equal(t, "params does not have key: operationId", jsonError.(map[string]interface{})["data"])
}

func TestNewJsonRpcErrorNoExecutionCtx(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, jsonRpcHandler)
	server.Start()

	req, err := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noExecutionContext))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var jsonError = data[lib.METHOD_RESULT_ERROR]
	assert.Equal(t, "params does not have key: ctx", jsonError.(map[string]interface{})["data"])
}

func TestNewJsonRpcErrorNoExecutionInput(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, jsonRpcHandler)
	server.Start()

	req, err := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noInput))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var jsonError = data[lib.METHOD_RESULT_ERROR]
	assert.Equal(t, "params does not have key: input", jsonError.(map[string]interface{})["data"])
	server.Stop()
}

func TestJsonRpc11MismatchOperationIdentifier(t *testing.T) {

	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, jsonRpcHandler)
	server.Start()
	req, err := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noInput))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)
	req.Header["vapi-service"] = []string{"sample.first.toy_VM"}
	req.Header["vapi-operation"] = []string{"create1"}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var responseObject = make(map[string]interface{})
	responseString := string(body)
	d := json.NewDecoder(strings.NewReader(responseString))
	d.UseNumber()
	d.Decode(&responseObject)
	errorData := responseObject[lib.METHOD_RESULT_ERROR].(map[string]interface{})
	errorCode, _ := errorData[lib.ERROR_CODE].(json.Number).Int64()
	assert.Equal(t, errorCode, int64(-31001))
	assert.Equal(t, errorData[lib.ERROR_MESSAGE].(string),
		"Mismatching operation identifier in HTTP header and payload")
	server.Stop()
}

func TestJsonRpc11SecurityContextOverride(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	req, _ := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noInput))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)
	req.Header.Set(lib.VAPI_SESSION_HEADER, "abc")
	securityContext := security.NewSessionSecurityContext("xyz")
	jsonRpcHandler.copyHeadersToContexts(core.NewExecutionContext(nil, securityContext), req)
	assert.Equal(t, securityContext.Property("sessionId").(string), "abc")
}

func TestJsonRpc11ApplicationContextOverride(t *testing.T) {
	var jsonRpcHandler = NewJsonRpcHandler(getInitializedLocalProvider())
	req, _ := http.NewRequest("POST", protocol_test.TEST_SERVER_URL, bytes.NewBuffer(noInput))
	req.Header.Set(lib.HTTP_CONTENT_TYPE_HEADER, lib.JSON_CONTENT_TYPE)
	req.Header.Set(lib.HTTP_ACCEPT_LANGUAGE, "da, en-gb;q=0.8, en;q=0.7")
	req.Header["vapi-ctx-vcenter-route"] = []string{"header-value"}
	req.Header["vapi-ctx-opid"] = []string{"abc-def"}
	req.Header["vapi-ctx-format-locale"] = []string{"en"}
	req.Header["vapi-ctx-timezone"] = []string{"Europe/Berlin"}
	req.Header.Set(lib.HTTP_USER_AGENT_HEADER, "user-agent-header-value")
	req.Header.Set("vapi-ctx-$useragent", "vapi-ctx-value")
	appContextMap := map[string]*string{}
	var bodyValue = "body-value"
	appContextMap["vcenter-route"] = &bodyValue
	appContext := core.NewApplicationContext(appContextMap)
	jsonRpcHandler.copyHeadersToContexts(core.NewExecutionContext(appContext, nil), req)
	assert.Equal(t, *appContext.GetProperty("vcenter-route"), "header-value")
	assert.Equal(t, *appContext.GetProperty(lib.OPID), "abc-def")
	assert.Equal(t, *appContext.GetProperty(lib.HTTP_ACCEPT_LANGUAGE), "da, en-gb;q=0.8, en;q=0.7")
	assert.Equal(t, *appContext.GetProperty(lib.VAPI_L10N_FORMAT_LOCALE), "en")
	assert.Equal(t, *appContext.GetProperty(lib.VAPI_L10N_TIMEZONE), "Europe/Berlin")
	assert.Equal(t, "user-agent-header-value", *appContext.GetProperty("$userAgent"))
}

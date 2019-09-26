package protocol_test

import (
	"fmt"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"testing"
)

//func TestProtocolHandler(t *testing.T) {
//	//Server with jsonrpc handler
//	var server = rpc.NewServer()
//	server.Start()
//	time.Sleep(100 * time.Second)
//	resp, err := http.Get(serverURL)
//	if err != nil {
//		t.Error(err)
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode == http.StatusOK {
//		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
//		if err2 != nil {
//			// handle error
//		}
//		bodyString := string(bodyBytes)
//		assert.Equal(t, "Hello Json RPC", bodyString)
//
//	}
//	server.Stop()
//}

func TestJsonRpcRequest(t *testing.T) {

	//var server = rpc.NewServer()
	//server.Start()
	//var requestBody = buildVapiJsonRpcRequestBody("1", "myinterface", "hello", nil, nil)
	//data, err := json.Marshal(requestBody)
	//fmt.Println(string(data))
	//if err != nil {
	//	t.Error(err)
	//}
	//resp, err := http.Post(serverURL,
	//	"application/json", strings.NewReader(string(data)))
	//if err != nil {
	//	t.Error(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	t.Error(err)
	//}
	//result := make(map[string]interface{})
	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	t.Error(err)
	//}
	//log.Println(result)
}

//func TestPetStoreApiInterface(t *testing.T) {
//
//	var localProvider, _ = local.NewLocalProvider("localprovider", true)
//	localProvider.AddInterface(provider.NewPetStoreApiInterface())
//	var jsonRpcHandler = msg.NewJsonRpcHandler(localProvider)
//	var server = rpc.NewServer(":8001", jsonRpcHandler)
//	server.Start()
//	time.Sleep(10000 * time.Second)
//}

func TestGetStructValueForUser(t *testing.T) {

	var structValue = data.NewStructValue("User", nil)
	structValue.SetField("id", data.NewIntegerValue(1))
	structValue.SetField("userName", data.NewStringValue("sreenidhi"))
	structValue.SetField("firstName", data.NewStringValue("Sreenidhi"))
	structValue.SetField("lastName", data.NewStringValue("Sreesha"))
	structValue.SetField("email", data.NewStringValue("sreeshas@vmware.com"))
	structValue.SetField("password", data.NewStringValue("password"))
	structValue.SetField("phone", data.NewStringValue("000"))
	structValue.SetField("userStatus", data.NewIntegerValue(1))
	var jsonEncoder = msg.NewJsonRpcEncoder()
	var result, _ = jsonEncoder.Encode(structValue)
	fmt.Println(string(result))
}

//func TestToyVm(t *testing.T) {
//	var localProvider, _ = local.NewLocalProvider("localprovider", true)
//	localProvider.AddInterface(toyVM.NewToyVMApiInterface(impls.NewToyVMServiceImpl()))
//	var jsonRpcHandler = msg.NewJsonRpcHandler(localProvider)
//	var server = rpc.NewServer(":8001", jsonRpcHandler)
//	server.Start()
//	time.Sleep(10000 * time.Second)
//}

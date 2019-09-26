package introspection_test
//
//import (
//	"crypto/tls"
//	"github.com/stretchr/testify/assert"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
//	//"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection/test/operation"
//	//"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection/test/service"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/local"
//	"net/http"
//	"sort"
//	"testing"
//)
//
//
//func getClientConnector(url string, httpClient http.Client) client.Connector{
//	connector := client.NewJsonRpcConnector(url, httpClient)
//	return connector
//}
//
//func startServer() {
//	var localProvider, _ = local.NewLocalProvider("localprovider", true)
//
//	var jsonRpcHandler = msg.NewJsonRpcHandler(localProvider)
//	var server = rpc.NewServer(":8001", jsonRpcHandler)
//	server.Start()
//}
//func TestServiceList(t *testing.T) {
//		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//		startServer()
//		httpClient := http.Client{}
//		connector := getClientConnector("http://localhost:8001", httpClient)
//		svc_client := service.NewServiceClientImpl(connector)
//
//		svc_set, err := svc_client.List()
//		assert.Nil(t,err)
//		result := []string{}
//		for key, _ := range svc_set {
//			result = append(result, key)
//		}
//		sort.Strings(result)
//		expected := []string{"com.vmware.vapi.std.introspection.service", "com.vmware.vapi.std.introspection.operation", "com.vmware.vapi.std.introspection.provider"}
//		sort.Strings(expected)
//		assert.Equal(t, expected, result)
//}
//
//func TestServiceGet(t *testing.T) {
//	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	startServer()
//	httpClient := http.Client{}
//	connector := getClientConnector("http://localhost:8001", httpClient)
//	svc_client := service.NewServiceClientImpl(connector)
//	info, err := svc_client.Get("com.vmware.vapi.std.introspection.service")
//	assert.Nil(t, err)
//	assert.True(t, info.Operations["list"])
//	assert.True(t, info.Operations["get"])
//
//}
//
////func TestServiceNotFound(t *testing.T) {
////	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
////	startServer()
////	httpClient := http.Client{}
////	connector := getClientConnector("http://localhost:8001", httpClient)
////	svc_client := service.NewServiceClientImpl(connector)
////	info, err := svc_client.Get("bogus")
////	assert.NotNil(t, err)
////	assert.Nil(t, info)
////}
//
//func TestOperationList(t *testing.T) {
//	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	startServer()
//	httpClient := http.Client{}
//	connector := getClientConnector("http://localhost:8001", httpClient)
//	op_client := operation.NewOperationClientImpl(connector)
//	list, err := op_client.List("com.vmware.vapi.std.introspection.service")
//	assert.Nil(t, err)
//	assert.True(t, list["list"])
//	assert.True(t, list["get"])
//}
//
////func TestOperationListNotFound(t *testing.T) {
////	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
////	startServer()
////	httpClient := http.Client{}
////	connector := getClientConnector("http://localhost:8001", httpClient)
////	op_client := operation.NewOperationClientImpl(connector)
////	list, err := op_client.List("bogus")
////	fmt.Println(err)
////	assert.NotNil(t, err)
////	assert.Nil(t, list)
////}
//
//func TestOperationGet(t *testing.T) {
//	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	startServer()
//	httpClient := http.Client{}
//	connector := getClientConnector("http://localhost:8001", httpClient)
//	op_client := operation.NewOperationClientImpl(connector)
//	opInfo, err := op_client.Get("com.vmware.vapi.std.introspection.service", "get")
//	assert.Nil(t, err)
//	assert.Equal(t, "operation-input", *opInfo.InputDefinition.Name)
//	assert.Equal(t, "com.vmware.vapi.std.errors.internal_server_error", *opInfo.ErrorDefinitions[0].Name)
//	assert.Equal(t, "com.vmware.vapi.std.introspection.service.info", *opInfo.OutputDefinition.Name)
//
//
//}
//
//
//

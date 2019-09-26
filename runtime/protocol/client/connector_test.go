package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/test"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"

	"net/http"
	"reflect"
	"testing"
)

type MockHandler struct {
	RequestArgument *http.Request
}

func (m *MockHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.RequestArgument = req
}

func TestRequestHeaders(t *testing.T) {
	var mockHandler = MockHandler{}
	var server = rpc.NewServer(protocol_test.TEST_SERVER_PORT, &mockHandler)
	server.Start()
	connector := NewJsonRpcConnector(protocol_test.TEST_SERVER_URL, http.Client{})
	serviceId := "myservice"
	operationId := "myoperation"
	opId := "123abc"
	sessionId := "mysessionID"
	al := "de"
	fl := "en"
	tz := "Europe/Berlin"
	appContextMap := map[string]*string{
		lib.HTTP_ACCEPT_LANGUAGE:    &al,
		lib.VAPI_L10N_FORMAT_LOCALE: &fl,
		lib.VAPI_L10N_TIMEZONE:      &tz,
		lib.OPID:                    &opId,
	}

	ctx := core.NewExecutionContext(core.NewApplicationContext(appContextMap),
		security.NewSessionSecurityContext(sessionId))

	connector.Invoke(serviceId, operationId, getTestInputDataValue(), ctx)

	req := mockHandler.RequestArgument

	assert.Equal(t, lib.JSON_CONTENT_TYPE, req.Header[lib.HTTP_CONTENT_TYPE_HEADER][0])
	fmt.Println(req.Header)
	assert.Equal(t, serviceId, req.Header["Vapi-Service"][0])
	assert.Equal(t, operationId, req.Header["Vapi-Operation"][0])
	assert.Equal(t, opId, req.Header["Vapi-Ctx-Opid"][0])
	assert.Equal(t, sessionId, req.Header["Vmware-Api-Session-Id"][0])
	assert.Equal(t, al, req.Header["Accept-Language"][0])
	assert.Equal(t, fl, req.Header["Vapi-Ctx-Format-Locale"][0])
	assert.Equal(t, tz, req.Header["Vapi-Ctx-Timezone"][0])
}

func getTestInputDataValue() data.DataValue {
	typeConverter := bindings.NewTypeConverter()
	fields := map[string]bindings.BindingType{
		"vm_id": bindings.NewStringType(),
	}
	fieldNameMap := map[string]string{
		"vm_id": "VmId",
	}
	inputBindingType := bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, nil)
	sv := bindings.NewStructValueBuilder(inputBindingType, typeConverter)
	sv.AddStructField("VmId", "vm-123")
	inputDataValue, _ := sv.GetStructValue()
	return inputDataValue
}

package local_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntrospectorGetServices(t *testing.T) {
	localProviderIntrospector, _ := introspection.NewLocalProviderIntrospector("test1")
	localProviderIntrospector.AddService("api1", NewMockApiInterface())
	ids := localProviderIntrospector.GetServices()
	assert.Equal(t, len(ids), 1)
	assert.Equal(t, ids[0], "api1")
}

func TestIntrospectorGetServiceInfo(t *testing.T) {
	localProviderIntrospector, _ := introspection.NewLocalProviderIntrospector("test1")
	localProviderIntrospector.AddService("api1", NewMockApiInterface())
	serviceInfo := localProviderIntrospector.GetServiceInfo("api1")
	assert.Nil(t, serviceInfo.Error())
	var serviceInfoSV = serviceInfo.Output().(*data.StructValue)
	assert.Equal(t, serviceInfoSV.Name(), "com.vmware.vapi.std.introspection.service.info")
	assert.Equal(t, serviceInfoSV.HasField("operations"), true)
	var listValue, _ = serviceInfoSV.Field("operations")
	assert.Equal(t, len(listValue.(*data.ListValue).List()), 8)
	serviceInfo2 := localProviderIntrospector.GetServiceInfo("api2")
	assert.Equal(t, serviceInfo2.Error().Name(), "com.vmware.vapi.std.errors.not_found")
}

func TestIntrospectorGetOperations(t *testing.T) {
	localProviderIntrospector, _ := introspection.NewLocalProviderIntrospector("test1")
	apiInterface := NewMockApiInterface()
	localProviderIntrospector.AddService("api1", apiInterface)
	serviceInfo := localProviderIntrospector.GetOperations("api1")
	assert.Nil(t, serviceInfo.Error())
	serviceInfo2 := localProviderIntrospector.GetOperations("api2")
	assert.Equal(t, serviceInfo2.Error().Name(), "com.vmware.vapi.std.errors.not_found")
}

func TestIntrospectorGetOperationInfo(t *testing.T) {
	localProviderIntrospector, _ := introspection.NewLocalProviderIntrospector("test1")
	methodResult := localProviderIntrospector.GetOperationInfo("serviceId0", "operationId0")
	assert.Equal(t, methodResult.Error().Name(), "com.vmware.vapi.std.errors.not_found")
	apiInterface := NewMockApiInterface()
	localProviderIntrospector.AddService("serviceId1", apiInterface)
	result := localProviderIntrospector.GetOperationInfo("serviceId1", "increment")
	assert.Nil(t, result.Error())
}

package introspection
//
//import (
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection/bindings/service"
//)
//
//type ServiceImpl struct {
//	introspector  *LocalProviderIntrospector
//}
//
//func NewServiceImpl(introspector *LocalProviderIntrospector) *ServiceImpl{
//    return &ServiceImpl{introspector: introspector}
//}
//
//func(ServiceImpl ServiceImpl) List(ctx *core.ExecutionContext) map[string]bool {
//    return ServiceImpl.introspector.GetServices()
//}
//func(ServiceImpl ServiceImpl) Get(id string, ctx *core.ExecutionContext) (service.Info, error) {
//	serviceInfo, err := ServiceImpl.introspector.GetServiceInfo(id)
//	return serviceInfo, err
//
//}
//

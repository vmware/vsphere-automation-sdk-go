module gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings

go 1.12

require (
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance v1.0.0
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging v1.0.0
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content v1.0.0
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/esx v1.0.0
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter v1.0.0
	gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vstats v1.0.0
)

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance v1.0.0 => ./appliance

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging v1.0.0 => ./cis/tagging

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content v1.0.0 => ./content

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/esx v1.0.0 => ./esx

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter v1.0.0 => ./vcenter

replace gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vstats v1.0.0 => ./vstats

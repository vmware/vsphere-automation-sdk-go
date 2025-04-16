module github.com/vmware/vsphere-automation-sdk-go/services/nsxt

go 1.17

replace github.com/vmware/vsphere-automation-sdk-go/runtime => /builds/vapi-sdk/unified-sdk-generator/vsphere-automation-sdk-go/runtime

replace github.com/vmware/vsphere-automation-sdk-go/lib => /builds/vapi-sdk/unified-sdk-generator/vsphere-automation-sdk-go/lib

require (
	github.com/vmware/vsphere-automation-sdk-go/lib v0.0.0-00010101000000-000000000000
	github.com/vmware/vsphere-automation-sdk-go/runtime v0.7.0
)

require (
	github.com/beevik/etree v1.1.0 // indirect
	github.com/gibson042/canonicaljson-go v1.0.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)

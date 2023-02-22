# vsphere-automation-sdk-go (Beta)

vSphere Automation SDK Beta for GO language.

[![License](https://img.shields.io/static/v1?&label=License&message=BSD-2-Clause&color=red&style=for-the-badge)](LICENSE.txt)

![Supported VMware Products](https://img.shields.io/static/v1?&label=Supported%20VMware%20Products&message=%20VMC%2C%20NSX-T%20&color=blue&style=for-the-badge)

# VMware vSphere Automation SDK for GO Language

## Table of Contents

- [Overview](#overview)
- [Table of Contents](#table-of-contents)
- [Quick Start Guide](#quick-start-guide)
- [API Documentation](#api-documentation)
- [Resource Maintenance](#resource-maintenance)
- [Repository Administrator Resources](#repository-administrator-resources)
  - [Board Members](#board-members)
- [VMware Resources](#vmware-resources)
- [License](#license)

## Overview

This document describes the vSphere Automation GO SDK services(client-bindings), that use the vSphere Automation
GO client library. Currently, vSphere Automation Go SDK supports VMC and NSX-T. Support for vSphere(vCenter) will be added soon.

## Quick Start Guide

This document will walk you through setting up the SDK to start writing your Sample.

### Environment Set Up

#### Pre-Requisites

- golang 1.13
- Recommneded use of [Go Modules](https://blog.golang.org/using-go-modules) for development,
  as it will be very convenient and easy to upgrade to major version releases.

#### Go Environment Variables

- **GO111MODULE** environment variable is recommended to be set as

  > **auto**: if all the Go projects in your development environment are not compatible with Go Modules.

  > **on**: if all your projects are compatible with Go Modules.

- **GOPATH** environment variable is required to be set only if "GO111MODULE" is set to "auto".

### Start Writing your Sample

#### Create Go Module

Start with creating Go Module by executing the command

```shell
go mod init sample
```

Executing this command will create go.mod file in current directory, as shown below

```golang
module sample

go 1.13

```

> Note

- If "GOPATH" is set in your development environment, use directory outside "GOPATH" as workspace for writing your sample.
- If the Go Module, you are developing, is to be stored on any public or private Code repository, use fully qualified URL as the module name like for module vsphere-automation-sdk-go on github.com has module name "github.com/vmware/vsphere-automation-sdk-go".

#### Writing your Sample

You can directly start writing your sample.go file by importing the required Go Modules.
For **VMC** services import

> github.com/vmware/vsphere-automation-sdk-go/services/vmc

For **NSX-T** services import

> github.com/vmware/vsphere-automation-sdk-go/services/nsxt

You can refer to the below code excerpt.

```golang
package main

import (
	"fmt"
	"os"

	vmc "github.com/vmware/vsphere-automation-sdk-go/services/vmc"
.
.
.
```

> Note

- If you are using IDE to develop and want to use auto-completion, navigation and other such features for fast and easy development, just write all your imports with "\_" (underscore) as alias as shown below

  ```golang
  package main

  import (
      _ "github.com/vmware/vsphere-automation-sdk-go/services/vmc"
  )
  ```

  and then execute below command in the directory with the go.mod file

  ```shell
  go test
  ```
  OR
  ```
  go build
  ```

This will fetch all the dependency modules and your IDE code-related features will start working.

## API Documentation

The API documentation can be accessed from here:

[VMC APIs](https://godoc.org/github.com/vmware/vsphere-automation-sdk-go/services/vmc).
[NSX-T APIs](https://godoc.org/github.com/vmware/vsphere-automation-sdk-go/services/nsxt).

## Resource Maintenance

### Filing Issues

Any bugs or other issues should be filed with appropriate tags, within GitHub by way of the repository’s Issue Tracker.

### Resolving Issues

This project contains auto-generated code, therefore code contributions will not be accepted. However, vsphere-automation-sdk-go team welcomes issue reports and we will do our best to ensure that future population of repository will contain a fix.

## Repository Administrator Resources

### Board Members

Board members are volunteers from the SDK community and VMware staff members.

Members:

- Deyan Popov (VMware)

## VMware Resources

- [vSphere Automation SDK Overview](http://pubs.vmware.com/vsphere-65/index.jsp#com.vmware.vapi.progguide.doc/GUID-AF73991C-FC1C-47DF-8362-184B6544CFDE.html)
- [VMware Code](https://code.vmware.com/home)
- [VMware Developer Community](https://communities.vmware.com/community/vmtn/developer)
- [VMware VMC Go API Reference documentation](https://godoc.org/github.com/vmware/vsphere-automation-sdk-go/services/vmc).
  <!---* [VMware NSX-T GOLang API Reference documentation](https://godoc.org/github.com/vmware/vsphere-automation-sdk-go/services/nsxt).--->
  <!---* [VMware vSphere GOLang API Reference documentation](https://godoc.org/github.com/vmware/vsphere-automation-sdk-go/services/vsphere).--->
  <!---* [VMware GOLang forum]()--->

## License

vsphere-automation-sdk-go is available under the [BSD-2 license](LICENSE.txt).

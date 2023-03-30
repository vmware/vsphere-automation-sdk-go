/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package main

import (
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	auth "github.com/vmware/vsphere-automation-sdk-go/runtime/examples/authentication/server/authentication_handlers"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/examples/authentication/server/bindings/src/generated/vmodl/examples/greeter"
	greeterImpl "github.com/vmware/vsphere-automation-sdk-go/runtime/examples/authentication/server/impl"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/provider/local"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
)

func initAuthenticationFilter(provider core.APIProvider) *security.AuthenticationFilter {
	var metadataDir = "metadata/"
	var metadataFile = "vmodl.examples.greeter_metadata.json"

	var authenticationMetadataDir = []string{metadataDir + metadataFile}

	var userPwdAuthenticationHandler = auth.NewUserPwdAuthenticationHandler()
	var authenticationHandlers = []security.AuthenticationHandler{userPwdAuthenticationHandler}

	authenticationFilter, err := security.NewAuthenticationFilter(authenticationHandlers, provider, authenticationMetadataDir)
	if err != nil {
		panic("Cannot create authentication filter")
	}

	return authenticationFilter
}

func main() {
	var localProvider, err = local.NewLocalProvider("localprovider", true)
	if err != nil {
		panic("Couldn't initialize LocalProvider\n")
	}
	localProvider.AddInterface(greeter.NewGreeterApiInterface(greeterImpl.NewGreeterServiceImpl()))

	var authenticationFilter = initAuthenticationFilter(localProvider)

	var jsonRpcHandler = msg.NewJsonRpcHandler(authenticationFilter)

	server := &http.Server{
		Handler: jsonRpcHandler,
		Addr:    ":8088",
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic("Error occurred while starting the server")
		}
	}()

	fmt.Println("Started server at localhost:8088")
	select {}

}

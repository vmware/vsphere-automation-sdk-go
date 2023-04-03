/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package main

import (
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/examples/stream/server/bindings/src/generated/vmodl/examples/greeter"
	greeterImpl "github.com/vmware/vsphere-automation-sdk-go/runtime/examples/stream/server/impl"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/provider/local"
	"net/http"
)

func main() {
	var localProvider, err = local.NewLocalProvider("localprovider", true)
	if err != nil {
		panic("Couldn't initialize LocalProvider\n")
	}
	localProvider.AddInterface(greeter.NewGreeterApiInterface(greeterImpl.NewGreeterServiceImpl()))

	var jsonRpcHandler = msg.NewJsonRpcHandler(localProvider)

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

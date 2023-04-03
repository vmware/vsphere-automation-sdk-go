/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package main

import (
	"crypto/tls"
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/examples/jwt/server/bindings/vmodl/examples/greeter"
	greeterImpl "github.com/vmware/vsphere-automation-sdk-go/runtime/examples/jwt/server/impl"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/provider/local"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/security"
	"net/http"
)

// Do not rely on the current address to be live
const vcHost = "10.185.17.192"

func initAuthenticationFilter(provider core.APIProvider) *security.AuthenticationFilter {
	var metadataDir = "metadata/"
	var metadataFile = "vmodl.examples.greeter_metadata.json"
	var authenticationMetadataDir = []string{metadataDir + metadataFile}

	var jwtHandler = buildJwtAuthenticationHandler()

	var authenticationHandlers = []security.AuthenticationHandler{jwtHandler}

	authenticationFilter, err := security.NewAuthenticationFilter(authenticationHandlers, provider, authenticationMetadataDir)
	if err != nil {
		panic("Cannot create authentication filter")
	}

	return authenticationFilter
}

func buildVerificationKeyProvider() *security.OidcJwksVerificationKeyCache {
	// For the sake of the example, the cache client will avoid ssl related security checks
	// DO NOT USE IN PRODUCTION
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// If running on VC utilize the default constructor which configures plain-text HTTP on localhost
	provider, _ := security.NewOidcJwksVerificationKeyCache(security.WithClient(client), security.WithHost(vcHost))
	return provider
}

func buildJwtAuthenticationHandler() *security.JwtAuthenticationHandler {
	provider := buildVerificationKeyProvider()
	handler, _ := security.NewJwtAuthenticationHandler(provider)
	return handler
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

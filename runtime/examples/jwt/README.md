## Setup JWT Authentication Handler guide

The `JwtAuthenticationHandler` is an `AuthenticationHandler` delivered by the goruntime,
which parses and validates JWT tokens. The handler requires a `VerificationKeyProvider`, responsible
for delivering the keys used for JWT's signature validation. 

### Setting up `VerificationKeyProvider`
The goruntime project provides an implementation of the contract - `OidcJwksVerificationKeyCache`.
The cache provider can be created as follows:

```go
return security.NewOidcJwksVerificationKeyCache()
```
which returns a key provider with a default client communicating via `HTTP` to `localhost:1080`.
Configuring the key provider with a specific client, vc address and `HTTPS` scheme is done via `CacheConfigOptions`
as follows:

```go
provider, _ := security.NewOidcJwksVerificationKeyCache(security.WithClient(client), security.WithHost(vcHost), security.UseHttps(true))
```
Refer to the documentation of `OidcJwksVerificationKeyCache` for further details.

### Setting up `JwtAuthenticationHandler`
Creating the handler is trivial when a `VerificationKeyProvider` is present:

```go
handler, _ := security.NewJwtAuthenticationHandler(verificationKeyCache)
```
The returned handler contains default values for max clock skew - 10 min., and acceptable
audiences - `["vmware-tes:vapi"]`. Specifying values for these parameters is done in a
similar way as `OidcJwksVerificationKeyCache` configuration, via `JwtHandlerConfigOption`
as follows:

```go
handler, _ := NewJwtAuthenticationHandler(verificationKeyCache, WithMaxClockSkew(3*time.Minute), WithAcceptableAudiences([]string{"aud"}))
```
Refer to the documentation of `JwtAuthenticationHandler` for further details.

### Attaching the `JwtAuthenticationHandler` to the `AuthenticationFilter`
The authentication filter requires authentication metadata, which provides information about
operations and their associated authentication schemes. The metadata is derived from an `authn.ini`
descriptor file such as: 

```
[metadata]
type = authn

[SchemeJwt]
Type=SessionLess
AuthenticationScheme=com.vmware.vapi.std.security.Oauth

[packages]
vmodl.examples.greeter = SchemeJwt
```
The authn.ini is fed to vmodl-compiler (https://confluence.eng.vmware.com/display/VAPI/VMODL2+Compiler)
for metadata generation.

Once metadata is present, along with the authentication handlers, the `AuthenticationFilter` is
instantiated as follows:

```go
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
```

## Running the example
### Starting the greeter server

For a correctly configured server deploy a VC and substitute `vcHost` with the correct address

1. cd goruntime/examples/jwt
2. go build ./...
3. cd server && go run server.go

### Testing out the server with a valid JWT
1. Install govc - https://github.com/vmware/govmomi/tree/master/govc
   (`go install github.com/vmware/govmomi/govc`)
2. Acquire a SAML token via govc and base64 encode it -
   `saml=$(./govc session.login -k -u <vc-user>:<vc-pass>@<your-testing-vc> -issue | base64 -w 0 )`
3. Access \<your-testing-vc\>/apiexplorer and find under `vcenter` API the `authentication/token` 
   operation and click the deprecated button (the deprecated is required because api-explorer does
   not support application/x-www-form-urlencoded)
4. Send the following request via the UI:
   ```json
   {
       "spec": {
           "audience": "vmware-tes:vapi",
           "grant_type": "urn:ietf:params:oauth:grant-type:token-exchange",
           "requested_token_type": "urn:ietf:params:oauth:token-type:id_token",
           "subject_token": "<saml>",
           "subject_token_type": "urn:ietf:params:oauth:token-type:saml2"
       }
   }
   ```
   where \<saml\> is substituted with the saml token from step 2
5. From the response copy the "access_token" (your JWT token)
6. Issue the following curl:
```curl
curl --header "Content-Type: application/json" --header "Accept: application/json" --header "Vapi-Operation: greet" --header "Vapi-Service: vmodl.examples.greeter.greeter" --data-binary '{"id":"e9d7a8cd-9cd0-4bf1-8598-fd9fb74c78c7","jsonrpc":"2.0","method":"invoke","params":{"ctx":{"appCtx":{"opId":"e9d7a8cd-9cd0-4bf1-8598-fd9fb74c78c7"},"securityCtx":{"accessToken":"<jwt>","schemeId":"com.vmware.vapi.std.security.oauth"}},"input":{"STRUCTURE":{"operation-input":{}}},"operationId":"greet","serviceId":"vmodl.examples.greeter.greeter"}}' http://localhost:8088/
```
where \<jwt\> is substituted with the token from step 5
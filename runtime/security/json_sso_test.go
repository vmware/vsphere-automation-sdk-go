package security_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractCertificate(t *testing.T) {
	hokTokenBytes, err := readFile("test/sso/hoktoken.txt")
	if err != nil {
		t.Fail()
	}
	cert, err := security.ExtractCertificate(string(hokTokenBytes))
	assert.Nil(t, err)
	assert.NotNil(t, cert)

}

func readFile(fileLocation string) ([]byte, error) {
	fileBeingRead, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer fileBeingRead.Close()
	readBytes, err := ioutil.ReadAll(fileBeingRead)
	if err != nil {
		return nil, err
	}
	return readBytes, nil
}

func prepareHokRequest() (map[string]interface{}, error) {
	//load request.json
	requestBytes, err := readFile("test/sso/request.json")
	if err != nil {
		return nil, err
	}
	requestJson := map[string]interface{}{}
	unmarshError := json.Unmarshal(requestBytes, &requestJson)
	if unmarshError != nil {
		return nil, unmarshError
	}

	//read private key
	privateKeyBytes, err := readFile("test/sso/key.pem")
	if err != nil {
		return nil, err
	}

	//read hok token
	hokTokenBytes, err := readFile("test/sso/hoktoken.txt")
	if err != nil {
		return nil, err
	}

	//populate security context
	params := requestJson["params"].(map[string]interface{})
	ctx := params["ctx"].(map[string]interface{})
	securityCtx := ctx["securityCtx"].(map[string]interface{})
	securityCtx["schemeId"] = "com.vmware.vapi.std.security.saml_hok_token"
	securityCtx["privateKey"] = string(privateKeyBytes)
	securityCtx["samlToken"] = string(hokTokenBytes)
	securityCtx["signatureAlgorithm"] = "RS256"
	return requestJson, nil
}

func prepareBearerRequest() (map[string]interface{}, error) {
	//load request.json
	requestBytes, err := readFile("test/sso/request.json")
	if err != nil {
		return nil, err
	}
	requestJson := map[string]interface{}{}
	unmarshError := json.Unmarshal(requestBytes, &requestJson)
	if unmarshError != nil {
		return nil, unmarshError
	}

	//read bearer token
	bearerTokenBytes, err := readFile("test/sso/bearertoken.xml")
	if err != nil {
		return nil, err
	}

	//populate security context
	params := requestJson["params"].(map[string]interface{})
	ctx := params["ctx"].(map[string]interface{})
	securityCtx := ctx["securityCtx"].(map[string]interface{})
	securityCtx["schemeId"] = "com.vmware.vapi.std.security.saml_bearer_token"
	securityCtx["samlToken"] = string(bearerTokenBytes)
	securityCtx[security.TIMESTAMP] = security.GenerateRequestTimeStamp()
	return requestJson, nil
}

func createBearerPayloadWithHokToken() (map[string]interface{}, error) {
	//load request.json
	requestBytes, err := readFile("test/sso/request.json")
	if err != nil {
		return nil, err
	}
	requestJson := map[string]interface{}{}
	unmarshError := json.Unmarshal(requestBytes, &requestJson)
	if unmarshError != nil {
		return nil, unmarshError
	}

	//read hok token
	hokTokenBytes, err := readFile("test/sso/hoktoken.txt")
	if err != nil {
		return nil, err
	}

	//populate security context
	params := requestJson["params"].(map[string]interface{})
	ctx := params["ctx"].(map[string]interface{})
	securityCtx := ctx["securityCtx"].(map[string]interface{})
	securityCtx["schemeId"] = "com.vmware.vapi.std.security.saml_bearer_token"
	securityCtx["samlToken"] = string(hokTokenBytes)
	securityCtx[security.TIMESTAMP] = security.GenerateRequestTimeStamp()
	return requestJson, nil
}

func TestSigningAndVerification(t *testing.T) {
	requestJson, err := prepareHokRequest()
	assert.Nil(t, err)
	jsonSsoSigner := security.NewJSONSsoSigner()
	err = jsonSsoSigner.Process(&requestJson)
	assert.Nil(t, err)
	jsonSsoVerifer := security.NewJSONSsoVerifier()
	err = jsonSsoVerifer.Process(&requestJson)
	assert.Nil(t, err)
}

func TestSigningAndVerificationFail(t *testing.T) {
	requestJson, err := prepareHokRequest()
	assert.Nil(t, err)
	jsonSsoSigner := security.NewJSONSsoSigner()
	err = jsonSsoSigner.Process(&requestJson)
	assert.Nil(t, err)
	//change requestJson content so hash is different
	requestJson["hello"] = "world"
	jsonSsoVerifer := security.NewJSONSsoVerifier()
	err = jsonSsoVerifer.Process(&requestJson)
	assert.NotNil(t, err)
}

func TestBearerSecurityCtxPass(t *testing.T) {

	requestJson, err := prepareBearerRequest()
	assert.Nil(t, err)
	bearerTokenVerifier := security.NewBearerTokenSCVerifier()
	err = bearerTokenVerifier.Process(&requestJson)
	assert.Nil(t, err)

}

func TestBearerSecurityCtxFail(t *testing.T) {

	// case 1: when hok request payload is passed
	requestJson, err := prepareHokRequest()
	assert.Nil(t, err)

	bearerTokenVerifier := security.NewBearerTokenSCVerifier()
	err = bearerTokenVerifier.Process(&requestJson)
	assert.Nil(t, err)

	// case 2: when bearer payload is used but with hok token
	requestJson, err = createBearerPayloadWithHokToken()
	assert.Nil(t, err)
	err = bearerTokenVerifier.Process(&requestJson)
	assert.Equal(t, err.Error(), "SAML token passed is not of type Bearer")

}

package refreshtoken

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
)

type info struct {
	cspURL       string
	refreshToken string
}

func (oad *info) CSPURL() string {
	return oad.cspURL
}

func (oad *info) RefreshToken() string {
	return oad.refreshToken
}

func (oad *info) GetName() string {
	return Name
}

func (oad *info) GetAuthInterface() interface{} {
	return oad
}

func (oad *info) GetSecurityContext() (core.SecurityContext, error) {
	payload := strings.NewReader("refresh_token=" + oad.refreshToken)

	req, _ := http.NewRequest("POST", oad.cspURL+CSPRefreshURLSuffix, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("status code %d trying to get csp auth token. %s", res.StatusCode, string(b))
	}

	defer res.Body.Close()

	var jsondata map[string]interface{}
	json.NewDecoder(res.Body).Decode(&jsondata)

	var accessToken string
	if token, ok := jsondata["access_token"]; ok {
		if accessTokenStr, ok := token.(string); ok {
			accessToken = accessTokenStr
		} else {
			errMsg := fmt.Sprintf("Invalid type for access_token, expected string, actual %s", reflect.TypeOf(token).String())
			return nil, errors.New(errMsg)
		}
	} else {
		return nil, errors.New("CSP auth response doesn't contain access token")
	}

	// log.Info("CSP Access token is " + accessToken)
	securityCtx := security.NewOauthSecurityContext(accessToken)
	return securityCtx, nil
}

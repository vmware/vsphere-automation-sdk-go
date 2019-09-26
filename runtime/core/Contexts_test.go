/* **********************************************************
 * Copyright 2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * *********************************************************
 */
package core_test

import (
	"github.com/stretchr/testify/mock"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetAndGetProperty(t *testing.T) {
	app := make(map[string]*string)
	appContext := core.NewApplicationContext(app)
	value1 := "value1"
	appContext.SetProperty("key1", &value1)
	assert.Equal(t, appContext.GetProperty("key1"), &value1)
}

func TestGetAllProperties(t *testing.T) {
	app := make(map[string]*string)
	appContext := core.NewApplicationContext(app)
	value1 := "value1"
	value2 := "value2"
	appContext.SetProperty("key1", &value1)
	appContext.SetProperty("key2", &value2)
	assert.Equal(t, appContext.GetAllProperties()["key1"], &value1)
	assert.Equal(t, appContext.GetAllProperties()["key2"], &value2)
}

type MessageFormatterMock struct {
	mock.Mock
}

func (m *MessageFormatterMock) GetFormatterForLocalizationParams(acceptLanguage *string,
																   formatLanguage *string,
																   timezone *string) (l10n.MessageFormatter, error) {
	// this records that the method was called and passes in the value
	// it was called with
	m.Called(acceptLanguage,formatLanguage,timezone)
	var empty l10n.MessageFormatter
	return empty,nil
}

func (m *MessageFormatterMock) GetDefaultFormatter() *l10n.MessageFormatter {
	var empty l10n.MessageFormatter
	m.Called()
	return &empty
}

func TestGetMessageFormatter(t *testing.T) {
	al := "de"
	fl := "en"
	tz := "Europe/Berlin"
	ctx := core.NewExecutionContext(getAppCtx(&al,&fl,&tz), nil)

	var emptyMf l10n.MessageFormatter

	mockMessageFormatter := MessageFormatterMock{}
	mockMessageFormatter.On("GetFormatterForLocalizationParams",
		mock.Anything,mock.Anything,mock.Anything).Return(emptyMf)

	// need not verify value returned by mock call
	ctx.GetMessageFormatter(&mockMessageFormatter)

	mockMessageFormatter.AssertCalled(t, "GetFormatterForLocalizationParams", &al, &fl, &tz)
}

func getAppCtx(al *string, fl *string, tz *string) *core.ApplicationContext {
	l10nHeaders := map[string]*string{
		lib.HTTP_ACCEPT_LANGUAGE: al,
		lib.VAPI_L10N_FORMAT_LOCALE: fl,
		lib.VAPI_L10N_TIMEZONE: tz,
	}

	appctx := core.NewApplicationContext(l10nHeaders)
	return appctx
}

func TestGetMessageFormatterNilCtx(t *testing.T) {
	var nilCtx *core.ExecutionContext

	var emptyMf l10n.MessageFormatter

	mockMessageFormatter := MessageFormatterMock{}
	mockMessageFormatter.On("GetDefaultFormatter").Return(emptyMf)

	// need not verify value returned by mock call
	nilCtx.GetMessageFormatter(&mockMessageFormatter)

	mockMessageFormatter.AssertCalled(t, "GetDefaultFormatter")
}

func TestGetMessageFormatterNilAppCtx(t *testing.T) {
	ctx := core.NewExecutionContext(nil, nil)

	var emptyMf l10n.MessageFormatter

	mockMessageFormatter := MessageFormatterMock{}
	mockMessageFormatter.On("GetFormatterForLocalizationParams",
		mock.Anything,mock.Anything,mock.Anything).Return(emptyMf)

	// need not verify value returned by mock call
	ctx.GetMessageFormatter(&mockMessageFormatter)

	var emptyStrPtr *string
	mockMessageFormatter.AssertCalled(t, "GetFormatterForLocalizationParams",
		emptyStrPtr, emptyStrPtr, emptyStrPtr)
}

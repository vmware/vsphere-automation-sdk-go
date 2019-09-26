//Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
package l10n_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"golang.org/x/text/language"
	"testing"
	"time"
)

var testBundleMessageFactory = getTestBundleMessageFactory()
var testDefaultFormatter = l10n.NewDefaultMessageFormatter(l10n.NewMessageFactory())

func TestDefaults(t *testing.T) {
	expectedFormatter := testFormatterWithSettings(
		language.English, language.English, "UTC", int64(2), l10n.DateTimeFormat_SHORT_DATE_TIME)

	assert.Equal(t, testDefaultFormatter, expectedFormatter)
}

func testFormatterWithSettings(al language.Tag, fl language.Tag, tz string, precision int64, df string) l10n.MessageFormatter {
	timezone, _ := time.LoadLocation(tz)
	mfWithSettings := l10n.NewMessageFactoryWithSettings(&df, &precision)
	return l10n.NewMessageFormatter(&al, &fl, timezone, mfWithSettings)
}

func TestFormatInterfaceParam(t *testing.T) {
	formatter := testFormatterWithSettings(
		language.English, language.English, "EST", int64(3), l10n.DateTimeFormat_FULL_DATE_TIME)

	mst, _ := time.LoadLocation("MST")
	assert.Equal(t, "2019 Apr 05, Fri 15:30:02 EST",
		formatter.FormatInterfaceParam(time.Date(2019, time.April, 5, 13, 30, 02, 01, mst)))

	assert.Equal(t, "1.235",
		formatter.FormatInterfaceParam(float32(1.234567)))

	assert.Equal(t, "1.235",
		formatter.FormatInterfaceParam(1.234567))

	assert.Equal(t, "abc",
		formatter.FormatInterfaceParam("abc"))

	assert.Equal(t, "123",
		formatter.FormatInterfaceParam(123))

	assert.Equal(t, "123",
		formatter.FormatInterfaceParam(int32(123)))

	assert.Equal(t, "[1 2 3 4]",
		formatter.FormatInterfaceParam([]int64{1, 2, 3, 4}))
}

func TestFormatFloat(t *testing.T) {
	p1 := int64(1)
	p2 := int64(2)
	p4 := int64(4)
	assert.Equal(t, "1.2",
		testDefaultFormatter.FormatFloat(1.234567, &p1))
	assert.Equal(t, "1.23",
		testDefaultFormatter.FormatFloat(1.234567, &p2))
	assert.Equal(t, "1.2346",
		testDefaultFormatter.FormatFloat(1.234567, &p4))
	assert.Equal(t, "1.23",
		testDefaultFormatter.FormatFloat(1.234567, nil))
}

func TestFormatDateTime(t *testing.T) {
	est, _ := time.LoadLocation("EST")
	testDate := time.Date(2019, time.April, 5, 13, 30, 02, 01, est)
	fullDt := l10n.DateTimeFormat_FULL_DATE_TIME
	shortDt := l10n.DateTimeFormat_SHORT_DATE
	assert.Equal(t, "2019 Apr 05, Fri 18:30:02 UTC",
		testDefaultFormatter.FormatDateTime(testDate, &fullDt))
	assert.Equal(t, "2019-04-05",
		testDefaultFormatter.FormatDateTime(testDate, &shortDt))
	assert.Equal(t, "2019-04-05 18:30",
		testDefaultFormatter.FormatDateTime(testDate, nil))

	formatter := testFormatterWithSettings(
		language.English, language.English, "MST", int64(3), l10n.DateTimeFormat_FULL_DATE_TIME)
	assert.Equal(t, "2019 Apr 05, Fri 11:30:02 MST",
		formatter.FormatDateTime(testDate, nil))
}

func TestPositionalParamsEnglish(t *testing.T) {
	formatter := getFormatter(language.English)

	assert.Equal(t, "Invalid input to method",
		formatter.GetPositionalArgMessage("msg.no.param"))
	assert.Equal(t, "Could not connect to 12",
		formatter.GetPositionalArgMessage("msg.one.param", 12))
	assert.Equal(t, "Invalid Type, expected atype, instead got btype",
		formatter.GetPositionalArgMessage("msg.two.param", "atype", "btype"))
	params := make([]interface{}, 2)
	params[0] = "atype"
	params[1] = "btype"
	assert.Equal(t, "Invalid Type, expected atype, instead got btype",
		formatter.GetPositionalArgMessage("msg.two.param", params...))
	assert.Equal(t, "JSON object has extra field(s)= \"fff\"",
		formatter.GetPositionalArgMessage("msg.eq.in.param", "fff"))
}

func TestPositionalParamsGerman(t *testing.T) {
	formatter := getFormatter(language.German)

	assert.Equal(t, "Ungültige Eingabemethode",
		formatter.GetPositionalArgMessage("msg.no.param"))
	assert.Equal(t, "Verbindung mit 12 nicht möglich",
		formatter.GetPositionalArgMessage("msg.one.param", 12))
	assert.Equal(t, "Ungültiger Typ: „atype“ erwartet, „btype“ erhalten",
		formatter.GetPositionalArgMessage("msg.two.param", "atype", "btype"))
	assert.Equal(t, "JSON-Objekt weist mindestens ein zusätzliches Feld auf: fff",
		formatter.GetPositionalArgMessage("msg.eq.in.param", "fff"))
}

func TestNamedParamsEnglish(t *testing.T) {
	formatter := getFormatter(language.English)
	oneParam := map[string]string{
		"intParam": "24",
	}
	assert.Equal(t, "Could not connect to 24",
		formatter.GetLocalizedMessage("named.one.param", oneParam))

	twoParams := map[string]string{
		"stringParam": "atype",
		"floatParam":  "1.6",
	}
	assert.Equal(t, "Invalid Type, expected atype, instead got 1.6",
		formatter.GetLocalizedMessage("named.two.param", twoParams))

}

func TestInvalidMessageIdEnglish(t *testing.T) {
	formatter := getFormatter(language.English)
	assert.Equal(t, "Unknown message ID invalid.key requested",
		formatter.GetPositionalArgMessage("invalid.key"))
}

func TestInvalidMessageIdGerman(t *testing.T) {
	formatter := getFormatter(language.German)
	assert.Equal(t, "Unknown message ID invalid.key requested",
		formatter.GetPositionalArgMessage("invalid.key"))
}

func TestGetFormatterForLocalizationParams(t *testing.T) {
	al := "de"
	fl := "en"
	tz := "Europe/Berlin"

	formatter, _ := testBundleMessageFactory.GetFormatterForLocalizationParams(&al, &fl, &tz)

	assert.Equal(t, "Verbindung mit 12 nicht möglich",
		formatter.GetPositionalArgMessage("msg.one.param", 12))

	assert.NotNil(t, formatter)
}

func getTestBundleMessageFactory() l10n.MessageFactory {
	bundles := map[string]string{
		"en": "tests/testBundle_en.properties",
		"de": "tests/testBundle_de.properties",
	}
	messageFactory := l10n.NewMessageFactory()
	messageFactory.AddBundles(bundles)

	return messageFactory
}

func getFormatter(tag language.Tag) l10n.MessageFormatter {
	return l10n.NewMessageFormatter(&tag, nil, nil, testBundleMessageFactory)
}

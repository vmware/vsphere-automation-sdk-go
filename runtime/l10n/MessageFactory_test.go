//Copyright 2019 VMware, Inc.  All rights reserved. -- VMware Confidential
package l10n_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"golang.org/x/text/language"
)

func loadFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func TestAddBundleInvalidLocale(t *testing.T) {
	reader, _ := loadFile("tests/testBundle_en.properties")
	err := l10n.NewMessageFactory().AddBundle("blah", reader)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "language: tag is not well-formed")
}

func TestAddBundleNoKey(t *testing.T) {
	reader, err1 := loadFile("tests/errorBundle_no_key_en.properties")
	assert.Nil(t, err1)
	err := l10n.NewMessageFactory().AddBundle("en", reader)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No key for val")
	defer reader.Close()
}

func TestAddBundleNoValue1(t *testing.T) {
	reader, _ := loadFile("tests/errorBundle_no_val1_en.properties")
	err := l10n.NewMessageFactory().AddBundle("en", reader)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No value for key")
	defer reader.Close()
}

func TestAddBundleNoValue2(t *testing.T) {
	reader, _ := loadFile("tests/errorBundle_no_val1_en.properties")
	err := l10n.NewMessageFactory().AddBundle("en", reader)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No value for key")
	defer reader.Close()
}

func TestNoDefaultBundle1(t *testing.T) {
	bundles := map[string]string{
		"de": "testBundle_de.properties",
	}
	_, err := l10n.NewMessageFactoryWithBundle(bundles, nil, nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No english messages available")
}

func TestNoDefaultBundle2(t *testing.T) {
	mf := l10n.NewMessageFactory()
	reader, _ := loadFile("tests/testBundle_de.properties")
	mf.AddBundle("de", reader)
	_, err := mf.GetFormatterForLocalizationParams(nil, nil, nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No english messages available")
}

func TestAddBundles(t *testing.T) {
	bundles := map[string]string{
		"en": "tests/testBundle_en.properties",
		"de": "tests/testBundle_de.properties",
	}
	messageFactory := l10n.NewMessageFactory()
	err := messageFactory.AddBundles(bundles)
	assert.Nil(t, err)

	catalog := messageFactory.Catalog()
	availableLanguages := catalog.Languages()
	assert.Equal(t, len(availableLanguages), 2)
	assert.Contains(t, availableLanguages, language.English)
	assert.Contains(t, availableLanguages, language.German)

	formatter, _ := messageFactory.GetFormatterForLocalizationParams(nil, nil, nil)

	assert.Equal(t, "Invalid input to method",
		formatter.GetPositionalArgMessage("msg.no.param"))

	al := "de"
	fl := "en"
	tz := "UTC"
	formatter, _ = messageFactory.GetFormatterForLocalizationParams(&al, &fl, &tz)
	assert.Equal(t, "Verbindung mit 12 nicht möglich",
		formatter.GetPositionalArgMessage("msg.one.param", 12))
}

func TestAddBundle(t *testing.T) {
	messageFactory := l10n.NewMessageFactory()
	reader, _ := loadFile("tests/testBundle_en.properties")
	err := messageFactory.AddBundle("en", reader)
	assert.Nil(t, err)

	catalog := messageFactory.Catalog()
	availableLanguages := catalog.Languages()
	assert.Equal(t, len(availableLanguages), 1)
	assert.Contains(t, availableLanguages, language.English)

	formatter := l10n.NewMessageFormatter(nil, nil, nil, messageFactory)

	assert.Equal(t, "Invalid input to method",
		formatter.GetPositionalArgMessage("msg.no.param"))

	al := "de"
	fl := "en"
	tz := "UTC"
	formatter, _ = messageFactory.GetFormatterForLocalizationParams(&al, &fl, &tz)

	assert.Equal(t, "Invalid input to method",
		formatter.GetPositionalArgMessage("msg.no.param"))
}

package lang

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Lang struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

func Init() *Lang {
	b := newBundle()
	return &Lang{
		bundle:    b,
		localizer: i18n.NewLocalizer(b, "fa"),
	}
}

func newBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.Persian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("fa.json")
	return bundle
}

func (l *Lang) T(id string, data ...*map[string]any) string {
	config := i18n.LocalizeConfig{
		MessageID: id,
	}
	if len(data) > 0 {
		config.TemplateData = data[0]
	}
	text, _ := l.localizer.Localize(&config)
	return text
}

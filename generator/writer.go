package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func WriteLocale(localePath string, locale *Locale) {
	path := localePath + "/" + locale.Code

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			log.Panic(err)
		}
	}

	localeFilepath := path + "/" + locale.Code + ".go"

	if f, err := os.Create(localeFilepath); err != nil {
		log.Panic(err)
	} else {
		defer f.Close()

		if err := WriteLocaleGo(locale, f); err != nil {
			log.Panic(err)
		}
	}

	cmd := exec.Command("goimports", "-w", localeFilepath)
	if err := cmd.Run(); err != nil {
		log.Panic("failed execute \"goimports\" for file ", localeFilepath, ": ", err)
	}

	cmd = exec.Command("gofmt", "-s", "-w", localeFilepath)
	if err := cmd.Run(); err != nil {
		log.Panic("failed execute \"gofmt\" for file ", localeFilepath, ": ", err)
	}
}

func WriteLocaleGo(locale *Locale, w io.Writer) error {
	fs := GetEmbedFS()

	tpl := template.Must(template.ParseFS(fs, "templates/locale.tmpl"))

	ctx := map[string]interface{}{
		"Locale":      locale,
		"Code":        locale.Code,
		"Territories": locale.Territories,
		"Currencies":  locale.Currencies,
		"TimeZones":   locale.TimeZones,
	}

	return tpl.Execute(w, ctx)
}

func WriteTimezonesGo(basePath string, locale *Locale) error {
	fs := GetEmbedFS()

	var f *os.File
	var err error

	if f, err = os.Create(basePath + "/timezones.go"); err != nil {
		log.Panic(err)
	}

	defer f.Close()

	tpl := template.Must(template.ParseFS(fs, "templates/timezones.tmpl"))

	ctx := map[string]interface{}{
		"Locale": locale,
	}

	return tpl.Execute(f, ctx)

}

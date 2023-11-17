package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"text/template"
)

func WriteModule(localePath string, locale *Locale) {
	path := localePath + "/" + locale.Code

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			log.Panic(err)
		}
	}

	localeFilepath := path + "/" + locale.Code + ".go"

	if fs, err := os.Create(localeFilepath); err != nil {
		log.Panic(err)
	} else {
		if err := WriteGo(locale, fs); err != nil {
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

func WriteGo(locale *Locale, w io.Writer) error {
	fs := GetEmbedFS()

	tpl := template.Must(template.ParseFS(fs, "templates/data.tmpl"))

	ctx := map[string]interface{}{
		"Locale":      locale,
		"Code":        locale.Code,
		"Territories": locale.Territories,
		"Currencies":  locale.Currencies,
		"TimeZones":   locale.TimeZones,
		"Symbols":     locale.Symbols,
	}

	return tpl.Execute(w, ctx)
}

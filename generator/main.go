package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

//go:embed all:templates
var content embed.FS

func GetEmbedFS() embed.FS {
	return content
}

func main() {
	CldrPath := os.Getenv("CLDR_DIR")

	if CldrPath == "" {
		fmt.Printf("CLDR_DIR is not set\n")
		os.Exit(-1)
		return
	}

	LocalePath := os.Getenv("LOCALE_DIR")
	if LocalePath == "" {
		fmt.Printf("LOCALE_DIR is not set\n")
		os.Exit(-1)
		return
	}

	cldr := LoadCLDR(CldrPath)

	fmt.Printf("\nLoading root locale\n")
	cldr.RootLocale = LoadLocaleFromFile(CldrPath+"/main/root.xml", cldr)

	WriteModule(LocalePath, cldr.RootLocale)

	list := []string{"en.xml", "fr.xml", "fr_CA.xml", "sr.xml"}

	fmt.Printf("\nLoading locales\n")
	filepath.Walk(CldrPath+"/main", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() == "root.xml" {
			return nil
		}

		if !slices.Contains(list, info.Name()) {
			return nil
		}

		fmt.Printf("\n --- [ %s ] ---\n", info.Name())

		code := strings.Split(info.Name(), ".")[0]
		parts := strings.Split(code, "_")

		if len(parts) > 2 { // no supported for now
			fmt.Printf("> Skipping non supported case %s\n", code)
			return nil
		}

		fmt.Printf("> Parsing supported case %s\n", info.Name())

		if len(parts) == 2 {
			langCode := parts[0]
			// check if the base lang exist, if not created too
			langModulePath := LocalePath + "/" + langCode
			if _, err := os.Stat(langModulePath); os.IsNotExist(err) {
				locale := LoadLocaleFromFile(CldrPath+"/main/"+langCode+".xml", cldr)
				fmt.Printf("> Generate file base module %s\n", locale.Code)
				WriteModule(LocalePath, locale)

				cldr.Locales[locale.Code] = locale
			}
		}

		locale := LoadLocaleFromFile(path, cldr)
		cldr.Locales[locale.Code] = locale

		fmt.Printf("> Generate file module %s\n", locale.Code)
		WriteModule(LocalePath, locale)

		return nil
	})
}

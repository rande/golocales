package main

import (
	"embed"
	"fmt"
	"log"
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

	cldr := &CLDR{}
	cldr.Path = CldrPath

	// load validity files
	validityFiles := []string{
		"currency.xml",
		"language.xml",
		"region.xml",
		// "script.xml",
		// "subdivision.xml",
		// "unit.xml",
		// "variant.xml",
	}

	// validities are required to load root module
	for _, file := range validityFiles {
		fmt.Printf(" > Loading validity file: %s\n", file)

		supplemental := &SupplementalData{}
		if err := LoadXml(CldrPath+"/validity/"+file, supplemental); err != nil {
			log.Panic(err.Error())
		}

		AttachSupplemental(cldr, supplemental)
	}

	// load supplemental files
	supplementalFiles := []string{
		// "attributeValueValidity.xml",
		// "characters.xml",
		// "coverageLevels.xml",
		// "dayPeriods.xml",
		// "genderList.xml",
		// "grammaticalFeatures.xml",
		// "languageGroup.xml",
		// "languageInfo.xml",
		// "likelySubtags.xml",
		"metaZones.xml",
		// "numberingSystems.xml",
		// "ordinals.xml",
		// "pluralRanges.xml",
		// "plurals.xml",
		// "rgScope.xml",
		// "subdivisions.xml",
		// "supplementalData.xml",
		// "supplementalMetadata.xml",
		// "units.xml",
		// "windowsZones.xml",
	}

	for _, file := range supplementalFiles {
		fmt.Printf(" > Loading supplemental file: %s\n", file)

		supplemental := &SupplementalData{}
		if err := LoadXml(CldrPath+"/supplemental/"+file, supplemental); err != nil {
			log.Panic(err.Error())
		}

		AttachSupplemental(cldr, supplemental)
	}

	fmt.Printf("\nLoading root locale\n")
	cldr.RootLocale = LoadLocaleFromFile(CldrPath+"/main/root.xml", cldr)

	WriteModule(LocalePath, cldr.RootLocale)

	list := []string{"en.xml", "fr.xml", "root.xml"}

	fmt.Printf("\nLoading locales\n")
	filepath.Walk(CldrPath+"/main", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || info.Name() == "root.xml" {
			return nil
		}

		if !slices.Contains(list, info.Name()) {
			return nil
		}
		fmt.Printf("> Parsing supported case %s\n", info.Name())

		code := strings.Split(info.Name(), ".")[0]
		parts := strings.Split(code, "_")

		if len(parts) > 2 { // no supported for now
			fmt.Printf("Skipping non supported case %s\n", code)
			return nil
		}

		// fr, en, ... skip this will be created by another file, ie: fr_FR, en_GB
		// if len(parts) == 1 {
		// 	return nil
		// }

		// langCode := parts[0]

		// // check if the base lang exist, if not created too
		// langModulePath := LocalePath + "/" + langCode
		// if _, err := os.Stat(langModulePath); os.IsNotExist(err) {
		// 	locale := LoadLocaleFromFile(CldrPath+"/main/"+langCode+".xml", cldr)
		// 	// fmt.Printf("Generate file %s\n", locale.Code)
		// 	WriteModule(LocalePath, locale)
		// }

		locale := LoadLocaleFromFile(path, cldr)

		if !locale.IsBase {
			// fmt.Printf("Skipping territory %s for now\n", locale.Code)
			return nil
		}

		// fmt.Printf("Generate file %s\n", locale.Code)
		WriteModule(LocalePath, locale)

		return nil
	})
}

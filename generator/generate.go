package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	err := filepath.Walk(CldrPath+"/main", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ldml, err := LoadLdml(path)

		if err != nil {
			log.Panic(err.Error())
		}

		locale := BuildLocale(ldml)

		if locale.IsTerritory() {
			//fmt.Printf("Skipping territory %s\n", locale.Code)
			return nil
		}

		fmt.Printf("Generate file %s\n", locale.Code)

		localePath := LocalePath + "/" + locale.Code

		if _, err := os.Stat(localePath); os.IsNotExist(err) {
			if err := os.Mkdir(localePath, 0755); err != nil {
				log.Panic(err)
			}
		}

		localeFilepath := localePath + "/" + locale.Code + ".go"

		if fs, err := os.Create(localeFilepath); err != nil {
			log.Panic(err)
		} else {
			if err := WriteGo(locale, fs); err != nil {
				log.Panic(err)
			}
		}

		cmd := exec.Command("goimports", "-w", localeFilepath)
		if err = cmd.Run(); err != nil {
			log.Panic("failed execute \"goimports\" for file ", localeFilepath, ": ", err)
		}

		cmd = exec.Command("gofmt", "-s", "-w", localeFilepath)
		if err = cmd.Run(); err != nil {
			log.Panic("failed execute \"gofmt\" for file ", localeFilepath, ": ", err)
		}

		// // Create a new directory named "example"

		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println("Directory created successfully")
		// }

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", CldrPath+"/main", err)
		return
	}
}

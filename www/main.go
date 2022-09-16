// Handles some amount of voidedtech.com hosting
package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var (
	//go:embed index.html
	indexHTML string
	//go:embed main.css
	mainCSS []byte
)

type (
	hostData struct {
		mainIndex []byte
		mainCSS   []byte
	}
)

func createHostData(date, file string) (hostData, error) {
	d, err := os.ReadFile(file)
	if err != nil {
		return hostData{}, err
	}
	jFile := SiteData{}
	if err := json.Unmarshal(d, &jFile); err != nil {
		return hostData{}, err
	}
	obj := &SiteData{Date: date, IsPublic: true, Title: "voidedtech", Header: "About me:"}
	obj.Links = jFile.Links
	obj.Places = jFile.Places
	obj.Mail = jFile.Mail
	tmpl, err := template.New("t").Parse(indexHTML)
	if err != nil {
		return hostData{}, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		return hostData{}, err
	}
	return hostData{mainIndex: buf.Bytes(), mainCSS: mainCSS}, nil
}

func main() {
	config := flag.String("config", "", "site json definition")
	target := flag.String("target", "", "target output")
	date := flag.String("date", "", "setup date")
	flag.Parse()
	data, err := createHostData(*date, *config)
	if err != nil {
		die("unable to create host data", err)
	}
	dest := *target
	if err := os.WriteFile(filepath.Join(dest, "index.html"), data.mainIndex, 0644); err != nil {
		die("failed to write index", err)
	}
	if err := os.WriteFile(filepath.Join(dest, "main.css"), data.mainCSS, 0644); err != nil {
		die("failed to write css", err)
	}
}

func die(message string, err error) {
	fmt.Printf("%s: %v", message, err)
	os.Exit(1)
}
